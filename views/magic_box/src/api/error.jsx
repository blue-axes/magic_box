export default class MyError extends Error {
  code = ''
  msg = ''
  constructor (msg, code = 'unexpected') {
    super(msg)
    this.code = code
    this.msg = msg
  }

  GetCode = () => {
    return this.code
  }

  GetMessage = () => {
    return this.msg
  }

  ToString = () => {
    return '[' + this.code + '] ' + this.message
  }
}
