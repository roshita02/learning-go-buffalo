package actions

import (
	"learning_buffalo/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// TagsShow default implementation.
func TagsShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	tag := models.Tag{}
	tag_id := c.Param("id")

	err := tx.Find(&tag, tag_id)
	if err != nil {
		c.Flash().Add("warning", "Error handling tag.")
		c.Redirect(301, "/")
	}

	c.Set("tag", tag)
	return c.Render(http.StatusOK, r.HTML("tags/show.html"))
}

