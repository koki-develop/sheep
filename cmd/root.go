package cmd

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/sheep/internal/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "sheep",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := model.New(&model.Config{
			Duration: time.Second * 2,
		})
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			return err
		}
		if m.Aborted() {
			os.Exit(130)
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
