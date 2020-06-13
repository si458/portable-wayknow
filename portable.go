package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	_ "./statik"
	"github.com/kirsle/configdir"
	"github.com/rakyll/statik/fs"
)

func logStuff(mylog string) {
	if true {
		log.Println(mylog)
	}
}

func getAndWrite(statikFS http.FileSystem, getname string, writename string) {
	var thefile []byte
	var err error
	logStuff("Getting Packaged File")
	if thefile, err = fs.ReadFile(statikFS, getname); err != nil {
		log.Fatal(err)
	}
	logStuff("Writing Packaged File")
	if err = ioutil.WriteFile(writename, thefile, 0777); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var statikFS http.FileSystem
	var err error
	logStuff("Getting Wayk Roaming Folder")
	localRoaming := configdir.LocalConfig("Wayk")
	logStuff("Checking Wayk Roaming Folder")
	if _, err = os.Stat(localRoaming); os.IsNotExist(err) {
		logStuff("Wayk Roaming Folder Doesnt Exist")
		if err = os.MkdirAll(localRoaming, 0777); err != nil {
			log.Fatal(err)
		}
		logStuff("Created Wayk Roaming Folder")
	} else {
		logStuff("Wayk Roaming Folder Already Exists")
	}
	logStuff("Loading Packaged Files")
	if statikFS, err = fs.New(); err != nil {
		log.Fatal(err)
	}
	getAndWrite(statikFS, "/wayknow.exe", localRoaming+"\\wayknow.exe")
	getAndWrite(statikFS, "/branding.zip", localRoaming+"\\branding.zip")
	getAndWrite(statikFS, "/WaykNow.cfg", localRoaming+"\\WaykNow.cfg")
	logStuff("Start Portable Wayk")
	cmd := exec.Command(localRoaming + "\\wayknow.exe")
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
