package user

type userManager struct {
	users []*User
}

var UserManager userManager

func (this *userManager) AddUser(user *User) {
	this.users = append(this.users, user)
}

func (this *userManager) AllUser() []*User {
	return this.users
}
