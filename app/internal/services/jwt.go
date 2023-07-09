package services

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kingsbloc/scissor/internal/config"
	"github.com/kingsbloc/scissor/internal/utils"
	"gopkg.in/square/go-jose.v2"
)

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

type JWTAuthContext struct {
	m map[string]string
}

func (v JWTAuthContext) Get(key string) string {
	return v.m[key]
}

type JwtService interface {
	GenerateJWT(email string, id string, isRefresh bool) (tokenString string, error error)
	VerifyJWT(tokenString string, isRefresh bool) (*jwt.Token, error)
	EncryptJWT(jwtToken string) (any, error)
	JWTAuth(next http.Handler) http.Handler
	JWTAuthPassive(next http.Handler) http.Handler
	GetJWTAuthContext(r *http.Request) JWTAuthContext
}

type jwtService struct{}

func NewJwtService() JwtService {
	return &jwtService{}
}

var c = config.New()

var (
	JWT_ACCESS_SECRET  = c.Jwt.Access_secret
	JWT_REFRESH_SECRET = c.Jwt.Refresh_secret
	JWT_ISSUER         = c.Jwt.Issuer
	JWT_AUDIENCE       = c.Jwt.Audience
)

var jwtAccessSecret = []byte(c.Jwt.Access_secret)
var jwtRefreshSecret = []byte(c.Jwt.Refresh_secret)

type JWTClaim struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

func (j *jwtService) GenerateJWT(email string, id string, isRefresh bool) (tokenString string, error error) {
	var expirationTime time.Time = time.Now().Add(120 * time.Minute)
	if isRefresh {
		expirationTime = time.Now().Add(168 * time.Hour)
	}
	uuidStr := uuid.New()
	claims := &JWTClaim{
		Email: email,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			Id:        uuidStr.String(),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    JWT_ISSUER,
			Audience:  JWT_AUDIENCE,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if isRefresh {
		tokenString, err := token.SignedString(jwtRefreshSecret)
		return tokenString, err
	} else {
		tokenString, err := token.SignedString(jwtAccessSecret)
		return tokenString, err
	}
}

func (j *jwtService) VerifyJWT(tokenString string, isRefresh bool) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("invalid signing method"))
		}
		if _, ok := t.Claims.(jwt.MapClaims); !ok && !t.Valid {
			return nil, fmt.Errorf(("expired token"))
		}
		aud := JWT_AUDIENCE
		checkAudience := t.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAudience {
			return nil, fmt.Errorf(("invalid aud"))
		}
		iss := JWT_ISSUER
		checkIss := t.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return nil, fmt.Errorf(("invalid iss"))
		}
		if isRefresh {
			return jwtRefreshSecret, nil
		}
		return jwtAccessSecret, nil
	})
	return token, err
}

func (j *jwtService) EncryptJWT(jwtToken string) (any, error) {
	reciept := jose.Recipient{
		Algorithm:  jose.PBES2_HS256_A128KW,
		Key:        JWT_ACCESS_SECRET,
		PBES2Count: 4096,
		PBES2Salt:  []byte{},
	}
	enc, err := jose.NewEncrypter(jose.A128CBC_HS256, reciept, nil)

	if err != nil {
		return nil, err
	}
	encJWT, err := enc.Encrypt([]byte(jwtToken))
	if err != nil {
		return nil, err
	}

	key, err := encJWT.CompactSerialize()
	if err != nil {
		return nil, err
	}
	return key, err
}

func (s *jwtService) JWTAuthPassive(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwtAuthContext = JWTAuthContext{map[string]string{
			"email": "",
			"id":    "",
		}}

		tokenString := extractToken(r)
		token, err := s.VerifyJWT(tokenString, false)
		if err != nil {
			ctx := context.WithValue(r.Context(), config.JWTAuthContext, jwtAuthContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			ctx := context.WithValue(r.Context(), config.JWTAuthContext, jwtAuthContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		jwtAuthContext = JWTAuthContext{map[string]string{
			"email": claims["email"].(string),
			"id":    claims["id"].(string),
		}}

		ctx := context.WithValue(r.Context(), config.JWTAuthContext, jwtAuthContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *jwtService) JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwtAuthContext = JWTAuthContext{map[string]string{
			"email": "",
			"id":    "",
		}}

		tokenString := extractToken(r)
		token, err := s.VerifyJWT(tokenString, false)
		if err != nil {
			render.Render(w, r, &utils.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
				Success: false,
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			render.Render(w, r, &utils.ApiResponse{
				Status:  http.StatusUnauthorized,
				Message: "Invalid Token",
				Success: false,
			})
			return
		}

		jwtAuthContext = JWTAuthContext{map[string]string{
			"email": claims["email"].(string),
			"id":    claims["id"].(string),
		}}

		ctx := context.WithValue(r.Context(), config.JWTAuthContext, jwtAuthContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *jwtService) GetJWTAuthContext(r *http.Request) JWTAuthContext {
	return r.Context().Value(config.JWTAuthContext).(JWTAuthContext)
}
