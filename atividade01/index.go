package main

import "fmt"

func main() {
	// isso é um slice ( é tipo um array )
	ex1 := []string{"nony", "baianor", "bescoito"}

	// isso é um append ( é tipo um .push ou .concat )
	ex2 := append(ex1, "adran", "diogão")

	// isso aqui é um map ( tipo um objeto )
	ex3 := map[string]string{"nome": "bescoito"}

	// isso é um console.log humilde
	fmt.Println(ex1, ex2, ex3)
}

// pra usar esse tipo de "console.log" o go importa aquele fmt na linha 3
// toda aplicação tem que ter um package main pra poder rodar
