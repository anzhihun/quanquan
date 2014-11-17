// client means user interface client, such as web view or other traditional view client

package conn

import (
	"code.google.com/p/go.net/websocket"
)

type ClientConnector interface {
	SendMsg(msg string) error
	IsMe(id string) bool
}

type WebsocketClientConnector struct {
	conn *websocket.Conn
	id   string
}

func NewWebsocketClientConnector(conn *websocket.Conn, id string) *WebsocketClientConnector {
	return &WebsocketClientConnector{conn, id}
}
func (this *WebsocketClientConnector) SendMsg(msg string) error {
	return nil
}

func (this *WebsocketClientConnector) IsMe(id string) bool {
	return this.id == id
}

var clientConnectors = []*WebsocketClientConnector{}

func AddClientConnector(conn *WebsocketClientConnector) {
	// avoid add replicate conns
	old := FindClientConnector(conn.id)
	if old != nil {
		return
	}

	clientConnectors = append(clientConnectors, conn)
}

func FindClientConnector(id string) *WebsocketClientConnector {
	for i := 0; i < len(clientConnectors); i++ {
		if clientConnectors[i].IsMe(id) {
			return clientConnectors[i]
		}
	}
	return nil
}

func Broadcast2AllClient(content []byte) {
	for i := 0; i < len(clientConnectors); i++ {
		if err := websocket.Message.Send(clientConnectors[i].conn, string(content)); err != nil {
			// log.GetLogger().Warning(fmt.Sprintf("rtMsg Send failed\n remote address:%s\n clientId:%d\n message:%s\n error:%s\n", self.connection.RemoteAddr().String(), self.clientId, msg, err.Error()))
		}
	}

}
