package i18n

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	i18nFilesPath = "testdata"
	enTran := newTranslator("en-US")
	actual := enTran.T("testKey", map[string]interface{}{
		"Error": "test error",
	})
	expected := "test translation test error"
	if actual != expected {
		t.Fatal("failed to tranlate to en-US, actual: ", actual, " expected: ", expected)
	}

	zhTran := newTranslator("zh-CN")
	actual = zhTran.T("testKey", map[string]interface{}{
		"Error": "测试中文错误",
	})
	expected = "test translation 测试中文错误"
	if actual != expected {
		t.Fatal("failed to tranlate to zh-CN, actual: ", actual, " expected: ", expected)
	}
}
