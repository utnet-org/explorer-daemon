package fetch

import (
	"context"
	"encoding/base64"
	"explorer-daemon/service/remote"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tetratelabs/wazero"
)

func HandleContract() {
	res, err := remote.ViewContractCode("unc")
	if err != nil {
		return
	}
	base64Str := res.CodeBase64
	wasmFunctions, err := getWasmFunctions(base64Str)
	if err != nil {
		return
	}
	log.Debugf("wasm functions: %v", wasmFunctions)
}

func getWasmFunctions(wasmBase64 string) ([]string, error) {
	wasmBytes, err := base64.StdEncoding.DecodeString(wasmBase64)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %v", err)
	}
	ctx := context.Background()
	module, err := wazero.NewRuntime(ctx).CompileModule(ctx, wasmBytes)
	if err != nil {
		return nil, fmt.Errorf("wasm module parse failed: %v", err)
	}
	funcNames := getFunctionNames(module)

	return funcNames, nil
}

func getFunctionNames(module wazero.CompiledModule) []string {
	var funcNames []string
	exports := module.ExportedFunctions()
	for _, exp := range exports {
		funcNames = append(funcNames, exp.ExportNames()[0])
	}
	return funcNames
}
