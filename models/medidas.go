package models

import "gorm.io/gorm"

type Medidas struct {
	gorm.Model
	Cavalo         float64 `json:"cavalo"`
	Esterno        float64 `json:"esterno"`
	Braco          float64 `json:"braco"`
	Email          string  `json:"email"`
	Tronco         float64 `json:"tronco"`
	QuadroSpeed    float64 `json:"quadro_speed"`
	QuadroMTB      float64 `json:"quadro_mtb"`
	AlturaSelim    float64 `json:"altura_selim"`
	TopTubeEfetivo float64 `json:"top_tube_efetivo"`
}

func (m Medidas) CalculaTronco(esterno, cavalo float64) float64 {
	return esterno - cavalo
}

func (m Medidas) CalculaSpeed(cavalo float64) float64 {
	return cavalo * 0.67
}

func (m Medidas) CalculaMTB(cavalo float64) float64 {
	return (cavalo*0.67 - 10) * 0.393700787
}

func (m Medidas) CalculaSelim(cavalo float64) float64 {
	return cavalo * 0.883
}

func (m Medidas) CalculaTopTube(esterno, cavalo, braco float64) float64 {
	return ((esterno - cavalo + braco) / 2) + 4
}
