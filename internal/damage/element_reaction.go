package damage

import (
	"fmt"

	"github.com/laurencelizhixin/genshin-dps/internal/constant"
)

const StrongDescRate = 1.6 / float32(12)
const WeakDescRate = 0.8 / float32(9.5)
const ElementAdhesionRate = 0.8

type ElementReactionResult struct {
	DamageRate                float32
	LeftElement               constant.ElementType
	LeftElementAmount         float32
	IsLeftElementAmountStrong bool
}

func (e *ElementReactionResult) String() string {
	return fmt.Sprintf("%+v", *e)
}

// GetElementReactionResult 是一个静态方法，计算元素反应结果
func GetElementReactionResult(damageElement constant.ElementType, isDamageElementStrong bool,
	adhesionElement constant.ElementType, amount float32, isAdhesionElementStrong bool, updateTime, nowTime float32, em int) *ElementReactionResult {
	duration := nowTime - updateTime
	descRate := StrongDescRate
	if !isAdhesionElementStrong {
		descRate = WeakDescRate
	}
	// 计算元素随时间自然消耗
	adhesionElementNowAmount := amount - descRate*duration
	// 物理伤害，不参与元素反应
	if damageElement == constant.Physics {
		if adhesionElementNowAmount > 0 {
			// 保留原始附着元素
			return &ElementReactionResult{
				DamageRate:                1,
				LeftElement:               adhesionElement,
				LeftElementAmount:         adhesionElementNowAmount,
				IsLeftElementAmountStrong: isAdhesionElementStrong,
			}
		} else {
			// 无元素
			return &ElementReactionResult{
				DamageRate:                1,
				LeftElement:               constant.Physics,
				LeftElementAmount:         0,
				IsLeftElementAmountStrong: false,
			}
		}
	}
	// 造成伤害元素总量
	damageElementAmount := float32(2)
	if !isDamageElementStrong {
		damageElementAmount = float32(1)
	}
	// 已有元素已经消耗完，或者相同元素，按照造成伤害元素附着/刷新
	if adhesionElementNowAmount <= 0 || adhesionElement == damageElement {
		return &ElementReactionResult{
			DamageRate:                1,
			LeftElement:               damageElement,
			LeftElementAmount:         damageElementAmount * ElementAdhesionRate,
			IsLeftElementAmountStrong: isDamageElementStrong,
		}
	}

	// 计算元素反应消耗
	beConsumedRate := getElementReactionConsumingRate(damageElement, adhesionElement)
	beConsumedAmount := damageElementAmount * beConsumedRate

	// 计算元素反应结果
	leftElementAmount := adhesionElementNowAmount - beConsumedAmount
	leftElement := adhesionElement
	isLeftElementStrong := isAdhesionElementStrong
	if beConsumedAmount > adhesionElementNowAmount {
		// all adhesion element is consumed
		leftElementAmount = damageElementAmount * ElementAdhesionRate
		leftElement = damageElement
		//reactionAmountOfDamage = adhesionElementNowAmount / beConsumedAmount
		isLeftElementStrong = isDamageElementStrong
	}
	return &ElementReactionResult{
		DamageRate:                getBasicDamageRate(damageElement, adhesionElement, em),
		LeftElement:               leftElement,
		LeftElementAmount:         leftElementAmount,
		IsLeftElementAmountStrong: isLeftElementStrong,
	}
}

func getElementReactionConsumingRate(element, beConsumedElement constant.ElementType) float32 {
	if element == constant.Water && beConsumedElement == constant.Fire {
		return 2
	}

	if element == constant.Fire && beConsumedElement == constant.Water {
		return 0.5
	}

	if element == constant.Ice && beConsumedElement == constant.Fire {
		return 0.5
	}

	if element == constant.Fire && beConsumedElement == constant.Ice {
		return 2
	}
	return 1
}

// todo 剧变反应支持
func getBasicDamageRate(damageElement constant.ElementType, adhesionElement constant.ElementType, em int) float32 {
	if damageElement == constant.Water && adhesionElement == constant.Fire {
		// 蒸发
		return 2 * (1 + getZengFuReactionEMAdditionalRate(em))
	}

	if damageElement == constant.Fire && adhesionElement == constant.Water {
		// 蒸发
		return 1.5 * (1 + getZengFuReactionEMAdditionalRate(em))
	}

	if damageElement == constant.Ice && adhesionElement == constant.Fire {
		// 融化
		return 1.5 * (1 + getZengFuReactionEMAdditionalRate(em))
	}

	if damageElement == constant.Fire && adhesionElement == constant.Ice {
		// 融化
		return 2 * (1 + getZengFuReactionEMAdditionalRate(em))
	}
	return 1
}

func getZengFuReactionEMAdditionalRate(em int) float32 {
	// refers to http://www.8fe.com/gonglue/3574.html
	return 2.78 * float32(em) / (float32(em) + 1400)
}

// todo support jubian reaction
//func getJuBianReactionAdditionalEMRate(em int) float32 {
//	return 6.67 * float32(em) / (float32(em) + 1400)
//}
