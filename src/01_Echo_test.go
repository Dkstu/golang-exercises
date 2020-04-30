package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		text string
	}{
		{text: "Hello"},
		{text: "World"},
		{text: "Life"},
		{text: "is"},
		{text: "Good"},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			assert.Equal(tt.text, Echo(tt.text))
		})
	}
}
