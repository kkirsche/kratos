package management

import (
	"os"
	"syscall"

	"github.com/kkirsche/kratos/utilities/logs"
)

// DeleteKratos is used to allow kratos to delete itself on different hosts
func DeleteKratos() {
	executable, err := os.Executable()
	if err != nil {
		logs.PrintlnError("failed to locate kratos executable", err)
		return
	}
	err = syscall.Unlink(executable)
	if err != nil {
		logs.PrintlnError("failed to delete kratos", err)
		return
	}
}
