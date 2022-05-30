package pre_defined_artifacts

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/damage"
)

func GetMuoNv2ArtifactFilter() damage.DamageFilter {
	return &MuoNv2ArtifactFilter{}
}

type MuoNv2ArtifactFilter struct {
}

func (z *MuoNv2ArtifactFilter) TouchOff(dmgCtx *damage.DamageContext) {
	//dmgCtx.EMDamage += 0.15 already added in green value
}

func (z *MuoNv2ArtifactFilter) Duration() float32 {
	return -1
}
