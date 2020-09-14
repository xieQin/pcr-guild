import axios, { AxiosInstance } from 'axios'

declare module 'vue/types/vue' {
  interface Vue {
    $api: AxiosInstance;
  }
}

const api = axios.create({
  baseURL: "", //** 基础地址 要请求的url前缀  
  timeout: 5000 // 请求超时时间
});

api.interceptors.request.use(
  config => {
    config.baseURL = '/api/'
    if (!config.headers['Content-Type']) {
      config.headers['Content-Type'] = 'application/json'
    }
    return config
  }
)

api.interceptors.response.use(
  response => {
    return Promise.resolve(response)
  },
  error => {
    switch (error.response.status) {
      case 401:
        window.location.href = '/#/logout'
    }
    return Promise.reject(error)
  }
)

export default api
