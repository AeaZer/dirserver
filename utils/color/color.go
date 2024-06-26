package color

import "fmt"

// typeColor 颜色种类
type typeColor int

const (
	green  typeColor = 32
	blue   typeColor = 34
	red    typeColor = 31
	yellow typeColor = 33
	none   typeColor = -1 // none 无颜色

	formatStart = "\x1b[0;%dm" // formatStart 颜色开始
	formatTail  = "\x1b[0m"    // formatTail 颜色重置
)

type Color struct {
	color typeColor
}

func newColor(color typeColor) *Color {
	return &Color{color: color}
}

// Dyeing 染色
func (c *Color) Dyeing(format string, a ...any) string {
	x := fmt.Sprintf(format, a...)
	return fmt.Sprintf(formatStart+"%v"+formatTail, c.color, x)
}

var (
	GreenDA  = newColor(green)  // greenDA green Dyeing apparatus
	RedDA    = newColor(red)    // RedDA red Dyeing apparatus
	BlueDA   = newColor(blue)   // blueDA blue Dyeing apparatus
	YellowDA = newColor(yellow) // yellowDA yellow Dyeing apparatus
)
