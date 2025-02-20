package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/score-spec/score-humanitec/internal/version"
)

var (
	rootCmd = &cobra.Command{
		Use:   "score-humanitec",
		Short: "SCORE to Humanitec translator",
		Long: `SCORE is a specification for defining environment agnostic configuration for cloud based workloads.
This tool creates a Humanitec deployment from the SCORE specification.
Complete documentation is available at https://score.sh.`,
		Version: fmt.Sprintf("%s (build: %s; sha: %s)", version.Version, version.BuildTime, version.GitSHA),
	}
)

func init() {
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}
`)
}

func Execute() error {
	return rootCmd.Execute()
}
