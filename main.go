package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	monitor_config := getMonitorConfigs("monitor.json")

	variables_to_index := monitor_config["variables"].([]interface{})
	contractAddress := monitor_config["contractAddress"].(string)
	for _, element := range variables_to_index {
		element := element.(map[string]interface{})
		variable_snap := getVariable(element, contractAddress).(map[string]interface{})
		vs := variable_snap["data"].([]interface{})
		vs1 := vs[0].(map[string]interface{})
		log.Println(vs1["variableName"], vs1["variableType"], vs1["variableValue"])
	}
}

func getMonitorConfigs(filename string) map[string]interface{} {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var monitor_config map[string]interface{}
	json.Unmarshal([]byte(byteValue), &monitor_config)

	return monitor_config
}

func getVariable(variable_data map[string]interface{}, contractAddress string) interface{} {
	var key = ""
	var deep_key = ""
	var struct_var = ""
	if variable_data["key"] != nil {
		key = variable_data["key"].(string)
	}
	if variable_data["deep_key"] != nil {
		deep_key = variable_data["deep_key"].(string)
	}
	if variable_data["struct_var"] != nil {
		struct_var = variable_data["struct_var"].(string)
	}

	payload := strings.NewReader(`{
	"contractAddress": "` + fmt.Sprintf("%+v", contractAddress) + `",
	"targetVariable": "` + fmt.Sprintf("%+v", variable_data["name"]) + `",
	"key": "` + fmt.Sprintf("%+v", key) + `",
	"deepKey": "` + fmt.Sprintf("%+v", deep_key) + `",
	"structVar": "` + fmt.Sprintf("%+v", struct_var) + `"
	}`)
	get_variable_route := "/getVariable"
	resp := makePostRequest(get_variable_route, payload)

	var data interface{}
	err := json.Unmarshal(resp, &data)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func makePostRequest(route string, payload *strings.Reader) []byte {
	baseUrl := "https://7fdc-2405-201-2000-d1ff-e573-8e78-4e01-6af.in.ngrok.io/queriooor"
	url := baseUrl + route
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return body
}
