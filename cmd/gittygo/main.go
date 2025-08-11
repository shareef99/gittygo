package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("GittyGo"))
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	var incrementTag = new(bool)
	var decrementTag = new(bool)
	count := 0

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			for {
				ev, ok := gtx.Event(pointer.Filter{
					Target: incrementTag,
					Kinds:  pointer.Press,
				})
				if !ok {
					break
				}
				if x, ok := ev.(pointer.Event); ok && x.Kind == pointer.Press {
					count = count + 1
				}
			}

			for {
				ev, ok := gtx.Event(pointer.Filter{
					Target: decrementTag,
					Kinds:  pointer.Press,
				})
				if !ok {
					break
				}
				if x, ok := ev.(pointer.Event); ok && x.Kind == pointer.Press {
					count = count - 1
				}
			}

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Shareef Master!"+strconv.Itoa((count)))

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					rect := clip.Rect{Max: image.Pt(100, 100)}.Push(gtx.Ops)
					event.Op(gtx.Ops, incrementTag)

					paint.ColorOp{Color: color.NRGBA{G: 0xFF, A: 0xFF}}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)

					rect.Pop()

					return layout.Dimensions{Size: image.Pt(100, 100)}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					countText := material.Body1(theme, "Count: "+strconv.Itoa(count))
					countText.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
					return countText.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					rect := clip.Rect{Max: image.Pt(100, 100)}.Push(gtx.Ops)
					event.Op(gtx.Ops, decrementTag)

					paint.ColorOp{Color: color.NRGBA{R: 0xFF, A: 0xFF}}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)

					rect.Pop()
					return layout.Dimensions{Size: image.Pt(100, 100)}
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
