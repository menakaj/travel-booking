package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getPet(empId int32) (*Pet, error) {

	fmt.Printf("EndpointUrl %s", os.Getenv("PETSTORE_ENDPOINT_URL"))

	requestUrl := fmt.Sprintf("%s/pet/findByStatus?status=available", os.Getenv("PETSTORE_ENDPOINT_URL"))

	fmt.Println(requestUrl)

	getEmp, _ := http.NewRequest("GET", requestUrl, nil)
	getEmp.Header.Add("Content-Type", "application/json")

	empResp, e := http.DefaultClient.Do(getEmp)

	if e != nil {
		fmt.Println("error while getting employee details", e)
		return nil, fmt.Errorf("error while getting employee details %v", e)
	}

	if empResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("employee not found")
	}

	emp := &Pet{}

	body, _ := io.ReadAll(empResp.Body)

	fmt.Println(string(body))

	json.Unmarshal(body, &emp)
	return emp, nil
}
