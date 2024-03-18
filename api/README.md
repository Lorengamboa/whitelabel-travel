## How to connect to the database from pgadmin

Host name/address = host.docker.internal

## How to run the initial migration

migrate -path=./migrations -database="postgresql://postgres:postgres@localhost:5432/whitelabel-travel?sslmode=false" up
