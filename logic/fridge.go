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

// StartFridge inicia los sensores de la caja de vacunas
func StartFridge(visual chan VisualEvent) {
	go sensorTemperaturaInterior(visual)
	go sensorTemperaturaExterior(visual)
	go sensorHumedad(visual)
}

// sensorTemperaturaInterior simula lectura de temperatura interior
func sensorTemperaturaInterior(visual chan VisualEvent) {
	for {
		temp := rand.Intn(8) + 2 // 2..9 grados
		txt := strconv.Itoa(temp) + "°C INT"

		// Flujo: fridgeSensor → rasp → apiVaccines → computadora → apiVacunados
		visual <- VisualEvent{Text: txt, From: "fridgeSensor", To: "rasp"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "rasp", To: "apiVaccines"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "apiVaccines", To: "computadora"}
		time.Sleep(400 * time.Millisecond)

		// Enviar a API de vacunas
		go enviarSensorVacunaAPI("TempInterior-A1", float64(temp), "°C")

		time.Sleep(3 * time.Second)
	}
}

// sensorTemperaturaExterior simula lectura de temperatura exterior
func sensorTemperaturaExterior(visual chan VisualEvent) {
	for {
		temp := rand.Intn(15) + 20 
		txt := strconv.Itoa(temp) + "°C EXT"

		visual <- VisualEvent{Text: txt, From: "fridgeSensor", To: "rasp"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "rasp", To: "apiVaccines"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "apiVaccines", To: "computadora"}
		time.Sleep(400 * time.Millisecond)

		go enviarSensorVacunaAPI("TempExterior-A1", float64(temp), "°C")

		time.Sleep(4 * time.Second)
	}
}

// sensorHumedad simula lectura de humedad
func sensorHumedad(visual chan VisualEvent) {
	for {
		hum := rand.Intn(30) + 50 // 50..79%
		txt := strconv.Itoa(hum) + "%"

		// Flujo: fridgeSensor → rasp → apiVaccines → computadora
		visual <- VisualEvent{Text: txt, From: "fridgeSensor", To: "rasp"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "rasp", To: "apiVaccines"}
		time.Sleep(400 * time.Millisecond)

		visual <- VisualEvent{Text: txt, From: "apiVaccines", To: "computadora"}
		time.Sleep(400 * time.Millisecond)

		go enviarSensorVacunaAPI("Humedad-A1", float64(hum), "%")

		time.Sleep(5 * time.Second)
	}
}

// enviarSensorVacunaAPI envía datos a la API de vacunas
func enviarSensorVacunaAPI(sensor string, value float64, unit string) {
	payload := map[string]interface{}{
		"measurementUnit": unit,
		"nameSensor":      sensor,
		"information":     value,
		"idVaccineBox":    1,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8001/api/sensorsVaccine", bytes.NewReader(b))
	if err != nil {
		log.Printf("Error creando request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error enviando a API vacunas: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Enviado a API vacunas: %d", resp.StatusCode)
}
