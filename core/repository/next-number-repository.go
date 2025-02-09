package repository

import (
	"fmt"
	"generador-dominios/config"
	"log"
	"sort"
	"strconv"
)

func GetNextNumber(unidad, oficina string) (string, error) {
	nom := unidad + "-PC-" + oficina + "-%"

	query := `
		SELECT SUBSTRING_INDEX(e.nombre, '-', -1) 
		FROM equipo e 
		WHERE e.nombre LIKE ?
	`

	rows, err := config.DB.Query(query, nom)
	if err != nil {
		log.Fatal("❌ Error en la consulta:", err)
		return "", err
	}
	defer rows.Close()

	var numbers []int

	for rows.Next() {
		var numStr string
		if err := rows.Scan(&numStr); err == nil {
			if num, err := strconv.Atoi(numStr); err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	// Ordenar numeros en orden ascendente
	sort.Ints(numbers)

	// El menor numero faltante en el rango
	nextNum := 1
	for _, num := range numbers {
		if num != nextNum {
			break
		}
		nextNum++
	}

	return fmt.Sprintf("%03d", nextNum), nil
}
