package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getPet(empId int32) (*Pet, error) {
	accessToken, tokenErr := GetToken()

	if tokenErr != nil {
		return nil, tokenErr
	}

	requestUrl := fmt.Sprintf("%s/pets/%d", os.Getenv("PETSTORE_ENDPOINT_URL"), empId)

	getEmp, _ := http.NewRequest("GET", requestUrl, nil)
	getEmp.Header.Add("Authorization", "Bearer "+accessToken)

	empResp, e := http.DefaultClient.Do(getEmp)

	if e != nil {
		return nil, fmt.Errorf("error while getting employee details")
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
