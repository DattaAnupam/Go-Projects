# Go-Projects
This Repo contains small projects on go language to touch various aspects of Go Language

## Book-Management-CRUD

## About
This shows basic CRUD i.e. Create, Read (Get), Update and Delete operation of an Rest Api.
<br>MySql Db has been used.
<br>Db name: **mytestdb**
<br>table name: **books**
<br>**Note**
<br>Please make required changes for `root password` of MySql inside `main.go`.
<br>`github.com/gorilla/mux` is being used as request router and dispatcher for incoming request.
<br>*Postman* has been used and *Postman Collection* is also attahced with the project.
<br>**Note**
<br>Please make required changes for port number and book id if needed.

## Requirements
<ol>
    <li>MySql should be installed. You can download it from <a href="https://dev.mysql.com/downloads/installer/">here</a> and install it.</li>
    <li>Keep Root Password handy. You will get it during installation.</li>
    <li>Set `PATH` variable. You can learn it from <a href="https://www.tutorialspoint.com/adding-mysql-to-windows-path">here</a></li>
</ol>
#### To Build Book-Management-Crud project
> go to main directory, inside pkg directory.
> type "go build".

#### Error: `go: go.mod file not found in current directory or any parent directory;...`
Solution: inside terminal type `go env -w GO111MODULE=auto` (1 -> One, not L in small)

#### To Run 
> 1. Double click on .exe file
> 2. type `go run main.go`

By default the server runs on port :8080
<br>Sql server is running on port: 3306

## Routes
<table>
    <tr>
        <td>GET</td>
        <td>/api/book</td>
        <td>Get all books</td>
    </tr>
    <tr>
        <td>GET</td>
        <td>/api/book/{bookId}</td>
        <td>Get a book by its ID</td>
    </tr>
    <tr>
        <td>POST</td>
        <td>/api/book</td>
        <td>Create new book</td>
    </tr>
    <tr>
        <td>PUT</td>
        <td>/api/book/{bookId}</td>
        <td>Update a book's detail</td>
    </tr>
    <tr>
        <td>DELETE</td>
        <td>/api/book/{bookId}</td>
        <td>Delete a book</td>
    </tr>
</table>