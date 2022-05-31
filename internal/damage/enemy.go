package damage

import "github.com/laurencelizhixin/genshin-dps/internal/constant"

type Enemy struct {
	// 元素抗性
	AllElementDef *AllElementDef

	Level int
}

type AllElementDef struct {
	allElements   map[constant.ElementType]float32
	defFiltersMap map[string]ElementDefFilter
}

func NewAllElementDef(def map[constant.ElementType]float32) *AllElementDef {
	defaultDef := &AllElementDef{
		allElements: map[constant.ElementType]float32{
			constant.Physics:  0.1,
			constant.Ice:      0.1,
			constant.Fire:     0.1,
			constant.Water:    0.1,
			constant.Wind:     0.1,
			constant.Stone:    0.1,
			constant.Electric: 0.1,
			constant.Grass:    0.1,
		},
		defFiltersMap: map[string]ElementDefFilter{},
	}
	for k, v := range def {
		defaultDef.allElements[k] = v
	}
	return defaultDef
}

func (a *AllElementDef) GetElementDef(elementType constant.ElementType, now float32) float32 {
	copyDefMap := make(map[constant.ElementType]float32)
	for k, v := range a.allElements {
		copyDefMap[k] = v
	}

	// check map validation
	invalidFilterKeys := make([]string, 0)

	for k, v := range a.defFiltersMap {
		if v.ValidUntil() < now {
			invalidFilterKeys = append(invalidFilterKeys, k)
		}
	}
	for _, invalidKey := range invalidFilterKeys {
		delete(a.defFiltersMap, invalidKey)
	}
	for _, validFilter := range a.defFiltersMap {
		validFilter.TouchOff(copyDefMap)
	}
	return copyDefMap[elementType]
}

func (a *AllElementDef) AddFilter(name string, f ElementDefFilter) {
	a.defFiltersMap[name] = f
}

type ElementDefFilter interface {
	ValidUntil() float32
	TouchOff(a map[constant.ElementType]float32)
}
