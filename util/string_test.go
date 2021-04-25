package util

import "testing"
import "github.com/stretchr/testify/assert"

func TestMd5(t *testing.T) {
	hash := Md5("hello,word")
	assert.NotEmpty(t, hash)
}
