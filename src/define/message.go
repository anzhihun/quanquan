package define

const (
	MSG_TYPE_ONLINE  = "online"
	MSG_TYPE_OFFLINE = "offline"
	MSG_TYPE_JOIN    = "join"
	MSG_TYPE_TALK    = "talk"
)

type Message struct {
	MsgType  string
	From     string
	HeadImg  string
	To       string
	IsPublic bool
	Content  string
}
