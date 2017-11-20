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

var users map[int]string

func init() {
	users = map[int]string{
		1: "Maria",
		2: "João",
		3: "Joaquim",
		4: "Ana Clara",
	}
}

func name(userID int) string {
	name, ok := users[userID]
	if ok == false {
		return "Não identificado"
	}
	return name
}
