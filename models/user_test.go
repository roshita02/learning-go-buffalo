package models

func (ms *ModelSuite) Test_User() {
	u := &User{
		FirstName: "John",
		LastName: "Doe",
	}
	ms.Equal("John Doe", u.FullName(), "FullName returns user name.")
}
