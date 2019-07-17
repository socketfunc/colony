package runtime

//go:generate mkdir -p mock
//go:generate mockgen -source=runtime.go -package=mock -destination=mock/runtime.go
import (
	"strings"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	errors "golang.org/x/xerrors"
)

const (
	eof         byte = 0
	defaultSize      = 128
)

// Runtime WebAssembly runtime
// not thread safe and cannot be used in parallel across goroutines.
type Runtime interface {
	Close()
	Execute(name string, args []interface{}) (interface{}, error)
}

type runtime struct {
	instance *wasm.Instance
}

func New(file []byte) (Runtime, error) {
	instance, err := wasm.NewInstance(file)
	if err != nil {
		return nil, err
	}
	runtime := &runtime{
		instance: &instance,
	}
	return runtime, nil
}

func (r *runtime) Close() {
	r.instance.Close()
}

func (r *runtime) Execute(name string, args []interface{}) (interface{}, error) {
	exec, ok := r.instance.Exports[name]
	if !ok {
		return nil, errors.New("runtime: unexpected function")
	}
	ret, err := exec(args...)
	if err != nil {
		return nil, errors.Errorf("runtime: failed to execute: %w", err)
	}
	switch ret.GetType() {
	case wasm.TypeI32:
		ptr := ret.ToI32()
		data := r.instance.Memory.Data()
		buf := new(strings.Builder)
		buf.Grow(defaultSize)
		for _, b := range data[ptr:] {
			if b == eof {
				break
			}
			buf.WriteByte(b)
		}
		if buf.Len() > 0 {
			return buf.String(), nil
		}
		return ptr, nil
	case wasm.TypeI64:
		return ret.ToI64(), nil
	case wasm.TypeF32:
		return ret.ToF32(), nil
	case wasm.TypeF64:
		return ret.ToF64(), nil
	default:
		return nil, nil
	}
}
