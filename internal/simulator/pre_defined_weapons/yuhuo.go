package pre_defined_weapons

import (
	"github.com/laurencelizhixin/genshin-dps/internal/constant"
	"github.com/laurencelizhixin/genshin-dps/internal/damage"
)

func GetYuhuoFilter() damage.DamageFilter {
	return &YuhuoArtifactFilter{}
}

type YuhuoArtifactFilter struct {
}

func (z *YuhuoArtifactFilter) TouchOff(dmgCtx *damage.DamageContext) {
	if dmgCtx.SkillType == constant.QSkill {
		dmgCtx.CRITRate += 0.06
		dmgCtx.EMDamage += 0.16
	}
}

func (z *YuhuoArtifactFilter) Duration() float32 {
	return -1
}
