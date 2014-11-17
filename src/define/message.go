package define

const (
	MSG_TYPE_ONLINE     = "online"
	MSG_TYPE_ACK_ONLINE = "ackonline"
	MSG_TYPE_OFFLINE    = "offline"
	MSG_TYPE_JOIN       = "join"
	MSG_TYPE_TALK       = "talk"

	MSG_TYPE_USER_ADD = "addUser"
)

type Message struct {
	MsgType  string
	From     string
	HeadImg  string
	To       string
	IsPublic bool
	Content  string
}

type AddUserMessage struct {
	UserName string
	Password string
}

type ToClientMessage struct {
	MsgType string
	Content string
}
