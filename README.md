# CRM With Go Fiber and MySql

## About
This Program shows the working of CRUD operation, using **Go Fiber**, **GORM** and **MySql**.
<br>
**Go Fiber** is used to handle the request and response. For more details about **Go Fiber** click <a href="https://gofiber.io/" target="_blank">here</a>.
<br>
**GORM** is used to map object into database model. It is also handles the operations(Create, Query, Update, Delete) on database. For more details about **GORM** click <a href="https://gorm.io/" target="_blank">here</a>.
<br>
**MySql** is used as local database.

## Requirements
<ol>
    <li>Mysql should be installed in local system.</li>
    <li>Username and password of the database want to use</li>
</ol>

## Packages Used
<ol>
    <li>fmt</li>
    <li>log</li>
    <li>github.com/gofiber/fiber/v2</li>
    <li>gorm.io/driver/mysql</li>
    <li>gorm.io/gorm</li>
</ol>

## Build and Run
After cloning the repo
<ol>
    <li>type <code>go mod init github.com/anupam/crm-with-go-fiber-mysql</code></li>
    <li>type <code>go mod tidy</code></li>
    <li>type <code>go build</code></li>
    <li>type <code>.\crm-with-go-fiber-mysql.exe</code></li>
</ol>

**Note**
> Can change the name of mod file in _step1_ instead of _github.com/anupam/crm-with-go-fiber-mysql_ you can put any name of the package.

> if you change the name of mod file please do make changes to the respective files.

> After _step 2_ you can run <code>go get</code> if dependencies are not loaded successfully.

> Make required changes to the database username, password, tcp port number and database name.

> By default the server runs on port :8080
<br>MySql server is running on port: 3306

> Commands may differ depending on OS. Above are for Windows OS.

## Routes
<table>
    <tr>
        <th>Method</th>
        <th>Endpoint</th>
        <th>Detail</th>
        <th>Url for localhost:8080</th>
    </tr>
    <tr>
        <td>GET</td>
        <td>/api/lead</td>
        <td>Get All Leads</td>
        <td>localhost:8080/api/lead</td>
    </tr>
    <tr>
        <td>GET</td>
        <td>/api/lead/:id</td>
        <td>Get a Lead by its ID</td>
        <td>localhost:8080/api/lead/1</td>
    </tr>
    <tr>
        <td>POST</td>
        <td>/api/lead</td>
        <td>Create New Lead</td>
        <td>localhost:8080/api/lead</td>
    </tr>
    <tr>
        <td>PUT</td>
        <td>/api/lead/:id</td>
        <td>Update a Lead's Detail</td>
        <td>localhost:8080/api/lead/1</td>
    </tr>
    <tr>
        <td>DELETE</td>
        <td>/api/lead/:id</td>
        <td>Delete a Lead</td>
        <td>localhost:8080/api/lead/1</td>
    </tr>
</table>

<a href="https://github.com/DattaAnupam/Go-Projects">**Back to master branch**</a>