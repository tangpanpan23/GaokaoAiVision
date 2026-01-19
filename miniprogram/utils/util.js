/**
 * 工具函数
 */

/**
 * 格式化时间
 */
const formatTime = (date, format = 'YYYY-MM-DD HH:mm:ss') => {
  if (!date) return '';

  const d = new Date(date);
  const year = d.getFullYear();
  const month = d.getMonth() + 1;
  const day = d.getDate();
  const hour = d.getHours();
  const minute = d.getMinutes();
  const second = d.getSeconds();

  const formatMap = {
    'YYYY': year,
    'MM': month.toString().padStart(2, '0'),
    'DD': day.toString().padStart(2, '0'),
    'HH': hour.toString().padStart(2, '0'),
    'mm': minute.toString().padStart(2, '0'),
    'ss': second.toString().padStart(2, '0')
  };

  return format.replace(/YYYY|MM|DD|HH|mm|ss/g, (match) => formatMap[match]);
};

/**
 * 格式化分数显示
 */
const formatScore = (score) => {
  if (score === null || score === undefined || score < 0) {
    return '暂无';
  }
  return `${score}分`;
};

/**
 * 格式化位次显示
 */
const formatRank = (rank) => {
  if (rank === null || rank === undefined || rank <= 0) {
    return '暂无';
  }
  return `位次${rank}`;
};

/**
 * 计算分数等级
 */
const calculateScoreLevel = (score, scoreType) => {
  if (score === null || score === undefined) {
    return { level: '暂无', color: '#999999' };
  }

  let level = '';
  let color = '';

  if (scoreType === 1) { // 文科
    if (score >= 650) { level = 'A+'; color = '#e74c3c'; }
    else if (score >= 600) { level = 'A'; color = '#e67e22'; }
    else if (score >= 550) { level = 'B+'; color = '#f39c12'; }
    else if (score >= 500) { level = 'B'; color = '#27ae60'; }
    else if (score >= 450) { level = 'C+'; color = '#3498db'; }
    else if (score >= 400) { level = 'C'; color = '#9b59b6'; }
    else { level = 'D'; color = '#95a5a6'; }
  } else if (scoreType === 2) { // 理科
    if (score >= 680) { level = 'A+'; color = '#e74c3c'; }
    else if (score >= 630) { level = 'A'; color = '#e67e22'; }
    else if (score >= 580) { level = 'B+'; color = '#f39c12'; }
    else if (score >= 530) { level = 'B'; color = '#27ae60'; }
    else if (score >= 480) { level = 'C+'; color = '#3498db'; }
    else if (score >= 430) { level = 'C'; color = '#9b59b6'; }
    else { level = 'D'; color = '#95a5a6'; }
  } else { // 综合改革或其他
    if (score >= 650) { level = 'A+'; color = '#e74c3c'; }
    else if (score >= 600) { level = 'A'; color = '#e67e22'; }
    else if (score >= 550) { level = 'B+'; color = '#f39c12'; }
    else if (score >= 500) { level = 'B'; color = '#27ae60'; }
    else if (score >= 450) { level = 'C+'; color = '#3498db'; }
    else if (score >= 400) { level = 'C'; color = '#9b59b6'; }
    else { level = 'D'; color = '#95a5a6'; }
  }

  return { level, color };
};

/**
 * 志愿推荐等级颜色
 */
const getVolunteerLevelColor = (level) => {
  const colorMap = {
    '冲': '#dc3545', // 红色
    '稳': '#28a745', // 绿色
    '保': '#ffc107'  // 黄色
  };
  return colorMap[level] || '#666666';
};

/**
 * 截断字符串
 */
const truncateText = (text, maxLength, suffix = '...') => {
  if (!text || text.length <= maxLength) {
    return text;
  }
  return text.substring(0, maxLength - suffix.length) + suffix;
};

/**
 * 防抖函数
 */
const debounce = (func, delay) => {
  let timeoutId;
  return function(...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => func.apply(this, args), delay);
  };
};

/**
 * 节流函数
 */
const throttle = (func, limit) => {
  let inThrottle;
  return function(...args) {
    if (!inThrottle) {
      func.apply(this, args);
      inThrottle = true;
      setTimeout(() => inThrottle = false, limit);
    }
  };
};

/**
 * 深度克隆对象
 */
const deepClone = (obj) => {
  if (obj === null || typeof obj !== 'object') {
    return obj;
  }

  if (obj instanceof Date) {
    return new Date(obj.getTime());
  }

  if (obj instanceof Array) {
    return obj.map(item => deepClone(item));
  }

  if (typeof obj === 'object') {
    const clonedObj = {};
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        clonedObj[key] = deepClone(obj[key]);
      }
    }
    return clonedObj;
  }
};

/**
 * 获取当前页面路径
 */
const getCurrentPageUrl = () => {
  const pages = getCurrentPages();
  const currentPage = pages[pages.length - 1];
  return currentPage.route;
};

/**
 * 获取页面参数
 */
const getPageParams = () => {
  const pages = getCurrentPages();
  const currentPage = pages[pages.length - 1];
  return currentPage.options || {};
};

/**
 * 显示加载提示
 */
const showLoading = (title = '加载中...') => {
  wx.showLoading({
    title,
    mask: true
  });
};

/**
 * 隐藏加载提示
 */
const hideLoading = () => {
  wx.hideLoading();
};

/**
 * 显示消息提示
 */
const showToast = (title, icon = 'none', duration = 2000) => {
  wx.showToast({
    title,
    icon,
    duration
  });
};

/**
 * 显示确认对话框
 */
const showModal = (options) => {
  return new Promise((resolve) => {
    wx.showModal({
      title: options.title || '提示',
      content: options.content || '',
      showCancel: options.showCancel !== false,
      cancelText: options.cancelText || '取消',
      confirmText: options.confirmText || '确定',
      success: (res) => {
        resolve(res.confirm);
      }
    });
  });
};

/**
 * 验证手机号
 */
const validatePhone = (phone) => {
  const reg = /^1[3-9]\d{9}$/;
  return reg.test(phone);
};

/**
 * 验证邮箱
 */
const validateEmail = (email) => {
  const reg = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return reg.test(email);
};

/**
 * 获取设备信息
 */
const getSystemInfo = () => {
  try {
    const res = wx.getSystemInfoSync();
    return res;
  } catch (e) {
    console.error('获取设备信息失败:', e);
    return {};
  }
};

/**
 * 格式化数字（添加千分位）
 */
const formatNumber = (num) => {
  if (!num && num !== 0) return '0';
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
};

/**
 * 检查网络状态
 */
const checkNetworkStatus = () => {
  return new Promise((resolve) => {
    wx.getNetworkType({
      success: (res) => {
        resolve({
          isConnected: res.networkType !== 'none',
          networkType: res.networkType
        });
      },
      fail: () => {
        resolve({
          isConnected: false,
          networkType: 'unknown'
        });
      }
    });
  });
};

module.exports = {
  formatTime,
  formatScore,
  formatRank,
  calculateScoreLevel,
  getVolunteerLevelColor,
  truncateText,
  debounce,
  throttle,
  deepClone,
  getCurrentPageUrl,
  getPageParams,
  showLoading,
  hideLoading,
  showToast,
  showModal,
  validatePhone,
  validateEmail,
  getSystemInfo,
  checkNetworkStatus,
  formatNumber
};
