package actions

import (
	"fmt"
	"learning_buffalo/models"
)

func (as *ActionSuite) Test_Blogs_Show() {
	as.LoadFixture("example blogs")

	b := models.Blog{}
	err := as.DB.Last(&b)
	if err != nil {
		panic(err)
	}

	res := as.HTML(fmt.Sprintf("/blogs/%s", b.ID)).Get()
	body := res.Body.String()
	as.Contains(body, "First Blog post", "Blog title appears on Show page.")
	as.Contains(body, "Roshita Shakya", "Author name appears on blog Show page.")
}
