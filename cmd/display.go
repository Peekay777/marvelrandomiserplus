package cmd

import (
	"fmt"
	"os"

	"com.koutsios/marvelrandomiserplus/data"
	"com.koutsios/marvelrandomiserplus/random"
	"com.koutsios/marvelrandomiserplus/utils"
)

type ErrCode int

const (
	ERROR_LOADING_FILE ErrCode = iota + 1
	NO_SCHEME_DATA
)

type Display interface {
	ShowResults()
	ShowFatalError(err error, code int)
}

type Response struct {
	SchemeName string
	Villain    string
	Recommend  []string
	Sets       []string
	RandomSets []string
}

func ShowResults(scheme *data.Scheme, randomSets []string, randomHeroes []random.Hero) {
	sets := []string{}
	sets = append(sets, scheme.VillainSet)
	sets = append(sets, scheme.Required...)
	fmt.Println("Villain: ", scheme.Villain)
	fmt.Println("Scheme Name: ", scheme.SchemeName)
	fmt.Println("Sets: ", utils.PrintArr(sets))
	fmt.Println("Random sets: ", utils.PrintArr(randomSets))
	for idx, hero := range randomHeroes {
		fmt.Printf("Hero_%v: %v, Aspects: %v\n", idx+1, hero.HeroName, utils.PrintArr(hero.Aspects))
	}
	fmt.Println("Press enter to finsh...")
	fmt.Scanln()
}

func ShowFatalError(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
