package main

import (
	"fmt"
	"main/hoofs"
	"main/models"
	"main/operations"
	"sort"
)

func main() {
	nomes := []string{"victor", "alberto", "pasip", "igor", "vitinho"}
	numeros := []int{4, 2, 3, 1, 8}
	sort.Ints(numeros)
	sort.Strings(nomes)

	molde := make([]models.Tipos, 0)
	for i := range nomes{
		tudao := models.Tipos{Palavras: nomes[i], Numeros: numeros[i]}
		molde = append(molde, tudao)
	}

	fmt.Println(molde)
	fmt.Println(hoofs.Find("victor", nomes))
	fmt.Println(operations.Soma(numeros))
}
