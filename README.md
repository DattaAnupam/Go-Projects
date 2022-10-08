# Go-Projects
This Repo contains small projects on go language to touch various aspects of Go Language

#### To Build Go-SERVER project
> type "go build" inside terminal.

#### Error: `go: go.mod file not found in current directory or any parent directory;...`
Solution: inside terminal type `go env -w GO111MODULE=auto` (1 -> One, not L in small)

#### To Run 
> 1. Double click on .exe file
> 2. type `go run main.go`

By default the server runs on port :8080

To understand working of form handler function type `localhost:8080/form.html`
> it will first open the `form.html` file inside *static* folder, and then will submit it to *formHandler* function.
