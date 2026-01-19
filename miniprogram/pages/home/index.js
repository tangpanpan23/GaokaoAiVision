// pages/home/index.js
import { showLoading, hideLoading, showToast, formatNumber } from '../../utils/util.js';
import { userAPI } from '../../utils/api.js';

Page({
  data: {
    // 用户状态
    isLoggedIn: false,
    userInfo: null,

    // 功能模块
    features: [
      {
        id: 'volunteer-suggestion',
        title: '志愿推荐',
        subtitle: 'AI智能生成志愿方案',
        icon: '/images/icons/volunteer.png',
        page: '/pages/volunteer-suggestion/index',
        color: '#28a745'
      },
      {
        id: 'score-query',
        title: '分数查询',
        subtitle: '查询历年录取分数线',
        icon: '/images/icons/score.png',
        page: '', // 暂时空置，表示开发中
        color: '#007bff'
      },
      {
        id: 'ai-chat',
        title: 'AI咨询',
        subtitle: '专业志愿规划指导',
        icon: '/images/icons/ai.png',
        page: '', // 暂时空置，表示开发中
        color: '#ffc107'
      },
      {
        id: 'senior-share',
        title: '学长分享',
        subtitle: '真实校园生活体验',
        icon: '/images/icons/share.png',
        page: '', // 暂时空置，表示开发中
        color: '#dc3545'
      },
      {
        id: 'career-test',
        title: '职业测评',
        subtitle: '探索职业兴趣方向',
        icon: '/images/icons/career.png',
        page: '', // 暂时空置，表示开发中
        color: '#6f42c1'
      },
      {
        id: 'college-info',
        title: '学校专业',
        subtitle: '全面了解高校信息',
        icon: '/images/icons/college.png',
        page: '', // 暂时空置，表示开发中
        color: '#17a2b8'
      }
    ],

    // 公告通知
    announcements: [
      {
        id: 1,
        title: '2024高考志愿填报时间表',
        content: '各省高考志愿填报时间已公布，请及时关注',
        time: '2024-06-01',
        type: 'important'
      },
      {
        id: 2,
        title: '新功能上线：AI志愿推荐',
        content: '基于大数据的智能志愿推荐功能已上线',
        time: '2024-05-20',
        type: 'update'
      }
    ],

    // 统计数据
    stats: {
      totalUsers: 0,
      totalColleges: 0,
      totalMajors: 0,
      todayQueries: 0
    }
  },

  onLoad() {
    console.log('Home page loaded');
    this.checkLoginStatus();
    this.loadStats();
  },

  onShow() {
    // 每次显示页面时检查登录状态
    this.checkLoginStatus();
  },

  // 检查登录状态
  checkLoginStatus() {
    const app = getApp();
    const isLoggedIn = app.isLoggedIn();
    const userInfo = app.getUserInfo();

    this.setData({
      isLoggedIn,
      userInfo
    });

    if (isLoggedIn) {
      console.log('User is logged in:', userInfo);
    } else {
      console.log('User is not logged in');
    }
  },

  // 加载统计数据
  loadStats() {
    // 这里可以调用API获取统计数据
    // 暂时使用模拟数据
    this.setData({
      stats: {
        totalUsers: formatNumber(125680),
        totalColleges: 1280,
        totalMajors: 580,
        todayQueries: formatNumber(12890)
      }
    });
  },

  // 功能卡片点击
  onFeatureTap(e) {
    const { page } = e.currentTarget.dataset;

    // 检查页面路径是否为空
    if (!page) {
      wx.showToast({
        title: '功能开发中，敬请期待',
        icon: 'none',
        duration: 2000
      });
      return;
    }

    // 检查是否需要登录（暂时只对志愿推荐要求登录）
    if (!this.data.isLoggedIn && page === '/pages/volunteer-suggestion/index') {
      wx.showModal({
        title: '提示',
        content: '请先登录后再使用此功能',
        confirmText: '去登录',
        success: (res) => {
          if (res.confirm) {
            this.goToLogin();
          }
        }
      });
      return;
    }

    wx.navigateTo({
      url: page,
      fail: (err) => {
        console.error('Navigate to page failed:', err);
        showToast('页面跳转失败', 'none');
      }
    });
  },

  // 公告点击
  onAnnouncementTap(e) {
    const { id } = e.currentTarget.dataset;
    const announcement = this.data.announcements.find(item => item.id === id);

    if (announcement) {
      wx.showModal({
        title: announcement.title,
        content: announcement.content,
        showCancel: false
      });
    }
  },

  // 微信授权登录
  onLoginTap() {
    this.goToLogin();
  },

  // 跳转到登录
  goToLogin() {
    wx.login({
      success: (res) => {
        if (res.code) {
          this.loginWithCode(res.code);
        } else {
          console.error('Login failed:', res.errMsg);
          showToast('登录失败，请重试', 'none');
        }
      },
      fail: (err) => {
        console.error('wx.login failed:', err);
        showToast('登录失败，请重试', 'none');
      }
    });
  },

  // 使用code登录
  async loginWithCode(code) {
    try {
      showLoading('登录中...');

      const response = await userAPI.login(code);

      if (response.code === 200) {
        const { user_id, open_id, token, need_profile } = response.data;

        // 保存登录状态
        const app = getApp();
        app.setUserLogin(token, {
          userId: user_id,
          openId: open_id,
          needProfile: need_profile
        });

        this.setData({
          isLoggedIn: true,
          userInfo: {
            userId: user_id,
            openId: open_id,
            needProfile: need_profile
          }
        });

        showToast('登录成功', 'success');

        // 如果需要完善档案，跳转到个人中心
        if (need_profile) {
          setTimeout(() => {
            wx.navigateTo({
              url: '/pages/profile/index?tab=profile'
            });
          }, 1500);
        }
      } else {
        throw new Error(response.msg || '登录失败');
      }
    } catch (error) {
      console.error('Login error:', error);
      getApp().handleError(error, 'login');
    } finally {
      hideLoading();
    }
  },

  // 跳转到个人中心
  onProfileTap() {
    if (this.data.isLoggedIn) {
      wx.switchTab({
        url: '/pages/profile/index'
      });
    } else {
      this.onLoginTap();
    }
  },

  // 分享小程序
  onShareAppMessage() {
    return {
      title: '灯塔志愿 - 高考志愿填报智能助手',
      path: '/pages/home/index',
      imageUrl: '/images/share.png'
    };
  },

  // 下拉刷新
  onPullDownRefresh() {
    this.loadStats();
    setTimeout(() => {
      wx.stopPullDownRefresh();
    }, 1000);
  }
});
