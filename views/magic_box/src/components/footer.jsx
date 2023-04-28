import React, { Component } from 'react'

class Footer extends Component {
  properties = {
    selfStyle: {
      width: '100%',
      height: '100px',
      backgroundColor: 'black'
    },
    iconStyle: {
      color: 'white',
      width: '100px',
      height: '100px',
      float: 'left'
    },
    authorStyle: {
      color: 'white',
      width: '30%',
      height: '100%',
      float: 'left'
    },
    docStyle: {
      color: 'white',
      width: '50px',
      height: '30px',
      float: 'left'
    }
  }

  constructor (props) {
    super(props)
    this.properties = {
      ...this.properties,
      ...props
    }
  }

  render = () => {
    const { selfStyle, iconStyle, authorStyle, docStyle } = this.properties

    return (
        <div style={selfStyle}>
          <div style={iconStyle}>图标</div>
          <div style={authorStyle}>author: zouxinjiang</div>
          <div style={docStyle}>文档</div>
        </div>
    )
  }
}

export default Footer
