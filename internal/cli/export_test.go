package cli

// RootCmd is exported for testing purposes only.
var RootCmd = rootCmd

func ResetGlobs() {
	cfgFile = ""
	service = ""
	outputDir = ""
	check = true
	noDownload = false
	downloadProgress = true

	for k := range formatFlags {
		*formatFlags[k] = false
	}
}
