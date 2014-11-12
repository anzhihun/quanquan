package user

type userManager struct {
	users []*User
}

var UserManager = userManager{[]*User{&Self}}

func (this *userManager) AddUser(user *User) {
	if this.users == nil {
		this.users = make([]*User, 0)
	}
	oldUser := this.FindUser(user.Name)
	if oldUser == nil {
		this.users = append(this.users, user)
	} else {
		oldUser.HeadImg = user.HeadImg
		oldUser.IP = user.IP
		oldUser.Name = user.Name
		oldUser.Online = user.Online
	}
}

func (this *userManager) RemoveUser(user *User) {
	if this.users == nil || len(this.users) == 0 {
		return
	}

	for index := 0; index < len(this.users); index++ {
		if this.users[index].Name == user.Name {
			this.users = append(this.users[:index], this.users[index+1:]...)
			return
		}
	}
	return
}

func (this *userManager) FindUser(userName string) *User {
	if this.users == nil || len(this.users) == 0 {
		return nil
	}

	for index := 0; index < len(this.users); index++ {
		if this.users[index].Name == userName {
			return this.users[index]
		}
	}
	return nil
}

func (this *userManager) AllUser() []*User {
	return this.users
}

func (this *userManager) Clear() {
	this.users = nil
}
