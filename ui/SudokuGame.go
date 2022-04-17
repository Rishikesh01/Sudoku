package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	_ "image/png"
	"log"
)

type Game struct {
	img        *ebiten.Image
	keyMapping map[ebiten.Key]string
	key        []ebiten.Key
	font       font.Face
	x          int
	y          int
}

func (g *Game) Update() error {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	g.font, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {

	}
	mouse := ebiten.MouseButtonLeft
	if inpututil.IsMouseButtonJustPressed(mouse) {
		g.x, g.y = ebiten.CursorPosition()
	}
	inpututil.AppendPressedKeys(g.key[:0])

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.img, op)

	if value, exists := g.keyMapping[g.key[len(g.key)-1]]; exists {
		text.Draw(g.img, value, g.font, g.x, g.y, color.Black)
		log.Println(value)
		g.key = append(g.key[:0], ebiten.Key0)
	}

}
func (g *Game) Layout(outWidth, outHeight int) (screenWidth, screenHeight int) {
	return outWidth, outHeight
}

func (g *Game) init() {
	g.keyMapping = make(map[ebiten.Key]string)
	g.key = make([]ebiten.Key, 1)
	g.keyMapping[ebiten.Key1] = "1"
	g.keyMapping[ebiten.Key2] = "2"
	g.keyMapping[ebiten.Key3] = "3"
	g.keyMapping[ebiten.Key4] = "4"
	g.keyMapping[ebiten.Key5] = "5"
	g.keyMapping[ebiten.Key6] = "6"
	g.keyMapping[ebiten.Key7] = "7"
	g.keyMapping[ebiten.Key8] = "8"
	g.keyMapping[ebiten.Key9] = "9"
}

func Start() {
	g := &Game{}
	var err error
	g.img, _, err = ebitenutil.NewImageFromFile("./resources/blank-sudoku-grid.png")
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(504, 504)
	g.Layout(504, 540)
	g.init()
	if err := ebiten.RunGame(g); err != nil {
		println(err)
	}
}
