package main

import (
	"fmt"
	"greenfetch/internal/api"
	"greenfetch/internal/power"
	"greenfetch/internal/ui"
	"os"
)

func main() {
	// Passo a passo para execução principal:
	// 1. Pegar watts gerados pelo power draw
	// 2. Usar API para pegar informações de consumo de acordo com a grid pertencente
	// 3. Calcular consumo - Formula: (Watts / 1000) * (gCO2 / kWh) = gCO2 / hora
	// 4. Renderizar a UI
	watts := power.GetPowerDrawW()

	gridData, err := api.GetGridIntensity()
	if err != nil {
		fmt.Printf("Error fetching grid data: %v\n", err)
		os.Exit(1)
	}

	kiloWatts := watts / 1000.0
	gCO2perHour := kiloWatts * gridData.Intensity

	ui.PrintDashboard(watts, gridData.Intensity, gCO2perHour, gridData.CountryCode)
}
