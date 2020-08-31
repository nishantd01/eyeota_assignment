package main

import "log"

//DataMap to keep track of data where keys are org,paramName,paramVal & value is array of segmentIds
type DataMap struct {
	dataMap map[string]map[string]map[string][]string
}

func main() {
	log.Println("Eyeota Cache Implementation")

}
