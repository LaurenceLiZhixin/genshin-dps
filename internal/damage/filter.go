package damage

type DamageFilter interface {
	TouchOff(dmgCtx *DamageContext)
	Duration() float32
}

type EnvFilter interface {
	TouchOff(dmgCtx *EnvContext)
	Duration() float32
}
