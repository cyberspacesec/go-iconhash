package cmd

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once.
func Execute() {
	// Initialize all commands and flags
	Initialize()

	// Print logo (before any Cobra output)
	PrintLogo()

	// Execute the root command
	err := RootCmd.Execute()
	if err != nil {
		HandleError(err)
	}
}
