package models

import "fmt"

type User struct {
	ID      int64   `json:"id"`
	Address Address `json:"address"`
	Age     int     `json:"age"`
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %d, Address: %v, Age: %d.", u.ID, u.Address, u.Age)
}

type Address struct {
	CityID int64  `json:"city_id"`
	Street string `json:"street"`
}

func (a Address) String() string {
	return fmt.Sprintf("CityID: %d, Street: %s", a.CityID, a.Street)
}
