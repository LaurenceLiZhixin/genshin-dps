package damage

import (
	"fmt"
	"github.com/alibaba/ioc-golang/autowire/normal"
	"github.com/laurencelizhixin/genshin-dps-simulator/internal/constant"
)

// +ioc:autowire=true
// +ioc:autowire:paramType=ContextParam
// +ioc:autowire:type=normal
// +ioc:autowire:constructFunc=Init

type DamageContext struct {
	ContextParam
	// extra items
	ExtraAttachRate float32

	// env
	EnvCtx *EnvContext

	// trailer damages
	TrailerDamages []*DamageContext
}

func (d *DamageContext) BeforeDamageSnapshot() {
	fmt.Printf("\n--------\nTime = %f, Charactor = %s, Skill = %d\n", d.EnvCtx.AbsoluteTime, d.Character.Name, d.SkillType)
	fmt.Printf("\nBefore: %+v\nEnv: %+v\n", d, d.EnvCtx)
}

func (d *DamageContext) AfterDamageSnapshot(lowDamage, maxDamage int, elementReactinoResult *ElementReactionResult) {
	fmt.Printf("\nAfter: %+v\nEnv: %+v\n", d, d.EnvCtx)
	fmt.Printf("\nTotalAttack = %d\n", int(d.getTotalAttack()))
	fmt.Printf("ElementAdditionalRate = %f\n", d.EMDamage)
	fmt.Printf("TalentRate = %f\n", d.TalentSkillRate)
	fmt.Printf("ElementReactionResult = %s\n", elementReactinoResult.String())
	fmt.Printf("BackgroundLockDashboard = %t\n", d.BackgroundSkillLockDashboard)
	fmt.Printf("Damage = %d, Max = %d\n", lowDamage, maxDamage)
}

// MaxDamage is （基础攻击力×百分比额外攻击力+数值额外攻击力）× （暴击伤害  + 1）×（元素（物理）伤害加成+1）×技能倍率×（增幅类元素反应倍率+1）×（造成伤害加成+1）
func (d *DamageContext) MaxDamage() (int, *ElementReactionResult) {
	return d.getDamageWithTargetCRITRate(1)
}

// AvgDamage is （基础攻击力×百分比额外攻击力+数值额外攻击力）× （暴击伤害 × 暴击率 + 1）×（元素（物理）伤害加成+1）×技能倍率×（增幅类元素反应倍率+1）×（造成伤害加成+1）
func (d *DamageContext) AvgDamage() (int, *ElementReactionResult) {
	return d.getDamageWithTargetCRITRate(d.CRITRate)
}

// LowDamage is （基础攻击力×百分比额外攻击力+数值额外攻击力）×（元素（物理）伤害加成+1）×技能倍率×（增幅类元素反应倍率+1）×（造成伤害加成+1）
func (d *DamageContext) LowDamage() (int, *ElementReactionResult) {
	return d.getDamageWithTargetCRITRate(0)
}

func (d *DamageContext) getTotalAttack() float32 {
	return float32(d.BaseAttack)*(1+d.ExtraAttachRate) + float32(d.ExtraAttack)
}

func (d *DamageContext) getElementAndTalentRate() float32 {
	return (d.EMDamage + 1) * d.TalentSkillRate
}

