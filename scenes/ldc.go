//lcd.go
package scenes

import(
	"github.com/hajimehoshi/ebiten/v2"
	"multi/shared"
)

var(
	ldc *ebiten.Image
)

func InitLcd(){
	lcd = shared.LoadImage("public/lcd.png")
}

func DrawModalLcd(ebiten *ebiten.Image){
	if lcd != nil{
		shared.DrawCenteredImage(ebiten, lcd, 0.5)
	}
}