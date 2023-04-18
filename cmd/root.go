package cmd

import (
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/sheep/internal/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "sheep [time]",
	Short:        "Sleep with Sheep",
	Long:         "Sleep with Sheep.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}

		base := time.Second
		m := model.New(&model.Config{
			Duration: time.Duration(t * float64(base)),
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
