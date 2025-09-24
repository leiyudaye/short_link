package idgen

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const DefaultHashLength = 6

type HashGenerator struct {
}

func NewHashGenerator() Generator {
	rand.Seed(time.Now().UnixNano())
	return &HashGenerator{}
}

func (g *HashGenerator) GeneratorShortLink(ctx context.Context, url string) (string, error) {
	if url == "" {
		return "", errors.New("idgen: url is empty")
	}

	hasher := sha256.New()
	hasher.Write([]byte(url))
	hasher.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	hasher.Write([]byte(fmt.Sprintf("%d", rand.Int63())))
	haseBytes := hasher.Sum(nil)
	encode := base64.URLEncoding.EncodeToString(haseBytes)

	if len(encode) <= DefaultHashLength {
		return encode, nil
	}
	return encode[:DefaultHashLength], nil
}
