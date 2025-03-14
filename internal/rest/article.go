package rest

import (
	"context"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/bxcodec/go-clean-arch/domain"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleService represent the article's usecases
//
//go:generate mockery --name ArticleService
type ArticleService interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]domain.Article, string, error)
	GetByID(ctx context.Context, id int64) (domain.Article, error)
	Update(ctx context.Context, ar *domain.Article) error
	GetByTitle(ctx context.Context, title string) (domain.Article, error)
	Store(context.Context, *domain.Article) error
	Delete(ctx context.Context, id int64) error
}

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	Service ArticleService
}

const defaultNum = 10

// NewArticleHandler will initialize the articles/ resources endpoint

// FetchArticle will fetch the article based on given params


// GetByID will get article by given id

func isRequestValid(m *domain.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body


// Delete will delete article by given param

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
