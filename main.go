package main

import (
	"fmt"
	"github.com/arzh/clu"
	"os"
	"osext"
	"text/template"
)

const (
	helpMsg = `alias - Windows command line aliasing tool
----------------------------------
Create simple cmd aliases that are avaiable without restarting cmd.
An alias is given the arguments from the cmd (%1-%9)
If a alias has arguments they must be wrapped in quotes

ex:
alias gob "go build"
alias rd rmdir`

	errorMsg = "Not enough arguments"

	cmdTemplate = `@echo off
if "%1"=="/?" (
	goto :HELP
)

:STANDARD_CMD
{{.Command}} %1 %2 %3 %4 %5 %6 %7 %8 %9
goto :END_SCRIPT

:HELP
echo. 
echo 	{{.File}}
echo -----------------------------------------------------
echo	Alias for {{.Command}}
echo.

:END_SCRIPT`
)

type Alias struct {
	File    string
	Command string
}

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
		cmd_str += e + "\n"
	}

	var a Alias
	a.File = alias
	a.Command = cmd_str

	alias_dir, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Println("Couldn't get run directory: ", err)
	}

	alias_file := alias_dir + alias + ".cmd"
	//fmt.Println(alias_file)

	f, err := os.Create(alias_file)
	if err != nil {
		fmt.Printf("Error creating file [%s]: %s\n", alias_file, err)
	}

	t, err := template.New("cmd").Parse(cmdTemplate)
	if err != nil {
		fmt.Println("Error creating template: ", err)
	}

	t.Execute(f, a)
}
