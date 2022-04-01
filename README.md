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

| METHOD  | ROUTE               | FUNCTION                                                                                                             |
|---------|---------------------|----------------------------------------------------------------------------------------------------------------------|
| GET     | /authors/           | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).GetAuthors       |
| GET     | /authors/:id        | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).GetAuthor        |
| POST    | /authors/           | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).CreateAuthor     |
| PUT     | /authors/:id        | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).UpdateAuthor     |
| DELETE  | /authors/:id        | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).DeleteAuthor     |
| GET     | /authors/:id/books  | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).ListAuthorsBooks |
| GET     | /books/             | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).GetBooks         |
| GET     | /books/:id          | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).GetBook          |
| POST    | /books/             | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).CreateBook       |
| PUT     | /books/:id          | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).UpdateBook       |
| DELETE  | /books/:id          | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).DeleteBook       |
| POST    | /books/:id/buy      | github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller.(*Controller).BuyBook          |

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