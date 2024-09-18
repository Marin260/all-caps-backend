package acidentity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateJWT(t *testing.T) {
	testUser1 := "test_user1@gmail.com"
	testUser2 := "test_user2@gmail.com"


	token1, err := CreateJWT(testUser1)
	if err != nil {
		t.Fatalf("Failed to create a JWT token: %s", err)
	}
	token2, err := CreateJWT(testUser2)
	if err != nil {
		t.Fatalf("Failed to create a JWT token: %s", err)
	}
	time.Sleep(1 * time.Second)

	token3, err := CreateJWT(testUser1)
	
	if err != nil {
		t.Fatalf("Failed to create a JWT token: %s", err)
	}

	require.NotEqual(t, token1, token2)
	require.NotEqual(t, token1, token3)
}
