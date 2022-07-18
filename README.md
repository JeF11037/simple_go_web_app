## Simple application having CRUD operations under customer object

## How to run ?
- Download or copy repository
- The project uses environment variable "connection" as postgresql string connection
  - Something like this: *user=postgres password=12345678 host=localhost port=5432 dbname=database_name connect_timeout=20 sslmode=disable*
  - Accordingly, you need to have a pre-prepared database to which there will be a connection
    - Database only. The table will generate itself
- No external dependencies
- go run main.go :D
- url: localhost:8080
