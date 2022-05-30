package damage

import (
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
)

type ContextParam struct {
	CharacterLevel  int
	BaseAttack      int
	ExtraAttack     int
	CRITRate        float32
	CRITDamage      float32
	EMDamage        float32
	TalentSkillRate float32

	// SkillType
	SkillType constant.SkillType

	// Character
	Character *Character

	ElementType     constant.ElementType
	IsStrongElement bool

	// Enemy
	Enemy *Enemy

	// Skill
	Skill *Skill

	// Background Skill
	BackgroundSkillLastUpdateTime      float32
	BackgroundSkillLastTouchOffTime    int
	BackgroundSkillLockDashboard       bool
	BackgroundSkillLockDashboardDamage int
}

func (c *ContextParam) CopyTo() ContextParam {
	return ContextParam{
		CharacterLevel:  c.CharacterLevel,
		BaseAttack:      c.BaseAttack,
		ExtraAttack:     c.ExtraAttack,
		CRITRate:        c.CRITRate,
		CRITDamage:      c.CRITDamage,
		EMDamage:        c.EMDamage,
		TalentSkillRate: c.TalentSkillRate,

		SkillType: c.SkillType,

		Character: c.Character,

		ElementType:     c.ElementType,
		IsStrongElement: c.IsStrongElement,

		Enemy: c.Enemy,

		Skill: c.Skill,

		BackgroundSkillLastUpdateTime:      c.BackgroundSkillLastUpdateTime,
		BackgroundSkillLastTouchOffTime:    c.BackgroundSkillLastTouchOffTime,
		BackgroundSkillLockDashboard:       c.BackgroundSkillLockDashboard,
		BackgroundSkillLockDashboardDamage: c.BackgroundSkillLockDashboardDamage,
	}
}

func (c *ContextParam) Init(impl *DamageContext) (*DamageContext, error) {
	impl.ContextParam = *c
	return impl, nil
}
