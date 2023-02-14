package main

type ProgramArgs struct {
	Path       string `arg:"required, positional" help:"The path to the directory to create"`
	Permission string `arg:"positional" default:"0755" help:"The permission to set on the directories. This should be a string of octal digits. For example, 0755."`
}
