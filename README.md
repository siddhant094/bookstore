# Bookstore

This Project implements CRUD functionality. It is a Book Management system where users can view all books, view books by ID, create books, edit books, and delete books.

Database Used: MySQL

**The Postman Collection for Testing has been included.**

## Prerequisites

1. Go Installed
2. Docker Installed


## Project Structure
<a><img src="https://i.ibb.co/19MgdX0/uml.png" alt="uml" border="0"></a><br /><a target='_blank' href='https://nonprofitlight.com/mi/manistee/filer-credit-union'></a><br />

In Postman, created 5 Endpoints are: 
    | Request Name | Method | URL |
    | --- | --- | --- |
    | GET ALL | GET | localhost:9010/book/ |
    | GET BY ID | GET | localhost:9010/book/{id} |
    | CREATE | POST | localhost:9010/book/ |
    | UPDATE | PUT | localhost:9010/book/{id} |
    | DELETE | DELETE | localhost:9010/book/{id} |

    

## Build and run

1. First Run this command:
```
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 2001:3306 -d mysql:8.0.30
```

2. Now, run this command:
```
docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE books (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL, publication VARCHAR(255) NOT NULL);"
```

3. Now, run this:
```
go run main.go
```

4. Now Download Postman API Collection and Send requests to [http://localhost:9000/](http://localhost:9000/) 

## Testing

For Testing, Run:
```
go test
```
