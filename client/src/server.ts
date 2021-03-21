import sirv from 'sirv'
import polka from 'polka'
import compression from 'compression'
import * as sapper from '@sapper/server'
import session from 'express-session'
import { createProxyMiddleware } from 'http-proxy-middleware'

const { PORT, NODE_ENV } = process.env
const dev = NODE_ENV === 'development'

const server = polka() // You can also use Express

if (dev) {
  // @ts-ignore
  server.use(
    '/api',
    createProxyMiddleware('/api', { target: 'http://localhost:8080' })
  )

  // @ts-ignore
  server.use(
    '/auth',
    createProxyMiddleware('/auth', { target: 'http://localhost:8080' })
  )
}

server
  .use(
    compression({ threshold: 0 }),
    sirv('static', { dev }),
    session({
      secret: process.env.ACCESS_SECRET || 'Go Lakers',
      resave: false,
      saveUninitialized: false,
    }),
    (req, res, next) => {
      return sapper.middleware({
        // @ts-ignore
        session: () => ({ user: !!req.session.accessToken }),
      })(req, res, next)
    }
  )
  .listen(PORT, (err) => {
    if (err) console.log('error', err)
  })
