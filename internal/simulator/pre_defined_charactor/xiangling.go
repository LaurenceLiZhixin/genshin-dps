package pre_defined_charactor

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/damage"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/simulator/pre_defined_artifacts"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/simulator/pre_defined_weapons"
)

func GetXiangLing() *damage.Character {
	return damage.NewCharacter(&damage.CharacterParam{
		Level:                   80,
		Name:                    "xiangling",
		BaseAttack:              648,
		ExtraAttack:             593,
		EM:                      168,
		CRITRate:                0.178,
		CRITDamage:              1.306,
		AdditionalElementDamage: 0.616,
		AdditionalDamageElement: constant.Fire,
		EnergyIncreaseRate:      2.089,

		ArtifactFilters: []damage.DamageFilter{
			pre_defined_artifacts.GetZongshi2ArtifactFilter(),
			pre_defined_artifacts.GetMuoNv2ArtifactFilter(),
			pre_defined_weapons.GetYuhuoFilter(),
		},

		ASkill: *damage.NewSkill(&damage.SkillParam{
			CD:       0,
			Rates:    []float32{0.5, 0.5, 0.5, 0.5, 0.5},
			Element:  constant.Physics,
			CostTime: 0.7,
			Type:     constant.ASkill,
		}),
		ESkill: *damage.NewSkill(&damage.SkillParam{}),
		QSkill: *damage.NewSkill(&damage.SkillParam{
			CD:              20,
			Rates:           []float32{1.3, 1.58, 1.97},
			CostTime:        0.3,
			IsStrongElement: false,
			Element:         constant.Fire,
			Type:            constant.QSkill,
			BackgroundSkills: []*damage.Skill{
				damage.NewSkill(&damage.SkillParam{
					Rates:                     []float32{2.02},
					Type:                      constant.QSkill,
					Element:                   constant.Fire,
					CostTime:                  1.25,
					IsStrongElement:           false,
					AsBackgroundTouchOffTimes: 8,
					LockDashboard:             true,
				}),
			},
		}),
	})
}
