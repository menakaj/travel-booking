package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getEmployee(empId int32) (*Employee, error) {
	accessToken, tokenErr := GetToken()

	if tokenErr != nil {
 		return nil, tokenErr
	}

	choreoApiKey := os.Getenv("CHOREO_CONN2_CHOREOAPIKEY")
	requestUrl := fmt.Sprintf("%s/employees/%d", os.Getenv("HR_SERVICE_URL"), empId)
	fmt.Println("sending request to", requestUrl)

	getEmp, _ := http.NewRequest("GET", requestUrl, nil)
	getEmp.Header.Add("Choreo-API-Key", choreoApiKey)
	getEmp.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	empResp, e := http.DefaultClient.Do(getEmp)

	if e != nil {
		fmt.Println(e.Error())
		return nil, fmt.Errorf("error while getting employee details")
	}

	fmt.Println("response is not error")

	fmt.Printf("response status code %d", empResp.StatusCode)

	if empResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("employee not found")
	}

	emp := &Employee{}

	body, _ := io.ReadAll(empResp.Body)

	fmt.Println(string(body))

	json.Unmarshal(body, &emp)
	return emp, nil
}
