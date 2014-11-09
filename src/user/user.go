package user

type User struct {
	Name    string
	IP      string
	HeadImg string
}

var Self User = User{"anzhihun", "127.0.0.1", "/images/anzhihun.png"}
