package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/internal/app/blog/entity"
	"github.com/trikrama/Depublic/internal/app/blog/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type BlogHandler struct {
	blogService service.BlogServiceInterface
}

func NewBlogHandler(cfg *config.Config, blogService service.BlogServiceInterface) *BlogHandler {
	return &BlogHandler{
		blogService: blogService,
	}
}

func (h *BlogHandler) GetAllBlog(c echo.Context) error {
	blogs, err := h.blogService.GetAllBlog(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	return c.JSON(200, echo.Map{"blogs": blogs})
}

func (h *BlogHandler) GetBlogByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	blog, err := h.blogService.GetBlogByID(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"blog": blog})
}

func (h *BlogHandler) CreateBlog(c echo.Context) error {
	blogRequest := entity.BlogRequest{}
	if err := c.Bind(&blogRequest); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	newBlog := entity.NewBlog(blogRequest)
	blog, err := h.blogService.CreateBlog(c.Request().Context(), newBlog)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"blog": blog,
	})
}

func (h *BlogHandler) UpdateBlog(c echo.Context) error {
	blog := entity.BlogRequestUpdate{}
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	blogRequest := entity.NewBlogUpdate(blog)
	blogUpdate, err := h.blogService.UpdateBlog(c.Request().Context(), blogRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"message": err.Error()})
	}
	blogResp := entity.NewBlogRespose(*blogUpdate)
	return c.JSON(http.StatusOK, echo.Map{
		"blog":    blogResp,
	})
}

func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.blogService.DeleteBlog(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "blog deleted",
	})
}
