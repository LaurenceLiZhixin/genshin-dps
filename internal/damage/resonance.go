package damage

func GetFireResonanceFilter() DamageFilter {
	return &FireResonaceFilter{}
}

type FireResonaceFilter struct {
}

func (f FireResonaceFilter) TouchOff(dmgCtx *DamageContext) {
	dmgCtx.ExtraAttachRate += 0.25
}

func (f FireResonaceFilter) Duration() float32 {
	return -1
}
