import axios, { AxiosRequestConfig } from 'axios'

export default abstract class Request<T, D> {
  protected httpRequest (config: AxiosRequestConfig, inSpec: boolean): Promise<T> {
    return new Promise((resolve, reject) => {
      axios(config).then(res => {
        resolve(inSpec ? res.data.spec : res.data.metadata.response)
      }).catch(err => {
        reject(err.response.status)
      })
    })
  }

  abstract handle (opts: D): Promise<T>
}
