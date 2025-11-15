package cmd

import (
	"fmt"
	"strings"

	"github.com/hakkim/takwin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list-targets",
	Short: "List available build targets",
	Long:  `Display all available build targets defined in the configuration file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load configuration
		cfg, err := config.Load(viper.ConfigFileUsed())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if len(cfg.Targets) == 0 {
			fmt.Println("No targets defined in configuration")
			return nil
		}

		fmt.Printf("Available targets (%d):\n\n", len(cfg.Targets))

		for _, target := range cfg.Targets {
			fmt.Printf("  %s\n", target.Name)
			fmt.Printf("    Type: %s\n", target.Type)
			if len(target.Sources) > 0 {
				fmt.Printf("    Sources: %s\n", strings.Join(target.Sources, ", "))
			}
			if target.Output != "" {
				fmt.Printf("    Output: %s\n", target.Output)
			}
			fmt.Println()
		}

		return nil
	},
}

//nolint:gochecknoinits // init required for cobra command registration
func init() {
	rootCmd.AddCommand(listCmd)
}
