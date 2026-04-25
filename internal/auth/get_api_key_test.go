package auth

import(
	"net/http"
	"testing"
)

func TestEmptyHeader(t *testing.T) {
    headers := http.Header{}
    
    _, err := GetAPIKey(headers)
    
    if err != ErrNoAuthHeaderIncluded {
        t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
    }
}
