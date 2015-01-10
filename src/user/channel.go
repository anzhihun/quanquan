package user

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
}

func (this *Channel) RemoveUser(u *User) {
	for index := 0; index < len(this.Users); index++ {
		if this.Users[index].Name == u.Name {
			this.Users = append(this.Users[:index], this.Users[index+1:]...)
			return
		}
	}
}

var channels = make([]*Channel, 0)

func AddChannel(channel *Channel) {
	channels = append(channels, channel)
}

func RemoveChannelByName(chanName string) {
	for i := 0; i < len(channels); i++ {
		if channels[i].Name == chanName {
			channels = append(channels[:i], channels[i+1:]...)
			break
		}
	}
}

func FindChannelByName(chanName string) *Channel {
	for i := 0; i < len(channels); i++ {
		if channels[i].Name == chanName {
			return channels[i]
		}
	}
	return nil
}
