import FetchUtil from '../util/fetch'
import MyError from './error'

const ServerAddr = 'http://localhost'

export const Fetch = new FetchUtil(ServerAddr)

export function Post (path, data) {
  const res = Fetch.Post(path, data).then(resp => {
    if (!resp.ok) {
      throw new MyError('server error')
    }
    if (resp.headers && resp.headers.get('Content-Type').startsWith('application/json')) {
      return resp.json()
    }
    throw new MyError('unsupported response content-type')
  }).then(data => {
    /**
     * {
     *     "trace_id":"",
     *     "code":"ok",
     *     "msg":"",
     *     "data":{}
     * }
     */
    if (data.code.toLowerCase() !== 'ok') {
      throw new MyError(data.msg, data.code)
    }
    return data.data
  })
  return res
}

export const Path = {
  Login: '/user/login',
  RefreshAccessToken: '/user/refresh_access_token'
}
