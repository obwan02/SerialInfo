package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func finishResponse(out http.ResponseWriter, result map[string]interface{}, status int) {
	jsonOut := json.NewEncoder(out)
	out.WriteHeader(status)
	jsonOut.Encode(result)
	fmt.Println("")
}

//StartServer func
func StartServer(port string) error {
	http.HandleFunc("/", func(out http.ResponseWriter, req *http.Request) {
		out.Header().Set("Content-Type", "application/json")
		status := http.StatusOK

		fmt.Printf("New request from %s\n", req.Host)
		result := make(map[string]interface{})
		defer finishResponse(out, result, status)

		if req.Method != "POST" {
			fmt.Println("ERROR: Recieved request that isn't POST")
			status = http.StatusMethodNotAllowed
			return
		}

		decoder := json.NewDecoder(req.Body)
		serials := make([]string, 1)
		err := decoder.Decode(&serials)
		if err != nil {
			fmt.Println("ERROR: Invalid JSON in serials param")
			fmt.Printf("ERROR: %s", err.Error())
			status = http.StatusBadRequest
			return
		}

		if serials == nil {
			status = http.StatusBadRequest
			fmt.Println("ERROR: Recieved request with invalid serials param")
			return
		}

		fmt.Printf("Looking up %d serials\n", len(serials))
		for _, item := range serials {
			var err error
			fmt.Printf("Looking up serial %s\n", item)
			result[item], err = GetSerialInfo(item)
			if err != nil {
				result[item] = "ERR"
			}
		}
	})

	return http.ListenAndServe(":"+port, nil)
}
