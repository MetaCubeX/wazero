package binaryencoding

import (
	"github.com/metacubex/wazero/internal/wasm"
)

func encodeConstantExpression(expr wasm.ConstantExpression) (ret []byte) {
	ret = append(ret, expr.Data...)
	return
}
