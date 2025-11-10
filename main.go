// main.go
package main

import (
	"log"
	"math"

	_ "image/jpeg"

	"multi/assets"
	"multi/logic"
	"multi/scenes"
	"multi/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type App struct {
	fondo      *ebiten.Image
	visualChan chan logic.VisualEvent
}

func (a *App) Update() error {
	scenes.DetectarClickLcd()
	scenes.DetectarClickBtnVisual()

	select {
	case evt := <-a.visualChan:
		scenes.AgregarEventoVisual(evt.Text, evt.From, evt.To)
	default:
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	escalaX := float64(screen.Bounds().Dx()) / float64(a.fondo.Bounds().Dx())
	escalaY := float64(screen.Bounds().Dy()) / float64(a.fondo.Bounds().Dy())
	escala := math.Min(escalaX, escalaY)

	shared.DrawCenteredImage(screen, a.fondo, escala)
	if scenes.MostrarBrigada {
		scenes.DrawSistema(screen)
	} else {
		scenes.DrawFlujo(screen)
	}

	ebitenutil.DebugPrint(screen, "¡Bienvenido a brigadas de vacunación!")
	if scenes.MostrarLcd {
		scenes.DrawModalLcd(screen)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1200, 687
}

func main() {
	fondo := shared.LoadImage("public/fondoBlanco.jpg")

	shared.FontLoad(nil, assets.RobotoTTF, 18)

	scenes.InitSimulationScene()
	scenes.InitFlujoProyecto()
	scenes.InitLcd()

	// Crear canal de eventos visuales
	visualChan := make(chan logic.VisualEvent, 100)

	// Iniciar pipeline y sensores de vacunas
	logic.StartPipeline(visualChan)
	logic.StartFridge(visualChan)

	ebiten.SetWindowSize(1200, 687)
	ebiten.SetWindowTitle("Brigadas de vacunación - Simulación")

	app := &App{
		fondo:      fondo,
		visualChan: visualChan,
	}
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
