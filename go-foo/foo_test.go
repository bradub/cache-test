package foo_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aarondl/opt/null"
	foo "github.com/bradub/cache-test/go-foo"
	"github.com/stretchr/testify/require"
)

func TestGreet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   null.Val[string]
		want string
	}{
		{name: "empty", in: null.Val[string]{}, want: "hello, stranger"},
		{name: "with name", in: null.From("alice"), want: "hello, alice"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, foo.Greet(tc.in))
		})
	}
}

func TestRouterHello(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(foo.NewRouter())
	t.Cleanup(server.Close)

	resp, err := http.Get(server.URL + "/hello?name=bob")
	require.NoError(t, err)
	require.NotNil(t, resp)
	t.Cleanup(func() { _ = resp.Body.Close() })

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "hello, bob", string(body))
}

func TestLogger(t *testing.T) {
	t.Parallel()

	logger := foo.Logger()
	logger.Info().Msg("hello from test")
}
