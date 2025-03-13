package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getEmployee(empId int32) (*Employee, error) {
	// accessToken, tokenErr := GetToken()

	// if tokenErr != nil {
	// 	return nil, tokenErr
	// }

	choreoApiKey := os.Getenv("CHOREO_CONN2_APIKEY")
	requestUrl := fmt.Sprintf("%s/employees/%d", os.Getenv("CHOREO_CONN2_SERVICEURL"), empId)
	fmt.Println("sending request to", requestUrl)

	getEmp, _ := http.NewRequest("GET", requestUrl, nil)
	getEmp.Header.Add("Choreo-API-Key", choreoApiKey)

	empResp, e := http.DefaultClient.Do(getEmp)

	if e != nil {
		fmt.Println(e.Error())
		return nil, fmt.Errorf("error while getting employee details")
	}

	if empResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("employee not found")
	}

	emp := &Employee{}

	body, _ := io.ReadAll(empResp.Body)

	fmt.Println(string(body))

	json.Unmarshal(body, &emp)
	return emp, nil
}
