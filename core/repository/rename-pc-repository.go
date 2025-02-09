package repository

import (
	"fmt"
	"os/exec"
)

type RenamePcParams struct {
	NewName       string
	UserAdmin     string
	PasswordAdmin string
}

func RenamePC(params RenamePcParams) error {

	powershellCmd := fmt.Sprintf(`$password = ConvertTo-SecureString "%s" -AsPlainText -Force; 
    $credential = New-Object System.Management.Automation.PSCredential ("%s", $password);
    Rename-Computer -NewName "%s" -Credential $credential -Force -Restart`, params.PasswordAdmin, params.UserAdmin, params.NewName)

	cmd := exec.Command("powershell", "-Command", powershellCmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al cambiar el nombre del PC: %v\nSalida del comando: %s", err, output)
	}
	return nil
}
