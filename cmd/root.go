package cmd

import (
	"os"
	"runtime/debug"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/sheep/internal/model"
	"github.com/spf13/cobra"
)

var (
	version string

	flagSecond bool
	flagMinute bool
	flagHour   bool
)

var rootCmd = &cobra.Command{
	Use:          "sheep [time]",
	Short:        "Sleep with Sheep",
	Long:         "Sleep with Sheep.",
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var t float64
		if len(args) > 0 {
			f, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return err
			}
			t = f
		}

		base := time.Second
		switch {
		case flagSecond:
			base = time.Second
		case flagMinute:
			base = time.Minute
		case flagHour:
			base = time.Hour
		}

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

func init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version

	rootCmd.Flags().BoolVar(&flagSecond, "second", false, "set the time unit to seconds (default)")
	rootCmd.Flags().BoolVar(&flagMinute, "minute", false, "set the time unit to minutes")
	rootCmd.Flags().BoolVar(&flagHour, "hour", false, "set the time unit to hours")
	rootCmd.MarkFlagsMutuallyExclusive("second", "minute", "hour")
}
