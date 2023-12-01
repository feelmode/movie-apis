# Movies API

This repository contains the REST APIs for `/movies` entity. The layout of this repository is based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Requirements
- Go (works well with Go v1.20)
- PostgreSQL (works well with PostgreSQL v14)

## Deployment (Dev Mode)

### Clone the Repository

```
git clone <repo_url>
```

### Create Database

Create a new database named `movies`.

### Execute DDLs

To create the table, copy the contents of `/scripts/schema.sql`, then paste and execute in a PostgreSQL console.

### Run the Tests

This command will also shows the test coverage. Expect it to be 100%.

> WARNING: It will truncate the `movies` table.

```
go test -cover ./internal/movie/http/handler/ -count=1
```

### Run the HTTP Server

From inside the repo directory, run the following command:
```
go run main.go
```

## List of APIs and Examples

Below are the list of `curl` commands to execute the APIs. To test an API with Postman, copy the command then do the import from Postman app.

### Create a Movie

```
curl --location 'localhost:8000/movies' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Title 1",
    "description": "Desc 1",
    "rating": 7,
    "image": "image1.jpg"
}'
```

Get an error because required fields not provided:

```
curl --location 'localhost:8000/movies' \
--header 'Content-Type: application/json' \
--data '{
    "rating": 7,
    "image": "image1.jpg"
}'
```

### Get all Movies
```
curl --location 'localhost:8000/movies'
```

### Get a Movie with ID = 1
```
curl --location 'localhost:8000/movies/1'
```

Get a 404 error because movie with ID = 1000 not found:
```
curl --location 'localhost:8000/movies/1000'
```

### Update a Movie

```
curl --location --request PATCH 'localhost:8000/movies/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Title 1a",
    "description": "Desc 1a",
    "rating": 9,
    "image": "image1a.jpg"
}'
```

### Delete a Movie
```
curl --location --request DELETE 'localhost:8000/movies/1'
```
