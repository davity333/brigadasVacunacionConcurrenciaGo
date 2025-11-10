// flujoProyecto.go
package scenes

import (
	"image/color"
	"multi/shared"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	raspberry        *ebiten.Image
	flecha           *ebiten.Image
	consumer         *ebiten.Image
	apiWeb           *ebiten.Image
	baseDatos        *ebiten.Image
	apiVacunas       *ebiten.Image
	dobleFlecha      *ebiten.Image
	computadora      *ebiten.Image
	apiVacunados     *ebiten.Image
	flechaLarga      *ebiten.Image
	flechaArriba     *ebiten.Image
	dht21            *ebiten.Image
	ds18b20          *ebiten.Image
	mq3              *ebiten.Image
	mlx90614         *ebiten.Image
	btnMostrarVisual *ebiten.Image

	MostrarBrigada bool = false
)

func InitFlujoProyecto() {
	raspberry = shared.LoadImage("public/rasp.png")
	flecha = shared.LoadImage("public/flechaDerecha.png")
	consumer = shared.LoadImage("public/consumer.png")
	apiWeb = shared.LoadImage("public/apiWebsocket.png")
	baseDatos = shared.LoadImage("public/bd.png")
	apiVacunas = shared.LoadImage("public/api.png")
	dobleFlecha = shared.LoadImage("public/dobleFlecha.png")
	computadora = shared.LoadImage("public/computadora.png")
	apiVacunados = shared.LoadImage("public/apiVaccines.png")
	flechaLarga = shared.LoadImage("public/flechaLarga.png")
	flechaArriba = shared.LoadImage("public/flechaArriba.png")
	dht21 = shared.LoadImage("public/dht21.png")
	ds18b20 = shared.LoadImage("public/ds18b20.png")
	mq3 = shared.LoadImage("public/mq3.png")
	mlx90614 = shared.LoadImage("public/mlx90614.png")
	btnMostrarVisual = shared.LoadImage("public/btnMostrarVisual.png")
}

func DrawFlujo(screen *ebiten.Image) {
	if raspberry != nil {
		shared.DrawImagen(screen, raspberry, 0.27, 850, 270)
	}

	if flecha != nil {
		shared.DrawImagen(screen, flecha, 0.3, 740, 500)
	}

	if consumer != nil {
		shared.DrawImagen(screen, consumer, 0.25, 610, 470)
	}

	if flecha != nil {
		shared.DrawImagen(screen, flecha, 0.3, 499, 500)
	}

	if apiWeb != nil {
		shared.DrawImagen(screen, apiWeb, 0.25, 380, 470)
	}

	if flecha != nil {
		shared.DrawImagen(screen, flecha, 0.3, 270, 500)
	}

	if baseDatos != nil {
		shared.DrawImagen(screen, baseDatos, 1, 185, 470)
	}

	if apiVacunas != nil {
		shared.DrawImagen(screen, apiVacunas, 0.3, 636, 270)
	}

	if dobleFlecha != nil {
		shared.DrawImagen(screen, dobleFlecha, 0.3, 529, 300)
	}

	if computadora != nil {
		shared.DrawImagen(screen, computadora, 0.8, 360, 270)
	}

	if dobleFlecha != nil {
		shared.DrawImagen(screen, dobleFlecha, 0.3, 263, 300)
	}

	if apiVacunados != nil {
		shared.DrawImagen(screen, apiVacunados, 0.3, 115, 270)
	}

	if flechaLarga != nil {
		shared.DrawImagen(screen, flechaLarga, 0.3, 30, 190)
	}

	if flechaArriba != nil {
		shared.DrawImagen(screen, flechaArriba, 0.4, 220, 370)
	}

	//sensores
	if dht21 != nil {
		shared.DrawImagen(screen, dht21, 1, 1020, 290)
	}

	if ds18b20 != nil {
		shared.DrawImagen(screen, ds18b20, 0.3, 980, 210)
	}

	if mq3 != nil {
		shared.DrawImagen(screen, mq3, 1, 888, 205)
	}

	if mlx90614 != nil {
		shared.DrawImagen(screen, mlx90614, 1, 820, 340)
	}

	if dht21 != nil {
		shared.DrawImagen(screen, dht21, 1, 1020, 290)
		text.Draw(screen, "DHT21", shared.LsdNumero, 1020, 280, color.White)
	}

	if ds18b20 != nil {
		shared.DrawImagen(screen, ds18b20, 0.3, 980, 210)
		text.Draw(screen, "DS18B20", shared.LsdNumero, 980, 200, color.White)
	}

	if mq3 != nil {
		shared.DrawImagen(screen, mq3, 1, 888, 205)
		text.Draw(screen, "MQ3", shared.LsdNumero, 888, 195, color.White)
	}

	if mlx90614 != nil {
		shared.DrawImagen(screen, mlx90614, 1, 820, 340)
		text.Draw(screen, "MLX90614", shared.LsdNumero, 820, 330, color.White)
	}

	if btnMostrarVisual != nil {
		shared.DrawImagen(screen, btnMostrarVisual, 0.2, 1058, 610)
	}

	// IMPORTANTE: Dibujar los valores animados también en esta vista
	dibujarValoresAnimados(screen)
}

func DetectarClickBtnVisual() {
	if btnMostrarVisual == nil {
		return
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		escala := 0.2
		ancho := int(float64(btnMostrarVisual.Bounds().Dx()) * escala)
		alto := int(float64(btnMostrarVisual.Bounds().Dy()) * escala)

		inicioX := 45 - (ancho / 2) + 40
		inicioY := 50 - (alto / 2)
		finX := inicioX + ancho
		finY := inicioY + alto

		if x >= inicioX && x <= finX && y >= inicioY && y <= finY {
			MostrarBrigada = !MostrarBrigada
		}
	}
}
