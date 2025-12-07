package cli

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	service string
)

var rootCmd = &cobra.Command{
	Use:   "download-geofabrik",
	Short: "A command-line tool for downloading OSM files",
	Long:  `download-geofabrik is a CLI tool that downloads OpenStreetMap data from Geofabrik.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	initCLI()

	RegisterDownloadCmd()
	RegisterGenerateCmd()
	RegisterListCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func initCLI() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is geofabrik.yml)")
	rootCmd.PersistentFlags().StringVarP(&service, "service", "s", config.DefaultService,
		"Service to use (geofabrik, geofabrik-parse, openstreetmap.fr, geo2day, bbbike)")
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	// Bind flags to viper
	if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding config flag: %v\n", err)
	}

	if err := viper.BindPFlag("service", rootCmd.PersistentFlags().Lookup("service")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding service flag: %v\n", err)
	}

	if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding verbose flag: %v\n", err)
	}

	if err := viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding quiet flag: %v\n", err)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".download-geofabrik" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		if service != "" {
			viper.SetConfigName(service)
		} else {
			viper.SetConfigName(config.DefaultConfigFile)
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		slog.Info("Using config file", "file", viper.ConfigFileUsed())
	}
}
