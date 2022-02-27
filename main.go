package main

import "com.koutsios/marvelrandomiserplus/cmd"

const (
	dataFileName = "data.json"
)

func main() {
	cli := cmd.NewCli(dataFileName)
	cli.Start()
}
