package riot

import "os"

var (
	Region       = os.Getenv("REGION")
	GlobalRegion = os.Getenv("GLOBAL_REGION")
)
