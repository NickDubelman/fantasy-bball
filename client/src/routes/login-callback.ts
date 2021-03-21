import type polka from 'polka'
import type http from 'http'
import type session from 'express-session'

interface UserSession extends session.Session {
  accessToken?: string
  refreshToken?: string
}

export async function get(
  request: polka.Request & { session: UserSession },
  response: http.ServerResponse,
  next: () => void
) {
  const params = new URLSearchParams(request.search)

  request.session.accessToken = params.get('accessToken')
  request.session.refreshToken = params.get('refreshToken')
  const nextURLEncoded = params.get('state')

  // Redirect user to where they were originally trying to go
  response.writeHead(301, {
    Location: Buffer.from(nextURLEncoded, 'base64').toString(),
  })
  response.end()
}
