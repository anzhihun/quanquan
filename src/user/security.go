package user

func Validate(userName, password string) bool {

	existUser := FindUser(userName)
	if existUser == nil {
		return false
	}

	if password != existUser.Password {
		return false
	}
	return true
}
