package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"os"
	"strings"

	log "github.com/inconshreveable/log15"
)

func main() {
	log.Info("Starting simulation...")

	config := getMonitorConfig("monitor.json")

	variables := config["variables"].([]interface{})
	contractAddress := config["contractAddress"].(string)

	for _, element := range variables {
		element := element.(map[string]interface{})
		variableSnap, err := getVariable(element, contractAddress)
		if err != nil {
			continue
		}
		snap := variableSnap.(map[string]interface{})
		vs := snap["data"].([]interface{})
		vs1 := vs[0].(map[string]interface{})

		log.Info("Indexed variable", "name", vs1["variableName"], "type", vs1["variableType"], "value", vs1["variableValue"])
	}
}

func getMonitorConfig(filename string) map[string]interface{} {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Error("Error in reading file", "err", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config map[string]interface{}
	json.Unmarshal([]byte(byteValue), &config)

	return config
}

func getVariable(variableData map[string]interface{}, contractAddress string) (interface{}, error) {
	var key = ""
	var deep_key = ""
	var struct_var = ""
	if variableData["key"] != nil {
		key = variableData["key"].(string)
	}
	if variableData["deep_key"] != nil {
		deep_key = variableData["deep_key"].(string)
	}
	if variableData["struct_var"] != nil {
		struct_var = variableData["struct_var"].(string)
	}

	payload := strings.NewReader(`{
		"contractAddress": "` + fmt.Sprintf("%+v", contractAddress) + `",
		"targetVariable": "` + fmt.Sprintf("%+v", variableData["name"]) + `",
		"key": "` + fmt.Sprintf("%+v", key) + `",
		"deepKey": "` + fmt.Sprintf("%+v", deep_key) + `",
		"structVar": "` + fmt.Sprintf("%+v", struct_var) + `"
	}`)
	get_variable_route := "/getVariable"
	resp, err := makePostRequest(get_variable_route, payload)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Error("Error in json unmarshal", "err", err)
		return nil, err
	}

	return data, err
}

func makePostRequest(route string, payload *strings.Reader) ([]byte, error) {
	baseUrl := "https://7fdc-2405-201-2000-d1ff-e573-8e78-4e01-6af.in.ngrok.io/queriooor"
	url := baseUrl + route
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Error("Error in making request", "err", err)
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Error("Error in making request", "err", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("Error in reading body", "err", err)
		return nil, err
	}

	return body, nil
}
