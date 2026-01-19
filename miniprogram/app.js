// app.js
import { checkNetworkStatus, showToast } from './utils/util.js';

App({
  // 全局数据
  globalData: {
    userInfo: null,
    token: null,
    systemInfo: null,
    isConnected: true,
    version: '1.0.0'
  },

  // 小程序启动
  onLaunch(options) {
    console.log('App onLaunch', options);

    // 获取设备信息
    this.getSystemInfo();

    // 检查网络状态
    this.checkNetwork();

    // 初始化用户状态
    this.initUserStatus();

    // 检查更新
    this.checkForUpdate();
  },

  // 小程序显示
  onShow(options) {
    console.log('App onShow', options);
  },

  // 小程序隐藏
  onHide() {
    console.log('App onHide');
  },

  // 小程序错误
  onError(msg) {
    console.error('App onError', msg);
  },

  // 页面不存在
  onPageNotFound(res) {
    console.log('Page not found', res);
    wx.redirectTo({
      url: '/pages/home/index'
    });
  },

  // 获取设备信息
  getSystemInfo() {
    try {
      const systemInfo = wx.getSystemInfoSync();
      this.globalData.systemInfo = systemInfo;
      console.log('System info:', systemInfo);
    } catch (e) {
      console.error('Get system info failed:', e);
    }
  },

  // 检查网络状态
  async checkNetwork() {
    try {
      const networkStatus = await checkNetworkStatus();
      this.globalData.isConnected = networkStatus.isConnected;

      if (!networkStatus.isConnected) {
        showToast('网络连接异常', 'none', 3000);
      }

      console.log('Network status:', networkStatus);
    } catch (e) {
      console.error('Check network failed:', e);
    }
  },

  // 初始化用户状态
  initUserStatus() {
    try {
      const token = wx.getStorageSync('token');
      const userInfo = wx.getStorageSync('userInfo');

      if (token) {
        this.globalData.token = token;
      }

      if (userInfo) {
        this.globalData.userInfo = userInfo;
      }

      console.log('User status initialized:', { hasToken: !!token, hasUserInfo: !!userInfo });
    } catch (e) {
      console.error('Init user status failed:', e);
    }
  },

  // 设置用户登录状态
  setUserLogin(token, userInfo) {
    this.globalData.token = token;
    this.globalData.userInfo = userInfo;

    try {
      wx.setStorageSync('token', token);
      wx.setStorageSync('userInfo', userInfo);
    } catch (e) {
      console.error('Save user login status failed:', e);
    }
  },

  // 清除用户登录状态
  clearUserLogin() {
    this.globalData.token = null;
    this.globalData.userInfo = null;

    try {
      wx.removeStorageSync('token');
      wx.removeStorageSync('userInfo');
    } catch (e) {
      console.error('Clear user login status failed:', e);
    }
  },

  // 检查用户是否登录
  isLoggedIn() {
    return !!(this.globalData.token && this.globalData.userInfo);
  },

  // 获取用户信息
  getUserInfo() {
    return this.globalData.userInfo;
  },

  // 获取用户token
  getToken() {
    return this.globalData.token;
  },

  // 检查版本更新
  checkForUpdate() {
    if (wx.getUpdateManager) {
      const updateManager = wx.getUpdateManager();

      updateManager.onCheckForUpdate((res) => {
        console.log('Check for update result:', res);
        if (res.hasUpdate) {
          updateManager.onUpdateReady(() => {
            wx.showModal({
              title: '更新提示',
              content: '新版本已经准备好，是否重启应用？',
              success: (res) => {
                if (res.confirm) {
                  updateManager.applyUpdate();
                }
              }
            });
          });

          updateManager.onUpdateFailed(() => {
            showToast('新版本下载失败', 'none');
          });
        }
      });
    }
  },

  // 全局错误处理
  handleError(error, context = '') {
    console.error(`Error in ${context}:`, error);

    let message = '发生未知错误';
    if (error.message) {
      message = error.message;
    } else if (typeof error === 'string') {
      message = error;
    }

    showToast(message, 'none');
  },

  // 全局分享配置
  onShareAppMessage() {
    return {
      title: '灯塔志愿 - 高考志愿填报智能助手',
      path: '/pages/home/index',
      imageUrl: '/images/share.png'
    };
  },

  onShareTimeline() {
    return {
      title: '灯塔志愿 - 高考志愿填报智能助手',
      imageUrl: '/images/share.png'
    };
  }
});
