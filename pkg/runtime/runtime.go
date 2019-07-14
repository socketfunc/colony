package runtime

//go:generate mkdir -p mock
//go:generate mockgen -source=runtime.go -package=mock -destination=mock/runtime.go
import (
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	errors "golang.org/x/xerrors"
)

// Runtime WebAssembly runtime
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
		return ret.ToI32(), nil
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
