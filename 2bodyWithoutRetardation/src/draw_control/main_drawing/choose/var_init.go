package choose

import sh "2bodyBinary/draw_control/shapes"

var (
	IsWarning = false
	shapes []sh.Shape
	continue0 = false
	numFrom *sh.TextField
	numTo *sh.TextField
	selectMode *bool
	selected *map[int]bool
	length int
	isInterval = false
)

const (
	notRead = "I can not read %v number"
)