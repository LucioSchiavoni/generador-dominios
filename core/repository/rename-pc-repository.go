package repository

import (
	"fmt"
	"os/exec"
)

type RenamePcParams struct {
	NewName string
}

func RenamePC(params RenamePcParams) error {
	if params.NewName == "" {
		return fmt.Errorf("❌ Error: el nuevo nombre del PC no puede estar vacío")
	}

	fmt.Printf("🔹 Cambiando nombre del PC a: %s\n", params.NewName)

	powershellCmd := fmt.Sprintf(`Rename-Computer -NewName "%s" -Force -Restart`, params.NewName)

	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-Command", powershellCmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("❌ Error al cambiar el nombre del PC: %v\nSalida del comando: %s", err, output)
	}

	fmt.Println("✅ Cambio de nombre exitoso, reiniciando...")
	return nil
}
