package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID    string `json:"id,omitempty"`
	name  string `json:"name,omitempty"`
	email string `json:"email,omitempty"`
}

var users []User

func saveUser(c echo.Context) error {
	u := new(User)
	fmt.Println(u)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, c)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func main() {
	e := echo.New()
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":3333"))
}
