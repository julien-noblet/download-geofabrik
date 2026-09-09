package cli

import (
	"cmp"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var (
	// Version is the CLI version injected at build time.
	Version = "dev"

	cfgFile string
	service string
)

var rootCmd = &cobra.Command{
	Use:           "download-geofabrik",
	Short:         "A command-line tool for downloading OSM files",
	Long:          `download-geofabrik is a CLI tool for downloading OpenStreetMap data and extracts from multiple providers.`,
	Version:       Version,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

var setupCLI = sync.OnceFunc(func() {
	initCLI()

	RegisterDownloadCmd()
	RegisterGenerateCmd()
	RegisterListCmd()
	RegisterMCPCmd()
})

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	setupCLI()

	rootCmd.Version = Version

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("root cmd execution failed: %w", err)
	}

	return nil
}

func initCLI() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is geofabrik.yml)")
	rootCmd.PersistentFlags().StringVarP(&service, "service", "s", catalog.DefaultService,
		"Service to use (geofabrik, geofabrik-parse, openstreetmap.fr, geo2day, bbbike, "+
			"movisda, planet.osm.ch, osm.kewl.lu, osm.fit.vutbr.cz, osmit-estratti, osm.kcwu.csie.org)")
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose mode")
	rootCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	_ = rootCmd.RegisterFlagCompletionFunc("service", func(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		provider.RegisterDefaultProviders()

		services := provider.List()

		var matches []string

		for _, s := range services {
			if strings.HasPrefix(s, toComplete) {
				matches = append(matches, s)
			}
		}

		return matches, cobra.ShellCompDirectiveNoFileComp
	})

	// Bind flags to viper
	flags := []string{"config", "service", "verbose", "quiet"}
	for _, flag := range flags {
		if err := viper.BindPFlag(flag, rootCmd.PersistentFlags().Lookup(flag)); err != nil {
			slog.Error("Failed to bind CLI flag", "flag", flag, "error", err)
		}
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
		// Search config in current directory and system directory.
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/download-geofabrik")
		viper.SetConfigType("yaml")

		viper.SetConfigName(cmp.Or(service, catalog.DefaultConfigFile))
	}

	viper.SetEnvPrefix("DOWNLOAD_GEOFABRIK")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFound) && cfgFile != "" {
			slog.Warn("Error reading config file", "file", cfgFile, "error", err)
		}
	} else {
		slog.Info("Using config file", "file", viper.ConfigFileUsed())
	}
}
