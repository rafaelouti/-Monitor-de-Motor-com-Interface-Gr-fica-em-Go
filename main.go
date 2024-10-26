package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"math/rand"
)

// Funções para simular valores de sensores
func getRPM() int {
	return rand.Intn(5000) // Simula um valor de RPM entre 0 e 5000
}

func getTemperature() float64 {
	return rand.Float64()*100 + 20 // Simula uma temperatura entre 20 e 120°C
}

func getVibration() float64 {
	return rand.Float64()*10 + 1 // Simula vibração entre 1 e 10 G
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Monitoramento de Motor")

	// Botão para exibir RPM
	rpmButton := widget.NewButton("Verificar RPM", func() {
		rpm := getRPM()
		alertMsg := fmt.Sprintf("RPM Atual: %d RPM", rpm)
		dialog.ShowInformation("RPM do Motor", alertMsg, myWindow)
	})

	// Botão para exibir Temperatura
	tempButton := widget.NewButton("Verificar Temperatura", func() {
		temperature := getTemperature()
		alertMsg := fmt.Sprintf("Temperatura Atual: %.2f°C", temperature)
		dialog.ShowInformation("Temperatura do Motor", alertMsg, myWindow)
	})

	// Botão para exibir Vibração
	vibrationButton := widget.NewButton("Verificar Vibração", func() {
		vibration := getVibration()
		alertMsg := fmt.Sprintf("Vibração Atual: %.2f G", vibration)
		dialog.ShowInformation("Vibração do Motor", alertMsg, myWindow)
	})

	// Layout da interface
	content := container.NewVBox(
		widget.NewLabel("Sistema de Monitoramento de Motor"),
		rpmButton,
		tempButton,
		vibrationButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
