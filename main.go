package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/brunorsch/pregs-tool/internal/tui"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "pregs-tool",
		Short: "Ferramenta CLI e TUI para pregs-tool",
	}

	rootCmd.AddCommand(tuiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Inicia a interface TUI",
	Run: func(cmd *cobra.Command, args []string) {
		tui.RunTUI()
	},
}
