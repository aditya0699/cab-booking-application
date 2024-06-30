package models

type User struct {
	ID   int
	Name string
}

type Driver struct {
	ID            string
	Name          string
	TotalEarnings float64
	Status        string
}
