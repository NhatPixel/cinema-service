package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/NhatPixel/cinema-service/internal/service"
	"github.com/NhatPixel/cinema-service/internal/dto"
	"github.com/NhatPixel/cinema-service/internal/validation"
)

type CinemaHandler struct {
	service *service.CinemaService
}

func NewCinemaHandler(s *service.CinemaService) *CinemaHandler {
	return &CinemaHandler{service: s}
}

func (h *CinemaHandler) Create(c *gin.Context) {
	var req dto.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validation.TranslateValidationError(err).Error()})
		return
	}

	if err := h.service.Create(req.ToModel()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Tạo rạp chiếu thất bại!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tạo rạp chiếu thành công."})
}

func (h *CinemaHandler) Get(c *gin.Context) {
	var req dto.GetRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": validation.TranslateValidationError(err).Error(),
		})
		return
	}

	cinemas, total, err := h.service.Get(
		req.Status,
		req.Keyword,
		req.Page,
		req.Limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lấy danh sách rạp chiếu thất bại!",
		})
		return
	}

	var resps []dto.GetResponse
	for _, cinema := range cinemas {
		var resp dto.GetResponse
		resp.FromModel(&cinema)
		resps = append(resps, resp)
	}

	totalPages := (total + req.Limit - 1) / req.Limit
	if totalPages == 0 {
		totalPages = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resps,
		"pagination": gin.H{
			"page":        req.Page,
			"limit":       req.Limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}


func (h *CinemaHandler) Update(c *gin.Context) {
	var req dto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validation.TranslateValidationError(err).Error()})
		return
	}

	if err := h.service.Update(req.ToModel()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cập nhật rạp chiếu thất bại!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật rạp chiếu thành công."})
}

func (h *CinemaHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID không được để trống!"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Xóa rạp chiếu thất bại!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa rạp chiếu thành công."})
}