package controller

import (
	"fmt"
	"github.com/gocraft/web"
	"net/http"
	"strings"
	"utils"
)

type resController struct {
	*RootContext
}

func (this *resController) getLanguage(rw web.ResponseWriter, req *web.Request) {
	langId := getLanguageId(req.Header)
	// get language file
	filePath := utils.MakePath([]string{utils.GetRootDir(), "www", "nls", langId, "strings.js"})
	var result []byte
	var err error
	fmt.Println("file path: ", filePath)
	if !utils.IsFileExist(filePath) {
		filePath = utils.MakePath([]string{utils.GetRootDir(), "www", "nls", "root", "strings.js"})
	}

	if result, err = utils.ReadFile(filePath); err != nil {
		http.Error(rw, "failed to read language file "+langId, 500)
	} else {
		rw.Write(result)
	}

}

func getLanguageId(header http.Header) string {
	return strings.TrimSpace(strings.Split(header["Accept-Language"][0], ",")[0])
}
