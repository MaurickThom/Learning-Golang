package main

import "fmt"

type Contact struct {
	ID    int    `json: "id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	Phone string `json:"phone"`
}

type Contacts []Contact

func (c *Contact) toString() string {
	return fmt.Sprintf("%d,%s,%s,%s", c.ID, c.Name, c.City, c.Phone)
}
