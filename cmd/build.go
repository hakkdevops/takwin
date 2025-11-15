package cmd

import (
	"fmt"

	"github.com/hakkdevops/takwin/internal/config"
	"github.com/hakkdevops/takwin/internal/engine"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	target string
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project targets",
	Long: `Build one or more targets defined in the build configuration.
If no target is specified, builds the default target.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load configuration
		cfg, err := config.Load(viper.ConfigFileUsed())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Create build engine
		buildEngine := engine.NewEngine(cfg)

		// Build target
		if target != "" {
			return buildEngine.BuildTarget(target)
		}

		return buildEngine.BuildDefault()
	},
}

//nolint:gochecknoinits // init required for cobra command registration
func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&target, "target", "t", "", "specific target to build")
}
