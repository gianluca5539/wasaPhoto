package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type CustomJWTClaims struct {
	UserID int `json:"userid"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Username string `json:"username"`
}

const (
	// JWTSigningKey is the key used to sign the JWT token
	JWTSigningKey = "wasaphoto_secret" // in production should be put in .env or something
)

// Function to generate a JWT token
func GenerateJWTToken(UserID int) (string, error) {
	// Create the Claims
	claims := CustomJWTClaims{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Token will expire after 72 hours
			Issuer:    "WASAPhoto",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(JWTSigningKey))
	if err != nil {
		return "", fmt.Errorf("error in GenerateJWTToken while signing token: %w", err)
	}

	return ss, nil
}


func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := loginReq.Username

	// Check if the username exists
	user,found, err := rt.db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If the user doesn't exist, create a new user in the database
	if !found {
		user, err = rt.db.CreateUser(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} 
		token, err := GenerateJWTToken(user.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := map[string]interface{}{
			"user": user,
			"token": token,
		}
		// Return the user and the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	} else {
		token, err := GenerateJWTToken(user.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error" + err.Error()))
			return
		}

		res := map[string]interface{}{
			"user": user,
			"token": token,
		}
		// Return the user and the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}; 