package runtime

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRuntime_Execute(t *testing.T) {
	file, err := ioutil.ReadFile("testdata/sum.wasm")
	require.NoError(t, err)

	runtime, err := New(file)
	require.NoError(t, err)
	defer runtime.Close()

	args := []interface{}{int32(20), int32(80)}
	ret, err := runtime.Execute("sum", args)
	assert.NoError(t, err)
	assert.Equal(t, int32(100), ret)
}
