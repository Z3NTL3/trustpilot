package client_test

import (
	"testing"

	"github.com/z3ntl3/trustpilot/client"
)

// go test -timeout 30s -run ^TestCient$ github.com/z3ntl3/trustpilot/tests -v
func TestCient(t *testing.T) {
	client := client.New()

	result, err := client.Execute("https://trustpilot.com/review/www.ikea.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Resulting date model: %+v", result)
}
/*
~\Documents\trustpilot via 🐹 v1.22.2 
❯ go  test -timeout 30s -run ^TestCient$ github.com/z3ntl3/trustpilot/tests -v
=== RUN   TestCient
    client_test.go:17: Resulting date model: &{Company:IKEA  Reviews:{Total:1.5 Text:24,080   •   Bad} ImageSRC:https://cdn.trustpilot.net/brand-assets/4.1.0/stars/stars-1.5.svg}
--- PASS: TestCient (0.39s)
PASS
ok      github.com/z3ntl3/trustpilot/tests      
*/