package main

type tipos struct{
	palavras string
	numeros int
}

func soma(dales ...tipos)(totaln int, totals string){
	for _, v := range dales{
		totaln += v.numeros
		totals += v.palavras
	}
	return
}

func main() {
	dales := make([]tipos, 4)
	nomes := []string{"victor", "alberto", "adson", "igor"}
	numeros := []int{5, 7, 3, 5}

	for i, _ := range nomes{
		temp := tipos{palavras: nomes[i], numeros: numeros[i]}
		dales = append(dales, temp)
	}
}