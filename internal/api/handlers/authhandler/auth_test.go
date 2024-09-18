package authhandler

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Auth routes not tested, I will blindly believe that google knows how to implement auth
// TODO: implement this one day
func TestGetFrontendURL(t *testing.T) {

	expectedURL := "http://localhost:4200/"
	dir, _ := os.Getwd()

	fmt.Println(dir)

	requestdURL := getFrontendURL()
	require.Equal(t, expectedURL, requestdURL)
}
