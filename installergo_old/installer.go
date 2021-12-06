package main

import (
	instmemory "fhir-isih/installer/pckinst"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const strInfo string = "Hello installer\nnouvelle ligne"
const str_install string = "apt-get update\napt-get -y upgrade\napt-get -y install docker-compose"

func main() {

	// tester version du système
	//doit être une ubuntu version >= 20.04 LTS
	instmemory.MemoryShow()

	if !instmemory.Arch_test("linux") {
		fmt.Println("Le système ne correspond pas a un système Linux Compatible...")
		os.Exit(1)
	}

	//installer docker
	writetodisk(str_install,"i_inst")
	running("/bin/sh","./i_inst")
	erasetodisk("i_inst")

	//download script docker
	writetodisk(download("https://get.docker.com"), "d_inst")
	running("/bin/sh", "./d_inst")
	erasetodisk("d_inst")

	//afficher mode de lancement...
	fmt.Println(strInfo)
}

/***********************
	Supprimer le fichier
************************/
func erasetodisk(s string) {
	err := os.Remove(s)
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	} else {
		log.Printf("Suppress file %s\n", s)
	}
}

/******************************
	Lancer une commande...
*******************************/
func running(commande string, arguments string) {
	cmd := exec.Command(commande, arguments)
	log.Printf("Running command %s %s and waiting for it to finish...",commande, arguments)
	err := cmd.Run()
	if err == nil {
		log.Printf("Command finished with no error")
	} else {
		log.Printf("Command finished with error: %v", err)
		os.Exit(2) //pas nécessaire de continuer...
	}
}

/*******************************************
	Download package from
********************************************/
func download(url string) string {

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		//convertir byArray en string
		return string(contents[:])
	}
	return ""
}

/**************************************
	Ecrire fichier sur disque...
**************************************/
func writetodisk(contenu string, destination string) {
	err := ioutil.WriteFile(destination, []byte(contenu), 0644)
	if err != nil {
		log.Printf("Error writing file on local Disk...")
	}
	log.Printf("File %s write on disk\n", destination)
}
