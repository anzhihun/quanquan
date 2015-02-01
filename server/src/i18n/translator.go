package i18n

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"utils"
)

type Translator struct {
	filePath       string
	languageId     string
	usedLanguageId string
	translateFunc  i18n.TranslateFunc
}

var i18nFilesPath = utils.MakePath([]string{utils.GetRootDir(), "res", "nls"})

func newTranslator(languageId string) *Translator {
	filePath := utils.MakePath([]string{i18nFilesPath, languageId + ".json"})
	usedLanguageId := languageId
	if !utils.IsFileExist(filePath) {
		// set en-US as default language
		filePath = utils.MakePath([]string{i18nFilesPath, "en-US.json"})
		usedLanguageId = "en-US"
	}

	translator := &Translator{
		filePath:       filePath,
		languageId:     languageId,
		usedLanguageId: usedLanguageId,
	}
	translator.load()
	return translator
}

func (this *Translator) load() {
	i18n.MustLoadTranslationFile(this.filePath)
	this.translateFunc, _ = i18n.Tfunc(this.usedLanguageId)
}

func (this *Translator) T(key string, params map[string]interface{}) string {
	return this.translateFunc(key, params)
}
