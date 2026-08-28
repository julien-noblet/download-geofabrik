package cli

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Version is the CLI version injected at build time.
	Version = "dev"

	cfgFile string
	service string
)

var rootCmd = &cobra.Command{
	Use:     "download-geofabrik",
	Short:   "A command-line tool for downloading OSM files",
	Long:    `download-geofabrik is a CLI tool for downloading OpenStreetMap data and extracts from multiple providers.`,
	Version: Version,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

var once sync.Once

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	once.Do(func() {
		initCLI()

		RegisterDownloadCmd()
		RegisterGenerateCmd()
		RegisterListCmd()
	})

	rootCmd.Version = Version

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("root cmd execution failed: %w", err)
	}

	return nil
}

func initCLI() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is geofabrik.yml)")
	rootCmd.PersistentFlags().StringVarP(&service, "service", "s", config.DefaultService,
		"Service to use (geofabrik, geofabrik-parse, openstreetmap.fr, geo2day, bbbike, "+
			"movisda, planet.osm.ch, osm.kewl.lu, osm.fit.vutbr.cz, osmit-estratti, osm.kcwu.csie.org)")
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

func setupLogging() {
	var level slog.Level

	switch {
	case viper.GetBool("quiet"):
		level = slog.LevelError
	case viper.GetBool("verbose"):
		level = slog.LevelDebug
	default:
		level = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(handler))
}

func initConfig() {
	setupLogging()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory.
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
