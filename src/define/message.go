package define

const (
	MSG_TYPE_HI = "hi"
)

type Message struct {
	MsgType string
	From    string
	To      string
	Content string
}
