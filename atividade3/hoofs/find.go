package hoofs

func Find(buscar string, array []string)(result interface{}) {
	for _, v := range array{
		if v == buscar {
			result = v
		}
	}
	return
}