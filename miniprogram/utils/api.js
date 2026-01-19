/**
 * API接口定义 - 微信小程序云开发版本
 */

// 云环境ID
const CLOUD_ENV_ID = 'lighthouse-volunteer-dev'

// 云函数调用封装
const callCloudFunction = async (functionName, data = {}) => {
  try {
    console.log(`调用云函数: ${functionName}`, data)

    const result = await wx.cloud.callFunction({
      name: functionName,
      data
    })

    console.log(`云函数 ${functionName} 返回:`, result)

    const { code, message, data: responseData, error } = result.result

    if (code !== 200) {
      throw new Error(message || '云函数调用失败')
    }

    return responseData
  } catch (error) {
    console.error(`云函数 ${functionName} 调用失败:`, error)

    // 如果是网络错误或云函数不存在，返回模拟数据
    if (error.errCode === -1 || error.errMsg?.includes('cloud function')) {
      console.log('云函数不可用，返回模拟数据')
      return getMockData(functionName, data)
    }

    throw error
  }
}

// 获取模拟数据（云函数不可用时使用）
const getMockData = (functionName, params) => {
  console.log(`使用模拟数据: ${functionName}`)

  const mockData = {
    getAdmissionScores: {
      list: [
        {
          id: 'score_2024_10001_080901',
          year: 2024,
          province: '北京',
          collegeCode: '10001',
          collegeName: '清华大学',
          majorCode: '080901',
          majorName: '计算机科学与技术',
          batch: '一本',
          scoreType: 2,
          minScore: 685,
          minRank: 150,
          avgScore: 690,
          enrollmentCount: 50,
          dataSource: '北京市教育考试院',
          dataQuality: 5,
          scoreTypeText: '理科',
          minRankFormatted: '位次150',
          dataQualityStars: '⭐⭐⭐⭐⭐'
        }
      ],
      pagination: {
        total: 1,
        page: 1,
        pageSize: 20,
        totalPages: 1,
        hasNext: false,
        hasPrev: false
      },
      searchSummary: {
        province: params.province || '北京',
        scoreType: params.scoreType || 2,
        collegeKeyword: params.collegeName || null,
        majorKeyword: params.majorName || null
      }
    },

    generateVolunteerSuggestion: {
      categories: [
        {
          category: '冲',
          reason: '分数有竞争力，值得冲刺',
          categoryColor: '#dc3545',
          colleges: [
            {
              collegeCode: '10001',
              collegeName: '清华大学',
              majorName: '计算机科学与技术',
              minScore: 685,
              minRank: 150,
              matchingScorePercent: 85,
              advantages: '综合实力强，就业前景好',
              considerations: '竞争激烈，需要优秀表现'
            }
          ]
        },
        {
          category: '稳',
          reason: '分数匹配度高，录取把握大',
          categoryColor: '#28a745',
          colleges: [
            {
              collegeCode: '10003',
              collegeName: '复旦大学',
              majorName: '计算机科学与技术',
              minScore: 645,
              minRank: 300,
              matchingScorePercent: 75,
              advantages: '人文底蕴深厚，综合实力强',
              considerations: '专业竞争较为激烈'
            }
          ]
        },
        {
          category: '保',
          reason: '分数优势明显，确保录取',
          categoryColor: '#ffc107',
          colleges: [
            {
              collegeCode: '10004',
              collegeName: '上海交通大学',
              majorName: '计算机科学与技术',
              minScore: 620,
              minRank: 500,
              matchingScorePercent: 65,
              advantages: '工科实力雄厚，科研水平高',
              considerations: '建议结合兴趣选择专业方向'
            }
          ]
        }
      ],
      analysisSummary: '根据您的分数和省份情况，建议重点关注前两类学校，同时做好保底准备。计算机科学与技术专业就业前景良好，但竞争较为激烈。',
      recommendations: [
        '重点关注清华大学、北京大学等顶尖高校',
        '结合个人兴趣选择专业方向',
        '关注学校就业数据和专业实力',
        '合理安排冲稳保志愿梯度',
        '建议咨询专业老师获取更多建议'
      ]
    },

    getUserProfile: {
      hasProfile: true,
      profile: {
        id: 'user_test',
        userId: 'test_open_id',
        graduationYear: 2024,
        province: '北京',
        scoreType: 2,
        scoreTypeText: '理科',
        totalScore: 650,
        rank: 5000,
        subjects: '物理+历史',
        interestTags: ['计算机', '人工智能'],
        personalityType: 'INTJ',
        targetCollege: '清华大学',
        targetMajor: '计算机科学与技术',
        scoreLevel: '良好',
        rankLevel: '中等',
        completeness: 85
      }
    },

    getCollegeInfo: {
      list: [
        {
          _id: 'college_10001',
          collegeCode: '10001',
          collegeName: '清华大学',
          province: '北京',
          city: '北京',
          level: '985',
          type: '综合',
          nature: '公办',
          website: 'https://www.tsinghua.edu.cn',
          description: '清华大学是中国著名高等学府，位于北京市海淀区。',
          ranking: 1,
          employmentRate: 98.5,
          admissionStats: {
            hasData: true,
            avgScore: 680,
            minScore: 650,
            maxScore: 700,
            scoreRange: '650-700',
            difficulty: {
              level: '极高',
              color: '#dc3545',
              description: '顶尖名校，竞争激烈'
            }
          },
          levelText: '985工程',
          typeText: '综合大学',
          natureText: '公办'
        }
      ],
      pagination: {
        total: 1,
        page: 1,
        pageSize: 20,
        totalPages: 1,
        hasNext: false,
        hasPrev: false
      }
    }
  }

  return mockData[functionName] || {}
}

