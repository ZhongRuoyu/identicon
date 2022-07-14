package internal

import (
	"image/color"
	"math"
)

type HSL struct {
	Hue float32
	Sat float32
	Lum float32
}

func (hsl *HSL) RGBA() *color.RGBA {
	hue := hsl.Hue / 360.0
	sat := hsl.Sat / 100.0
	lum := hsl.Lum / 100.0

	var b float32
	if lum <= 0.5 {
		b = lum * (sat + 1.0)
	} else {
		b = lum + sat - lum*sat
	}
	a := lum*2.0 - b

	red := HueToRGB(a, b, hue+1.0/3.0)
	green := HueToRGB(a, b, hue)
	blue := HueToRGB(a, b, hue-1.0/3.0)

	return &color.RGBA{
		R: uint8(math.Round(float64(red * 255.0))),
		G: uint8(math.Round(float64(green * 255.0))),
		B: uint8(math.Round(float64(blue * 255.0))),
		A: 255,
	}
}

func HueToRGB(a float32, b float32, hue float32) float32 {
	h := hue
	if h < 0.0 {
		h += 1.0
	} else if h > 1.0 {
		h -= 1.0
	}

	if h < 1.0/6.0 {
		return a + (b-a)*6.0*h
	}
	if h < 1.0/2.0 {
		return b
	}
	if h < 2.0/3.0 {
		return a + (b-a)*(2.0/3.0-h)*6.0
	}
	return a
}
