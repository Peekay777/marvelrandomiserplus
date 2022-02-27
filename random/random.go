package random

import (
	"errors"
	"math/rand"

	"com.koutsios/marvelrandomiserplus/data"
	"com.koutsios/marvelrandomiserplus/utils"
)

type Randoms interface {
	RandomScheme() (*data.Scheme, error)
	RemainingModulars(scheme *data.Scheme) []data.Modular
	RandomSets(scheme *data.Scheme, remainingModulars []data.Modular) *[]string
}

type RandomSetup struct {
	data *data.Data
}

func NewRandomSetup(data *data.Data, seed int64) *RandomSetup {
	rand.Seed(seed)
	return &RandomSetup{data: data}
}

func (ms *RandomSetup) RandomScheme() (*data.Scheme, error) {
	i := len(ms.data.Schemes)
	if i <= 0 {
		return nil, errors.New("no schemes in data")
	}
	schemeIdx := rand.Intn(i)
	return &ms.data.Schemes[schemeIdx], nil
}

func (ms *RandomSetup) RemainingModulars(scheme *data.Scheme) []data.Modular {
	var remainingModulars []data.Modular
	for _, modular := range ms.data.Modulars {
		foundRequired := utils.IsInSlice(scheme.Required, modular.ModularName)
		foundRecommended := utils.IsInSlice(scheme.Recommended, modular.ModularName)
		if foundRequired || foundRecommended {
			continue
		}
		remainingModulars = append(remainingModulars, modular)
	}
	return remainingModulars
}

func (ms *RandomSetup) RandomSets(scheme *data.Scheme, remainingModulars []data.Modular) *[]string {
	var randomSets []string
	requiredNumberOfSets := len(scheme.Recommended)
	for i := 0; i < requiredNumberOfSets; i++ {
		rndModularIdx := rand.Intn(len(remainingModulars))
		randomSets = append(randomSets, remainingModulars[rndModularIdx].ModularName)
		remainingModulars = append(remainingModulars[:rndModularIdx], remainingModulars[rndModularIdx+1:]...)
	}
	return &randomSets
}
