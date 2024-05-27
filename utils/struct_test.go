package utils_test

import (
	"net/url"
	"testing"

	"github.com/qiruos/jupiter/utils"
	"github.com/stretchr/testify/require"
)

func TestStructToUrlValues(t *testing.T) {
	type testStruct struct {
		Hello string `url:"hello"`
		World string `url:"world"`
	}

	t.Run("nil struct", func(t *testing.T) {
		uv, err := utils.StructToUrlValues(nil)
		require.Error(t, err)
		require.Nil(t, uv)
	})

	t.Run("url.Values struct", func(t *testing.T) {
		data := url.Values{
			"hello": []string{"hello"},
			"world": []string{"world"},
		}

		uv, err := utils.StructToUrlValues(data)
		require.NoError(t, err)
		require.NotNil(t, uv)
		require.Equal(t, "hello", uv.Get("hello"))
		require.Equal(t, "world", uv.Get("world"))
	})

	t.Run("struct with url tags", func(t *testing.T) {
		ts := testStruct{
			Hello: "hello",
			World: "world",
		}

		uv, err := utils.StructToUrlValues(ts)
		require.NoError(t, err)
		require.NotNil(t, uv)
		require.Equal(t, "hello", uv.Get("hello"))
		require.Equal(t, "world", uv.Get("world"))
	})
}
