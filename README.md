# basic-go-CRUD-sql-service

# Run a local sql service in docker before running the application
cd to root directory of this repo

docker build -t my-local-sql .

docker run -d -p 3306:3306 --name mysql-container my-local-sql
