package rest

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/mtdx/keyc/common"
)

var ts *httptest.Server
var body, expected string

func TestMain(m *testing.M) {
	r := StartRouter()
	ts = httptest.NewServer(r)
	defer ts.Close()

	code := m.Run()
	os.Exit(code)
}

func TestHome(t *testing.T) {
	t.Parallel()

	_, body = common.TestRequest(t, ts, "GET", "/", nil)
	expected = `{"elapsed":0}`
	if strings.Compare(strings.TrimSpace(body), expected) != 0 {
		t.Fatalf("expected:%s got:%s", expected, body)
	}
}
