package ds

const (
	// GrowthFactor ...
	growthFactor float32 = float32(2.0)  // growth by 100%
	shrinkFactor float32 = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
	NullString   string  = ""
)
