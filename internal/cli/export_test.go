package cli

// RootCmd is exported for testing purposes only.
var RootCmd = rootCmd

// DownloadCmd is exported for testing purposes only.
var DownloadCmd = downloadCmd

func ResetGlobs() {
	cfgFile = ""
	service = ""
	outputDir = ""
	check = true
	noDownload = false
	force = false
	downloadProgress = true
	downloadRetries = defaultDownloadRetries

	for _, flagDef := range formatFlagList {
		_ = downloadCmd.Flags().Set(flagDef.key, "false")
	}
}
