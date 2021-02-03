// Based on https://github.com/shomali11/util/blob/master/xcompressions/xcompressions.go

package elcompress

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGzipCompress(t *testing.T) {
	data, err := GzipCompress([]byte("Raed Shomali"))
	assert.Nil(t, err)

	decompressedData, err := GzipDecompress(data)
	assert.Nil(t, err)

	assert.Equal(t, string(decompressedData), "Raed Shomali")
}

func TestGzipDecompress(t *testing.T) {
	_, err := GzipDecompress([]byte("Kaka"))
	assert.NotNil(t, err)
}
