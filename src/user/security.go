package user

func Validate(userName, password string) bool {
	if userName == "admin" && password == "admin" {
		return true
	}
	return false
}
