package storage

import (
	"context"
	"errors"
)

var (
	ErrNotFound       = errors.New("storage: not found link")
	ErrShortLinkExist = errors.New("storage: short link exist")
)

type Link struct {
	ShortLink  string // 短链
	LongUrl    string // 原始长链
	VisitCount int64  // 访问次数
	CreatedAt  int64  // 创建时间
}

type Store interface {
	// Save 保存一个短链
	Save(ctx context.Context, link *Link) error

	// FindByShortLink 根据短链查找Link信息
	FindByShortLink(ctx context.Context, shortLink string) (*Link, error)

	// IncrementVisitCount 增加短链访问次数
	IncrementVisitCount(ctx context.Context, shortLink string) error
}