// 用户相关接口
const userAPI = {
  // 用户登录（微信授权）
  login: async (code) => {
    // 这里通常不需要云函数，直接处理微信授权
    return {
      code: 200,
      msg: 'success',
      data: {
        user_id: Date.now(),
        open_id: 'wx_' + code,
        token: 'wx_token_' + Date.now(),
        need_profile: true
      }
    }
  },

  // 获取用户档案
  getProfile: async () => {
    return await callCloudFunction('getUserProfile')
  },

  // 保存用户档案
  saveProfile: async (profileData) => {
    return await callCloudFunction('saveUserProfile', profileData)
  }
}

// 志愿填报相关接口
const volunteerAPI = {
  // 查询录取分数线
  getScores: async (params) => {
    return await callCloudFunction('getAdmissionScores', params)
  },

  // 生成志愿推荐
  getSuggestions: async (params) => {
    return await callCloudFunction('generateVolunteerSuggestion', params)
  }
}

// 学校相关接口
const collegeAPI = {
  // 获取学校列表
  getList: async (params) => {
    return await callCloudFunction('getCollegeInfo', params)
  },

  // 获取学校详情
  getDetail: async (collegeCode) => {
    const result = await callCloudFunction('getCollegeInfo', {
      collegeCode: collegeCode
    })
    return result.list?.[0] || null
  }
}

// 专业相关接口
const majorAPI = {
  // 获取专业列表（暂时使用模拟数据）
  getList: async (params) => {
    return {
      list: [
        {
          majorCode: '080901',
          majorName: '计算机科学与技术',
          category: '工学',
          degree: '学士',
          duration: 4,
          description: '计算机科学与技术专业...',
          employmentDirection: '软件开发、系统分析...'
        }
      ],
      total: 1
    }
  }
}

// 统计相关接口
const statsAPI = {
  // 获取统计数据
  getOverview: async () => {
    return {
      totalUsers: 123456,
      totalColleges: 2894,
      totalMajors: 738,
      todayQueries: 8765
    }
  }
}

// 数据库初始化接口（仅管理员使用）
const initAPI = {
  // 初始化数据库
  initDatabase: async (action, data) => {
    return await callCloudFunction('initDatabase', { action, data })
  }
}

module.exports = {
  userAPI,
  volunteerAPI,
  collegeAPI,
  majorAPI,
  statsAPI,
  initAPI
}