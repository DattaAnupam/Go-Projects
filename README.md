# Go-Projects
This Repo contains small projects on go language to touch various aspects of Go Language

## Book-Management-CRUD

## About
This shows basic CRUD i.e. Create, Read (Get), Update and Delete operation of an Rest Api.
<br>MySql Db has been used.
<br>`github.com/gorilla/mux` is being used as request router and dispatcher for incoming request.
<br>*Postman* has been used and *Postman Collection* is also attahced with the project.
<br>**Note**
<br>Please make required changes for port number and book id if needed.

#### To Build Book-Management-Crud project
> go to main directory, inside pkg directory.
> type "go build".

#### Error: `go: go.mod file not found in current directory or any parent directory;...`
Solution: inside terminal type `go env -w GO111MODULE=auto` (1 -> One, not L in small)

#### To Run 
> 1. Double click on .exe file
> 2. type `go run main.go`

By default the server runs on port :8080

### **Note**
The project is buildable but not executable.
<br>Not able to connect to MySql.