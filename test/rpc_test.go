package test

import (
	"explorer-daemon/service/remote"
	"testing"
)

func TestRPC(t *testing.T) {
	remote.Experimental()
}

func TestHttp(t *testing.T) {
	remote.ExperimentalHttp()
}
