package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"

// const mongoURI = "mongodb://localhost:27017" + dbName
// Uri for my mongo atlas cluster
const mongoURI = "mongodb+srv://dummyUser:Yma6hMB7DmuvUjcN@cluster0.hdt7puu.mongodb.net/test"

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Connect to Mongo DB
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

// GET
// Get all Employees
var GetAllEmployee = func(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []Employee = make([]Employee, 0)

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

// GET
// Get Employee by ID
var GetEmployeeByID = func(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(Employee)

	query := bson.D{{Key: "_id", Value: id}}

	cursor := mg.Db.Collection("employees").FindOne(c.Context(), query)

	if err := cursor.Decode(employee); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employee)
}

// POST
// Create a new Employee
var CreateEmployee = func(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")
	employee := new(Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	employee.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), employee)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &Employee{}
	createRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}

// PUT
// Update an Existing Employee
var UpdateEmployee = func(c *fiber.Ctx) error {
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: employeeID}}

	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "name", Value: employee.Name},
			{Key: "age", Value: employee.Age},
			{Key: "salary", Value: employee.Salary},
		},
	}}

	err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), filter, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}

	employee.ID = id

	return c.Status(200).JSON(employee)
}

// DELETE
// Delete an Employee by ID
var DeleteEmployeeByID = func(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: id}}
	result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("record deleted")
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err.Error())
	}
	app := fiber.New()

	// Get all employees
	app.Get("/employee", GetAllEmployee)

	// Get employee by ID
	app.Get("/employee/:id", GetEmployeeByID)

	// Create new Employee
	app.Post("/employee", CreateEmployee)

	// Update an existing employee
	app.Put("/employee/:id", UpdateEmployee)

	// Delete employee by ID
	app.Delete("/employee/:id", DeleteEmployeeByID)

	// Run the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
