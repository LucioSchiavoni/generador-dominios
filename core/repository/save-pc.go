package repository

import "generador-dominios/config"

func SavePcRepository(namePc string) error {

	query := `INSERT INTO equipo (nombre) VALUES (?)`

	_, err := config.DB.Exec(query, namePc)

	if err != nil {
		println("Error al guardar el nombre del PC en la base de datos:", err)
	}
	return nil
}
