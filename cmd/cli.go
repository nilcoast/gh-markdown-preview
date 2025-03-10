package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var Version string

var verbose = false

type Param struct {
	filename       string
	markdownMode   bool
	reload         bool
	forceLightMode bool
	forceDarkMode  bool
	autoOpen       bool
	mermaidVersion string
}

var rootCmd = &cobra.Command{
	Use:   "gh-markdown-preview",
	Short: "GitHub CLI extension to preview Markdown",
	Run: func(cmd *cobra.Command, args []string) {

		showVerionFlag, _ := cmd.Flags().GetBool("version")
		if showVerionFlag {
			showVersion()
			os.Exit(0)
		}

		filename := ""
		if len(args) > 0 {
			filename = args[0]
		}

		verbose, _ = cmd.Flags().GetBool("verbose")

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")

		server := Server{host: host, port: port}

		disableReload, _ := cmd.Flags().GetBool("disable-reload")
		reload := true
		if disableReload {
			reload = false
		}

		forceLightMode, _ := cmd.Flags().GetBool("light-mode")
		forceDarkMode, _ := cmd.Flags().GetBool("dark-mode")

		markdownMode, _ := cmd.Flags().GetBool("markdown-mode")

		mermaidVersion, _ := cmd.Flags().GetString("mermaid-version")

		disableAutoOpen, _ := cmd.Flags().GetBool("disable-auto-open")
		autoOpen := true
		if disableAutoOpen {
			autoOpen = false
		}

		param := &Param{
			filename:       filename,
			markdownMode:   markdownMode,
			reload:         reload,
			forceLightMode: forceLightMode,
			forceDarkMode:  forceDarkMode,
			autoOpen:       autoOpen,
			mermaidVersion: mermaidVersion,
		}

		err := server.Serve(param)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("port", "p", 3333, "TCP port number of this server")
	rootCmd.Flags().StringP("host", "", "localhost", "Hostname this server will bind")
	rootCmd.Flags().BoolP("version", "", false, "Show the version")
	rootCmd.Flags().BoolP("disable-reload", "", false, "Disable live reloading")
	rootCmd.Flags().BoolP("markdown-mode", "", false, "Force \"markdown\" mode (rather than default \"gfm\")")
	rootCmd.Flags().BoolP("disable-auto-open", "", false, "Disable auto opening your browser")
	rootCmd.Flags().BoolP("verbose", "", false, "Show verbose output")
	rootCmd.Flags().BoolP("light-mode", "", false, "Force light mode")
	rootCmd.Flags().BoolP("dark-mode", "", false, "Force dark mode")
	rootCmd.Flags().StringP("mermaid-version", "", "11.3.0", "Set the mermaidjs version")
}

func showVersion() {
	fmt.Printf("gh-markdown-preview version %s\n", Version)
}
