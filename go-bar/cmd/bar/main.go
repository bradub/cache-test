// Package main is the bar CLI entrypoint used to exercise CI caching.
package main

import (
	"os"

	bar "github.com/bradub/cache-test/go-bar"
	"github.com/spf13/cobra"
)

const exitFailure = 1

func main() {
	logger := bar.Logger()

	rootCmd := &cobra.Command{
		Use:   "bar",
		Short: "bar CLI",
		RunE: func(_ *cobra.Command, _ []string) error {
			total := bar.Sum([]int{1, 2, 3})
			logger.Info().Int("total", total).Msg("computed sum")

			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		logger.Error().Err(err).Msg("command failed")
		os.Exit(exitFailure)
	}
}
