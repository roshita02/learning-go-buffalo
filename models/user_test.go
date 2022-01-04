package models

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
