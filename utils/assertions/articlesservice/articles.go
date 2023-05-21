package articlesservicechecks

import (
	"fmt"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
	solutions "sample_go_grpc_testing/utils/assertions/common"
)

func CheckArticle(g gomega.Gomega, actualArticle, expectedArticle *articlesservice.Article) {
	allure.Step(allure.Description("Checking article"), allure.Action(func() {
		solutions.AssertToEqual(g, actualArticle.Id, expectedArticle.Id, "Article Id")
		solutions.AssertToEqual(g, actualArticle.Title, expectedArticle.Title, "Article Title")
		solutions.AssertToEqual(g, actualArticle.Author, expectedArticle.Author, "Article Author")
		solutions.AssertToEqual(g, actualArticle.Description, expectedArticle.Description, "Article Description")
	}))
}

func CheckArticleError(g gomega.Gomega, actualError, expectedError *articlesservice.Error) {
	allure.Step(allure.Description("Checking article error"), allure.Action(func() {
		solutions.AssertToEqual(g, actualError.Type, expectedError.Type, "Error Type")
		solutions.AssertToEqual(g, actualError.Message, expectedError.Message, "Error Message")
	}))
}

func CheckArticleNotFoundError(g gomega.Gomega, actualError *articlesservice.Error, articleId string) {
	expectedError := articlesservice.Error{
		Type:    articlesservice.ErrorType_NOT_FOUND,
		Message: fmt.Sprintf("Article with Id %s not found", articleId),
	}
	CheckArticleError(g, actualError, &expectedError)
}
