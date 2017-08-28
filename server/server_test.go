package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/testutil"
)

//const healthMsg = `{"alive":true}`

func init() {
	testutil.PopFakeNhc()
}

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)
	handler.ServeHTTP(rr, req)
	if rr.Body.String() != healthMsg {
		t.Errorf("health test failed: got %v, expect: %v", rr.Body.String(), healthMsg)
	}

}

func Test_getNhcItem(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/nhc/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getNhcItem)
	handler.ServeHTTP(rr, req)
	expected := "light"
	var res nhc.Item
	json.Unmarshal(rr.Body.Bytes(), &res)
	if res.Name != expected {
		t.Errorf("getNhcItem failed: got %v, expect: %v", res, expected)
	}
}

func Test_getNhcItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/nhc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getNhcItems)
	handler.ServeHTTP(rr, req)
	var found bool
	expected := "light"
	var res []nhc.Item
	json.Unmarshal(rr.Body.Bytes(), &res)
	for _, val := range res {
		if val.ID == 0 {
			if val.Name == expected {
				found = true
			}
		}
	}
	if !found {
		t.Error("getNhcItems failed, expected light record not found")
	}
}
