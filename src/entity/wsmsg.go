package entity

const (
	WS_MSGTYPE_TALK = "talk"
)

type WSAckUser struct {
	Name     string `json:"name"`
	IconUrl  string `json:"iconUrl"`
	Online   bool   `json:"online"`
	ServerId uint64 `json:"serverId"`
}
