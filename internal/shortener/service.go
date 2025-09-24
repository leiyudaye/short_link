package shortener

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"short_link/internal/idgen"
	"short_link/internal/storage"
	"time"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

type Options struct {
	Store          storage.Store
	Generator      idgen.Generator
	MaxGenAttempts int
}

type Service struct {
	store          storage.Store
	generator      idgen.Generator
	maxGenAttempts int
}

func NewService(opts *Options) *Service {
	return &Service{
		store:          opts.Store,
		generator:      opts.Generator,
		maxGenAttempts: opts.MaxGenAttempts,
	}
}

func (s *Service) CreateShortLink(ctx context.Context, longUrl string) (string, error) {
	if longUrl == "" {
		return "", fmt.Errorf("%w: longUrl is empty", ErrInvalidInput)
	}

	var shortLink string
	var err error
	for i := 0; i < s.maxGenAttempts; i++ {
		shortLink, err = s.generator.GeneratorShortLink(ctx, longUrl)
		if err != nil {
			return "", fmt.Errorf("attempt %d: failed to generate short link: %w", i+1, err)
		}

		link := &storage.Link{
			ShortLink:  shortLink,
			LongUrl:    longUrl,
			VisitCount: 0,
			CreatedAt:  time.Now().UnixMilli(),
		}
		err = s.store.Save(ctx, link)
		if err == nil {
			return shortLink, nil
		}
	}

	return "", fmt.Errorf("failed to generate short link: %w", err)
}

func (s *Service) GetLongUrl(ctx context.Context, shortLink string) (string, error) {
	if shortLink == "" {
		return "", fmt.Errorf("%w: shortLink is empty", ErrInvalidInput)
	}

	link, err := s.store.FindByShortLink(ctx, shortLink)
	if err != nil {
		return "", fmt.Errorf("failed to find link '%s': %w", shortLink, err)
	}

	go func() {
		err := s.store.IncrementVisitCount(ctx, shortLink)
		if err != nil {
			slog.Error("failed to increment visit count", "err", err)
			return
		}
		slog.Error("success to increment visit count", "shortLink", shortLink,
			"visitCount", link.VisitCount+1)
	}()

	return link.LongUrl, nil
}
