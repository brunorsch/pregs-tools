package cli

import (
	"fmt"
	"strconv"
	"time"

	"brunorsch/pregs-tools/internal/db"
	"github.com/spf13/cobra"
)

// Aqui você pode adicionar funções auxiliares para a CLI futuramente. 

// ComandoAdicionarCompra cria um comando para adicionar uma nova compra
func ComandoAdicionarCompra() *cobra.Command {
	var descricao, categoria, observacoes string
	var valor float64
	var data string

	cmd := &cobra.Command{
		Use:   "adicionar",
		Short: "Adiciona uma nova compra",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse da data
			var dataCompra time.Time
			var err error
			if data == "" {
				dataCompra = time.Now()
			} else {
				dataCompra, err = time.Parse("2006-01-02", data)
				if err != nil {
					fmt.Printf("Erro ao parsear data: %v\n", err)
					return
				}
			}

			compra := db.Compra{
				Descricao:   descricao,
				Valor:       valor,
				Data:        dataCompra,
				Categoria:   categoria,
				Observacoes: observacoes,
			}

			if err := db.InserirCompra(compra); err != nil {
				fmt.Printf("Erro ao inserir compra: %v\n", err)
				return
			}

			fmt.Println("Compra adicionada com sucesso!")
		},
	}

	cmd.Flags().StringVarP(&descricao, "descricao", "d", "", "Descrição da compra (obrigatório)")
	cmd.Flags().Float64VarP(&valor, "valor", "v", 0, "Valor da compra (obrigatório)")
	cmd.Flags().StringVarP(&data, "data", "t", "", "Data da compra (formato: YYYY-MM-DD, padrão: hoje)")
	cmd.Flags().StringVarP(&categoria, "categoria", "c", "", "Categoria da compra")
	cmd.Flags().StringVarP(&observacoes, "observacoes", "o", "", "Observações adicionais")

	cmd.MarkFlagRequired("descricao")
	cmd.MarkFlagRequired("valor")

	return cmd
}

// ComandoListarCompras cria um comando para listar todas as compras
func ComandoListarCompras() *cobra.Command {
	return &cobra.Command{
		Use:   "listar",
		Short: "Lista todas as compras",
		Run: func(cmd *cobra.Command, args []string) {
			compras, err := db.ListarCompras()
			if err != nil {
				fmt.Printf("Erro ao listar compras: %v\n", err)
				return
			}

			if len(compras) == 0 {
				fmt.Println("Nenhuma compra encontrada.")
				return
			}

			fmt.Printf("\n%-5s %-30s %-10s %-12s %-15s %s\n", "ID", "Descrição", "Valor", "Data", "Categoria", "Observações")
			fmt.Println("----------------------------------------------------------------------------------------")
			
			for _, compra := range compras {
				fmt.Printf("%-5d %-30s R$ %-8.2f %-12s %-15s %s\n",
					compra.ID,
					truncateString(compra.Descricao, 28),
					compra.Valor,
					compra.Data.Format("02/01/2006"),
					truncateString(compra.Categoria, 13),
					truncateString(compra.Observacoes, 50))
			}
			fmt.Println()
		},
	}
}

// ComandoBuscarCompra cria um comando para buscar uma compra específica
func ComandoBuscarCompra() *cobra.Command {
	return &cobra.Command{
		Use:   "buscar [ID]",
		Short: "Busca uma compra pelo ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("ID deve ser um número inteiro")
				return
			}

			compra, err := db.BuscarCompraPorID(id)
			if err != nil {
				fmt.Printf("Erro ao buscar compra: %v\n", err)
				return
			}

			fmt.Printf("\nCompra #%d:\n", compra.ID)
			fmt.Printf("Descrição: %s\n", compra.Descricao)
			fmt.Printf("Valor: R$ %.2f\n", compra.Valor)
			fmt.Printf("Data: %s\n", compra.Data.Format("02/01/2006 15:04"))
			fmt.Printf("Categoria: %s\n", compra.Categoria)
			fmt.Printf("Observações: %s\n", compra.Observacoes)
		},
	}
}

// ComandoDeletarCompra cria um comando para deletar uma compra
func ComandoDeletarCompra() *cobra.Command {
	return &cobra.Command{
		Use:   "deletar [ID]",
		Short: "Deleta uma compra pelo ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("ID deve ser um número inteiro")
				return
			}

			if err := db.DeletarCompra(id); err != nil {
				fmt.Printf("Erro ao deletar compra: %v\n", err)
				return
			}

			fmt.Printf("Compra #%d deletada com sucesso!\n", id)
		},
	}
}

// truncateString trunca uma string para o tamanho especificado
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
} 