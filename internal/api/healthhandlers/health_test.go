package healthhandlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	port int
}

// setup test server
func NewTestServer(r *chi.Mux) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func executeRequest(req *http.Request, s *http.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Handler.ServeHTTP(rr, req)

	return rr
}

func checkHelloResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestHello(t *testing.T) {
	// create new chi router
	r := chi.NewRouter()

	// mount all health routes
	MountHealthRoutes(r)

	// run test server
	s := NewTestServer(r)

	req, _ := http.NewRequest("GET", "/hello", nil)
	
	response := executeRequest(req, s)
	
	fmt.Println("RESPONSE: ", response)


	checkHelloResponseCode(t, http.StatusOK, response.Code)

	// TODO CHECK RESPONSE BODY
	//require.Equal(t, "Hello World", response.Body)
}
