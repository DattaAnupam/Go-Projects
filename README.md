## GO-WITH-MYSQL

## About
This shows the basic working process with MySql and Go.
<br>MySql Db has been used.
<br>Db name: **mytestdb**
<br>table name: **table1**
<br>**Note**
<br>Please make required changes for `root password` of MySql inside `main.go`.

## Requirements
<ol>
    <li>MySql should be installed. You can download it from <a href="https://dev.mysql.com/downloads/installer/">here</a> and install it.</li>
    <li>Keep Root Password handy. You will get it during installation.</li>
    <li>Set `PATH` variable. You can learn it from <a href="https://www.tutorialspoint.com/adding-mysql-to-windows-path">here</a></li>
</ol>

#### To Build GO-WITH-MYSQL project
type 
> `go get`
<br>It will bring all the go dependencies used in the project.
<br>Note: _Althout for this project `/vendor` folder is provided, hence you can igonre the above command_
> <br>`go init mod <Name_of_module>`
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
<br>Note
<br>name_of_the_.exe = name_of_the_module
<br>please make changes accordingly.
