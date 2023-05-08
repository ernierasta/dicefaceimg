package diceimg

import (
	"fmt"

	"github.com/tdewolff/canvas"
)

// here are all dice shapes

// D6 dot positions. Counted as shown:
// 3   7
// 2 4 6
// 1   5
type d6dot struct {
	x, y, w, h, r float64
	ctx           *canvas.Context
}

func (d *d6dot) i1() {
	d.dot(d.w/4, d.h/4)
}
func (d *d6dot) i2() {
	d.dot(d.w/4, d.h/2)
}
func (d *d6dot) i3() {
	d.dot(d.w/4, (d.h/4)*3)
}
func (d *d6dot) i4() {
	d.dot(d.w/2, d.h/2)
}
func (d *d6dot) i5() {
	d.dot((d.w/4)*3, d.h/4)
}
func (d *d6dot) i6() {
	d.dot((d.w/4)*3, d.h/2)
}
func (d *d6dot) i7() {
	d.dot((d.w/4)*3, (d.h/4)*3)
}

// dot renders one dot related to rectangle
func (d *d6dot) dot(x, y float64) {
	fmt.Printf("dot(rad: %v) at %v, %v\n", d.r, x, y) //DEBUG
	c := canvas.Circle(d.r)
	d.ctx.DrawPath(d.x+x, d.y+y, c)

}
func D6Rec(ctx *canvas.Context, x, y, w, h, r float64) {
	fmt.Printf("rect (rad: %v): x %v,y %v, size: w %v, h %v\n", r, x, y, w, h)
	rec := canvas.RoundedRectangle(w, h, r)
	ctx.DrawPath(x, y, rec)
}

// 3   7
// 2 4 6
// 1   5
func Dots(ctx *canvas.Context, x, y, w, h, r float64, i int) {
	d := d6dot{x: x, y: y, w: w, h: h, r: r, ctx: ctx}
	switch i {
	case 1:
		d.i4()
	case 2:
		d.i1()
		d.i7()
	case 3:
		d.i1()
		d.i4()
		d.i7()
	case 4:
		d.i1()
		d.i3()
		d.i5()
		d.i7()
	case 5:
		d.i1()
		d.i3()
		d.i5()
		d.i7()
		d.i4()
	case 6:
		d.i1()
		d.i2()
		d.i3()
		d.i5()
		d.i6()
		d.i7()
	}
}
