package idgen_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"short_link/internal/idgen"
	"testing"
)

func TestHashIdGen(t *testing.T) {
	generator := idgen.NewHashGenerator()

	// 生成单个短链不报错，长度符合要求
	shortLink, err := generator.GeneratorShortLink(context.Background(), "https://www.baidu.com")
	assert.NoError(t, err)
	assert.Equal(t, idgen.DefaultHashLength, len(shortLink))
	t.Logf("short link: %s", shortLink)

	// 生成一批短链，重复次数不超过2次
	shortLinkMap := make(map[string]int, 100000)
	for range 100000 {
		shortLink, err := generator.GeneratorShortLink(context.Background(), "https://www.baidu.com")
		assert.NoError(t, err)
		shortLinkMap[shortLink]++
		if shortLinkMap[shortLink] > 2 {
			t.Errorf("short link: %s exist", shortLink)
		}
	}

}
