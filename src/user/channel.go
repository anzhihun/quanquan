package user

type Channel struct {
	name  string
	users []*User
}

func (this *Channel) AddUser(newUser *User) {
	if this.users == nil {
		this.users = make([]*User, 0)
	}
	this.users = append(this.users, newUser)
}

func (this *Channel) RemoveUser(u *User) {
	for index := 0; index < len(this.users); index++ {
		if this.users[index].Name == u.Name {
			this.users = append(this.users[:index], this.users[index+1:]...)
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
		if channels[i].name == chanName {
			channels = append(channels[:i], channels[i+1:]...)
			break
		}
	}
}

func FindChannelByName(chanName string) *Channel {
	for i := 0; i < len(channels); i++ {
		if channels[i].name == chanName {
			return channels[i]
		}
	}
	return nil
}
