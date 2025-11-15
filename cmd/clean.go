package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hakkdevops/takwin/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean build artifacts",
	Long:  `Remove all build artifacts and output directories.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load configuration to get output directory
		cfg, err := config.Load(viper.ConfigFileUsed())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		outputDir := cfg.Build.OutputDir
		if outputDir == "" {
			outputDir = "build"
		}

		// Check if directory exists
		if _, statErr := os.Stat(outputDir); os.IsNotExist(statErr) {
			fmt.Printf("Build directory '%s' does not exist\n", outputDir)
			return nil
		}

		// Remove the directory
		if removeErr := os.RemoveAll(outputDir); removeErr != nil {
			return fmt.Errorf("failed to remove build directory: %w", removeErr)
		}

		absPath, err := filepath.Abs(outputDir)
		if err != nil {
			absPath = outputDir
		}
		fmt.Printf("Cleaned build directory: %s\n", absPath)
		return nil
	},
}

//nolint:gochecknoinits // init required for cobra command registration
func init() {
	rootCmd.AddCommand(cleanCmd)
}
