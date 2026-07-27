package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type GridData struct {
	CountryCode string  `json:"country_code"`
	Intensity   float64 `json:"intensity"` // gCO2/kWh
	Timestamp   int64   `json:"timestamp"`
}

type IPResponse struct {
	CountryCode string `json:"countryCode"`
}

func GetGridIntensity() (GridData, error) {
	cacheDir, _ := os.UserCacheDir()
	cacheFile := filepath.Join(cacheDir, "greenfetch_cache.json")

	if data, err := os.ReadFile(cacheFile); err == nil {
		var cachedData GridData
		if json.Unmarshal(data, &cachedData) == nil {
			// Se dados ainda estiverem em cache nos últimos 30 minutos, utilizar esses dados
			if time.Now().Unix()-cachedData.Timestamp < 1800 {
				return cachedData, nil
			}
		}
	}

	// Localização do usuário
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return GridData{}, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	var ipData IPResponse
	json.NewDecoder(resp.Body).Decode(&ipData)

	// Utilizando mock para dados de carbono, poderia ser substituido por uma API real como CO2Signal/Electricity Maps (porém é paga)
	apiKey := os.Getenv("CO2_API_KEY")
	var intensity float64

	if apiKey != "" {
		// Valor simbólico
		intensity = 250.0
	} else {
		// Mock de dados baseado em localização para demo
		if ipData.CountryCode == "FR" || ipData.CountryCode == "BR" {
			intensity = 60.0 // Baixa utilização de carbono (Utilizam energia nuclear/hidroelétrica)
		} else {
			intensity = 350.0 // Média para utilização com combustíveis fósseis
		}
	}

	newData := GridData{
		CountryCode: ipData.CountryCode,
		Intensity:   intensity,
		Timestamp:   time.Now().Unix(),
	}

	os.MkdirAll(cacheDir, 0755)
	fileData, _ := json.Marshal(newData)
	os.WriteFile(cacheFile, fileData, 0644)

	return newData, nil
}
