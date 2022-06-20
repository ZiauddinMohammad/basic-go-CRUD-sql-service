# Base image
FROM mysql:latest

# Add a database name
ENV MYSQL_DATABASE bookstore

ENV MYSQL_ROOT_PASSWORD=supersecret

# Copy all the contents of the local scripts folder to docker-entrypoint-initdb.d
# All scripts in docker-entrypoint-initdb.d will be executed during container startup
COPY ./sqlfiles ./docker-entrypoint-initdb.d