func (d *DamageContext) getDamageWithTargetCRITRate(critRate float32) (int, *ElementReactionResult) {
	// totalAttack is（基础攻击力×百分比额外攻击力+数值额外攻击力）
	totalAttack := d.getTotalAttack()

	// avgCRITDamageRate is（暴击伤害 × 暴击率 + 1)
	avgCRITDamageRate := d.CRITDamage*critRate + 1

	//（元素（物理）伤害加成+1））×技能倍率
	elementAndTalent := d.getElementAndTalentRate()

	// 等级减伤系数
	levelDamageDescRate := d.getLevelDescDamageRate()

	// 元素减伤系数:
	emDamageDefValue := d.Enemy.AllElementDef.GetElementDef(d.ElementType, d.EnvCtx.AbsoluteTime)
	emDamageDefRateParam := getDefRateParam(emDamageDefValue)

	// 主元素反应 from d.EnvCtx, todo: 三元素反应
	// 注意，元素反应部分会操作envctx，
	// GetElementReactionResult 为静态方法
	elementReactionResult := GetElementReactionResult(d.ElementType, d.IsStrongElement, d.EnvCtx.AdhesionElements, d.EnvCtx.ElementAmount, d.EnvCtx.IsAdhesionElementStrong, d.EnvCtx.ElementUpdateAbsoluteTime, d.EnvCtx.AbsoluteTime, d.Character.EM) // https://bbs.mihoyo.com/ys/article/11183533=
	if d.EnvCtx.AdhesionElements != d.ElementType {
		// 触发了元素反应，计算反应时共存元素
		if elementReactionResult.LeftElement != d.EnvCtx.AdhesionElements {
			// 共存元素为原有附着元素
			d.EnvCtx.CoexistenceElement = d.EnvCtx.AdhesionElements
		} else {
			// 共存元素为伤害元素
			d.EnvCtx.CoexistenceElement = d.ElementType
		}
		d.EnvCtx.CoexistenceElementUpdateTime = d.EnvCtx.AbsoluteTime
	} else {
		// 没有触发元素反应，尝试与共存元素反应，如共存元素超时，则删掉
		if d.EnvCtx.CoexistenceElement != constant.Physics {
			if d.EnvCtx.AbsoluteTime-d.EnvCtx.CoexistenceElementUpdateTime > 0.2 {
				// 共存元素超时
				d.EnvCtx.CoexistenceElement = constant.Physics
			} else {
				// 共存元素未超时，尝试与其反应, 共存元素量非常小，设为0.01
				elementReactionResult = GetElementReactionResult(d.ElementType, d.IsStrongElement, d.EnvCtx.CoexistenceElement, 0.01, false, d.EnvCtx.ElementUpdateAbsoluteTime, d.EnvCtx.AbsoluteTime, d.Character.EM) // https://bbs.mihoyo.com/ys/article/11183533=
			}
		}
	}

	// （基础攻击力×百分比额外攻击力+数值额外攻击力）× （暴击伤害 × 暴击率 + 1）×（元素（物理）伤害加成+1）×技能倍率
	return int(totalAttack * avgCRITDamageRate * elementAndTalent * levelDamageDescRate * emDamageDefRateParam * elementReactionResult.DamageRate), elementReactionResult
}

// getDefRateParam 抗性转抗性系数 https://baijiahao.baidu.com/s?id=1719648085505022227&wfr=spider&for=pc
func getDefRateParam(emDamageDefValue float32) float32 {
	if emDamageDefValue < 0 {
		return 1 - emDamageDefValue/2
	}
	if emDamageDefValue < 0.75 {
		return 1 - emDamageDefValue
	}
	return 1 / (1 + 4*emDamageDefValue)
}

func (d *DamageContext) getLevelDescDamageRate() float32 {
	charactorParam := d.Character.Level*5 + 500
	enemyParam := d.Enemy.Level*5 + 500
	return float32(charactorParam) / float32(charactorParam+enemyParam)
}

func (d *DamageContext) SetEnvCtx(envCtx *EnvContext) {
	d.EnvCtx = envCtx
}

func (d *DamageContext) AddTrailerDamage(trailerDamageContext *DamageContext) {
	if d.TrailerDamages == nil {
		d.TrailerDamages = make([]*DamageContext, 0)
	}
	d.TrailerDamages = append(d.TrailerDamages, trailerDamageContext)
}

func NewDamageCtx(param *ContextParam) *DamageContext {
	impl, _ := normal.GetImpl("DamageContext-DamageContext", param)
	return impl.(*DamageContext)
}

func (d *DamageContext) CopyTo() *DamageContext {
	copiedTrailerDamages := make([]*DamageContext, 0)
	for _, v := range d.TrailerDamages {
		copiedTrailerDamages = append(copiedTrailerDamages, v.CopyTo())
	}
	return &DamageContext{
		ContextParam:    d.ContextParam.CopyTo(),
		ExtraAttachRate: d.ExtraAttachRate,
		EnvCtx:          d.EnvCtx,
		TrailerDamages:  copiedTrailerDamages,
	}
}

// +ioc:autowire=true
// +ioc:autowire:type=normal
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=Init

type EnvContext struct {
	Enemy *Enemy

	avgTPS   int
	totalDmg int

	DeployedDamageFilters []*DeployedDamageFilter

	AdhesionElements          constant.ElementType
	ElementAmount             float32
	ElementUpdateAbsoluteTime float32
	IsAdhesionElementStrong   bool

	CoexistenceElement           constant.ElementType
	CoexistenceElementUpdateTime float32

	AbsoluteTime float32 // 从模拟开始到现在的绝对时间

	// 脱手伤害
	BackgroundDamageCtxs   []*DamageContext
	BackgroundSkillContext *DamageContext

	// 共鸣
	Resonance []DamageFilter
}

