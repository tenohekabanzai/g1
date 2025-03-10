package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Salary int64              `bson:"salary" json:"salary"`
	Age    int64              `bson:"age" json:"age"`
}

var client *mongo.Client
var empCollection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
	mongoURI := os.Getenv("MONGO_URI")
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()
	client, err := mongo.Connect(ctxWithTimeout, options.Client().ApplyURI(mongoURI))

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Fatal("MongoDB connection timed out !!")
		} else {
			log.Fatal("MongoDB connection failed !!")
		}
	}

	empCollection = client.Database("Employees").Collection("Employees")
	fmt.Println("Connected to DB !!")
}

func CreateEmp(c *fiber.Ctx) error {
	var emp Employee
	err := c.BodyParser(&emp)
	if err != nil {
		fmt.Println("Error in Parsing Body")
		return nil
	}
	emp.ID = primitive.NewObjectID()
	_, err = empCollection.InsertOne(context.TODO(), emp)
	if err != nil {
		c.Send([]byte("failed to add employee"))
		return nil
	}
	c.Send([]byte("Added Employee successfully"))
	return nil
}

func GetEmp(c *fiber.Ctx) error {
	var emp []Employee
	cursor, err := empCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.Send([]byte("failed to get employees"))
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var e Employee
		cursor.Decode(&e)
		emp = append(emp, e)
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{"employees": emp})
	return nil
}

func GetEmpById(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send([]byte("Incorrect ID format"))
		return nil
	}
	var emp Employee
	err = empCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&emp)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send([]byte("Employee with given id does not exist"))
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"employee": emp})
	return nil
}

func UpdateEmp(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send([]byte("Incorrect ID format"))
		return nil
	}

	var emp Employee
	err = c.BodyParser(&emp)
	if err != nil {
		fmt.Println("Error in Parsing Body")
		return nil
	}
	emp.ID = id
	_, err = empCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": emp})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte("Failed to Update User"))
		return nil
	}
	c.Status(fiber.StatusOK).Send([]byte("Updated User Successfully"))
	return nil
}
func DeleteEmp(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send([]byte("Incorrect ID format"))
		return nil
	}
	_, err = empCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte("Failed to Update User"))
		return nil
	}
	c.Status(fiber.StatusOK).Send([]byte("Deleted User Successfully"))
	return nil
}

func main() {
	app := fiber.New()
	app.Get("/employee", GetEmp)
	app.Get("/employee/:id", GetEmpById)
	app.Post("/employee", CreateEmp)
	app.Put("/employee/:id", UpdateEmp)
	app.Delete("/employee/:id", DeleteEmp)
	app.Listen(":5001")
	fmt.Println("Server running at port 5001")
}
