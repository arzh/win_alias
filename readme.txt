Win_Alias:
A simple alias tool that I made so I didn't have to write a cmd everytime I wanted a quick alias

'go build' with create the win_alias.exe I usually rename and copy this to a place I have all my other cmd helpers

alias [name] [command]

alias creates a [name].cmd file in the same location as the alias.exe

if a command has arguments wrap the command and the arguments in double quotes
so 'alias gob "go build"' will create a file containing
	@echo off
	go build %1 %2 %3 %4 %5 %6 %7 %8 %9

