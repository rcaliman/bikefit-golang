package utils

import (
	"bikefit/databases"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Contador() string {
	var quantidade int
	var textoContador string
	databases.DB.Raw("select count(*) from medidas where DATE(created_at) = DATE_FORMAT(now(), '%Y-%m-%d');").Scan(&quantidade)

	if quantidade > 1 {
		textoContador = fmt.Sprintf("Seja bem vindo(a). Ja fizemos %v cálculos hoje! :)", quantidade)
	} else {
		textoContador = ""
	}
	return textoContador
}

func StringParaFloat(info string) float64 {
	valor, _ := strconv.ParseFloat(info, 64)
	return valor
}

func FormataData(t time.Time) string {
	return t.Format("02/01/2006")
}

func FormataFloat(v float64) string {
	s := fmt.Sprintf("%.1f", v)
	stringFormatada := strings.Replace(s, ".", ",", 1)
	return stringFormatada
}
