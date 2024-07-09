package test

import (
	"explorer-daemon/es"
	"explorer-daemon/service/fetch"
	"fmt"
	"testing"
)

func TestCompleteTxn(t *testing.T) {
	es.Init()
	err := fetch.CompleteTxnDetails()
	if err != nil {
		fmt.Println("[TestCompleteTxn] err:", err)
	}
}
