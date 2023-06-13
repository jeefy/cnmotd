package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	_ "net/http/pprof"
)

var Cmd = &cobra.Command{
	Use:  "cncfmotd",
	Long: "An ATOM MOTD generator for Cloud Native",
	RunE: run,
}

var args struct {
	debug        bool
	validate     bool
	entryDir     string
	cacheDir     string
	maxCacheAge  int
	cronSchedule string
	httpPort     int
}

var fullFeed MOTDItems

func init() {
	flags := Cmd.Flags()

	flags.StringVar(
		&args.entryDir,
		"entry-dir",
		"entries/",
		"Directory of entries to use for generating the feed",
	)
	flags.StringVar(
		&args.cacheDir,
		"cache-dir",
		"cache/",
		"Directory to cache common responses",
	)
	flags.IntVar(
		&args.maxCacheAge,
		"max-cache-age",
		0,
		"Max age (in hours) of files. Value of 0 means no files will be deleted (default 0)",
	)
	flags.IntVar(
		&args.httpPort,
		"http-port",
		8080,
		"Port to use for the HTTP server",
	)

	flags.BoolVar(
		&args.debug,
		"debug",
		false,
		"Enable debug logging",
	)

	flags.BoolVar(
		&args.validate,
		"validate",
		false,
		"Run a vlidation check on the entries",
	)

	flags.StringVar(
		&args.cronSchedule,
		"cron-schedule",
		"* */1 * * *",
		"Cron schedule to use for cleaning up cache files",
	)

	Cmd.RegisterFlagCompletionFunc("output-format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "prom"}, cobra.ShellCompDirectiveDefault
	})
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run(cmd *cobra.Command, argv []string) error {

	log.Println("Starting CNMOTD!")

	errors := LoadFeed()
	if args.validate {
		if len(errors) > 0 {
			log.Println("Errors found:")
			for _, err := range errors {
				log.Println(err)
			}
			os.Exit(1)
		}
		log.Println("No errors found")
		os.Exit(0)
	}

	StartCron()
	StartMetrics()
	StartHTTP()

	return nil
}
