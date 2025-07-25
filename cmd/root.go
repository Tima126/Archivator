package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{

	Short: "Simple archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Handler(err)
	}

}

func Handler(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
