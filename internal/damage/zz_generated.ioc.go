//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package damage

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: &Character{},
		Factory: func() interface{} {
			return &Character{}
		},
		ParamFactory: func() interface{} {
			var _ characterParamInterface = &CharacterParam{}
			return &CharacterParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(characterParamInterface)
			impl := i.(*Character)
			return param.Init(impl)
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: &DamageContext{},
		Factory: func() interface{} {
			return &DamageContext{}
		},
		ParamFactory: func() interface{} {
			var _ contextParamInterface = &ContextParam{}
			return &ContextParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(contextParamInterface)
			impl := i.(*DamageContext)
			return param.Init(impl)
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: &EnvContext{},
		Factory: func() interface{} {
			return &EnvContext{}
		},
		ParamFactory: func() interface{} {
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*EnvContext)
			return param.Init(impl)
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Interface: &Skill{},
		Factory: func() interface{} {
			return &Skill{}
		},
		ParamFactory: func() interface{} {
			var _ skillParamInterface = &SkillParam{}
			return &SkillParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(skillParamInterface)
			impl := i.(*Skill)
			return param.Init(impl)
		},
	})
}

type characterParamInterface interface {
	Init(impl *Character) (*Character, error)
}
type contextParamInterface interface {
	Init(impl *DamageContext) (*DamageContext, error)
}
type paramInterface interface {
	Init(impl *EnvContext) (*EnvContext, error)
}
type skillParamInterface interface {
	Init(impl *Skill) (*Skill, error)
}