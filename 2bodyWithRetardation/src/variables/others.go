package variables

import (
	"github.com/gonutz/prototype/draw"
)

const (
	Configs           = "configs"
	SavedDir          = "./saved_config"
	EulerConfig       = "euler.json"
	EulerOriginConfig = "origin_euler.json"
	Config = "config.txt"
)

var (
	Colors  = []draw.Color{draw.Red, draw.Cyan, draw.LightBlue, draw.Yellow, draw.LightGreen, draw.White}
	Warning = ""
	Quick   = make(chan byte)
)
