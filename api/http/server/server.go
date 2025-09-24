package server

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"short_link/api/http/handler"
	"short_link/internal/shortener"
)

type Server struct {
	handle *handler.Handler
	r      *gin.Engine
}

func NewServer(service *shortener.Service) *Server {
	return &Server{
		handle: handler.NewHandler(service),
		r:      gin.Default(),
	}
}

func (s *Server) Init() {
	s.r.POST("/create_short_link", s.handle.CreateShortLink)
	s.r.POST("/get_long_url", s.handle.GetLongLink)
}

func (s *Server) Run() {
	err := s.r.Run() // 等价于 r.Run(":8080")
	if err != nil {
		slog.Error("server run failed", "err", err)
	}
}
