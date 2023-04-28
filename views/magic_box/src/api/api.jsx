import { Path, Post, Fetch } from './config'

export default class Api {
  Login = (username, password, callback = (data, err) => {}) => {
    Post(Path.Login, { username, password }).then(data => {
      if (typeof (data.access_token) !== 'undefined' && data.access_token !== '') {
        Fetch.SetCommonHeader('AccessToken', data.access_token)
      }
      // 登陆成功
      callback(data, null)
    }).catch(err => {
      callback(null, err)
    })
  }

  RefreshAccessToken = (oldAccessToken, callback = (data, err) => {}) => {
    Post(Path.RefreshAccessToken, { access_token: oldAccessToken }).then(data => {
      if (typeof (data.access_token) !== 'undefined' && data.access_token !== '') {
        Fetch.SetCommonHeader('AccessToken', data.access_token)
      }
      callback(data, null)
    }).catch(err => {
      callback(null, err)
    })
  }
}
