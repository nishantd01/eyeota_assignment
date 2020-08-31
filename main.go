package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func initializeCache(fileName string) DataMap {
	// Read file
	rawFile, _ := ioutil.ReadFile(fileName)
	cache := DataMap{}
	// Initialize maps
	cache.dataMap = make(map[string]map[string]map[string][]string)
	jsonDataStr := string(rawFile)
	var jsonData []interface{}
	json.Unmarshal([]byte(jsonDataStr), &jsonData)

	// Top level loop
	for i := range jsonData {
		orgMapNA := jsonData[i]
		// Type assertion
		orgMap := orgMapNA.(map[string]interface{})
		// OrgKey level loop
		for org := range orgMap {
			// Type assertion
			pNameInd := orgMap[org].([]interface{})
			for j := range pNameInd {
				// Type assertion
				pNameMap := pNameInd[j].(map[string]interface{})
				// paramName level loop
				for pName := range pNameMap {
					// Type assertion
					pValArray := pNameMap[pName].([]interface{})
					for k := range pValArray {
						// Type assertion
						pValMap := pValArray[k].(map[string]interface{})
						// paramVal level loop
						for pVal := range pValMap {
							if strings.Contains(pVal, "\n") {
								// Take care of compound keys
								pValSlice := strings.Split(pVal, "\n")
								// Type assertion
								segMap := pValMap[pVal].(map[string]interface{})
								// Segment level loop
								for _, segValNA := range segMap {
									// Type assertion
									segVal := segValNA.(string)
									if cache.dataMap[org] == nil {
										// Initialize orgKey level map
										cache.dataMap[org] = make(map[string]map[string][]string)
									}

									if cache.dataMap[org][pName] == nil {
										// Initialize parameterName level map
										cache.dataMap[org][pName] = make(map[string][]string)
									}

									for _, pv := range pValSlice {
										// Split compound keys and reassign values
										cache.dataMap[org][pName][pv] = append(cache.dataMap[org][pName][pv], segVal)
									}
								}
							} else {
								segMap := pValMap[pVal].(map[string]interface{})
								// Segment level loop
								for _, segValNA := range segMap {
									segVal := segValNA.(string)
									if cache.dataMap[org] == nil {
										// Initilize orgKey level map
										cache.dataMap[org] = make(map[string]map[string][]string)
									}

									if cache.dataMap[org][pName] == nil {
										// Initialize parameterName level map
										cache.dataMap[org][pName] = make(map[string][]string)
									}
									// Assign for non-compound keys
									cache.dataMap[org][pName][pVal] = append(cache.dataMap[org][pName][pVal], segVal)
								}
							}
						}
					}
				}
			}
		}
	}

	return cache

}

func (cache DataMap) getSegment(w http.ResponseWriter, r *http.Request) {
	var requestParams RequestParams

	err := json.NewDecoder(r.Body).Decode(&requestParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestParams.OrgName == "" || requestParams.ParamName == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if requestParams.ParamVal == "" {
		//invoke GetSegmentForOrgAndKey
		ids := cache.GetSegmentForOrgAndKey(requestParams.OrgName, requestParams.ParamName)
		json.NewEncoder(w).Encode(ids)

	} else {
		//invoke GetSegmentForOrgAndKeyAndVal
		ids := cache.GetSegmentForOrgAndKeyAndVal(requestParams.OrgName, requestParams.ParamName, requestParams.ParamVal)
		json.NewEncoder(w).Encode(ids)
	}

}

func main() {
	log.Println("Eyeota Cache Implementation")

	fileName := "./data/data.json"

	cacheData := initializeCache(fileName)

	http.HandleFunc("/api/v1/getSegmentByQuery", cacheData.getSegment)

	http.ListenAndServe(":7020", nil)

}
