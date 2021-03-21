# fantasy-bball

Fantasy basketball with daily drafts. Create leagues and invite friends.

## Auth

1. Sapper client sends _HTTP only cookies_ to sapper server
1. Sapper server uses [express-session](https://github.com/expressjs/session) to associate cookies with user sessions, which contain the access token and refresh token (_JWTs_) for the given user. Sapper server sends these tokens to the API as necessary.
