package diceimg

import (
	"image/color"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type Dice struct {
	x, y        float64
	canv        *canvas.Canvas
	ctx         *canvas.Context
	strokeWidth float64
	strokeCol   color.RGBA
	valueCol    color.RGBA
	diceCol     color.RGBA
	radius      float64
}

// NewBlank creates blank canvas.
// At the end, image is rendered.
// w, h are canvas WxH (size of image)
func NewBlank(w, h, r, strokeWidth float64, strokeCol, valueCol, diceCol color.RGBA) *Dice {
	canv := canvas.New(w, h)

	return &Dice{
		x:           strokeWidth / 2,
		y:           strokeWidth / 2,
		canv:        canv,
		ctx:         canvas.NewContext(canv),
		strokeWidth: strokeWidth,
		strokeCol:   strokeCol,
		valueCol:    valueCol,
		diceCol:     diceCol,
		radius:      r,
	}
}

// New allows you to reuse existing ctx (canvas). x, y are coordinates of the dice to be rendered at.
func New(ctx *canvas.Context, x, y, r, strokeWidth float64, strokeCol, valueCol, diceCol color.RGBA) *Dice {

	return &Dice{
		x:           x,
		y:           y,
		ctx:         ctx,
		strokeWidth: strokeWidth,
		strokeCol:   strokeCol,
		valueCol:    valueCol,
		diceCol:     diceCol,
		radius:      r,
	}

}

// D6 will generate D6 dice face. Numeric determines if pips are used
// or just numeric value (like 1, 2, ...)
func (d *Dice) D6(face int, w, h float64, numeric bool) {
	d.ctx.SetStrokeWidth(d.strokeWidth)
	d.ctx.SetStrokeColor(d.strokeCol)
	d.ctx.SetFillColor(d.diceCol)
	D6Rec(d.ctx, d.x, d.y, w, h, d.radius)

	d.ctx.StrokeWidth = 0
	d.ctx.SetStrokeColor(d.valueCol)
	d.ctx.SetFillColor(d.valueCol)
	Dots(d.ctx, d.x, d.y, w, h, d.radius, face)
}

// WriteOutput renders dice into actual file.
func (d *Dice) WriteOutput(filename string) {
	if d.canv != nil {
		renderers.Write(filename, d.canv)
	}
}
