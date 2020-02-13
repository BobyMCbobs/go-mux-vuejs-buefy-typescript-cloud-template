import Request from '@/frontend/requests/request'

export default class GetHello extends Request<string, undefined> {
  handle (): Promise<string> {
    return this.httpRequest({
      url: '/api/hello',
      method: 'GET'
    }, false)
  }
}
