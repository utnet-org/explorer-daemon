package test

import (
	"explorer-daemon/es"
	"explorer-daemon/service/fetch"
	"testing"
)

func TestEsInsert(t *testing.T) {
	es.Init()
	fetch.InitFetchData()
}
