package cleanit

func returnUsername(userId int) string {
	if userId == 1 {
		return "Maria"
	} else if userId == 2 {
		return "João"
	} else if userId == 3 {
		return "Joaquim"
	} else if userId == 4 {
		return "Ana Clara"
	}
	return "Não identificado"
}
