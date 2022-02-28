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
	RandomHeroes(numberOfHeroes int) *[]string
}

type RandomSetup struct {
	data *data.Data
}

type Hero struct {
	HeroName string
	Aspects  []string
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

func (ms *RandomSetup) RandomHeroes(numberOfHeroes int) *[]Hero {
	var randomHeroes []Hero
	remaindingHeroes := make([]data.Hero, len(ms.data.Heroes))
	copy(remaindingHeroes, ms.data.Heroes)
	for i := 0; i < int(numberOfHeroes); i++ {
		randomHero := Hero{}
		rndHeroIdx := rand.Intn(len(remaindingHeroes))
		hero := remaindingHeroes[rndHeroIdx]
		randomHero.HeroName = hero.HeroName

		remaindingAspects := make([]string, len(ms.data.Aspects))
		copy(remaindingAspects, ms.data.Aspects)
		for j := 0; j < hero.HeroAspect; j++ {
			rndAspectIdx := rand.Intn(len(remaindingAspects))
			randomHero.Aspects = append(randomHero.Aspects, remaindingAspects[rndAspectIdx])
			remaindingAspects = append(remaindingAspects[:rndAspectIdx], remaindingAspects[rndAspectIdx+1:]...)

		}
		randomHeroes = append(randomHeroes, randomHero)
		remaindingHeroes = append(remaindingHeroes[:rndHeroIdx], remaindingHeroes[rndHeroIdx+1:]...)
	}
	return &randomHeroes
}
