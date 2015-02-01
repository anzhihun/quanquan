package controller

import (
	"github.com/gocraft/web"
	"net/http"
	"strings"
	"utils"
)

type resController struct {
	*RootContext
}

// get language strings file
func (this *resController) getLanguage(rw web.ResponseWriter, req *web.Request) {
	langId := getLanguageId(req.Header)
	filePath := utils.MakePath([]string{utils.GetRootDir(), "www", "nls", langId, "strings.js"})
	if !utils.IsFileExist(filePath) {
		filePath = utils.MakePath([]string{utils.GetRootDir(), "www", "nls", "root", "strings.js"})
	}

	if result, err := utils.ReadFile(filePath); err != nil {
		http.Error(rw, "failed to read language strings file "+langId, 500)
	} else {
		rw.Write(result)
	}

}

func getLanguageId(header http.Header) string {
	return strings.TrimSpace(strings.Split(header["Accept-Language"][0], ",")[0])
}
