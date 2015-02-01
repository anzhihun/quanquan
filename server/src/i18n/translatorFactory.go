package i18n

// cache translator
var translators map[string]*Translator

func GetTranslator(languageId string) *Translator {
	if _, ok := translators[languageId]; !ok {
		translators[languageId] = newTranslator(languageId)
	}

	result, _ := translators[languageId]
	return result
}
