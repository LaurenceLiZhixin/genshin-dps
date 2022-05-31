package damage

import (
	"github.com/alibaba/ioc-golang/autowire/normal"

	"github.com/laurencelizhixin/genshin-dps/internal/constant"
)

// +ioc:autowire=true
// +ioc:autowire:type=normal
// +ioc:autowire:paramType=SkillParam
// +ioc:autowire:constructFunc=Init

type Skill struct {
	SkillParam
}

type SkillParam struct {
	Rates                     []float32 // 多段伤害倍率, 如为脱手技能len必须为1
	Type                      constant.SkillType
	Element                   constant.ElementType
	IsStrongElement           bool
	CD                        int
	CostTime                  float32
	DmgFilters                []DamageFilter // 脱手增伤机制，伤害不由本技能触发，而是有后续技能触发，例如行秋Q班尼特Q
	BackgroundSkills          []*Skill       // 脱手攻击机制，伤害由本技能脱手触发，而是有后续技能触发，例如香菱Q皇女E丽莎Q甘雨Q神里Q
	AsBackgroundTouchOffTimes int            // 当前技能作为脱手攻击技能，可以持续的次数
	LockDashboard             bool           // 是否锁面板
}

func (sp *SkillParam) Init(impl *Skill) (*Skill, error) {
	impl.SkillParam = *sp
	return impl, nil
}

func NewSkill(param *SkillParam) *Skill {
	impl, _ := normal.GetImpl("Skill-Skill", param)
	return impl.(*Skill)
}
