// brigadaEscena.go
package scenes

import (
	_"fmt"
	//"multi/scenes"
	"multi/shared"

	"github.com/hajimehoshi/ebiten/v2"
)

var(
    davity *ebiten.Image
	persona *ebiten.Image
    mesa *ebiten.Image
    caja *ebiten.Image
    tubo *ebiten.Image
    sensorAlcohol *ebiten.Image
    sensorTemperatura *ebiten.Image
    nevera *ebiten.Image
    lcd *ebiten.Image
)

var(
    MostrarLcd bool
)

func InitSimulationScene() {
    davity = shared.LoadImage("public/davity.png")
    persona = shared.LoadImage("public/persona.png")
    mesa = shared.LoadImage("public/pc.png")
    caja = shared.LoadImage("public/cajaRasp.png")
    tubo = shared.LoadImage("public/tubo.png")
    sensorAlcohol = shared.LoadImage("public/sensorAlcohol.png")
    sensorTemperatura = shared.LoadImage("public/mlx90614.png")
    nevera = shared.LoadImage("public/cajaVacunas.png")
    lcd = shared.LoadImage("public/lcd.png")
    
}


func DrawSistema(screen *ebiten.Image) {

    // mostrar imagen
    if persona != nil {
        shared.DrawImagen(screen, davity, 0.17, 250, 90)
                                        //size   X   Y
    }

    if mesa != nil{
        shared.DrawImagen(screen, mesa, 0.25, 580, 120)
    }

    if caja != nil{
        shared.DrawImagen(screen, caja, 0.25, 620, 210)
    }

    if tubo != nil{
        shared.DrawImagen(screen, tubo, 0.25, 500, 100)
    }

    if sensorAlcohol != nil{
        shared.DrawImagen(screen, sensorAlcohol, 0.28, 490, 398)
    }

    if sensorTemperatura != nil{
        shared.DrawImagen(screen, sensorTemperatura, 0.18, 555, 390)
    }

    if nevera != nil{
        shared.DrawImagen(screen, nevera, 0.20, 700, 90)
    }

    if lcd != nil{
        shared.DrawImagen(screen, lcd, 0.4, 815, 90)
    }
}


func DetectarClickLcd() {
    x, y := ebiten.CursorPosition()
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        if x >= 720 && x <= 720+lcd.Bounds().Dx()/5 && y >= 106 && y <= 106+lcd.Bounds().Dy()/5 {
            MostrarLcd = true
        }
    }
}
