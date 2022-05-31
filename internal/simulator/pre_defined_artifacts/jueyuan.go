package pre_defined_artifacts

import (
	"github.com/laurencelizhixin/genshin-dps/internal/constant"
	"github.com/laurencelizhixin/genshin-dps/internal/damage"
)

func GetJueyuan2ArtifactFilter() damage.DamageFilter {
	return &Jueyuan2ArtifactFilter{}
}

type Jueyuan2ArtifactFilter struct {
}

func (z *Jueyuan2ArtifactFilter) TouchOff(dmgCtx *damage.DamageContext) {

}

func (z *Jueyuan2ArtifactFilter) Duration() float32 {
	return -1
}

func GetJueyuan4ArtifactFilter(energyIncreaseRate float32) damage.DamageFilter {
	attachAddtionRate := 0.25 * energyIncreaseRate
	if attachAddtionRate > 0.75 {
		attachAddtionRate = 0.75
	}
	return &Jueyuan4ArtifactFilter{
		AttachAdditionRate: attachAddtionRate,
	}
}

type Jueyuan4ArtifactFilter struct {
	AttachAdditionRate float32
}

func (z *Jueyuan4ArtifactFilter) Duration() float32 {
	return -1
}

func (z *Jueyuan4ArtifactFilter) TouchOff(dmgCtx *damage.DamageContext) {
	if dmgCtx.SkillType == constant.QSkill {
		dmgCtx.EMDamage += z.AttachAdditionRate
	}
}
