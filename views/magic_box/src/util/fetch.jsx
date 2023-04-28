class FetchUtil {
  #commonHeader = {}
  serverAddr = ''

  constructor (serverAddr) {
    this.serverAddr = serverAddr
  }

  SetCommonHeader = (k, v) => {
    this.#commonHeader[k] = v
  }

  GetCommonHeader = (k) => {
    return this.#commonHeader[k]
  }

  CleanCommonHeader = () => {
    this.#commonHeader = {}
  }

  SendRequest = (path, method = 'GET', headers = {}, body = null, options = {}) => {
    headers = {
      ...this.#commonHeader,
      ...headers
    }

    const properties = {
      method,
      headers,
      body,
      mode: 'cors', // cors、no-cors 或者 same-origin
      credentials: 'include', // omit、same-origin 或者 include
      cache: 'default', // default、 no-store、 reload 、 no-cache、 force-cache 或者 only-if-cached。
      redirect: 'follow', // follow (自动重定向), error (如果产生重定向将自动终止并且抛出一个错误），或者 manual (手动处理重定向)
      referrer: 'client', // no-referrer、client 或一个 URL
      referrerPolicy: 'no-referrer', // no-referrer、 no-referrer-when-downgrade、origin、origin-when-cross-origin、 unsafe-url
      ...options
    }

    const addr = this.serverAddr + path

    return fetch(addr, properties)
  }

  Get = (path, options = {}) => {
    return this.SendRequest(path, 'GET', {}, null, options)
  }

  PostForm = (path, body, options = {}) => {
    const header = {
      'Content-Type': 'application/x-www-form-urlencoded'
    }
    return this.SendRequest(path, 'POST', header, body, options)
  }

  Post = (path, data, options = {}) => {
    const header = {
      'Content-Type': 'application/json'
    }
    const jsonData = JSON.stringify(data)

    return this.SendRequest(path, 'POST', header, jsonData, options)
  }

  PostWithHeader = (path, data, header = {}, options = {}) => {
    header = {
      'Content-Type': 'application/json',
      ...header
    }
    const jsonData = JSON.stringify(data)

    return this.SendRequest(path, 'POST', header, jsonData, options)
  }

  DoJsonRequest = (method, path, body, options = {}) => {
    const header = {
      'Content-Type': 'application/json'
    }
    return this.SendRequest(path, method, header, body, options)
  }

  DoRequest = (method, path, body, options = {}) => {
    return this.SendRequest(path, method, {}, body, options)
  }
}

export default FetchUtil
