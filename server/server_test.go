package server_test

// TODO: review the http testing
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mch1307/gomotics/nhc"
	. "github.com/mch1307/gomotics/server"
	"github.com/mch1307/gomotics/testutil"
	"github.com/mch1307/gomotics/types"
)

var baseUrl string

//const healthMsg = `{"alive":true}`

func init() {
	baseUrl = "http://" + testutil.ConnectHost + ":8081"
	testutil.InitStubNHC()
}

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", baseUrl+"/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)
	handler.ServeHTTP(rr, req)
	if rr.Body.String() != HealthMsg {
		t.Errorf("health test failed: got %v, expect: %v", rr.Body.String(), HealthMsg)
	}

}

func Test_getNhcItem(t *testing.T) {
	req, err := http.NewRequest("GET", baseUrl+"/api/v1/nhc/99", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(nhc.GetNhcItem)
	handler.ServeHTTP(rr, req)
	expected := "light"
	var res types.Item
	json.Unmarshal(rr.Body.Bytes(), &res)
	if res.Name != expected {
		t.Errorf("getNhcItem failed: got %v, expect: %v", res, expected)
	}
}

func Test_getNhcItems(t *testing.T) {
	req, err := http.NewRequest("GET", baseUrl+"/api/v1/nhc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(nhc.GetNhcItems)
	handler.ServeHTTP(rr, req)
	var found bool
	expected := "light"
	var res []types.Item
	json.Unmarshal(rr.Body.Bytes(), &res)
	for _, val := range res {
		if val.ID == 0 {
			if val.Name == expected {
				found = true
			}
		}
	}
	if !found {
		t.Error("GetNhcItems failed, expected light record not found")
	}
}

func Test_nhcCmd(t *testing.T) {
	expected := "Success"
	url := baseUrl + "/api/v1/nhc/action?id=1&value=100"
	hCli := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	//	req.Header.Set("User-Agent", "Test_nhcCmd")
	rsp, getErr := hCli.Do(req)
	if getErr != nil {
		fmt.Println(err)
	}
	got, readErr := ioutil.ReadAll(rsp.Body)
	if readErr != nil {
		fmt.Println("Read err: ", readErr)
	}
	defer rsp.Body.Close()
	if string(got) != expected {
		t.Errorf("Test_nhcCmd failed, expecting %v, got %v", expected, string(got))
	}

}
