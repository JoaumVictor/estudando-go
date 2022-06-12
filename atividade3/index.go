package main

import (
	"fmt"
	"sort"
)

type tipos struct{
	palavras string
	numeros int
	valid bool
}

func soma(dales ...tipos)(totaln int, totals string){
	for _, v := range dales{
		totaln += v.numeros
		totals += v.palavras
	}
	return
}


func find(buscar int, array []int)(result interface{}) {
	for _, v := range array{
		if v == buscar {
			result = v
		}
	}
	return
}

func main() {
	molde := make([]tipos, 4)

	nomes := []string{"victor", "alberto", "pasip", "igor"}
	numeros := []int{4, 2, 3, 1}
	boolss := []bool{true, false, false, true}

	sort.Ints(numeros)
	sort.Strings(nomes)

	for i := range nomes{
		temp := tipos{palavras: nomes[i], numeros: numeros[i], valid: boolss[i]}
		molde = append(molde, temp)
	}

	fmt.Println(find(9, numeros))
}
