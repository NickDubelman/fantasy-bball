package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/NickDubelman/fantasy-bball/db"
	"github.com/NickDubelman/fantasy-bball/db/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	// PathLogin is the path to handle logins
	PathLogin = "/auth/login"

	// PathCallback is the path to handle the callback from OAuth backend (Google)
	PathCallback = "/auth/google/callback"

	// PathError is redirected to when the user has an auth error
	PathError = "/auth/error"

	codeRedirect = http.StatusFound
)

// GoogleUserInfo represents a response from the Google userinfo API
type GoogleUserInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Locale  string `json:"locale"`
}

// NotAuthorized is an error for when a user is not authorized to do something
type NotAuthorized struct{}

func (e NotAuthorized) Error() string {
	return "not authorized"
}

// TokenExpired is an error for when a user's access token is expired
type TokenExpired struct{}

func (e TokenExpired) Error() string {
	return "access token is expired"
}

// GoogleAuthFromConfig returns handlers that can be used for OAuth with Google
func GoogleAuthFromConfig() gin.HandlerFunc {
	configPath := os.Getenv("OAUTH_CONFIG_PATH")
	config, err := getGoogleAuthConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			switch c.Request.URL.Path {
			case PathLogin:
				// Redirect the user to Google to authenticate
				http.Redirect(
					c.Writer,
					c.Request,
					config.AuthCodeURL(
						extractPath(c.Request.URL.Query().Get("next")),
						oauth2.SetAuthURLParam("prompt", "login"),
					),
					codeRedirect,
				)

			case PathCallback:
				// User succesfully authenticated with Google
				handleOAuth2Callback(config, c)

			case PathError:
				c.String(http.StatusInternalServerError, "Error logging in")

			}
		}
	}
}

// UserFromContext takes a context and returns the UserInfo for the user making the
// request
func UserFromContext(ctx context.Context) (UserInfo, error) {
	tokenStr, err := AccessTokenFromContext(ctx)
	if err != nil {
		return UserInfo{}, err
	}

	claims := &UserInfo{}
	tkn, err := parseToken(tokenStr, claims)
	if err != nil {
		return UserInfo{}, err
	}

	now := time.Now()
	expires := time.Unix(claims.IssuedAt, 0).Add(accessTokenDuration)

	if !tkn.Valid {
		return UserInfo{}, NotAuthorized{} // Token invalid
	}

	if now.After(expires) {
		return UserInfo{}, TokenExpired{} // Token expired
	}

	return *claims, nil
}

// ContextWithAccessToken takes a context and an access token and returns a new
// context with the access token attached
func ContextWithAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKey{"accessToken"}, token)
}

// AccessTokenFromContext takes a context and returns the attached accessToken
func AccessTokenFromContext(ctx context.Context) (string, error) {
	ctxValue := ctx.Value(contextKey{"accessToken"})
	if ctxValue == nil {
		return "", NotAuthorized{}
	}

	accessToken, ok := ctxValue.(string)
	if !ok {
		return "", fmt.Errorf("Expected accessToken to have type string")
	}

	return accessToken, nil
}

// handleOAuth2Callback will be executed after the user authenticates with Google and
// consents to the scopes our app requires. We will retrieve their name, email, and
// picture from the Google userinfo endpoint. Lastly, we generate an access token and
// a refresh token for the user (both are JWTs)
func handleOAuth2Callback(cfg *oauth2.Config, ginCtx *gin.Context) {
	handleErr := func(err error) {
		log.Println(err)
		ginCtx.Redirect(http.StatusFound, PathError)
	}

	code := ginCtx.Request.URL.Query().Get("code")
	t, err := cfg.Exchange(oauth2.NoContext, code)
	if err != nil {
		handleErr(err)
		return
	}

	client := cfg.Client(oauth2.NoContext, t)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		handleErr(err)
		return
	}
	defer userinfo.Body.Close()

	data, err := ioutil.ReadAll(userinfo.Body)
	if err != nil {
		handleErr(err)
		return
	}

	userInfo := GoogleUserInfo{}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		handleErr(err)
		return
	}

	ctx := ginCtx.Request.Context()

	dbClient := db.FromContext(ctx)
	if dbClient == nil {
		err := fmt.Errorf("could not retrieve db client from context")
		handleErr(err)
		return
	}

	// Check if the user exists in our database
	var userID int

	authUser, err := dbClient.User.
		Query().
		Where(
			user.Email(userInfo.Email),
		).
		Only(ctx)

	if err != nil {
		if _, notFoundErr := err.(*db.NotFoundError); notFoundErr {
			// If user that is logging in is not yet in our db, add them
			authUser, err := dbClient.User.
				Create().
				SetName(userInfo.Name).
				SetEmail(userInfo.Email).
				SetPicture(userInfo.Picture).
				Save(ctx)
			if err != nil {
				log.Println(err)
				ginCtx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			userID = authUser.ID
		} else {
			handleErr(err)
			return
		}
	}

	// If user that is logging in already exists in our db, just update their info
	if authUser != nil {
		userID = authUser.ID
		_, err = authUser.
			Update().
			SetName(userInfo.Name).
			SetPicture(userInfo.Picture).
			SetLastActive(time.Now()).
			Save(ctx)
		if err != nil {
			handleErr(err)
			return
		}
	}

	// Generate an access token
	accessToken, err := createAccessToken(ctx, userID, userInfo)
	if err != nil {
		handleErr(err)
		return
	}

	// Generate a refresh token
	refreshToken, err := createRefreshToken(userID)
	if err != nil {
		handleErr(err)
		return
	}

	state := ginCtx.Request.URL.Query().Get("state")

	nextURL, err := url.Parse("http://localhost:3000/login-callback")
	if err != nil {
		handleErr(err)
		return
	}

	// We need to add the user's tokens to the redirect URL so that the frontend can
	// pluck it from the URL and store the tokens in localStorage
	query := nextURL.Query()
	query.Add("accessToken", accessToken)
	query.Add("refreshToken", refreshToken)
	query.Add("state", state)
	nextURL.RawQuery = query.Encode()

	http.Redirect(ginCtx.Writer, ginCtx.Request, nextURL.String(), codeRedirect)
}

// getGoogleAuthConfig attempts to read from a given oauthConfigPath and return a
// corresponding *auth2.Config. The config file should be obtained from the Google
// Developers Console's "Credentials" page
func getGoogleAuthConfig(oauthConfigPath string) (*oauth2.Config, error) {
	jsonKey, err := ioutil.ReadFile(oauthConfigPath)
	if err != nil {
		return nil, err
	}

	conf, err := google.ConfigFromJSON(jsonKey, "profile")
	if err != nil {
		return nil, err
	}

	conf.Scopes = []string{"profile", "email"} // the scopes we need for our app
	return conf, nil
}

func extractPath(next string) string {
	nextURL, err := url.Parse(next)
	if err != nil {
		return "/"
	}
	return nextURL.Path
}

type contextKey struct{ name string }
