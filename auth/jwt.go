package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	accessTokenDuration  = 5 * time.Minute
	refreshTokenDuration = 7 * 24 * time.Hour // 7 days
)

// UserInfo represents the claims contained within an access token
type UserInfo struct {
	GoogleUserInfo
	jwt.StandardClaims
}

// ID converts the user's subject claim to an int
func (ui UserInfo) ID() int {
	id, err := strconv.Atoi(ui.Subject)
	if err != nil {
		return 0
	}
	return id
}

// refreshClaims represents the claims contained within a refresh token
type refreshClaims struct {
	jwt.StandardClaims
}

func init() {
	// FIXME: this should be a secret
	os.Setenv("ACCESS_SECRET", "my name a borat! very nice!")
}

// RefreshAccessToken takes a user's accessToken and their refreshToken and generates
// a new accessToken for the user
func RefreshAccessToken(
	ctx context.Context,
	accessToken string,
	refreshToken string,
) (string, error) {
	// Parse the access token
	aClaims := &UserInfo{}
	tkn, err := parseToken(accessToken, aClaims)
	if err != nil {
		return "", err
	}

	// Ensure the access token is valid
	if !tkn.Valid {
		return "", NotAuthorized{}
	}

	// Parse the refresh token
	rClaims := &refreshClaims{}
	tkn, err = parseToken(refreshToken, rClaims)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expires := time.Unix(rClaims.IssuedAt, 0).Add(refreshTokenDuration)

	// Ensure the refresh token is valid
	if !tkn.Valid {
		return "", NotAuthorized{}
	}

	// Ensure the refresh token isnt expired
	if now.After(expires) {
		return "", TokenExpired{}
	}

	userID, err := strconv.Atoi(aClaims.Subject)
	if err != nil {
		return "", err
	}

	// Generate a new token with the same claims
	return createAccessToken(ctx, userID, aClaims.GoogleUserInfo)
}

func parseToken(tokenStr string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		},
	)
}

// createAccessToken takes some userInfo and creates a JWT access token
func createAccessToken(
	ctx context.Context,
	userID int,
	userInfo GoogleUserInfo,
) (string, error) {
	claims := jwt.MapClaims{
		// Reserved claims
		"sub": strconv.Itoa(userID),
		"iat": time.Now().Unix(),

		// Our claims
		"name":    userInfo.Name,
		"email":   userInfo.Email,
		"picture": userInfo.Picture,
	}

	return jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

// createRefreshToken takes a userID and creates a JWT refresh token
func createRefreshToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		// Reserved claims
		"sub": strconv.Itoa(userID),
		"iat": time.Now().Unix(),
	}

	return jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}
