import { message } from 'antd'

class Message {
  props = {
    duration: 3
  }

  Error = (msg) => {
    message.error({
      ...this.props,
      content: msg
    })
  }
}

export default new Message()
