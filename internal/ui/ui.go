package ui

import (
	"fmt"
)

// Cores
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Bold   = "\033[1m"
	Cyan   = "\033[36m"
)

const asciiArt = `
      %s.%s
    %s...%s
   %s.....%s
 %s........%s
   %s....%s
     %s|%s
`

func PrintDashboard(watts float64, intensity float64, gCO2perHour float64, country string) {
	// Determinar cores baseado na intensidade do grid
	gridColor := Yellow
	if intensity < 150 {
		gridColor = Green
	} else if intensity > 400 {
		gridColor = Red
	}

	fmt.Printf("\n" + Bold + Cyan + "   GreenFetch" + Reset + "\n")
	fmt.Println("   ---------------------")

	fmt.Printf("   Hardware Power Draw:  %.2f Watts\n", watts)
	fmt.Printf("   Grid Location:        %s\n", country)
	fmt.Printf("   Grid Intensity:       %s%.2f gCO2/kWh%s\n", gridColor, intensity, Reset)
	fmt.Println("   ---------------------")
	fmt.Printf("   Live Carbon Cost:     %s%.2f gCO2/hour%s\n\n", Bold+gridColor, gCO2perHour, Reset)
}
