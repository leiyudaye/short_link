package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short_link/internal/shortener"
)

type CreateShortLinkReq struct {
	LongUrl string `json:"long_url"`
}

type CreateShortLinkResp struct {
	ShortLink string `json:"short_link"`
}

type GetLongLinkReq struct {
	ShortLink string `json:"short_link"`
}

type GetLongLinkResp struct {
	LongUrl string `json:"long_url"`
}

type Handler struct {
	service *shortener.Service
}

func NewHandler(service *shortener.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateShortLink(ctx *gin.Context) {
	var req CreateShortLinkReq
	// 绑定JSON数据到结构体
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortLink, err := h.service.CreateShortLink(ctx, req.LongUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &CreateShortLinkResp{
		ShortLink: shortLink,
	})
}

func (h *Handler) GetLongLink(ctx *gin.Context) {
	var req GetLongLinkReq
	// 绑定JSON数据到结构体
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	longUrl, err := h.service.GetLongUrl(ctx, req.ShortLink)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &GetLongLinkResp{
		LongUrl: longUrl,
	})
}
