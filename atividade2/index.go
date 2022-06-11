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

func findStructByID(ID int, db []Database)(findUserByID Database){
	for _, user := range db{
		if(user.ID == ID) {
			findUserByID = user
		}
	}
	return
}

func printSalve(db Database)(result string){
	result = fmt.Sprintf("Opa salve %s turu baum?", db.Name)
	return
}

func removeUser(name string, db []Database)(newDB []Database) {
	for i, v := range db {
    if v.Name == name {
			newDB = append(db[:i], db[i+1:]...)
    }
	}
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
		Email: "pasipinho@gmail.com",
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
	
	findUser1 := findStructByName("João", db)

	findUser2 := findStructByID(5, db)

	newListUser := removeUser("Alberto", db)

	fmt.Println(findUser1)

	fmt.Println(findUser2)

	fmt.Println(newListUser)
}
