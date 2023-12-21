package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/trikrama/Depublic/internal/app/blog/entity"
	"gorm.io/gorm"
)

type BlogRepositoryInterface interface {
	GetAllBlog(c context.Context) ([]*entity.Blog, error)
	GetBlogByID(c context.Context, id int) (*entity.Blog, error)
	CreateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error)
	UpdateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error)
	DeleteBlog(c context.Context, id int) error
}

const (
	BlogKey = "blogs:all"
)

type BlogRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewBlogRepository(db *gorm.DB, redisClient *redis.Client) *BlogRepository {
	return &BlogRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *BlogRepository) GetAllBlog(c context.Context) ([]*entity.Blog, error) {
	blogs := make([]*entity.Blog, 0)
	val, err := r.redisClient.Get(context.Background(), BlogKey).Result()
	if err != nil {
		err := r.db.WithContext(c).Find(&blogs).Error
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(blogs)
		if err != nil {
			return nil, err
		}
		err = r.redisClient.Set(c, BlogKey, val, time.Duration(1)*time.Minute).Err()
		if err != nil {
			return nil, err
		}
		return blogs, nil
	}
	err = json.Unmarshal([]byte(val), &blogs)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepository) GetBlogByID(c context.Context, id int) (*entity.Blog, error) {
	blog := new(entity.Blog)
	err := r.db.WithContext(c).First(&blog, id).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) CreateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error) {
	err := r.db.WithContext(c).Create(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) UpdateBlog(c context.Context, blog *entity.Blog) (*entity.Blog, error) {
	err := r.db.WithContext(c).Model(&entity.Blog{}).Where("id = ?", blog.ID).Updates(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *BlogRepository) DeleteBlog(c context.Context, id int) error {
	err := r.db.WithContext(c).Delete(&entity.Blog{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
