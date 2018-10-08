package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// GetGladiusBase - Returns the base directory
func GetGladiusBase() (string, error) {
	var m string
	var err error

	cmdArg := ""

	if len(os.Args) > 1 {
		if os.Args[1] != "start" && os.Args[1] != "stop" && os.Args[1] != "install" && os.Args[1] != "uninstall" {
			cmdArg = os.Args[1]
		}
	}

	if cmdArg != "" {
		m = cmdArg
	} else if os.Getenv("GLADIUSBASE") != "" {
		m = os.Getenv("GLADIUSBASE")
	} else {
		switch runtime.GOOS {
		case "windows":
			m = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".gladius")
		case "linux":
			m = os.Getenv("HOME") + "/.gladius"
		case "darwin":
			m = os.Getenv("HOME") + "/.gladius"
		default:
			m = ""
			err = errors.New("unknown operating system, can't find gladius base directory. Set the GLADIUSBASE environment variable, or supply the directory as the first argument to add it manually")
		}
	}

	return m, err
}
