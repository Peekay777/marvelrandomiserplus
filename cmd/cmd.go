package cmd

import (
	"time"

	"com.koutsios/marvelrandomiserplus/data"
	"com.koutsios/marvelrandomiserplus/random"
)

type Cli struct {
	filename string
}

func NewCli(filename string) *Cli {
	return &Cli{filename: filename}
}

func (cli *Cli) Start() {
	// read data
	data, err := data.LoadData(cli.filename)
	failOnError(err, int(ERROR_LOADING_FILE))

	// select random data
	r := random.NewRandomSetup(data, time.Now().UnixNano())
	scheme, err := r.RandomScheme()
	failOnError(err, 2)
	modulars := r.RemainingModulars(scheme)
	sets := r.RandomSets(scheme, modulars)

	// display results
	ShowResults(scheme, *sets)
}

func failOnError(err error, code int) {
	if err != nil {
		ShowFatalError(err, code)
	}
}
