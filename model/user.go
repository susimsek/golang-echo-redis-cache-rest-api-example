package model

type User struct {
	ID string `json:"id" xml:"id"`
	*UserInput
}

type UserInput struct {
	FirstName string `json:"firstName" xml:"firstName"`
	LastName  string `json:"lastName" xml:"lastName"`
	Email     string `json:"email" xml:"email""`
}
