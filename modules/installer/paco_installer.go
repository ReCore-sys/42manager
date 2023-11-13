package installer

import (
	"fmt"
	"os"
	"os/user"
)

var InstallScript string = `
#!/bin/bash

cd "$SGOINFRE" || exit

mkdir temp_____

cd temp_____ || exit
rm -rf francinette

# download github
git clone --recursive https://github.com/xicodomingues/francinette.git

cp -r francinette "$SGOINFRE"/42manager/paco

cd "$SGOINFRE" || exit
rm -rf temp_____

cd "$SGOINFRE"/42manager/paco || exit

# start a venv inside francinette
if ! python3 -m venv venv ; then
	echo "Please make sure than you can create a python virtual environment"
	echo 'Contact me if you have no idea how to proceed (fsoares- on slack)'
	exit 1
fi

# activate venv
. venv/bin/activate

# install requirements
if ! pip3 install -r requirements.txt ; then
	echo 'Problem launching the installer. Contact me (fsoares- on slack)'
	exit 1
fi


# print help
"$SGOINFRE"/42manager/paco/tester.sh --help

`

func CheckPaco() (bool, error) {
	usr, err := user.Current()
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(fmt.Sprintf("Users/%s/sgoinfre/students/%s/42manager/paco", usr, usr)); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func InstallPaco() (bool, error) {
	usr, err := user.Current()
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(fmt.Sprintf("Users/%s/sgoinfre/students/%s/42manager/paco", usr, usr)); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("Installing Paco"))
		err := os.Mkdir(fmt.Sprintf("Users/%s/sgoinfre/students/%s/42manager/paco", usr, usr), 0700)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func UninstallPaco() (bool, error) {
	usr, err := user.Current()
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(fmt.Sprintf("Users/%s/sgoinfre/students/%s/42manager/paco", usr, usr)); os.IsNotExist(err) {
		return false, nil
	}
	err = os.RemoveAll(fmt.Sprintf("Users/%s/sgoinfre/students/%s/42manager/paco", usr, usr))
	return true, err
}
