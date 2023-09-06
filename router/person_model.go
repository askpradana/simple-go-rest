package router

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var persons []Person

type UpdatePerson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var updatePersons []UpdatePerson
