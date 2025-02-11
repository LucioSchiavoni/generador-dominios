package usecases

import (
	"fmt"
	"generador-dominios/core/repository"
	"log"
)

func RenamePCService() error {
	fmt.Println("Seleccione la Unidad:")
	var unidad string
	fmt.Scanln(&unidad)

	fmt.Println("Seleccione la Oficina:")
	var oficina string
	fmt.Scanln(&oficina)

	nextNumber, err := repository.GetNextNumber(unidad, oficina)
	if err != nil {
		log.Fatal("❌ Error al obtener el número incremental:", err)
	}

	newPCName := fmt.Sprintf("%s-PC-%s-%s", unidad, oficina, nextNumber)

	confirm := repository.AskForConfirmation(newPCName)
	if !confirm {
		fmt.Println("❌ Cambio de nombre del PC cancelado.")
		return nil
	}
	err = repository.SavePcRepository(newPCName)
	if err != nil {
		log.Fatal("❌ No se pudo guardar el nombre del PC:", err)
	}
	err = repository.RenamePC(repository.RenamePcParams{NewName: newPCName})

	if err != nil {
		log.Fatal("❌ No se pudo cambiar el nombre del PC:", err)
	}

	fmt.Println("✅ Nombre del PC cambiado exitosamente a:", newPCName)
	return nil
}
