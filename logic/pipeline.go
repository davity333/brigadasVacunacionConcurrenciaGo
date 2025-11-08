package logic

import(
	"math/rand"
	"time"
)

func SensorTemperatura(temperatura chan int) {
    for {
        nueva := rand.Intn(5) + 36 // entre 36 y 40
        temperatura <- nueva
        time.Sleep(2 * time.Second)
    }
}

func SensorAlcohol(alcoholemia chan bool) {
    for {
        lectura := rand.Intn(2) == 1 // true o false aleatorio
        alcoholemia <- lectura
        time.Sleep(1 * time.Second) // espera entre lecturas
    }
}
