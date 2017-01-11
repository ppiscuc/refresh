package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile    string
	approot    string
	igndir     []string
	ext        []string
	buildpath  string
	builddelay string
	binaryname string
	cmdflags   []string
	colors     bool
	logname    string
)

var RootCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh is a command line tool that builds and (re)starts your Go application everytime you save a Go or template file.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Refresh (%s)\n\n", Version)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	var empty []string
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "refresh.yml", "path to configuration file")
	RootCmd.PersistentFlags().StringVar(&approot, "approot", "", "app root")
	RootCmd.PersistentFlags().StringSliceVar(&igndir, "igndir", empty, "ignore directories")
	RootCmd.PersistentFlags().StringSliceVar(&ext, "ext", empty, "included extensions")
	RootCmd.PersistentFlags().StringVar(&buildpath, "buildpath", "", "build path")
	RootCmd.PersistentFlags().StringVar(&builddelay, "builddelay", "", "build delay")
	RootCmd.PersistentFlags().StringVar(&binaryname, "binaryname", "", "binary name")
	RootCmd.PersistentFlags().StringSliceVar(&cmdflags, "cmdflags", empty, "command flags")
	RootCmd.PersistentFlags().BoolVar(&colors, "colors", false, "colors")
	RootCmd.PersistentFlags().StringVar(&logname, "logname", "", "log name")
}
