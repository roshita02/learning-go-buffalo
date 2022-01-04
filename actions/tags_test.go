package actions

import (
	"fmt"
	"learning_buffalo/models"
)

func (as *ActionSuite) Test_Tags_Show() {
	as.LoadFixture("example tags")

	t := models.Tag{}
	err := as.DB.Last(&t)
	if err != nil {
		panic(err)
	}

	res := as.HTML(fmt.Sprintf("/tags/%s", t.ID)).Get()
	body := res.Body.String()
	as.Contains(body, "Travel and Lifestyle", "Tag name appears on Show page.")
}
