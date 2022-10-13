# Go-Projects
This Repo contains small projects on go language to touch various aspects of Go Language

## GO-WITH-MYSQL

## About
This shows the basic working process with MySql and Go.
<br>MySql Db has been used.
<br>Db name: **mytestdb**
<br>table name: **table1**
<br>**Note**
<br>Please make required changes for `root password` of MySql insde `main`.

## Requirements
<ol>
    <li>MySql should be installed. You can download it from <a href="https://dev.mysql.com/downloads/installer/">here</a> and install it.</li>
</ol>

#### To Build GO-WITH-MYSQL project
type 
> `go get`
> <br>`go build`

#### Error: `go: go.mod file not found in current directory or any parent directory;...`
Solution: inside terminal type `go env -w GO111MODULE=auto` (1 -> One, not L in small)

#### To Run 
> 1. Double click on .exe file
<br> type
> 2. `go run main.go`

#### To Build + Run
type
> `go build && ./go-with-sql.exe`
