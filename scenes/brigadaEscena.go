// brigadaEscena.go
package scenes

import (
	"image/color"
	"multi/shared"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	verTopologia      *ebiten.Image
)

// Estructura para eventos visuales en movimiento
type AnimatedValue struct {
	Text      string
	X         float64
	Y         float64
	TargetX   float64
	TargetY   float64
	Active    bool
	StartTime time.Time
}

var (
	// Cola de valores animados
	animatedValues = make([]AnimatedValue, 0, 10)
	// Contador de eventos para debug
	eventosRecibidos = 0
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

	// Dibujar contador de eventos en pantalla (esquina superior derecha)
	ebitenutil.DrawRect(screen, 1000, 50, 180, 40, color.RGBA{0, 0, 255, 200})
	ebitenutil.DebugPrintAt(screen, "Eventos: "+strconv.Itoa(eventosRecibidos), 1010, 60)

	// Dibujar valores animados
	dibujarValoresAnimados(screen)
}

// dibujarValoresAnimados actualiza y dibuja los valores en movimiento
func dibujarValoresAnimados(screen *ebiten.Image) {
	now := time.Now()

	// Debug: mostrar cuántas animaciones hay activas
	if len(animatedValues) > 0 {
		println("📊 Dibujando", len(animatedValues), "animaciones")
	}

	for i := len(animatedValues) - 1; i >= 0; i-- {
		av := &animatedValues[i]
		if !av.Active {
			continue
		}

		// Calcular progreso (0 a 1) en 2 segundos (MÁS lento para verlo mejor)
		elapsed := now.Sub(av.StartTime).Seconds()
		progress := elapsed / 2.0
		if progress >= 1.0 {
			// Animación completa, eliminar
			println("✅ Animación completada:", av.Text)
			animatedValues = append(animatedValues[:i], animatedValues[i+1:]...)
			continue
		}

		// Interpolación lineal suavizada
		av.X = av.X + (av.TargetX-av.X)*0.08
		av.Y = av.Y + (av.TargetY-av.Y)*0.08

		// Dibujar cuadro pequeño y elegante
		ancho := float64(60)
		alto := float64(20)

		// Fondo amarillo con borde negro
		ebitenutil.DrawRect(screen, av.X-2, av.Y-2, ancho+4, alto+4, color.RGBA{0, 0, 0, 255})
		ebitenutil.DrawRect(screen, av.X, av.Y, ancho, alto, color.RGBA{255, 255, 0, 255})

		// Texto negro pequeño
		ebitenutil.DebugPrintAt(screen, av.Text, int(av.X)+2, int(av.Y)+2)
		println("  -> Dibujando en", int(av.X), int(av.Y), ":", av.Text)
	}
}

func DetectarClickLcd() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		inicioX := 390.0
		inicioY := 540.0
		ancho := 118.0
		alto := 38.0
		finX := inicioX + ancho
		finY := inicioY + alto

		if float64(x) >= inicioX && float64(x) <= finX &&
			float64(y) >= inicioY && float64(y) <= finY {
			MostrarLcd = !MostrarLcd
		} else {
		}
	}
}

// ProcesarEventosVisuales consume eventos del pipeline y los convierte en animaciones
func ProcesarEventosVisuales(visualChan chan interface{}) {
	for {
		select {
		case evt := <-visualChan:
			// evt es un VisualEvent del pipeline
			if ve, ok := evt.(map[string]interface{}); ok {
				text := ve["Text"].(string)
				from := ve["From"].(string)
				to := ve["To"].(string)
				crearAnimacion(text, from, to)
			}
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// crearAnimacion crea una nueva animación basada en origen y destino
func crearAnimacion(text, from, to string) {
	var startX, startY, targetX, targetY float64

	positions := map[string][2]float64{
		"tempSensor":    {150, 450},
		"alcoholSensor": {100, 500},
		"fridgeSensor":  {120, 400},

		"rasp": {280, 350},

		"consumer": {550, 150},
		"apiWeb":   {750, 150},
		"db":       {950, 150},

		"apiVaccines":  {500, 350},
		"computadora":  {750, 350},
		"apiVacunados": {1000, 350},
	}

	if start, ok := positions[from]; ok {
		startX, startY = start[0], start[1]
	}
	if target, ok := positions[to]; ok {
		targetX, targetY = target[0], target[1]
	}

	av := AnimatedValue{
		Text:      text,
		X:         startX,
		Y:         startY,
		TargetX:   targetX,
		TargetY:   targetY,
		Active:    true,
		StartTime: time.Now(),
	}

	animatedValues = append(animatedValues, av)
}

func AgregarEventoVisual(text, from, to string) {
	println("Evento visual recibido:", text, "desde", from, "hacia", to)
	println("   Total animaciones activas:", len(animatedValues))
	eventosRecibidos++
	crearAnimacion(text, from, to)
}
