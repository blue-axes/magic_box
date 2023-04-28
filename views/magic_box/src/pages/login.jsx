import React, { Component } from 'react'
import { Button, Form, Input } from 'antd'
import Footer from '../components/footer'
import Header from '../components/header'
import CheckLogin from '../components/checkLogin'
import Api from '../api/api'
import MyError from '../api/error'
import Message from '../components/message'
import { SaveLoginState } from '../components/util'

class Login extends Component {
  constructor (props) {
    super(props)
    this.api = new Api()
  }

  Login = (values) => {
    this.api.Login(values.username, values.password, (data, err) => {
      if (typeof err !== 'undefined' && err instanceof MyError) {
        Message.Error(err.ToString)
      } else {
        const now = new Date()
        const end = new Date(now.getTime() + data.expire_in * 1000)

        SaveLoginState(data.access_token, end)
        this.setState({ isLogin: true })
      }
    })
  }

  render = () => {
    return (
        <div style={{ width: '100vw' }}>
            <CheckLogin />
          <Header/>
            <Form
                name="basic"
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
                style={{ maxWidth: 600, backgroundColor: 'grey' }}
                autoComplete="off"
                onFinish={this.Login}
            >
                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input/>
                </Form.Item>
                <Form.Item
                    label="Password"
                    name="password"
                    rules={[{ required: true, message: 'Please input your password!' }]}
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Login
                    </Button>
                </Form.Item>
            </Form>
          <Footer/>
        </div>
    )
  }
}

export default Login
