package cli

// RootCmd is exported for testing purposes only.
var RootCmd = rootCmd

func ResetGlobs() {
	cfgFile = ""
	service = ""
}
