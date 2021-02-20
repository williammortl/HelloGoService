package db

// Person is the type of record that the db stores
type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

// initial db
var personDB = map[int]Person{
	0: {Name: "William Mortl", Address: "9th St", Phone: "6305551212"},
	1: {Name: "Linda Mortl", Address: "Navarro St", Phone: "4805551212"},
	2: {Name: "Kara Linnersund", Address: "Denver West Ct", Phone: "3035551212"},
}

// GetPersonByID retrieves a person (by ID) from the db
func GetPersonByID(id int) *Person {
	_, exists := personDB[id]
	var ret *Person = nil
	if exists {
		person := personDB[id]
		ret = &person
	}
	return ret
}

// AddPerson adds a person to the db
func AddPerson(id int, person Person) {
	personDB[id] = person
}
