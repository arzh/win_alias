package main

import (
	"fmt"
	"github.com/arzh/clu"
	"os"
	"osext"
)

func ArgInit(conf clu.ArgConf) {
	conf.AddFlag("help", "?", "displays help")
}

func main() {
	args := clu.Parse(ArgInit)

	// TODO: Something with help here

	loosies := args.Loosies()
	if len(loosies) != 2 {
		// TODO: Display help message here as wel
		return
	}

	alias := loosies[0]
	cmd_str := loosies[1]

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
