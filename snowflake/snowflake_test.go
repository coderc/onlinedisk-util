package snowflake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnowFlake(t *testing.T) {
	id, err := GetId(1, 1)
	assert.NoError(t, err)
	assert.NotEqual(t, id, 0)
}
