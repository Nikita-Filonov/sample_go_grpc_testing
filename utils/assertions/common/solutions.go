package solutions

import (
	"fmt"
	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
)

func AssertToEqual(g gomega.Gomega, actual interface{}, expected interface{}, description string) {
	step := fmt.Sprintf("Checking that '%s' equals to '%s'", description, PrettifyValue(expected))

	allure.Step(allure.Description(step), allure.Action(func() {
		g.Expect(actual).To(gomega.Equal(expected), step)
	}))
}
