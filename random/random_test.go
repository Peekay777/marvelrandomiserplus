package random

import (
	"reflect"
	"testing"

	"com.koutsios/marvelrandomiserplus/data"
)

func getEmptyData() *data.Data {
	return &data.Data{
		Schemes:  []data.Scheme{},
		Modulars: []data.Modular{},
	}
}

func getData() *data.Data {
	return &data.Data{
		Schemes: []data.Scheme{
			{
				SchemeId:    1,
				SchemeName:  "The Break-In!",
				Villain:     "Rhino",
				VillainSet:  "Rhino",
				Required:    []string{"Standard"},
				Recommended: []string{"Bomb Scare"},
			},
			{
				SchemeId:   6,
				SchemeName: "Attack on Mount Athena",
				Villain:    "Crossbones",
				VillainSet: "Crossbones",
				Required:   []string{"Experimental Weapons", "Standard"},
				Recommended: []string{
					"Hydra Assault",
					"Weapon Master",
					"Legions of Hydra",
				},
			},
		},
		Modulars: []data.Modular{
			{
				ModularId:   0,
				ModularName: "Standard",
			},
			{
				ModularId:   1,
				ModularName: "Bomb Scare",
			},
			{
				ModularId:   2,
				ModularName: "Masters of Evil",
			},
			{
				ModularId:   3,
				ModularName: "Under Attack",
			},
			{
				ModularId:   5,
				ModularName: "Experimental Weapons",
			},
			{
				ModularId:   6,
				ModularName: "Hydra Assault",
			},
			{
				ModularId:   7,
				ModularName: "Weapon Master",
			},
			{
				ModularId:   8,
				ModularName: "Legions of Hydra",
			},
		},
		Heroes: []data.Hero{
			{
				HeroId:     1,
				HeroName:   "Spider-Man",
				HeroAspect: 1,
			},
			{
				HeroId:     2,
				HeroName:   "Spider-Woman",
				HeroAspect: 2,
			},
		},
		Aspects: []string{
			"Aggression",
			"Justice",
			"Leadership",
			"Protection",
		},
	}
}

func getScheme() *data.Scheme {
	return &data.Scheme{
		SchemeId:   6,
		SchemeName: "Attack on Mount Athena",
		Villain:    "Crossbones",
		VillainSet: "Crossbones",
		Required:   []string{"Experimental Weapons", "Standard"},
		Recommended: []string{
			"Hydra Assault",
			"Weapon Master",
			"Legions of Hydra",
		},
	}
}

func getRemainingModulars() []data.Modular {
	return []data.Modular{
		{
			ModularId:   1,
			ModularName: "Bomb Scare",
		},
		{
			ModularId:   2,
			ModularName: "Masters of Evil",
		},
		{
			ModularId:   3,
			ModularName: "Under Attack",
		},
	}
}

func TestRandomSetup_RandomScheme(t *testing.T) {
	tests := []struct {
		name    string
		ms      *RandomSetup
		want    *data.Scheme
		wantErr bool
	}{
		{
			name:    "Empty data",
			ms:      NewRandomSetup(getEmptyData(), 1),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Success",
			ms:      NewRandomSetup(getData(), 1),
			want:    getScheme(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ms.RandomScheme()
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomSetup.RandomScheme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomSetup.RandomScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomSetup_RemainingModulars(t *testing.T) {
	type args struct {
		scheme *data.Scheme
	}
	tests := []struct {
		name string
		ms   *RandomSetup
		args args
		want []data.Modular
	}{
		{
			name: "Success",
			ms:   NewRandomSetup(getData(), 1),
			args: args{getScheme()},
			want: getRemainingModulars(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ms.RemainingModulars(tt.args.scheme); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomSetup.RemainingModulars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomSetup_RandomSets(t *testing.T) {
	type args struct {
		scheme            *data.Scheme
		remainingModulars []data.Modular
	}
	tests := []struct {
		name string
		ms   *RandomSetup
		args args
		want *[]string
	}{
		{
			name: "Success",
			ms:   NewRandomSetup(getData(), 1),
			args: args{
				getScheme(),
				getRemainingModulars(),
			},
			want: &[]string{
				"Under Attack",
				"Masters of Evil",
				"Bomb Scare",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ms.RandomSets(tt.args.scheme, tt.args.remainingModulars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomSetup.RandomSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomSetup_RandomHeroes(t *testing.T) {
	type args struct {
		numberOfHeroes int
	}
	tests := []struct {
		name string
		ms   *RandomSetup
		args args
		want *[]Hero
	}{
		{
			name: "Success",
			ms:   NewRandomSetup(getData(), 1),
			args: args{
				numberOfHeroes: 2,
			},
			want: &[]Hero{
				{
					"Spider-Woman",
					[]string{"Protection", "Leadership"},
				},
				{
					"Spider-Man",
					[]string{"Justice"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ms.RandomHeroes(tt.args.numberOfHeroes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomSetup.RandomHeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
