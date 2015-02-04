package entity

type User struct {
	Name     string
	Password string
	IP       string
	HeadImg  string
	Online   bool
}

type Channel struct {
	Name    string
	Users   []*User
	Creator string
}

func (this *Channel) AddUser(newUser *User) {
	if this.Users == nil {
		this.Users = make([]*User, 0)
	}
	this.Users = append(this.Users, newUser)
	// event.Trigger("channel:user:change", nil, nil)
}

func (this *Channel) RemoveUser(u *User) {
	for index := 0; index < len(this.Users); index++ {
		if this.Users[index].Name == u.Name {
			this.Users = append(this.Users[:index], this.Users[index+1:]...)
			// event.Trigger("channel:user:change", nil, nil)
			return
		}
	}
}

func (this *Channel) ContainsUser(userName string) bool {
	for index := 0; index < len(this.Users); index++ {
		if this.Users[index].Name == userName {
			return true
		}
	}

	return false
}
