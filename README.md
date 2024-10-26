# Monitor de Motor com Interface Gráfica em Go

Este projeto é uma aplicação em Go que usa a biblioteca `fyne` para criar uma interface gráfica de monitoramento de um motor. A aplicação exibe botões interativos para verificar valores simulados de RPM, temperatura e vibração do motor, mostrando alertas de cada métrica.

## Tecnologias Utilizadas

- **Go**: Linguagem principal utilizada no projeto.
- **Fyne**: Biblioteca para criar interfaces gráficas em Go, permitindo que a aplicação rode em modo gráfico (GUI) e não no terminal.

## Pré-requisitos

- **Go**: Certifique-se de que o Go está instalado (versão 1.16 ou superior). [Instale o Go](https://golang.org/doc/install)
- **MinGW** (apenas para Windows): Para compilar o código Go com suporte a GUI no Windows. [MinGW-w64 Download](https://sourceforge.net/projects/mingw-w64/)

## Configuração para Windows

1. **Instale o MinGW**:
   - Baixe e instale o MinGW para Windows, adicionando o caminho do diretório `bin` do MinGW ao seu `PATH` do Windows.
   - Verifique a instalação com o comando:
     ```bash
     gcc --version
     ```

2. **Instale a biblioteca `fyne`**:
   No terminal do projeto, execute:
   ```bash
   go get fyne.io/fyne/v2


CODIGO PRINCIPAL

package main

import (
    "fmt"
    "math/rand"
    "time"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
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


Funcionalidades
Verificar RPM: Exibe um alerta com o valor simulado de RPM do motor.
Verificar Temperatura: Exibe um alerta com a temperatura simulada do motor.
Verificar Vibração: Exibe um alerta com o nível de vibração simulado do motor.
Observações
Este é um exemplo simples e utiliza valores simulados. Para capturar dados reais, você precisaria integrar sensores ao sistema e capturar as leituras.

Referências
Documentação do Fyne
