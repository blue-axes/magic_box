import Api from '../api/api'
import MyError from '../api/error'
import Message from './message'
import { LoginUrlPath } from '../router'

const keyAccessToken = 'access_token'
const keyExpireAt = 'access_token_expire_at'

export const SaveLoginState = (accessToken, expireAt) => {
  localStorage.setItem(keyAccessToken, accessToken)
  localStorage.setItem(keyExpireAt, expireAt)
}

export const IsLogin = () => {
  const accessToken = localStorage.getItem(keyAccessToken)
  let expireAt = localStorage.getItem(keyExpireAt)
  if (expireAt !== '') {
    expireAt = new Date(expireAt)
  }

  const now = new Date()
  if (accessToken !== '' && expireAt.getTime() > 0) { // 之前登陆过
    if (expireAt.getTime() > now.getTime() && expireAt.getTime() - now.getTime() < 10) { // 还有10s过期
      console.log('need refresh access_token')
      // 刷新access_token
      refreshAccessToken(accessToken)
    } else if (now.getTime() > expireAt.getTime()) { // 已经过期
      if (window.location.pathname !== LoginUrlPath) {
        Message.Error('session timeout. please login first.')
      }
      return false
    } else { // 没有过期
      return true
    }
  } else { // 之前没有登陆过
    return false
  }
}

const refreshAccessToken = (oldAccessToken) => {
  const api = new Api()
  api.RefreshAccessToken(oldAccessToken, (data, err) => {
    if (typeof err !== 'undefined' && err instanceof MyError) {
      Message.Error(err.ToString)
    } else {
      localStorage.setItem('access_key', data.access_token)
      const now = new Date()
      const end = new Date(now.getTime() + data.expire_in * 1000)
      localStorage.setItem('access_key_expire_at', end)
    }
  })
}
