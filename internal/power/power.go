package power

import (
	"os"
	"strconv"
	"strings"
)

// GetPowerDrawW tenta ler a utilização de Power Draw usando uma syscall
func GetPowerDrawW() float64 {
	const powerPath = "/sys/class/power_supply/BAT0/power_now"

	data, err := os.ReadFile(powerPath)
	if err != nil {
		// Fallback para sistemas sem filesystem BAT0
		return 35.0
	}

	rawStr := strings.TrimSpace(string(data))
	microWatts, err := strconv.ParseFloat(rawStr, 64)
	if err != nil {
		return 35.0
	}

	watts := microWatts / 1_000_000.0

	// Se a bateria do computador estiver carregada completamente ou estiver ligado na tomada, power_now pode retornar 0
	if watts <= 0 {
		return 15.0 // Estimativa
	}

	return watts
}
