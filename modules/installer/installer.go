package installer

import (
	"fmt"
	"os"
	"os/user"
)

type Installer struct {
	Name              string
	Description       string
	Source            string
	ZshrcAlias        string
	InstallFunction   func() (bool, error)
	UninstallFunction func() (bool, error)
	CheckFunction     func() (bool, error)
}

func EnsureSgoinfre() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	if _, err := os.Stat(fmt.Sprintf("Users/%s/sgoinfre/students/%s", usr, usr)); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("Sgoninfre not set up, creating it now"))
		err := os.Mkdir(fmt.Sprintf("Users/%s/sgoinfre/students/%s", usr, usr), 0700)
		if err != nil {
			return err
		}
	}
	return nil
}

var Installers = []Installer{
	{
		Name:        "Francinette",
		Description: "A tester for 42 projects",
		Source:      "https://github.com/xicodomingues/francinette",
		ZshrcAlias: `
		alias francinette="$SGOINFRE/42manager/paco/tester.sh"
		alias paco="$SGOINFRE/42manager/paco/tester.sh"`,
		InstallFunction:   InstallPaco,
		UninstallFunction: UninstallPaco,
		CheckFunction:     CheckPaco,
	},
}
