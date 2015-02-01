package controller

import (
	"fmt"
	"github.com/gocraft/web"
	"io/ioutil"
	"utils"
)

func (this *RootContext) GetIndex(rw web.ResponseWriter, req *web.Request) {
	// return index.html as index page
	indexFileName := utils.MakePath([]string{utils.GetRootDir(), "www", "index.html"})
	fmt.Println("index file: ", indexFileName)
	indexContent, _ := ioutil.ReadFile(indexFileName)
	rw.Write(indexContent)
}
