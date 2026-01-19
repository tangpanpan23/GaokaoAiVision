/**
 * 网络请求工具
 */
const BASE_URL = 'http://localhost:8888'; // 开发环境API地址

class Request {
  constructor() {
    this.baseURL = BASE_URL;
    this.timeout = 30000;
    this.interceptors = {
      request: [],
      response: []
    };
  }

  // 添加请求拦截器
  addRequestInterceptor(interceptor) {
    this.interceptors.request.push(interceptor);
  }

  // 添加响应拦截器
  addResponseInterceptor(interceptor) {
    this.interceptors.response.push(interceptor);
  }

  // 请求方法
  async request(options) {
    // 执行请求拦截器
    for (const interceptor of this.interceptors.request) {
      options = await interceptor(options);
    }

    return new Promise((resolve, reject) => {
      wx.request({
        url: this.baseURL + options.url,
        method: options.method || 'GET',
        data: options.data,
        header: {
          'Content-Type': 'application/json',
          'Authorization': wx.getStorageSync('token') || '',
          ...options.header
        },
        timeout: options.timeout || this.timeout,
        success: async (res) => {
          // 执行响应拦截器
          let response = res;
          for (const interceptor of this.interceptors.response) {
            response = await interceptor(response);
          }

          if (response.statusCode === 200) {
            if (response.data.code === 200) {
              resolve(response.data);
            } else {
              reject(new Error(response.data.msg || '请求失败'));
            }
          } else {
            reject(new Error(`网络错误: ${response.statusCode}`));
          }
        },
        fail: (err) => {
          reject(new Error(err.errMsg || '网络请求失败'));
        }
      });
    });
  }

  // GET请求
  get(url, data = {}, options = {}) {
    return this.request({
      ...options,
      url,
      method: 'GET',
      data
    });
  }

  // POST请求
  post(url, data = {}, options = {}) {
    return this.request({
      ...options,
      url,
      method: 'POST',
      data
    });
  }

  // PUT请求
  put(url, data = {}, options = {}) {
    return this.request({
      ...options,
      url,
      method: 'PUT',
      data
    });
  }

  // DELETE请求
  delete(url, data = {}, options = {}) {
    return this.request({
      ...options,
      url,
      method: 'DELETE',
      data
    });
  }
}

// 创建请求实例
const request = new Request();

// 请求拦截器 - 添加token
request.addRequestInterceptor((options) => {
  const token = wx.getStorageSync('token');
  if (token) {
    options.header = {
      ...options.header,
      'Authorization': `Bearer ${token}`
    };
  }
  return options;
});

// 响应拦截器 - 处理token过期
request.addResponseInterceptor((response) => {
  if (response.data && response.data.code === 401) {
    // token过期，清除本地token并跳转到登录页
    wx.removeStorageSync('token');
    wx.removeStorageSync('userInfo');

    wx.showToast({
      title: '登录已过期，请重新登录',
      icon: 'none',
      duration: 2000
    });

    setTimeout(() => {
      wx.reLaunch({
        url: '/pages/home/index'
      });
    }, 2000);
  }
  return response;
});

module.exports = request;
