package i18n

import (
	"testing"
)

func TestGetTranslator(t *testing.T) {
	i18nFilesPath = "testdata"
	translators = map[string]*Translator{}
	enTran := GetTranslator("en-US")
	if enTran == nil || len(translators) != 1 {
		t.Fatal("failed to add translator en-US instance when getting translator!")
	}

	// get again
	enTran = GetTranslator("en-US")
	if enTran == nil || len(translators) != 1 {
		t.Fatal("failed to add translator en-US instance when getting translator!")
	}

	cnTran := GetTranslator("zh-CN")
	if cnTran == nil || len(translators) != 2 {
		t.Fatal("failed to add translator zh-CN instance when getting translator!")
	}
}
