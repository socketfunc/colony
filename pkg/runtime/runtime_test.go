package runtime

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRuntime_Execute(t *testing.T) {
	t.Parallel()

	file, err := ioutil.ReadFile("testdata/sample.wasm")
	require.NoError(t, err)

	tests := []struct {
		title     string
		name      string
		args      []interface{}
		expect    interface{}
		expectErr error
	}{
		{
			title: "add fn",
			name:  "add",
			args: []interface{}{
				int32(20),
				int32(80),
			},
			expect:    int32(100),
			expectErr: nil,
		},
		{
			title:     "hello fn",
			name:      "hello",
			args:      nil,
			expect:    "Hello World",
			expectErr: nil,
		},
	}
	for _, testcase := range tests {
		testcase := testcase
		t.Run(testcase.title, func(t *testing.T) {
			t.Parallel()
			runtime, err := New(file)
			require.NoError(t, err)
			defer runtime.Close()

			ret, err := runtime.Execute(testcase.name, testcase.args)
			if testcase.expectErr == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
			assert.Equal(t, testcase.expect, ret)
		})
	}
}
