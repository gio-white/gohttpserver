package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	// Setup test data
	secret := "my-super-secret-key-12345"
	userID := uuid.New()
	duration := time.Hour

	// Case 1: Valid JWT creation and validation
	t.Run("Valid JWT", func(t *testing.T) {
		token, err := MakeJWT(userID, secret, duration)
		if err != nil {
			t.Fatalf("failed to make JWT: %v", err)
		}

		returnedID, err := ValidateJWT(token, secret)
		if err != nil {
			t.Fatalf("failed to validate valid JWT: %v", err)
		}

		if returnedID != userID {
			t.Errorf("expected ID %v, got %v", userID, returnedID)
		}
	})

	// Case 2: Expired JWT should fail
	t.Run("Expired JWT", func(t *testing.T) {
		// Create a token that expired 1 hour ago
		expiredDuration := -time.Hour
		token, err := MakeJWT(userID, secret, expiredDuration)
		if err != nil {
			t.Fatalf("failed to make JWT: %v", err)
		}

		_, err = ValidateJWT(token, secret)
		if err == nil {
			t.Error("expected error for expired token, but got nil")
		}
	})

	// Case 3: Wrong secret should fail
	t.Run("Wrong Secret", func(t *testing.T) {
		token, err := MakeJWT(userID, secret, duration)
		if err != nil {
			t.Fatalf("failed to make JWT: %v", err)
		}

		wrongSecret := "definitely-the-wrong-secret"
		_, err = ValidateJWT(token, wrongSecret)
		if err == nil {
			t.Error("expected error for wrong secret, but got nil")
		}
	})
}