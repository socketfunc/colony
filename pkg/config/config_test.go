package config

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	file, err := ioutil.ReadFile("./testdata/test.yaml")
	require.NoError(t, err)

	config, err := Parse(file)
	require.NoError(t, err)
	fmt.Println(config)
	function := config.Spec.GetFunction("test_func")
	fmt.Println(function)
	fmt.Println(function.Resources.Requests.Timeout())
}
