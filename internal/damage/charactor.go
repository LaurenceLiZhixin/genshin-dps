package damage

import (
	"github.com/alibaba/ioc-golang/autowire/normal"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
)

// +ioc:autowire=true
// +ioc:autowire:type=normal
// +ioc:autowire:paramType=CharacterParam
// +ioc:autowire:constructFunc=Init

type Character struct {
	CharacterParam
}

func (c *Character) TouchOff(envCtx *EnvContext, action constant.Action) {
	if action == constant.SAction {
		// switch
		return
	}

	// skill action
	var skill *Skill
	switch action {
	case constant.AAction:
		skill = &c.ASkill
	case constant.EAction:
		skill = &c.ESkill
	case constant.QAction:
		skill = &c.QSkill
	case constant.BAction:
		envCtx.Damage(envCtx.BackgroundSkillContext)
		return
	default:
		if int(action) > 100 {
			// costime action
			costTime := float32(action) / 1000
			envCtx.AbsoluteTime += costTime
			return
		}
	}

	emDamage := float32(0)
	if skill.Element == c.AdditionalDamageElement {
		emDamage = c.AdditionalElementDamage
	}
	// 多段伤害
	for _, r := range skill.Rates {
		// pre，当前非脱手技能释放前，先执行这段时间里需要执行的脱手技能
		runBackgroundIfNeeded(c, envCtx)

		damageCtx := NewDamageCtx(&ContextParam{
			CharacterLevel:  c.Level,
			BaseAttack:      c.BaseAttack,
			ExtraAttack:     c.ExtraAttack,
			CRITRate:        c.CRITRate,
			CRITDamage:      c.CRITDamage,
			EMDamage:        emDamage,
			TalentSkillRate: r,
			SkillType:       skill.Type,
			ElementType:     skill.Element,
			IsStrongElement: skill.IsStrongElement,
			Character:       c,
			Skill:           skill,
			Enemy:           envCtx.Enemy,
		})

		envCtx.Damage(damageCtx)
		for _, trailerDamageCtx := range damageCtx.TrailerDamages {
			envCtx.Damage(trailerDamageCtx.CopyTo())
		}
	}

	// 脱手攻击filter, 写入脱手攻击面板
	for _, skill := range skill.BackgroundSkills {
		damageCtx := NewDamageCtx(&ContextParam{
			CharacterLevel:                  c.Level,
			BaseAttack:                      c.BaseAttack,
			ExtraAttack:                     c.ExtraAttack,
			CRITRate:                        c.CRITRate,
			CRITDamage:                      c.CRITDamage,
			EMDamage:                        c.AdditionalElementDamage,
			TalentSkillRate:                 skill.Rates[0],
			SkillType:                       skill.Type,
			ElementType:                     skill.Element,
			IsStrongElement:                 skill.IsStrongElement,
			Character:                       c,
			Skill:                           skill,
			Enemy:                           envCtx.Enemy,
			BackgroundSkillLastUpdateTime:   envCtx.AbsoluteTime,
			BackgroundSkillLastTouchOffTime: skill.AsBackgroundTouchOffTimes,
		})
		envCtx.AddBackgroundDamageCtx(damageCtx)
	}
}

func runBackgroundIfNeeded(c *Character, envCtx *EnvContext) {
	if len(envCtx.BackgroundDamageCtxs) > 0 {
		// todo 目前不支持多个脱手技能
		// fixme 目前多段伤害按照一个Action，因此无法在一个技能的多段伤害之间插入脱手技能伤害/元素附着
		updatedBackgroundDamageCtxs := make([]*DamageContext, 0)
		for _, v := range envCtx.BackgroundDamageCtxs {
			for v.BackgroundSkillLastUpdateTime < envCtx.AbsoluteTime && v.BackgroundSkillLastTouchOffTime > 0 {
				envCtx.BackgroundSkillContext = v.CopyTo()
				// 注意！！！这里为了调回到脱手技能触发的环境上下文，调节了环境绝对时间
				envCtxOriginAbsoluteTime := envCtx.AbsoluteTime
				envCtx.AbsoluteTime = v.BackgroundSkillLastUpdateTime
				c.TouchOff(envCtx, constant.BAction)
				envCtx.AbsoluteTime = envCtxOriginAbsoluteTime

				v.BackgroundSkillLastUpdateTime += v.Skill.CostTime
				v.BackgroundSkillLastTouchOffTime -= 1
			}
			if v.BackgroundSkillLastTouchOffTime > 0 { // all damages are not done
				updatedBackgroundDamageCtxs = append(updatedBackgroundDamageCtxs, v)
			}
		}

		envCtx.BackgroundDamageCtxs = updatedBackgroundDamageCtxs
	}
}

type CharacterParam struct {
	Level int
	Name  string
	// dashboards
	BaseAttack         int // 白值攻击力
	ExtraAttack        int // 绿值攻击力
	EM                 int
	CRITRate           float32
	CRITDamage         float32
	EnergyIncreaseRate float32
	// fixme now we only support single element damage addition
	AdditionalElementDamage float32
	AdditionalDamageElement constant.ElementType

	ArtifactFilters []DamageFilter

	ASkill Skill
	ESkill Skill
	QSkill Skill
}

func (c *CharacterParam) Init(impl *Character) (*Character, error) {
	impl.CharacterParam = *c
	return impl, nil
}

func NewCharacter(param *CharacterParam) *Character {
	impl, _ := normal.GetImpl("Character-Character", param)
	return impl.(*Character)
}
