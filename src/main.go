// start a http server to comunicate with UI

package main

import (
	"fmt"
	"github.com/gocraft/web"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

type Context struct {
}

func (this *Context) getIndex(rw web.ResponseWriter, req *web.Request) {
	// return index.html as index page
	rootDir, _ := os.Getwd()
	indexFileName := rootDir + string(filepath.Separator) + "index.html"
	indexContent, _ := ioutil.ReadFile(indexFileName)
	rw.Write(indexContent)
}

func main() {
	startHttpServer()
	startCommunicateServer()
}

func startHttpServer() {
	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware). // Use some included middleware
					Middleware(web.ShowErrorsMiddleware)

	rootDir, _ := os.Getwd()
	router.Middleware(web.StaticMiddleware(rootDir))

	router.Get("/", (*Context).getIndex)

	http.ListenAndServe("localhost:53240", router) // Start the server!
}

// communicate with other quanquan clients throgh udp protocol
func startCommunicateServer() {

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 53241,
	})

	if err != nil {
		fmt.Println("failed to listen on port 53421!", err)
		return
	}
	defer socket.Close()

	for {
		// read data
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("Failed to read data from others!", err)
			continue
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n\n", data)

		// send data
		senddata := []byte("hello client!")
		_, err = socket.WriteToUDP(senddata, remoteAddr)
		if err != nil {
			return
			fmt.Println("Failed to send data to others!", err)
		}
	}
}
