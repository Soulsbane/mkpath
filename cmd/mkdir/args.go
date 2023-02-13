package main

type ProgramArgs struct {
	Path string `arg:"positional" help:"The path to the directory to create"`
}
