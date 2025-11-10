package logic

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type VisualEvent struct {
	Text string 
	From string 
	To   string
}

type Patient struct {
	ID   int
	Temp int
}

func StartPipeline(visual chan VisualEvent) {
	tempChan := make(chan Patient)
	alcoholChan := make(chan Patient)

	go generadorPacientes(tempChan)
	go etapaTemperatura(tempChan, alcoholChan, visual)
	go etapaAlcohol(alcoholChan, visual)
}

func generadorPacientes(out chan Patient) {
	for {
		out <- Patient{ID: 1}
		time.Sleep(2 * time.Second)
	}
}

func etapaTemperatura(in chan Patient, out chan Patient, visual chan VisualEvent) {
	for p := range in {
		temp := rand.Intn(6) + 35

		visual <- VisualEvent{Text: strconv.Itoa(temp) + "°C", From: "tempSensor", To: "rasp"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: strconv.Itoa(temp) + "°C", From: "rasp", To: "consumer"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: strconv.Itoa(temp) + "°C", From: "consumer", To: "apiWeb"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: strconv.Itoa(temp) + "°C", From: "apiWeb", To: "db"}
		time.Sleep(400 * time.Millisecond)

		// Enviar a API (siempre userID = 1)
		go enviarTemperaturaAPI(temp, 1)

		// Si temperatura <= 37, pasar a etapa de alcohol
		p.Temp = temp
		if temp <= 37 {
			out <- p
		} else {
			log.Printf("Paciente rechazado por temperatura alta: %d", temp)
		}
	}
}

// etapaAlcohol verifica alcoholemia y envía resultado
func etapaAlcohol(in chan Patient, visual chan VisualEvent) {
	for range in {
		// Simular lectura de alcohol
		alcoholOK := rand.Intn(2) == 0
		txt := "ALC_OK"
		if !alcoholOK {
			txt = "ALC_POS"
		}

		visual <- VisualEvent{Text: txt, From: "alcoholSensor", To: "rasp"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "rasp", To: "consumer"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "consumer", To: "apiWeb"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "apiWeb", To: "db"}
		time.Sleep(400 * time.Millisecond)

		// Enviar a API (siempre userID = 1)
		go enviarAlcoholAPI(alcoholOK, 1)

		if alcoholOK {
			log.Printf("Paciente aprobado para vacunación")
		} else {
			log.Printf("Paciente rechazado por alcohol positivo")
		}
	}
}

// enviarTemperaturaAPI envía temperatura a la API
func enviarTemperaturaAPI(temp int, userID int) {
	payload := map[string]interface{}{
		"measurementUnit":       "Celsius",
		"nameSensor":            "Temperatura Corporal",
		"information":           float64(temp),
		"UserCivil_idUserCivil": userID,
	}
	enviarJSON("http://localhost:8000/SensorCheck/", payload)
}

// enviarAlcoholAPI envía alcoholemia a la API
func enviarAlcoholAPI(ok bool, userID int) {
	val := 0.0
	if !ok {
		val = 1.0
	}
	payload := map[string]interface{}{
		"measurementUnit":       "unit",
		"nameSensor":            "Alcoholemia",
		"information":           val,
		"UserCivil_idUserCivil": userID,
	}
	enviarJSON("http://localhost:8000/SensorCheck/", payload)
}

func enviarJSON(url string, payload interface{}) {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		log.Printf("Error creando request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error enviando a %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Enviado a %s: %d", url, resp.StatusCode)
}
