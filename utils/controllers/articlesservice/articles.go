package articlesservicecontrollers

import (
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/google/uuid"
	"sample_go_grpc_testing/utils/fakers"
)

func GetRandomArticle() *articlesservice.Article {
	return &articlesservice.Article{
		Id:          uuid.NewString(),
		Title:       fakers.RandomString(30),
		Author:      fakers.RandomString(20),
		Description: fakers.RandomString(100),
	}
}
