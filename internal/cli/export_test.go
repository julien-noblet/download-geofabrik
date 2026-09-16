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
	downloadProgress = true

	for _, flagDef := range formatFlagList {
		_ = downloadCmd.Flags().Set(flagDef.key, "false")
	}
}
