# Go-Projects
This Repo contains small projects on go language to touch various aspects of Go Language

## About
This shows basic CRUD i.e. Create, Read (Get), Update and Delete operation of an Rest Api.
<br>Local structure has been used as storage.
<br>No Db has been used.
<br>`github.com/gorilla/mux` is being used as request router and dispatcher for incoming request.
<br>*Postman* has been used and *Postman Collection* is also attahced with the project.
<br>**Note**
<br>Please make required changes for port number and Bood id if needed.

#### To Build Book-Store-Crud project
> type "go build" inside terminal.

#### Error: `go: go.mod file not found in current directory or any parent directory;...`
Solution: inside terminal type `go env -w GO111MODULE=auto` (1 -> One, not L in small)

#### To Run 
> 1. Double click on .exe file
> 2. type `go run main.go`

By default the server runs on port :8080
