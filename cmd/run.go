package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ppiscuc/refresh/refresh"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "start", "build", "watch"},
	Short:   "watches your files and rebuilds/restarts your app accordingly.",
	Run: func(cmd *cobra.Command, args []string) {
		Run(args)
	},
}

func Run(args []string) {
	ctx := context.Background()
	RunWithContext(args, ctx)
}

func RunWithContext(args []string, ctx context.Context) {
	c := refresh.DefaultConfiguration()
	if cfgFile != "" {
		err := c.Load(cfgFile)
		if err != nil {
			log.Fatalln(err)
			os.Exit(-1)
		}
	} else {
		if approot != "" {
			c.AppRoot = approot
		}
		if len(igndir) > 0 {
			c.IgnoredFolders = igndir
		}
		if len(ext) > 0 {
			c.IncludedExtensions = ext
		}
		if buildpath != "" {
			c.BuildPath = buildpath
		}
		if builddelay != "" {
			dur, err := time.ParseDuration(builddelay)
			if err != nil {
				log.Fatalln(err)
				os.Exit(-1)
			}
			c.BuildDelay = dur
		}
		if binaryname != "" {
			c.BinaryName = binaryname
		}
		if len(cmdflags) > 0 {
			c.CommandFlags = cmdflags
		}
		if colors == true {
			c.EnableColors = true
		}
		if logname != "" {
			c.LogName = logname
		}
	}
	r := refresh.NewWithContext(&c, ctx)
	if err := r.Start(); err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}
