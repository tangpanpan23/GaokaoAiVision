// pages/volunteer-suggestion/index.js
import { showLoading, hideLoading, showToast, showModal } from '../../utils/util.js';
import { volunteerAPI } from '../../utils/api.js';

Page({
  data: {
    // 表单数据
    form: {
      province: '',
      scoreType: 1, // 1-文科，2-理科，3-综合改革
      score: '',
      rank: '',
      subjects: '',
      interestTags: []
    },

    // 省份列表
    provinces: [
      '北京', '上海', '天津', '重庆', '河北', '山西', '辽宁', '吉林',
      '黑龙江', '江苏', '浙江', '安徽', '福建', '江西', '山东', '河南',
      '湖北', '湖南', '广东', '广西', '海南', '四川', '贵州', '云南',
      '西藏', '陕西', '甘肃', '青海', '宁夏', '新疆'
    ],

    // 科目组合
    subjectCombinations: [
      '物理+历史', '物理+地理', '物理+政治', '物理+生物', '物理+化学',
      '历史+地理', '历史+政治', '历史+生物', '历史+化学',
      '思想政治', '地理', '历史', '生物', '物理', '化学'
    ],

    // 兴趣标签
    interestTags: [
      '计算机', '金融', '医学', '教育', '工程', '艺术', '法律', '管理',
      '农业', '体育', '旅游', '媒体', '科研', '创业', '公益', '其他'
    ],

    // 推荐结果
    suggestions: null,
    showResult: false,

    // 加载状态
    isLoading: false,

    // 表单验证
    formErrors: {},

    // 标签选中状态（用于WXML显示）
    tagSelectedStates: []
  },

  onLoad(options) {
    console.log('Volunteer suggestion page loaded', options);
    this.loadUserProfile();
    this.initTagSelectedStates();
  },

  // 初始化标签选中状态
  initTagSelectedStates() {
    const { interestTags, form } = this.data;
    const tagSelectedStates = interestTags.map(tag => form.interestTags.includes(tag));
    this.setData({ tagSelectedStates });
  },

  // 加载用户档案信息
  loadUserProfile() {
    const app = getApp();
    const userInfo = app.getUserInfo();

    if (userInfo && userInfo.needProfile) {
      showToast('请先完善个人档案', 'none');
      setTimeout(() => {
        wx.navigateTo({
          url: '/pages/profile/index?tab=profile'
        });
      }, 2000);
      return;
    }

    // 如果有用户档案，预填表单
    if (userInfo && userInfo.profile) {
      const profile = userInfo.profile;
      const interestTags = profile.interestTags ? profile.interestTags.split(',') : [];
      const { interestTags: allTags } = this.data;
      const tagSelectedStates = allTags.map(tag => interestTags.includes(tag));
      
      this.setData({
        form: {
          province: profile.province || '',
          scoreType: profile.scoreType || 1,
          score: profile.totalScore || '',
          rank: profile.rank || '',
          subjects: profile.subjects || '',
          interestTags: interestTags
        },
        tagSelectedStates: tagSelectedStates
      });
    }
  },

  // 省份选择
  onProvinceChange(e) {
    const province = this.data.provinces[e.detail.value];
    this.setData({
      'form.province': province
    });
    this.clearFormError('province');
  },

  // 分数类型选择
  onScoreTypeChange(e) {
    this.setData({
      'form.scoreType': parseInt(e.detail.value) + 1
    });
  },

  // 输入框变化
  onInputChange(e) {
    const { field } = e.currentTarget.dataset;
    const value = e.detail.value;
    this.setData({
      [`form.${field}`]: value
    });
    this.clearFormError(field);
  },

  // 科目选择
  onSubjectChange(e) {
    const subject = this.data.subjectCombinations[e.detail.value];
    this.setData({
      'form.subjects': subject
    });
    this.clearFormError('subjects');
  },

  // 兴趣标签选择
  onInterestTagTap(e) {
    const { tag, index } = e.currentTarget.dataset;
    const { interestTags, form, tagSelectedStates } = this.data;
    const currentTags = form.interestTags || [];
    const tagIndex = currentTags.indexOf(tag);

    if (tagIndex > -1) {
      // 取消选择
      currentTags.splice(tagIndex, 1);
      tagSelectedStates[parseInt(index)] = false;
    } else {
      // 选择标签（最多选择3个）
      if (currentTags.length >= 3) {
        showToast('最多只能选择3个兴趣标签', 'none');
        return;
      }
      currentTags.push(tag);
      tagSelectedStates[parseInt(index)] = true;
    }

    this.setData({
      'form.interestTags': currentTags,
      tagSelectedStates: tagSelectedStates
    });
  },

  // 处理推荐结果数据
  processSuggestions(data) {
    if (!data || !data.categories) {
      return data;
    }

    // 志愿等级颜色映射
    const levelColors = {
      '冲': '#dc3545',
      '稳': '#28a745',
      '保': '#ffc107'
    };

    // 处理每个分类
    const processedCategories = data.categories.map(category => {
      // 添加分类颜色
      category.categoryColor = levelColors[category.category] || '#666666';

      // 处理每个学校
      if (category.colleges) {
        category.colleges = category.colleges.map(college => {
          // 添加匹配度百分比
          if (college.matching_score !== undefined) {
            college.matching_score_percent = Math.round(college.matching_score * 100);
          }
          return college;
        });
      }

      return category;
    });

    return {
      ...data,
      categories: processedCategories
    };
  },

  // 生成推荐
  async onGenerateSuggestion() {
    if (!this.validateForm()) {
      return;
    }

    try {
      showLoading('AI分析中，请稍候...');
      this.setData({ isLoading: true });

      const params = {
        province: this.data.form.province,
        score_type: this.data.form.scoreType,
        score: parseInt(this.data.form.score),
        rank: parseInt(this.data.form.rank) || 0,
        subjects: this.data.form.subjects,
        interest_tags: this.data.form.interestTags
      };

      const response = await volunteerAPI.getSuggestions(params);

      if (response.code === 200) {
        // 处理推荐结果，添加颜色和格式化数据
        const suggestions = this.processSuggestions(response.data);
        
        this.setData({
          suggestions: suggestions,
          showResult: true
        });

        // 滚动到结果区域
        setTimeout(() => {
          wx.pageScrollTo({
            selector: '.result-section',
            duration: 300
          });
        }, 100);
      } else {
        throw new Error(response.msg || '获取推荐失败');
      }
    } catch (error) {
      console.error('Generate suggestion error:', error);
      getApp().handleError(error, 'volunteer_suggestion');
    } finally {
      hideLoading();
      this.setData({ isLoading: false });
    }
  },

  // 表单验证
  validateForm() {
    const { form } = this.data;
    const errors = {};

    if (!form.province) {
      errors.province = '请选择高考省份';
    }

    if (!form.score || isNaN(form.score) || form.score < 0 || form.score > 750) {
      errors.score = '请输入有效的分数（0-750）';
    }

    if (!form.subjects) {
      errors.subjects = '请选择选考科目';
    }

    if (form.interestTags.length === 0) {
      errors.interestTags = '请至少选择1个兴趣标签';
    }

    this.setData({ formErrors: errors });
    return Object.keys(errors).length === 0;
  },

  // 清除表单错误
  clearFormError(field) {
    const { formErrors } = this.data;
    if (formErrors[field]) {
      delete formErrors[field];
      this.setData({ formErrors });
    }
  },

  // 重置表单
  onResetForm() {
    showModal({
      title: '确认重置',
      content: '确定要清空所有输入内容吗？'
    }).then((confirm) => {
      if (confirm) {
        const { interestTags } = this.data;
        const tagSelectedStates = interestTags.map(() => false);
        
        this.setData({
          form: {
            province: '',
            scoreType: 1,
            score: '',
            rank: '',
            subjects: '',
            interestTags: []
          },
          formErrors: {},
          suggestions: null,
          showResult: false,
          tagSelectedStates: tagSelectedStates
        });
      }
    });
  },

  // 分享推荐结果
  onShareResult() {
    if (!this.data.suggestions) {
      showToast('暂无推荐结果可分享', 'none');
      return;
    }

    wx.showShareMenu({
      withShareTicket: true
    });
  },

  // 保存推荐结果
  onSaveResult() {
    if (!this.data.suggestions) {
      showToast('暂无推荐结果可保存', 'none');
      return;
    }

    // 这里可以调用API保存推荐结果到用户档案
    showToast('推荐结果已保存', 'success');
  },

  // 跳转到学校详情
  onCollegeTap(e) {
    const { collegeCode } = e.currentTarget.dataset;
    wx.navigateTo({
      url: `/pages/college-detail/index?code=${collegeCode}`
    });
  },

  // 重新生成
  onRegenerate() {
    this.setData({
      suggestions: null,
      showResult: false
    });

    // 滚动到表单区域
    wx.pageScrollTo({
      selector: '.form-section',
      duration: 300
    });
  },

  // 分享页面
  onShareAppMessage() {
    const { suggestions } = this.data;
    let title = '灯塔志愿 - 高考志愿推荐';

    if (suggestions && suggestions.categories && suggestions.categories.length > 0) {
      const firstCategory = suggestions.categories[0];
      if (firstCategory.colleges && firstCategory.colleges.length > 0) {
        const firstCollege = firstCategory.colleges[0];
        title = `我${firstCategory.category}的学校是${firstCollege.college_name}`;
      }
    }

    return {
      title,
      path: '/pages/volunteer-suggestion/index',
      imageUrl: '/images/share-volunteer.png'
    };
  }
});
