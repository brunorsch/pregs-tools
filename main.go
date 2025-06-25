package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"brunorsch/pregs-tools/internal/tui"
	"brunorsch/pregs-tools/internal/db"
)

func main() {
	// Inicializa o banco de dados
	if err := db.InitDB(); err != nil {
		fmt.Printf("Erro ao inicializar banco de dados: %v\n", err)
		os.Exit(1)
	}
	defer db.CloseDB()

	rootCmd := &cobra.Command{
		Use:   "pt",
		Short: "Ferramenta CLI e TUI para pregs-tools",
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
