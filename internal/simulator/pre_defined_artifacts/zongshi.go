package pre_defined_artifacts

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/damage"
)

func GetZongshi2ArtifactFilter() damage.DamageFilter {
	return &Zongshi2ArtifactFilter{}
}

type Zongshi2ArtifactFilter struct {
}

func (z *Zongshi2ArtifactFilter) TouchOff(dmgCtx *damage.DamageContext) {
	if dmgCtx.SkillType == constant.QSkill {
		dmgCtx.EMDamage += 0.2
	}
}

func (z *Zongshi2ArtifactFilter) Duration() float32 {
	return -1
}

func GetZongshi4ArtifactFilter() damage.DamageFilter {
	return &Zongshi4ArtifactFilter{}
}

type Zongshi4ArtifactFilter struct {
}

func (z *Zongshi4ArtifactFilter) Duration() float32 {
	return -1
}

func (z *Zongshi4ArtifactFilter) TouchOff(dmgCtx *damage.DamageContext) {
	if dmgCtx.SkillType == constant.QSkill {
		// todo check 不能叠加
		dmgCtx.EnvCtx.DeployedDamageFilters = append(dmgCtx.EnvCtx.DeployedDamageFilters, &damage.DeployedDamageFilter{
			DamageFilter: &Zongshi4DamageFilter{},
			InvalidTime:  dmgCtx.EnvCtx.AbsoluteTime + 12,
		})
	}
}

type Zongshi4DamageFilter struct {
}

func (z *Zongshi4DamageFilter) TouchOff(dmgCtx *damage.DamageContext) {
	dmgCtx.ExtraAttachRate += 0.2
}

func (z *Zongshi4DamageFilter) Duration() float32 {
	return 12
}
