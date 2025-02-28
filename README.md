# linkqu-be-go

1. cloe repo
2. create .env file and copy contents from .env.example file
3. change variable value to set connetion database in .env
4. create your database and insert this query to create table
   CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   age INT NOT NULL,
   city VARCHAR(255) NOT NULL
   );
5. go run main.go
