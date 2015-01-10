package entity

const (
	WS_MSGTYPE_TALK        = "talk"
	WS_MSGTYPE_ONLINE      = "online"
	WS_MSGTYPE_NEW_USER    = "newUser"
	WS_MSGTYPE_NEW_CHANNEL = "newChannel"

	WS_MSGCONTENT_TYPE_TEXT = "text"
)

type WSAckUser struct {
	Name     string `json:"name"`
	IconUrl  string `json:"iconUrl"`
	Online   bool   `json:"online"`
	ServerId uint64 `json:"serverId"`
}

type WSAckChannel struct {
	Name     string      `json:"name"`
	Creator  string      `json:"creator"`
	Users    []WSAckUser `json:"users"`
	ServerId uint64      `json:"serverId"`
}

type WSAckTalk struct {
	MsgType     string `json:"msgType"`
	ContentType string `json:"contentType"`
	Sender      string `json:"sender"`
	Is2P        bool   `json:"is2P"`
	Receiver    string `json:"receiver"`
	Content     string `json:"content"`
}

type WSAckNewUser struct {
	MsgType string    `json:"msgType"`
	User    WSAckUser `json:"user"`
}

type WSAckNewChannel struct {
	MsgType string       `json:"msgType"`
	Channel WSAckChannel `json:"channel"`
}
