package models

import "github.com/gofrs/uuid"

func (ms *ModelSuite) Test_User() {
	u := &User{
		FirstName: "John",
		LastName:  "Doe",
	}
	ms.Equal("John Doe", u.FullName(), "FullName returns user name.")

	db := ms.DB
	verrs, err := db.ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotNil(u.ID, "User ID is generated when saved to DB.")
	ms.True(verrs.HasAny(), "User cannot be created without age field.")
}

func (ms *ModelSuite) Test_UserAddress() {
	u := &User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		UserAddress: Address{
			Street: "1 Lexington Street",
			City:   "Lexington",
			State:  "CA",
			Zip:    "90011",
		},
	}

	db := ms.DB
	_, err := db.Eager().ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotEqual(uuid.Nil, u.UserAddress.ID, "Address saved along with User..")

	u2 := User{}
	err = db.Find(&u2, u.ID)
	if err != nil {
		panic(err)
	}

	ms.Empty(u2.UserAddress, "User address is not loaded by default.")
	u2.GetAdress(db)
	ms.NotEmpty(u2.UserAddress, "GetAddress loads the associated user address")
}

func (ms *ModelSuite) Test_UserBlogs() {
	u := &User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Blogs: Blogs{Blog{
			Title: "First blog",
			Body:  "<p>Hello, This is my first blog. </p>",
		}},
	}

	db := ms.DB
	_, err := db.Eager().ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	u2 := User{}
	err = db.Find(&u2, u.ID)
	if err != nil {
		panic(err)
	}

	ms.Empty(u2.Blogs, "Blogs are not loaded with user by default.")
	ms.Equal("John Doe", u2.FullName(), "Associated user is loaded.")

	err = u2.GetBlogs(db)
	if err != nil {
		panic(err)
	}
	ms.Len(u2.Blogs, 1, "The function loads all the user's blogs.")
}
