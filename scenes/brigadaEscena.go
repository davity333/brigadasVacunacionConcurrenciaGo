// brigadaEscena.go
package scenes

import (
    "multi/shared"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "fmt"
)

var (
    davity            *ebiten.Image
    persona           *ebiten.Image
    mesa              *ebiten.Image
    caja              *ebiten.Image
    tubo              *ebiten.Image
    sensorAlcohol     *ebiten.Image
    sensorTemperatura *ebiten.Image
    nevera            *ebiten.Image
    lcd               *ebiten.Image
    verTopologia *ebiten.Image
)

func InitSimulationScene() {
    davity = shared.LoadImage("public/davity.png")
    persona = shared.LoadImage("public/persona.png")
    mesa = shared.LoadImage("public/pc.png")
    caja = shared.LoadImage("public/cajaRasp.png")
    tubo = shared.LoadImage("public/tubo.png")
    sensorAlcohol = shared.LoadImage("public/sensorAlcohol.png")
    sensorTemperatura = shared.LoadImage("public/mlx90614A.png")
    nevera = shared.LoadImage("public/cajaVacunas.png")
    lcd = shared.LoadImage("public/lcd.png")
    verTopologia = shared.LoadImage("public/btnVerTopologia.png")
}

func DrawSistema(screen *ebiten.Image) {
    if persona != nil {
        shared.DrawImagen(screen, davity, 0.17, 250, 90)
    }
    if mesa != nil {
        shared.DrawImagen(screen, mesa, 0.25, 580, 120)
    }
    if caja != nil {
        shared.DrawImagen(screen, caja, 0.25, 620, 210)
    }
    if tubo != nil {
        shared.DrawImagen(screen, tubo, 0.25, 500, 100)
    }
    if sensorAlcohol != nil {
        shared.DrawImagen(screen, sensorAlcohol, 0.28, 490, 398)
    }
    if sensorTemperatura != nil {
        shared.DrawImagen(screen, sensorTemperatura, 0.18, 555, 390)
    }
    if nevera != nil {
        shared.DrawImagen(screen, nevera, 0.20, 700, 90)
    }
    if lcd != nil {
        shared.DrawImagen(screen, lcd, 0.03, 718, 108)
    }

    if verTopologia != nil {
		shared.DrawImagen(screen, verTopologia, 0.2, 1058, 610)
	}
}

func DetectarClickLcd() {
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        x, y := ebiten.CursorPosition()

        // Área del cuadro rojo "REGISTRO"
        inicioX := 390.0
        inicioY := 540.0
        ancho := 118.0
        alto := 38.0
        finX := inicioX + ancho
        finY := inicioY + alto

        fmt.Printf("Área REGISTRO: X[%.0f–%.0f] Y[%.0f–%.0f] Cursor: (%d,%d)\n", inicioX, finX, inicioY, finY, x, y)

        if float64(x) >= inicioX && float64(x) <= finX &&
           float64(y) >= inicioY && float64(y) <= finY {
            MostrarLcd = !MostrarLcd
        } else {
        }
    }
}
