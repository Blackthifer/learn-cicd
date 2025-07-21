package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T){
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded{
		t.Errorf("Expected %s, got %s", ErrNoAuthHeaderIncluded.Error(), err.Error())
	}
}