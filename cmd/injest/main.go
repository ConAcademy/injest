package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/conacademy/injest/pkg/jester"
	"github.com/spf13/pflag"
)

/////////////////////////////////////////////////////////////////////////////////////

var usageFormatShort string = `usage:  %s <options> [input]`

var usageFormat string = `usage:  %s <options> [input]

"injest" wraps input into a joke and emits it to stdout.

`

/////////////////////////////////////////////////////////////////////////////////////

var allTheJesters = map[string]jester.Jester{
	"knock":   jester.KnockKnockJoke{},
	"yomamma": jester.YoMammaJoke{},
}

// ///////////////////////////////////////////////////////////////////////////////////
func main() {
	var jokeKind string
	var useStdin bool
	var listKinds bool
	var showHelp bool

	pflag.StringVarP(&jokeKind, "joke", "j", "knock", "use joke kind (default: 'knock')")
	pflag.BoolVarP(&useStdin, "stdin", "s", false, "use stdin instead of command arguments")
	pflag.BoolVarP(&listKinds, "list", "l", false, "list joke kinds")
	pflag.BoolVarP(&showHelp, "help", "h", false, "show help")
	pflag.Parse()

	if showHelp {
		fmt.Fprintf(os.Stdout, usageFormat, os.Args[0])
		pflag.PrintDefaults()
		os.Exit(0)
	}

	if listKinds {
		for kind, _ := range allTheJesters {
			fmt.Fprintf(os.Stdout, "%s\n", kind)
		}
		os.Exit(0)
	}

	jojo, ok := allTheJesters[jokeKind]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown jester: %s.  Try --list\n", jokeKind)
		os.Exit(1)
	}

	var ingested string
	if useStdin {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading stdin: %s\n", err.Error())
			os.Exit(1)
		}
		ingested = string(bytes)
	} else {
		args := pflag.Args()
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "no input for joke, add arguments or use --stdin\n")
			os.Exit(1)
		}
		ingested = strings.Join(args, " ")
	}

	joke := jojo.MakeJoke(ingested)
	fmt.Fprintf(os.Stdout, "%s\n", joke)
}
