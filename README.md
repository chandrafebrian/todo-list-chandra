# Docker | chandrafebrian99@gmail.com - Making TODO App

This TODO app is a sample app to demonstrate the process of creating an app with Docker, consisting of:
- Frontend with [Vue JS framework]
- Backend with [Golang]
- Database with Postgres

All components are packaged with docker

## Running Postgres with Docker



Step 1: Run Postgres
```bash
docker run -p 5432:5432 --name todo-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=rahasia -e POSTGRES_DB=belajar -v $(pwd)/postgres/init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres
```

Step 2: Export database configuration as environment variable
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=rahasia
export DB_DATABASE=belajar
```

Step 3: Run Backend Golang app on local computer
```bash
go run backend/main.go
```

Step 4: Run Frontend JS app on local computer
```bash
cd frontend
yarn install
yarn serve
```

Step 5: Buka browser untuk mulai mengakses app TODO
```bash
http://localhost:8081
```

Step 6: Check the TODO input stored in the database
```bash
docker exec -it todo-postgres psql -U postgres -W belajar
SELECT * FROM todo;
```

## Running Everything as a Docker Container



Step 1: Create Docker Network
```bash
docker network create todo
```

Step 2: Run Docker postgres on the network
```bash
docker run -d \
-p 5432:5432 \
--name todo-postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=rahasia \
-e POSTGRES_DB=belajar \
-v $(pwd)/postgres/init.sql:/docker-entrypoint-initdb.d/init.sql \
--network todo \
postgres
```

Step 3: Create a docker image for the backend
```bash
cd backend
docker build -t todo-backend:v1 .
```

Step 4: Run the backend as a docker container on the network
```bash
docker run -d \
-p 8080:8080 \
--name todo-backend \
-e DB_USER=postgres \
-e DB_PASSWORD=rahasia \
-e DB_HOST=todo-postgres \
-e DB_PORT=5432 \
-e DB_DATABASE=belajar \
--network todo \
todo-backend:v1
```

Step 5: Create docker image for frontend
```bash2
cd frontend
docker build -t todo-frontend:v1 .
```

Step 6: Start frontend as docker container
```bash
docker run -d \
-p 8081:8080 \
--name todo-frontend \
--network todo \
todo-frontend:v1
```


