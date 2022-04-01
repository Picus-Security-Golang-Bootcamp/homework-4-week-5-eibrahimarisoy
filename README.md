## Homework | Week 5
# Book Store API

This repository contains Book Store API written by Go. It is a RESTful API that provides CRUD operations for books and authors.

# Using Tools
 - Gorilla Mux
 - Postgres
 - net/http

## Clone the project
```
$ git clone https://github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy.git
$ cd homework-4-week-5-eibrahimarisoy
```

# Database
App uses PostgreSQL database. You must set up the database before running the application.
Fill .env file with your database credentials.
```
    DB_HOST=
    DB_PORT=
    DB_USERNAME=
    DB_NAME=
    DB_PASSWORD=
```

 ## Run the project
```
$ go run main.go
```

## Routes
Default **BookStore API** routes are listed below. 

| METHOD  | ROUTE               | FUNCTIONS                      |
|---------|---------------------|--------------------------------|
| GET     | /authors/           | (*Controller).ListAuthors      |
| GET     | /authors/:id        | (*Controller).GetAuthor        |
| POST    | /authors/           | (*Controller).CreateAuthor     |
| PUT     | /authors/:id        | (*Controller).UpdateAuthor     |
| DELETE  | /authors/:id        | (*Controller).DeleteAuthor     |
| GET     | /authors/:id/books  | (*Controller).ListAuthorsBooks |
| GET     | /books/             | (*Controller).ListBooks        |
| GET     | /books/:id          | (*Controller).GetBook          |
| POST    | /books/             | (*Controller).CreateBook       |
| PUT     | /books/:id          | (*Controller).UpdateBook       |
| DELETE  | /books/:id          | (*Controller).DeleteBook       |
| POST    | /books/:id/buy      | (*Controller).BuyBook          |

# Models
## Book Model
```
{
    "name": string,
    "pages": uint,
    "stock_count": uint,
    "price": float64,
    "stock_code": string,
    "ISBN": string,
    "author_id": uint

}
```
## Author Model
```
{
    "name": string
}
```