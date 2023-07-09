## Scissor

Scissor is an app  for shorten url built with go, postgresql and redis

#### Steps
 1. Ensure you have go installed
 2. Ensure you have docker and docker compose installed
 3. create a .env file with the following variables: - 
     - PORT
     - SERVER_TIMEOUT_READ
     - SERVER_TIMEOUT_WRITE
     - SERVER_TIMEOUT_IDLE
     - SERVER_DEBUG
     - SERVER_URL
     - JWT_ACCESS_SECRET
     - JWT_REFRESH_SECRET
     - JWT_ISSUER
     - JWT_AUDIENCE
     - POSTGRES_DB
     - POSTGRES_USER
     - POSTGRES_PASSWORD
     - REDIS_URI
     - DATABASE_URL
     - APP_ENV 
4. Run ```docker-compose up``` to start
