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

	requestUrl := fmt.Sprintf("%s/employees/%d", os.Getenv("SERVICE_URL"), empId)

	getEmp, _ := http.NewRequest("GET", requestUrl, nil)
	getEmp.Header.Add("Authorization", "Bearer "+accessToken)

	empResp, _ := http.DefaultClient.Do(getEmp)

	if empResp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("employee not found")
	}

	emp := &Employee{}

	body, _ := io.ReadAll(empResp.Body)

	fmt.Println(string(body))

	json.Unmarshal(body, &emp)
	return emp, nil
}
