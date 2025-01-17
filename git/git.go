package git

import (
	"fmt"
	"os/exec"
)

func GetUserName() (string, error) {
	cmd := exec.Command("git", "config", "user.name")
	d, error := cmd.Output()
	if error != nil {
		return "", fmt.Errorf("foydalanuvchi nomi olinmadi: %v", error)
	}
	return string(d), nil
}

func GetUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "user.email")
	d, error := cmd.Output()
	if error != nil {
		return "", fmt.Errorf("foydalanuvchi emaili  olinmadi: %v", error)
	}
	return string(d), nil
}
