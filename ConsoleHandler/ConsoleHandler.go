package main

import (
	"log"
	"os"

	dh "../dataHandler"
)

var pathToInputFile string = "../input.txt"
var pathToOtputFile string = "../output.txt"

func main() {

	log := log.New(os.Stdout, "", log.Ldate)

	stringData, err := dh.ReadDataAsStringFromFile(pathToInputFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	sliceOfPoint, err := dh.ParseDataFromString(stringData)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	sliceOfPoint.Sort()
	SaveResults(sliceOfPoint.ToString(), pathToOtputFile)
}

func SaveResults(data string, pathToFile string) {
	file, err := os.Create(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(data)
}
