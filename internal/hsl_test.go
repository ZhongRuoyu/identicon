package internal

import (
	"image/color"
	"testing"
)

func TestHSLConvertsBlack(t *testing.T) {
	black := &color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
	rgb := (&HSL{
		Hue: 0.0,
		Sat: 0.0,
		Lum: 0.0,
	}).RGBA()
	if *black != *rgb {
		t.Fail()
	}
}

func TestHSLConvertsWhite(t *testing.T) {
	black := &color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	rgb := (&HSL{
		Hue: 0.0,
		Sat: 0.0,
		Lum: 100.0,
	}).RGBA()
	if *black != *rgb {
		t.Fail()
	}
}

func TestHSLConvertsRed(t *testing.T) {
	black := &color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	rgb := (&HSL{
		Hue: 0.0,
		Sat: 100.0,
		Lum: 50.0,
	}).RGBA()
	if *black != *rgb {
		t.Fail()
	}
}

func TestHSLConvertsGreen(t *testing.T) {
	black := &color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}
	rgb := (&HSL{
		Hue: 120.0,
		Sat: 100.0,
		Lum: 50.0,
	}).RGBA()
	if *black != *rgb {
		t.Fail()
	}
}

func TestHSLConvertsBlue(t *testing.T) {
	black := &color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
	rgb := (&HSL{
		Hue: 240.0,
		Sat: 100.0,
		Lum: 50.0,
	}).RGBA()
	if *black != *rgb {
		t.Fail()
	}
}
