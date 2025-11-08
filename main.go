// main.go
package main

import (
	"log"
	"math"

	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"multi/scenes"
	"multi/shared"
)

type App struct {
    fondo *ebiten.Image
}

func (a *App) Update() error {
    scenes.DetectarClickLcd()
    return nil
}

func (a *App) Draw(screen *ebiten.Image) {
    // Calcular escala para ajustar fondo a la ventana
    escalaX := float64(screen.Bounds().Dx()) / float64(a.fondo.Bounds().Dx())
    escalaY := float64(screen.Bounds().Dy()) / float64(a.fondo.Bounds().Dy())
    escala := math.Min(escalaX, escalaY)

    // Dibujar imagen centrada con escala
    shared.DrawCenteredImage(screen, a.fondo, escala)
    scenes.DrawSistema(screen)


    // Texto encima del fondo
    ebitenutil.DebugPrint(screen, "¡Bienvenido a brigadas de vacunación!")
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 1200, 687
}

func main() {
    // Cargar imagen de fondo
    fondo := shared.LoadImage("public/parque.jpg")
    scenes.InitSimulationScene()

    // Configurar ventana
    ebiten.SetWindowSize(1200, 687)
    ebiten.SetWindowTitle("Brigadas de vacunación - Simulación")

    // Ejecutar juego
    app := &App{fondo: fondo}
    if err := ebiten.RunGame(app); err != nil {
        log.Fatal(err)
    }
}
