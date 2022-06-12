package operations

func Soma(numbers []int)(totaln int){
	for _, v := range numbers{
		totaln += v
	}
	return
}