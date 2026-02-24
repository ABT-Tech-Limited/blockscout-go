package blockscout

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"resty.dev/v3"
)

func TestNew(t *testing.T) {
	c := New("https://eth.blockscout.com", "")
	assert.NotNil(t, c)
}

func TestNewWithOptions(t *testing.T) {
	c := New("https://eth.blockscout.com", "test-key", Options{
		Timeout: 30 * time.Second,
		Verbose: false,
	})
	assert.NotNil(t, c)
}

func TestNewWithClient(t *testing.T) {
	restyCli := resty.New()
	c := NewWithClient("https://eth.blockscout.com", "", restyCli)
	assert.NotNil(t, c)
}

func TestDebug(t *testing.T) {
	c := New("https://eth.blockscout.com", "")
	d := c.Debug()
	assert.NotNil(t, d)
}
