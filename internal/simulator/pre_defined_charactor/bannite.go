package pre_defined_charactor

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/damage"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/simulator/pre_defined_artifacts"
)

func GetBanNiTe() *damage.Character {
	return damage.NewCharacter(&damage.CharacterParam{
		Level:                   80,
		Name:                    "bannite",
		BaseAttack:              570,
		ExtraAttack:             516,
		EM:                      0,
		CRITRate:                0.081,
		CRITDamage:              0.5,
		AdditionalElementDamage: 0,
		AdditionalDamageElement: constant.Fire,
		EnergyIncreaseRate:      2.082,

		ArtifactFilters: []damage.DamageFilter{
			pre_defined_artifacts.GetZongshi2ArtifactFilter(),
			pre_defined_artifacts.GetZongshi4ArtifactFilter(),
		},

		ASkill: *damage.NewSkill(&damage.SkillParam{}),
		ESkill: *damage.NewSkill(&damage.SkillParam{}),
		QSkill: *damage.NewSkill(&damage.SkillParam{
			CD:              15,
			CostTime:        1.5,
			Rates:           []float32{3.72},
			Element:         constant.Fire,
			IsStrongElement: true, // 元素附着强弱 https://wiki.biligame.com/ys/%E5%85%83%E7%B4%A0%E9%99%84%E7%9D%80%E6%97%B6%E9%97%B4
			Type:            constant.QSkill,
			DmgFilters: []damage.DamageFilter{
				&BaNiTeQSkillFilter{},
			},
		}),
	})
}

type BaNiTeQSkillFilter struct {
}

func (b *BaNiTeQSkillFilter) Duration() float32 {
	return 12
}

func (b *BaNiTeQSkillFilter) TouchOff(dmgCtx *damage.DamageContext) {
	dmgCtx.ExtraAttack += 1.1 * 570
}
