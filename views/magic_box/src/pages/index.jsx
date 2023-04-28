import React, { Component } from 'react'
import Api from '../api/api'
import CheckLogin from '../components/checkLogin'

class Index extends Component {
  constructor (props) {
    super(props)
    this.api = new Api()
  }

  render = () => {
    return (<>
      <CheckLogin/>
      index
    </>)
  }
}

export default Index
