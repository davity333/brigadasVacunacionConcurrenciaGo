// main.go
package main

import (
	"log"
	"math"

	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "os"
	_ "fmt"
	"multi/scenes"
	"multi/shared"
)

type App struct {
    fondo *ebiten.Image
}

func (a *App) Update() error {
    scenes.DetectarClickLcd()
    scenes.DetectarClickBtnVisual()
    return nil
}

func (a *App) Draw(screen *ebiten.Image) {
    // Calcular escala para ajustar fondo a la ventana
    escalaX := float64(screen.Bounds().Dx()) / float64(a.fondo.Bounds().Dx())
    escalaY := float64(screen.Bounds().Dy()) / float64(a.fondo.Bounds().Dy())
    escala := math.Min(escalaX, escalaY)

    // Dibujar imagen centrada con escala
    shared.DrawCenteredImage(screen, a.fondo, escala)
    if scenes.MostrarBrigada {
        scenes.DrawSistema(screen)   // brigadaEscena.go
    } else {
        scenes.DrawFlujo(screen)     // flujoProyecto.go
    }

    // Texto encima del fondo
    ebitenutil.DebugPrint(screen, "¡Bienvenido a brigadas de vacunación!")
    if scenes.MostrarLcd {
        scenes.DrawModalLcd(screen)
    }
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 1200, 687
}

func main() {
    // Cargar imagen de fondo
    fondo := shared.LoadImage("public/fondoBlanco.jpg")

    fontData, err := os.ReadFile("public/tuFuente.ttf") // pon aquí tu archivo .ttf
    shared.FontLoad(err, fontData, 18) // tamaño 18 px

    scenes.InitSimulationScene()
    scenes.InitFlujoProyecto()
    scenes.InitLcd()

    // Configurar ventana
    ebiten.SetWindowSize(1200, 687)
    ebiten.SetWindowTitle("Brigadas de vacunación - Simulación")

    // Ejecutar juego
    app := &App{fondo: fondo}
    if err := ebiten.RunGame(app); err != nil {
        log.Fatal(err)
    }
}
