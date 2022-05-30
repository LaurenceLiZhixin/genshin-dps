package simulator

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/appearance"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/damage"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/simulator/pre_defined_charactor"
)

type DamageSimulator struct {
}

func (d *DamageSimulator) Run() {
	// 1. load all skill and charactor
	allCharactor := []*damage.Character{
		pre_defined_charactor.GetBanNiTe(),
		pre_defined_charactor.GetXingQiu(),
		pre_defined_charactor.GetXiangLing(),
	}

	// 2. load all actions
	//allActions := []constant.Action{
	//	constant.QAction, constant.SAction,
	//	constant.EAction, constant.QAction, constant.SAction,
	//	constant.EAction,
	//}

	// 3. create appearances
	allAppearances := []*appearance.Appearance{
		appearance.NewAppearance(&appearance.Param{
			Character: *allCharactor[0],
			Actions: []constant.Action{
				constant.QAction,
			},
		}),
		appearance.NewAppearance(&appearance.Param{
			Character: *allCharactor[1],
			Actions: []constant.Action{
				constant.EAction, constant.QAction, 2000,
			},
		}),
		appearance.NewAppearance(&appearance.Param{
			Character: *allCharactor[2],
			Actions: []constant.Action{
				constant.QAction, constant.AAction, constant.AAction, constant.AAction, constant.AAction, constant.AAction, constant.AAction, constant.AAction,
			},
		}),
	}

	// 4. load enemy
	enemy := &damage.Enemy{
		AllElementDef: damage.NewAllElementDef(nil), //https://lewan.baidu.com/lewanqapage?gameId=0&questionId=193144294251548672&idfrom=5015
		Level:         89,
	}

	// 5. create environment context
	envContext := &damage.EnvContext{
		Enemy:     enemy,
		Resonance: []damage.DamageFilter{},
	}

	// 6. each appearance run
	for _, appear := range allAppearances {
		appear.Run(envContext)
	}
}
