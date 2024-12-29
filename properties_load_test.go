package properties

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromBytes(t *testing.T) {
	data := []byte(`
# This is a comment
key1=value1
key2=value2
key3=123

# Another comment
`)

	props, err := LoadFromBytes(data)
	assert.NoError(t, err)

	val, err := props.Get("key1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	val, err = props.Get("key2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)

	val, err = props.Get("key3")
	assert.NoError(t, err)
	assert.Equal(t, "123", val)

	_, err = props.Get("missingKey")
	assert.Error(t, err)
}
