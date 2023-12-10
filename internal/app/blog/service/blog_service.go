package service

import (
	"context"

	"github.com/trikrama/Depublic/internal/app/blog/entity"
	"github.com/trikrama/Depublic/internal/app/blog/repository"
)

type BlogServiceInterface interface {
	GetAllBlog(c context.Context) ([]*entity.Blog, error)
	GetBlogByID(c context.Context, id int) (*entity.Blog, error)
	CreateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error)
	UpdateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error)
	DeleteBlog(c context.Context, id int) error
}

type BlogService struct {
	repo repository.BlogRepositoryInterface
}


func NewBlogService(repo repository.BlogRepositoryInterface) *BlogService {
	return &BlogService{
		repo: repo,
	}
}


func (s *BlogService) GetAllBlog(c context.Context) ([]*entity.Blog, error) {
	return s.repo.GetAllBlog(c)
}


func (s *BlogService) GetBlogByID(c context.Context, id int) (*entity.Blog, error) {
	return s.repo.GetBlogByID(c, id)
}


func (s *BlogService) CreateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error) {
	return s.repo.CreateBlog(c, blog)
}


func (s *BlogService) UpdateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error) {
	return s.repo.UpdateBlog(c, blog)
}


func (s *BlogService) DeleteBlog(c context.Context, id int) error {
	return s.repo.DeleteBlog(c, id)
}


