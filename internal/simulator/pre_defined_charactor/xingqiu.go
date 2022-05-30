package pre_defined_charactor

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/damage"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/simulator/pre_defined_artifacts"
)

var xingqiuSingleton *damage.Character

func GetXingQiu() *damage.Character {
	if xingqiuSingleton == nil {
		xingqiuSingleton = damage.NewCharacter(&damage.CharacterParam{
			Level:                   80,
			Name:                    "xingqiu",
			BaseAttack:              579,
			ExtraAttack:             1032,
			EM:                      0,
			CRITRate:                0.431,
			CRITDamage:              0.671,
			AdditionalElementDamage: 0.548,
			AdditionalDamageElement: constant.Water,
			EnergyIncreaseRate:      2.091,
			ArtifactFilters: []damage.DamageFilter{
				pre_defined_artifacts.GetJueyuan2ArtifactFilter(),
				pre_defined_artifacts.GetJueyuan4ArtifactFilter(2.091),
			},

			ASkill: *damage.NewSkill(&damage.SkillParam{
				CD:       0,
				Rates:    []float32{1, 1, 1},
				Element:  constant.Physics,
				CostTime: 1,
				Type:     constant.ASkill,
			}),
			ESkill: *damage.NewSkill(&damage.SkillParam{
				CD:              21,
				Rates:           []float32{2.52, 2.87},
				Element:         constant.Water,
				CostTime:        1.6,
				IsStrongElement: false,
				Type:            constant.ESkill,
			}),
			QSkill: *damage.NewSkill(&damage.SkillParam{
				CD:              20,
				Rates:           []float32{0},
				CostTime:        1,
				IsStrongElement: false,
				Element:         constant.Physics,
				Type:            constant.QSkill,
				DmgFilters: []damage.DamageFilter{
					&XingQiuQSkillFilter{
						lastTouchOffTime: 0,
					},
				},
			}),
		})
	}
	return xingqiuSingleton
}

type XingQiuQSkillFilter struct {
	counter          int
	lastTouchOffTime float32
}

func (b *XingQiuQSkillFilter) Duration() float32 {
	return 18
}

// todo 元素共存？三把雨剑都能蒸发
func (b *XingQiuQSkillFilter) TouchOff(dmgCtx *damage.DamageContext) {
	if dmgCtx.SkillType != constant.ASkill || dmgCtx.EnvCtx.AbsoluteTime-b.lastTouchOffTime < 1 {
		return
	}
	b.lastTouchOffTime = dmgCtx.EnvCtx.AbsoluteTime
	b.counter++
	num := 2
	if b.counter%2 == 0 {
		num += 1
	}

	for i := 0; i < num; i++ {
		damageCtx := damage.NewDamageCtx(&damage.ContextParam{
			CharacterLevel:  80,
			BaseAttack:      579,
			ExtraAttack:     887,
			CRITRate:        0.431,
			CRITDamage:      0.671,
			EMDamage:        0.548,
			TalentSkillRate: 0.977,
			Skill: damage.NewSkill(&damage.SkillParam{
				DmgFilters: []damage.DamageFilter{
					&XignqiuQDamageFilter{},
				},
			}),
			SkillType:       constant.QSkill,
			ElementType:     constant.Water,
			IsStrongElement: false,
			Character:       GetXingQiu(),
			Enemy:           dmgCtx.EnvCtx.Enemy,
		})
		dmgCtx.AddTrailerDamage(damageCtx)
	}
}

type XignqiuQDamageFilter struct {
}

func (x *XignqiuQDamageFilter) TouchOff(dmgCtx *damage.DamageContext) {
	dmgCtx.EnvCtx.Enemy.AllElementDef.AddFilter("XingqiuQWaterDefDescFilter", &XingqiuQWaterDefDescFilter{untilTime: dmgCtx.EnvCtx.AbsoluteTime + 4})
}

func (x *XignqiuQDamageFilter) Duration() float32 {
	return -1 // <0 means not duration filter
}

type XingqiuQWaterDefDescFilter struct {
	untilTime float32
}

func (x *XingqiuQWaterDefDescFilter) ValidUntil() float32 {
	return x.untilTime
}

func (x *XingqiuQWaterDefDescFilter) TouchOff(a map[constant.ElementType]float32) {
	a[constant.Water] -= 0.15
}
