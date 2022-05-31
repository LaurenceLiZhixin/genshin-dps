package appearance

import (
	"github.com/alibaba/ioc-golang/autowire/normal"

	"github.com/laurencelizhixin/genshin-dps/internal/constant"
	"github.com/laurencelizhixin/genshin-dps/internal/damage"
)

// +ioc:autowire=true
// +ioc:autowire:type=normal
// +ioc:autowire:constructFunc=Init
// +ioc:autowire:paramType=Param

type Appearance struct {
	Param
}

func (a *Appearance) Run(envCtx *damage.EnvContext) {
	for _, action := range a.Actions {
		a.Character.TouchOff(envCtx, action)
	}
}

type Param struct {
	Character damage.Character
	Actions   []constant.Action
}

func (p *Param) Init(a *Appearance) (*Appearance, error) {
	a.Param = *p
	return a, nil
}

func NewAppearance(param *Param) *Appearance {
	impl, _ := normal.GetImpl("Appearance-Appearance", param)
	return impl.(*Appearance)
}
