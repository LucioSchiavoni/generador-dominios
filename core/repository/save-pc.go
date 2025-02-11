package repository

import "generador-dominios/config"

func SavePcRepository(namePc string) error {

	query := `INSERT INTO equipo (nombre) VALUES (?)`

	_, err := config.DB.Exec(query, namePc)

	if err != nil {
		println("Se reporto un error al querer guardar el nombre en la base de datos:", err.Error())
		return err
	}
	return nil
}
