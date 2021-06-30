package chatwork

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(apiKey)
	refute(t, c, nil)
}
