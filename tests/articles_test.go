package tests

import (
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/dailymotion/allure-go"
	"github.com/dailymotion/allure-go/severity"
	articlesservicechecks "sample_go_grpc_testing/utils/assertions/articlesservice"
	"sample_go_grpc_testing/utils/common"
	articlesservicecontrollers "sample_go_grpc_testing/utils/controllers/articlesservice"
	"sample_go_grpc_testing/utils/reports"
	"testing"
)

func TestGetArticle(t *testing.T) {
	allure.Test(t, reports.ArticlesServiceFeature, reports.ArticlesSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.ArticlesTag),
		allure.Name("Get article"),
		allure.Action(func() {
			c, g := common.SetupTesting(t)

			article := articlesservicecontrollers.GetRandomArticle()
			c.ArticlesService.CreateArticle(g, &articlesservice.CreateArticleRequest{Article: article})

			response := c.ArticlesService.GetArticle(g, &articlesservice.GetArticleRequest{ArticleId: article.Id})

			articlesservicechecks.CheckArticle(g, response.Article, article)
		}),
	)
}

func TestCreateArticle(t *testing.T) {
	allure.Test(t, reports.ArticlesServiceFeature, reports.ArticlesSuite,
		allure.Severity(severity.Critical),
		allure.Tags(reports.ArticlesTag),
		allure.Name("Create article"),
		allure.Action(func() {
			c, g := common.SetupTesting(t)

			article := articlesservicecontrollers.GetRandomArticle()
			response := c.ArticlesService.CreateArticle(g, &articlesservice.CreateArticleRequest{Article: article})

			articlesservicechecks.CheckArticle(g, response.Article, article)
		}),
	)
}
