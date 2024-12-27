package services

import (
	"go_SayHi/models"
	"go_SayHi/repositories"

	"github.com/mlogclub/simple/sqls"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

func (s *articleService) Get(id int64) *models.Article {
	return repositories.ArticleRepository.Get(sqls.DB(), id)
}
