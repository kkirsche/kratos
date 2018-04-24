package logs

import "fmt"

// PrintlnError is used to print an error message to stdout
func PrintlnError(msg string, err error) (int, error) {
	ln := fmt.Sprintf("[!] %s with error: %s", msg, err.Error())
	return fmt.Println(ln)
}

// PrintlnInfo is used to print an informational message to stdout
func PrintlnInfo(msg string) (int, error) {
	ln := fmt.Sprintf("[*] %s", msg)
	return fmt.Println(ln)
}

// PrintlnDebug is used to print a debug message to stdout
func PrintlnDebug(msg string) (int, error) {
	ln := fmt.Sprintf("[+] %s", msg)
	return fmt.Println(ln)
}
