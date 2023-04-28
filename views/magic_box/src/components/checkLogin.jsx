import React, { Component } from 'react'
import { IsLogin } from './util'
import { Navigate } from 'react-router-dom'
import { LoginUrlPath } from '../router'

class CheckLogin extends Component {
  state = {
    isLogin: false,
    urlPath: ''
  }

  constructor (props) {
    super(props)
    this.state = { urlPath: window.location.pathname }
    if (IsLogin()) {
      this.state.isLogin = true
    }
  }

  render () {
    const { isLogin, urlPath } = this.state
    if (!isLogin && urlPath !== LoginUrlPath) {
      return <Navigate to={LoginUrlPath} />
    }
    return null
  }
}

export default CheckLogin
