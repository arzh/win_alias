package main

import (
	"fmt"
	"github.com/arzh/clu"
	"os"
	"osext"
)

const (
	helpMsg = `alias - Windows command line aliasing tool
----------------------------------
Create simple cmd aliases that are avaiable without restarting cmd.
An alias my have multiple calls, each call is given the arguments of the alias (%1-%9)
If a single call has spaces the call must be wrapped in quotes

ex:
alias gob "go build"
alias mcd md cd`

	errorMsg = "Not enough arguments"
)

func ArgInit(conf clu.ArgConf) {
	conf.AddFlag("help", "?", "displays help")
}

func main() {
	args := clu.Parse(ArgInit)

	// TODO: Something with help here
	if args.Flag("help") {
		fmt.Println(helpMsg)
		return
	}

	loosies := args.Loosies()
	if len(loosies) < 2 {
		fmt.Println(errorMsg)
		fmt.Println(helpMsg)
		return
	}

	alias := loosies[0]
	cmd_str := ""
	for _, e := range loosies[1:] {
		cmd_str += e + " %1 %2 %3 %4 %5 %6 %7 %8 %9\n"
	}

	alias_dir, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Printf("Couldn't get run directory: %s", err.Error())
	}

	alias_file := alias_dir + alias + ".cmd"
	fmt.Println(alias_file)

	f, err := os.Create(alias_file)
	if err != nil {
		fmt.Printf("Error creating file [%s]: %s", alias_file, err.Error())
	}

	alias_contents := fmt.Sprintf("@echo off\n%s", cmd_str)

	_, err = f.WriteString(alias_contents)
	if err != nil {
		fmt.Printf("Failed to write to file [%s]: ", alias_file, err.Error())
	}
}
