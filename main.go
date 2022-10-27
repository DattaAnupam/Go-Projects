package main

import (
	"context"
	"log"
	"time"

	"github.com/anupam/hrms-with-fiber-mongodb/pkg/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mg model.MongoInstance

// Name of the DB
const dbName = "fiber-hrms"

// const mongoURI = "mongodb://localhost:27017" + dbName
// Uri for (my) mongo atlas cluster
const mongoURI = "mongodb+srv://dummyUser:Yma6hMB7DmuvUjcN@cluster0.hdt7puu.mongodb.net/test"

// Configure: Connect to Mongo DB
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

	mg = model.MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

// GET
// Handler: Get all Employees
var GetAllEmployee = func(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []model.Employee = make([]model.Employee, 0)

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

// GET
// Handler: Get Employee by ID
var GetEmployeeByID = func(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(model.Employee)

	query := bson.D{{Key: "_id", Value: id}}

	cursor := mg.Db.Collection("employees").FindOne(c.Context(), query)

	if err := cursor.Decode(employee); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employee)
}

// POST
// Handler: Create a new Employee
var CreateEmployee = func(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")
	employee := new(model.Employee)
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

	createdEmployee := &model.Employee{}
	createRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}

// PUT
// Handler: Update an Existing Employee
var UpdateEmployee = func(c *fiber.Ctx) error {
	id := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(model.Employee)

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
// Handler: Delete an Employee by ID
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
	// Configure: Connection to DB
	if err := Connect(); err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	// Route: Get all employees
	app.Get("/api/employee", GetAllEmployee)

	// Route: Get employee by ID
	app.Get("/api/employee/:id", GetEmployeeByID)

	// Route: Create new Employee
	app.Post("/api/employee", CreateEmployee)

	// Route: Update an existing employee
	app.Put("/api/employee/:id", UpdateEmployee)

	// Route: Delete employee by ID
	app.Delete("/api/employee/:id", DeleteEmployeeByID)

	// Run the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
