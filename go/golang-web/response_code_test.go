package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		fmt.Fprintf(w, "Hello %s ", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	respose := recorder.Result()
	body, _ := io.ReadAll(respose.Body)

	fmt.Println(respose.StatusCode)
	fmt.Println(respose.Status)
	fmt.Println(string(body))

}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Made", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	respose := recorder.Result()
	body, _ := io.ReadAll(respose.Body)

	fmt.Println(respose.StatusCode)
	fmt.Println(respose.Status)
	fmt.Println(string(body))

}
