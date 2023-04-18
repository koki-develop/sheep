package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/sheep/internal/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "sheep",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := model.New()
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
