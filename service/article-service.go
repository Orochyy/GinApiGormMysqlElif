package service

import (
	"GinApiGormMysqlElif/dto"
	"GinApiGormMysqlElif/entity"
	"GinApiGormMysqlElif/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
)

type ArticleService interface {
	Insert(a dto.ArticleCreateDTO) entity.Articles
	Update(a dto.ArticleUpdateDTO) entity.Articles
	Delete(a entity.Articles)
	All() []entity.Articles
	FindByID(articleID uint64) entity.Articles
	IsAllowedToEdit(userID string, articleID uint64) bool
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) *articleService {
	return &articleService{
		articleRepository: articleRepo,
	}
}

func (service *articleService) Insert(a dto.ArticleCreateDTO) entity.Articles {
	article := entity.Articles{}
	err := smapping.FillStruct(&article, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.articleRepository.InsertArticle(article)
	return res
}

func (service *articleService) Update(a dto.ArticleCreateDTO) entity.Articles {
	article := entity.Articles{}
	err := smapping.FillStruct(&article, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.articleRepository.UpdateArticle(article)
	return res
}

func (service *articleService) Delete(a entity.Articles) {
	service.articleRepository.DeleteArticle(a)
}

func (sercice *articleService) All() []entity.Articles {
	return sercice.articleRepository.AllArticle()
}

func (service *articleService) FindByID(articleID uint64) entity.Articles {
	return service.articleRepository.FindArticleByID(articleID)
}

func (service *articleService) IsAllowedToEdt(userID string, articleID uint64) bool {
	a := service.articleRepository.FindArticleByID(articleID)
	id := fmt.Sprintf("%v", a.UserID)
	return userID == id
}
