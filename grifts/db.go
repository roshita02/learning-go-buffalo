package grifts

import (
	"bufio"
	"encoding/json"
	"learning_buffalo/models"
	"os"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		savedTags := map[string]models.Tag{}
		// Add DB seeding stuff here

		// Add tags
		tags := []string{"Technology", "Travel and Lifestyle", "Entertainment"}
		for _, v := range tags {
			t := models.Tag{Name: v}
			err := models.DB.Create(&t)
			if err != nil {
				panic(err)
			}
			savedTags[t.Name] = t
		}

		// Add a user
		user_1 := models.User{FirstName: "Roshita", LastName: "Shakya", Age: 35}
		err := models.DB.Create(&user_1)
		if err != nil {
			panic(err)
		}

		// Import blogs from file
		f, err := os.Open("blogs.json")
		if err != nil {
			panic(err)
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			var b models.Blog
			err = json.Unmarshal(s.Bytes(), &b)
			if err != nil {
				panic(err)
			}
			// Add author
			b.User = &user_1

			// Add tags
			t, ok := savedTags["Technology"]
			if ok {
				b.BlogTags = append(b.BlogTags, t)
			}
			err := models.DB.Create(&b)
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

})
