package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	version string
)

var rootCmd = &cobra.Command{
	Use:          "sheep [time]",
	Short:        "Sleep with Sheep",
	Long:         "Sleep with Sheep.",
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var duration time.Duration
		if len(args) == 0 {
			fmt.Print(simpleSheep.Awake)
			return nil
		}

		if seconds, err := strconv.ParseFloat(args[0], 64); err == nil {
			duration = time.Duration(seconds * float64(time.Second))
		} else {
			duration, err = time.ParseDuration(args[0])
			if err != nil {
				return err
			}
		}

		m := NewModel(&Config{
			Duration: time.Duration(duration),
		})
		p := tea.NewProgram(m, tea.WithOutput(os.Stderr))
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

func init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version

}
