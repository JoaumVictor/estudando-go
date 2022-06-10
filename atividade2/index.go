package main

import "fmt"

type Database struct {
 	Username string
	Name string
	Age int
	ID int
	Email string
}

func findStructByName(name string, db []Database)(findUserByName Database){
	for _, user := range db{
		if(user.Name == name) {
			findUserByName = user
		}
	}
	return
}

func printSalve(db Database)(result string){
	result = fmt.Sprintf("Opa salve %s turu baum?", db.Name)
	return
}

func main() {
	db := []Database{}

	db = append(db, Database{
		Username: "Bescoito",
		Name: "Victor",
		Age: 19,
		ID: 1,
		Email: "bescoito08@gmail.com",
	});

	db = append(db, Database{
		Username: "Pasip",
		Name: "João",
		Age: 22,
		ID: 2,
		Email: "Pasipinho@gmail.com",
	});

	db = append(db, Database{
		Username: "Baianor",
		Name: "Adson",
		Age: 24,
		ID: 3,
		Email: "baianor@gmail.com",
	});

	db = append(db, Database{
		Username: "Miranha",
		Name: "Cleb",
		Age: 19,
		ID: 4,
		Email: "miranharomantico@gmail.com",
	});

	db = append(db, Database{
		Username: "Bebeto",
		Name: "Alberto",
		Age: 38,
		ID: 5,
		Email: "bebetogameplay@gmail.com",
	});
	
	clebStruct := findStructByName("Alberto", db)

	fmt.Println(clebStruct)
	fmt.Println(printSalve(clebStruct))
}

// a meta agr é descobrir como tirar um objeto do array