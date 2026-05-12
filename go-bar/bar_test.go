package bar_test

import (
	"testing"

	"github.com/aarondl/opt/null"
	bar "github.com/bradub/cache-test/go-bar"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		want int
	}{
		{name: "empty", in: nil, want: 0},
		{name: "one", in: []int{7}, want: 7},
		{name: "many", in: []int{1, 2, 3, 4}, want: 10},
		{name: "negatives", in: []int{-1, -2, 3}, want: 0},
		{name: "all zeros", in: []int{0, 0, 0}, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, bar.Sum(tc.in))
		})
	}
}

func TestDescribe(t *testing.T) {
	t.Parallel()

	require.Equal(t, "anonymous", bar.Describe(null.Val[string]{}))
	require.Equal(t, "name=alice", bar.Describe(null.From("alice")))
}

func TestLogger(t *testing.T) {
	t.Parallel()

	logger := bar.Logger()
	logger.Info().Msg("hello from test")
}
