package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsUtil(t *testing.T) {
	_, err := GetFileBytes("os.go")
	assert.NoError(t, err)
	t.Log("xxxxxx")
}
