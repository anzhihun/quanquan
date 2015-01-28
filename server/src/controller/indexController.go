package controller

import (
	"fmt"
	"github.com/gocraft/web"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (this *RootContext) GetIndex(rw web.ResponseWriter, req *web.Request) {
	// return index.html as index page
	rootDir, _ := os.Getwd()
	pathSeparator := string(filepath.Separator)
	indexFileName := rootDir + pathSeparator + "www" + pathSeparator + "index.html"
	fmt.Println("index file: ", indexFileName)
	indexContent, _ := ioutil.ReadFile(indexFileName)
	rw.Write(indexContent)
}
