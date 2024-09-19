package acidentity

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestTokenHandling(t *testing.T) {
	testUser1 := "test_user1@gmail.com"
	testUser2 := "test_user2@gmail.com"

	token1, err := CreateToken(testUser1)
	if err != nil {
		t.Fatalf("Failed to create a JWT token: %s", err)
	}
	token2, err := CreateToken(testUser2)
	if err != nil {
		t.Fatalf("Failed to create a JWT token: %s", err)
	}
	time.Sleep(1 * time.Second)

	token3, err := CreateToken(testUser1) // if the same user requests the token should be different

	if err != nil {
		t.Fatalf("Failed to create a JWT token: %s", err)
	}

	require.NotEqual(t, token1, token2)
	require.NotEqual(t, token1, token3)

	require.True(t, VerifyToken(token1))
	require.True(t, VerifyToken(token2))
	require.True(t, VerifyToken(token3))

	parsed_token1, _, _ := new(jwt.Parser).ParseUnverified(token1, jwt.MapClaims{})
	parsed_token2, _, _ := new(jwt.Parser).ParseUnverified(token2, jwt.MapClaims{})
	parsed_token3, _, _ := new(jwt.Parser).ParseUnverified(token3, jwt.MapClaims{})
	claims1, _ := parsed_token1.Claims.(jwt.MapClaims)
	claims2, _ := parsed_token2.Claims.(jwt.MapClaims)
	claims3, _ := parsed_token3.Claims.(jwt.MapClaims)

	require.Equal(t, testUser1, claims1["user_email"])
	require.Equal(t, testUser2, claims2["user_email"])
	require.Equal(t, testUser1, claims3["user_email"])

	require.True(t, VerifyToken(token1))
	require.True(t, VerifyToken(token2))
	require.True(t, VerifyToken(token3))

	require.False(t, VerifyToken("random string that should return false"))
}
