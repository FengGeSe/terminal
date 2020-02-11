package util

import (
	"fmt"
)

var (
	greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	white        = string([]byte{27, 91, 51, 55, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	magenta      = string([]byte{27, 91, 51, 53, 109})
	cyan         = string([]byte{27, 91, 51, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

// 将header和tail包裹middle
func wrapSlice(header interface{}, tail interface{}, middle []interface{}) []interface{} {
	na := []interface{}{}
	na = append(na, header)
	na = append(na, middle...)
	na = append(na, tail)
	return na
}

func WrapRed(a ...interface{}) string {
	slc := wrapSlice(red, reset, a)
	return fmt.Sprint(slc...)
}
func WrapRedf(format string, a ...interface{}) string {
	str := fmt.Sprintf(format, a...)
	slc := wrapSlice(red, reset, []interface{}{str})
	return fmt.Sprint(slc...)
}

func WrapBlue(a ...interface{}) string {
	slc := wrapSlice(blue, reset, a)
	return fmt.Sprint(slc...)
}
func WrapBluef(format string, a ...interface{}) string {
	str := fmt.Sprintf(format, a...)
	slc := wrapSlice(blue, reset, []interface{}{str})
	return fmt.Sprint(slc...)
}

func WrapGreen(a ...interface{}) string {
	slc := wrapSlice(green, reset, a)
	return fmt.Sprint(slc...)
}
func WrapGreenf(format string, a ...interface{}) string {
	str := fmt.Sprintf(format, a...)
	slc := wrapSlice(green, reset, []interface{}{str})
	return fmt.Sprint(slc...)
}

func WrapYellow(a ...interface{}) string {
	slc := wrapSlice(yellow, reset, a)
	return fmt.Sprint(slc...)
}
func WrapYellowf(format string, a ...interface{}) string {
	str := fmt.Sprintf(format, a...)
	slc := wrapSlice(yellow, reset, []interface{}{str})
	return fmt.Sprint(slc...)
}
