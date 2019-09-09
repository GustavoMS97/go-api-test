package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientMongoDb mongo.Client
var userCollection *mongo.Collection

// User : Struct de usuário.
type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func saveUser(c echo.Context) error {
	var user User
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	user.Name = fmt.Sprintf("%v", m["name"])
	user.Email = fmt.Sprintf("%v", m["email"])

	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(200, insertResult.InsertedID)
}

func getUser(c echo.Context) error {
	var result User
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, result)
}

func show(c echo.Context) error {
	var results []*User
	cur, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return c.JSON(http.StatusOK, results)
}

func connectToMongodb() {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		fmt.Println("url mongo vazia.")
		mongoURL = "mongodb://localhost:27017"
	}
	fmt.Println(mongoURL)
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURL)
	// Connect to MongoDB
	clientMongoDb, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = clientMongoDb.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	userCollection = clientMongoDb.Database("GO_TEST_API").Collection("users")
}

func disconnectFromMongodb() {
	err := clientMongoDb.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func desafio1(c echo.Context) error {
	var result []string
	var naipes = []string{"C", "E", "P", "O"}
	var valores = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	for i := 0; i < len(naipes); i++ {
		for j := 0; j < len(valores); j++ {
			result = append(result, naipes[i]+valores[j]+",")
		}
	}
	return c.String(http.StatusOK, "["+strings.Join(result, " ")+"]")
}

func desafio2(c echo.Context) error {
	var result []string
	var naipes = []string{"C", "E", "P", "O"}
	var valores = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	for i := 0; i < len(naipes); i++ {
		for j := 0; j < len(valores); j++ {
			result = append(result, naipes[i]+valores[j])
		}
	}
	var carta1 = rand.Intn(len(result)-0) + 0
	var carta2 = rand.Intn(len(result)-0) + 0
	response := fmt.Sprintf("Carta 1: %s\nCarta 2: %s", result[carta1], result[carta2])
	return c.String(http.StatusOK, response)
}

func main() {
	connectToMongodb()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.GET("/users", show)
	e.GET("/desafio1", desafio1)
	e.GET("/desafio2", desafio2)
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Porta vazia.")
		port = "3000"
	}
	fmt.Println(port)
	e.Logger.Fatal(e.Start(":" + port))
	//MONGO_URL=mongodb+srv://GUSTAVOMS97:Gu4525214@cluster0-gvhbx.mongodb.net go run server.go
}
