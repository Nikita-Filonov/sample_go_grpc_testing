package reports

import (
	"encoding/json"
	"fmt"
	"github.com/dailymotion/allure-go"
)

func AddAllureAttachmentJSON(name string, a ...interface{}) {
	aJSON, _ := json.MarshalIndent(a, "", "\t")
	err := allure.AddAttachment(name, allure.ApplicationJson, []byte(fmt.Sprintf("%+v", string(aJSON))))

	if err != nil {
		panic(err)
	}
}

func AddAllureAttachmentText(name string, a ...interface{}) {
	err := allure.AddAttachment(name, allure.TextPlain, []byte(fmt.Sprintf("%+v", a)))

	if err != nil {
		panic(err)
	}
}
