package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

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
	return c.JSON(200, user)
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
			result = append(result, naipes[i]+valores[j]+",")
		}
	}
	return c.String(http.StatusOK, "["+strings.Join(result, " ")+"]")
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
	e.GET("/desafio", desafio)
	e.Logger.Fatal(e.Start(":3333"))
}
