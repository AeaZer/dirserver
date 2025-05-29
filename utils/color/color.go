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
)

// ColorStart and ColorEnd are the escape sequences for starting and ending a color.
var (
	formatStart = "\x1b[0;%dm"  // formatStart is the start of the format sequence.
	formatTail  = "\x1b[0m"     // formatTail is the end of the format sequence.
)

type Color struct {
	color typeColor
}

func newColor(color typeColor) *Color {
	return &Color{color: color}
}

// Dyeing applies a color to formatted text.
func (c *Color) Dyeing(format string, a ...any) string {
	x := fmt.Sprintf(format, a...)
	return fmt.Sprintf(formatStart+"%s"+formatTail, c.color, x)
}

var (
	GreenDA  = newColor(green)  // greenDA is the Green Dyeing apparatus.
	RedDA    = newColor(red)     // RedDA is the Red Dyeing apparatus.
	BlueDA   = newColor(blue)    // BlueDA is the Blue Dyeing apparatus.
	YellowDA = newColor(yellow)  // YellowDA is the Yellow Dyeing apparatus.
)
