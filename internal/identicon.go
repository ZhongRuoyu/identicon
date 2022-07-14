package internal

import (
	"image"
	"image/color"
)

type Identicon struct {
	source []uint8
	size   int
}

func NewIdenticon(source []uint8) *Identicon {
	identicon := Identicon{
		source: source,
		size:   420,
	}
	return &identicon
}

func mapValue(value uint32, vmin uint32, vmax uint32, dmin uint32, dmax uint32) float32 {
	return float32((value-vmin)*(dmax-dmin)) / float32((vmax-vmin)+dmin)
}

func (identicon *Identicon) foreground() *color.RGBA {
	h1 := (uint16(identicon.source[12]) & 0x0f) << 8
	h2 := uint16(identicon.source[13])

	h := uint32(h1 | h2)
	s := uint32(identicon.source[14])
	l := uint32(identicon.source[15])

	hue := mapValue(h, 0, 4095, 0, 360)
	sat := mapValue(s, 0, 255, 0, 20)
	lum := mapValue(l, 0, 255, 0, 20)

	return (&HSL{hue, 65.0 - sat, 75.0 - lum}).RGBA()
}

func rect(image *image.RGBA, x0 uint32, y0 uint32, x1 uint32, y1 uint32, color *color.RGBA) {
	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			image.Set(int(x), int(y), color)
		}
	}
}

func (identicon *Identicon) pixels() []bool {
	nibbler := NewNibbler(identicon.source)
	nibbles := make([]bool, 0)
	for next := nibbler.Next(); next != nil; next = nibbler.Next() {
		nibbles = append(nibbles, *next%2 == 0)
	}
	var i int
	pixels := make([]bool, 25)

	for col := 2; col >= 0; col-- {
		for row := 0; row < 5; row++ {
			ix := col + (row * 5)
			mirrorCol := 4 - col
			mirrorIx := mirrorCol + row*5
			paint := false
			if i < len(nibbles) {
				paint = nibbles[i]
				i++
			}
			pixels[ix] = paint
			pixels[mirrorIx] = paint
		}
	}

	return pixels
}

func (identicon *Identicon) Image() image.Image {
	const pixelSize = 70
	const spriteSize = 5
	const margin = pixelSize / 2

	background := &color.RGBA{240, 240, 240, 255}
	foreground := identicon.foreground()

	image := image.NewRGBA(image.Rect(0, 0, identicon.size, identicon.size))
	for x := 0; x < identicon.size; x++ {
		for y := 0; y < identicon.size; y++ {
			image.Set(x, y, background)
		}
	}

	pixels := identicon.pixels()
	for row := 0; row < spriteSize; row++ {
		for col := 0; col < spriteSize; col++ {
			if pixels[row*spriteSize+col] {
				x := col * pixelSize
				y := row * pixelSize
				rect(
					image,
					uint32(x+margin),
					uint32(y+margin),
					uint32(x+pixelSize+margin),
					uint32(y+pixelSize+margin),
					foreground,
				)
			}
		}
	}

	return image
}
