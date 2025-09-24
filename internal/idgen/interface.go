package idgen

import "context"

type Generator interface {
	// GeneratorShortLink 生成一个短链
	GeneratorShortLink(ctx context.Context, url string) (string, error)
}