func (e *EnvContext) AddBackgroundDamageCtx(ctx *DamageContext) {
	if e.BackgroundDamageCtxs == nil {
		e.BackgroundDamageCtxs = make([]*DamageContext, 0)
	}
	e.BackgroundDamageCtxs = append(e.BackgroundDamageCtxs, ctx)
}

func (e *EnvContext) Damage(dmgCtx *DamageContext) int {
	// 写入环境信息，伤害上下文信息快照
	dmgCtx.SetEnvCtx(e)
	dmgCtx.BeforeDamageSnapshot()

	// 1. run character artifact filter
	for _, f := range dmgCtx.Character.ArtifactFilters {
		f.TouchOff(dmgCtx)
	}

	// 2. add new damage filter if necessary
	onceFilters := make([]DamageFilter, 0)
	for _, f := range dmgCtx.Skill.DmgFilters {
		if f.Duration() < 0 {
			onceFilters = append(onceFilters, f)
		} else {
			e.DeployedDamageFilters = append(e.DeployedDamageFilters, &DeployedDamageFilter{
				DamageFilter: f,
				InvalidTime:  e.AbsoluteTime + f.Duration(),
			})
		}
	}

	// 3. run environment damage increase filter
	// 3.1 run once filter: e.g. xingqiu QSkill
	for _, f := range onceFilters {
		f.TouchOff(dmgCtx)
	}

	// 3.2 run duration filters, remove invalid filters
	validFilters := make([]*DeployedDamageFilter, 0)
	for _, f := range e.DeployedDamageFilters {
		if f.InvalidTime > e.AbsoluteTime {
			// delete invalid damage filters
			validFilters = append(validFilters, f)
		}
	}
	e.DeployedDamageFilters = validFilters
	for _, f := range e.DeployedDamageFilters {
		f.TouchOff(dmgCtx)
	}

	// 3.3 run resonance filter
	for _, f := range e.Resonance {
		f.TouchOff(dmgCtx)
	}

	// 4. calculate damage
	lowDmg, elementReactinoResult := dmgCtx.LowDamage()
	//avgDmt := dmgCtx.AvgDamage()

	// 5. add element adhesion
	if dmgCtx.ElementType != constant.Physics {
		e.AdhesionElements = elementReactinoResult.LeftElement
		e.ElementAmount = elementReactinoResult.LeftElementAmount
		e.ElementUpdateAbsoluteTime = e.AbsoluteTime
		e.IsAdhesionElementStrong = elementReactinoResult.IsLeftElementAmountStrong
	}

	// 6. 修改环境绝对时间
	if dmgCtx.Skill.AsBackgroundTouchOffTimes == 0 { // 非脱手技能才可以修改绝对时间
		e.AbsoluteTime += dmgCtx.Skill.CostTime
	}

	// 7. 总结伤害上下文和伤害数值
	// 锁面板脱手技能按照第一次伤害
	// fixme: 目前锁面板机制有问题，不能按照伤害锁，要按照filter来锁，例如锁面板后，元素增幅反应还需要正常计算
	if dmgCtx.BackgroundSkillLockDashboard {
		if dmgCtx.BackgroundSkillLockDashboardDamage != 0 {
			lowDmg = dmgCtx.BackgroundSkillLockDashboardDamage
		} else {
			dmgCtx.BackgroundSkillLockDashboardDamage = lowDmg
		}
	}
	// 暴击伤害
	maxDmt := int(float32(lowDmg) * (dmgCtx.Character.CRITDamage + 1))
	dmgCtx.AfterDamageSnapshot(lowDmg, maxDmt, elementReactinoResult)
	return lowDmg
}

type DeployedDamageFilter struct {
	DamageFilter
	InvalidTime float32 // 失效时间
}

type Param struct {
	Enemy Enemy
}

func (p *Param) Init(e *EnvContext) (*EnvContext, error) {
	e.Enemy = &p.Enemy
	e.DeployedDamageFilters = make([]*DeployedDamageFilter, 0)
	e.AbsoluteTime = 0
	return e, nil
}
