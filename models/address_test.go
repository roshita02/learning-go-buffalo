package models

import "github.com/gofrs/uuid"

func (ms *ModelSuite) Test_Address() {
	a := &Address{
		Street: "1 Lexington Street",
		City:   "Lexington",
		State:  "CA",
		Zip:    "90011",
	}

	u := &User{
		FirstName: "John",
		LastName:  "Doe",
		Age: 25,
		UserAddress: *a,
	}

	db := ms.DB
	verrs, err := db.Eager().ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotEqual(uuid.Nil, u.UserAddress.ID, "Address ID is generated when saved in DB")
	ms.False(verrs.HasAny(), "Address and user creation have no validation errors.")
}
