package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const str_stop string = "docker-compose -f dc.d down"
const str_docker_compose string = "version: '3.5'\nservices:\n    hapi-fhir-server:\n      image: hopsiia/fhir-isih:latest\n      container_name: hapi-fhir-server\n      hostname: hapi-fhir-server\n      restart: on-failure\n      ports:\n        - '8181:8080'\n    hapi-fhir-pg:\n      image: postgres:12\n      container_name: hapi-fhir-pg\n      hostname: hapi-db\n      restart: always\n      environment:\n        - POSTGRES_DB=hapi_r4\n        - POSTGRES_USER=postgres\n        - POSTGRES_PASSWORD=admin\n      command: postgres -c 'max_connections=200'\n      ports:\n        - '5432:5432'\n      volumes:\n       - './pgdatafhir:/var/lib/postgresql/data'\nnetworks:\n    hapi-bridge:\n      name: hapi-network"

func main() {
	message()
	writetodisk(str_stop, "s.d")
	writetodisk(str_docker_compose, "dc.d")
<<<<<<< HEAD
	//running("/bin/sh", "./s.d")
	erasetodisk("r.d")
=======
	running("/bin/sh", "./s.d")
>>>>>>> f7816b49352055c732417857c33f2af640b2dd24
	erasetodisk("s.d")
	erasetodisk("dc.d")
}

/******************************
	Lancer une commande...
*******************************/
func running(commande string, arguments string) {
	cmd := exec.Command(commande, arguments)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)
	cmd.Wait()
}

/***********************
	Ecrire sur la console
************************/
func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func message() {
	var str_message string = `
==============================
 Arret du serveur FHIR-ISIH
==============================`

	fmt.Println(str_message)
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
