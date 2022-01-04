package actions

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
)

type Payload struct {
	Data User `json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func AboutHandler(c buffalo.Context) error {
	c.Set("title", "About Page")
	c.Set("body", "A Go web development eco-system, designed to make your project easier.")
	return c.Render(http.StatusOK, r.HTML("home/about.plush.html"))
}

func UserHandler(c buffalo.Context) error {
	user_id := c.Param("user_id")
	user, err := getUserInfo(user_id)
	if err != nil {
		c.Flash().Add("warning", "Bad connection")
		c.Redirect(301, "/")
	}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("home/user.plush.html"))
}

func getUserInfo(user_id string) (User, error) {
	id, _ := strconv.Atoi(user_id)
	user := User{
		ID: id, Email: "roshitashakya" + user_id + "@gmail.com",
		FirstName: "Roshita", LastName: "Shakya",
	}
	return user, nil
}
