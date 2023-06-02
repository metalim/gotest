package gotest_test

import (
	"testing"

	"github.com/metalim/gotest"
)

func TestGo(t *testing.T) {
	t.Log(gotest.Hello())
}
