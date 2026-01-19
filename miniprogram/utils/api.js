/**
 * API接口定义 - 模拟版本
 */

// 模拟延迟
const delay = (ms = 1000) => new Promise(resolve => setTimeout(resolve, ms));

// 用户相关接口
const userAPI = {
  // 用户登录
  login: async (code) => {
    await delay(800);
    return {
      code: 200,
      msg: 'success',
      data: {
        user_id: 12345,
        open_id: 'test_open_id_' + code,
        token: 'test_jwt_token_' + Date.now(),
        need_profile: true
      }
    };
  },

  // 更新用户档案
  updateProfile: async (profileData) => {
    await delay(600);
    return {
      code: 200,
      msg: 'success'
    };
  }
};

// 志愿填报相关接口
const volunteerAPI = {
  // 查询录取分数线
  queryScores: async (params) => {
    await delay(1000);
    return {
      code: 200,
      msg: 'success',
      data: [
        {
          college_name: '清华大学',
          major_name: '计算机科学与技术',
          batch: '一本',
          min_score: 685,
          min_rank: 150,
          year: 2024
        },
        {
          college_name: '北京大学',
          major_name: '软件工程',
          batch: '一本',
          min_score: 680,
          min_rank: 200,
          year: 2024
        }
      ]
    };
  },

  // 获取志愿推荐
  getSuggestions: async (params) => {
    await delay(2000); // 模拟AI分析时间

    const baseScore = params.score;
    const baseRank = params.rank;

    return {
      code: 200,
      msg: 'success',
      data: {
        categories: [
          {
            category: '冲',
            reason: '分数有一定竞争力，建议冲刺理想学校',
            colleges: [
              {
                college_code: '10001',
                college_name: '清华大学',
                major_code: '080901',
                major_name: '计算机科学与技术',
                batch: '一本',
                min_score: baseScore + 10,
                min_rank: Math.max(1, baseRank - 500),
                year: 2024,
                matching_score: 0.85,
                advantages: '顶尖计算机专业，师资力量雄厚',
                considerations: '录取分数线较高，需要全力备考'
              },
              {
                college_code: '10002',
                college_name: '北京大学',
                major_code: '080902',
                major_name: '软件工程',
                batch: '一本',
                min_score: baseScore + 5,
                min_rank: Math.max(1, baseRank - 300),
                year: 2024,
                matching_score: 0.80,
                advantages: '综合性大学，学科交叉明显',
                considerations: '专业竞争激烈，建议多手准备'
              }
            ]
          },
          {
            category: '稳',
            reason: '分数较为稳定，建议选择有把握的学校',
            colleges: [
              {
                college_code: '10003',
                college_name: '上海交通大学',
                major_code: '080903',
                major_name: '信息工程',
                batch: '一本',
                min_score: Math.max(0, baseScore - 5),
                min_rank: baseRank + 200,
                year: 2024,
                matching_score: 0.90,
                advantages: '工科优势明显，就业前景良好',
                considerations: '需要保持良好发挥'
              }
            ]
          },
          {
            category: '保',
            reason: '确保录取，建议选择安全系数高的学校',
            colleges: [
              {
                college_code: '10004',
                college_name: '华东师范大学',
                major_code: '040101',
                major_name: '教育学',
                batch: '一本',
                min_score: Math.max(0, baseScore - 20),
                min_rank: baseRank + 1000,
                year: 2024,
                matching_score: 0.75,
                advantages: '师范类专业，就业稳定',
                considerations: '根据个人职业规划选择'
              }
            ]
          }
        ],
        analysis_summary: `根据你的${params.score_type === 1 ? '文科' : params.score_type === 2 ? '理科' : '综合改革'}分数${baseScore}分、位次${baseRank}，结合${params.subjects}选考科目和${params.interest_tags.join('、')}兴趣，推荐如下志愿方案：`,
        recommendations: [
          '建议按照冲、稳、保的原则合理分配志愿',
          '关注各省高考政策变化和院校调档规则',
          '保持良好心态，认真对待每一次模拟考试',
          '及时关注志愿填报时间节点，避免错过填报时间'
        ]
      }
    };
  },

  // AI志愿咨询
  getAdvice: async (params) => {
    await delay(1500);
    return {
      code: 200,
      msg: 'success',
      data: {
        answer: `关于"${params.query}"的问题，我的建议是：高考志愿填报需要综合考虑个人兴趣、专业前景、学校实力等多个因素。建议你根据自己的分数位次和兴趣方向，选择匹配度高的专业和学校。同时要关注专业的就业前景和未来的发展空间。`,
        session_id: 'session_' + Date.now(),
        sources: ['官方录取数据', '就业统计数据', '专业前景分析']
      }
    };
  }
};

// 学校专业信息接口
const collegeAPI = {
  // 查询学校信息
  getColleges: async (params) => {
    await delay(600);
    return {
      code: 200,
      msg: 'success',
      data: {
        total: 1280,
        page: params.page || 1,
        page_size: params.page_size || 20,
        pages: 64,
        list: [
          {
            college_code: '10001',
            college_name: '清华大学',
            province: '北京',
            level: '985',
            type: '综合',
            ranking: 1
          },
          {
            college_code: '10002',
            college_name: '北京大学',
            province: '北京',
            level: '985',
            type: '综合',
            ranking: 2
          }
        ]
      }
    };
  },

  // 查询专业信息
  getMajors: async (params) => {
    await delay(600);
    return {
      code: 200,
      msg: 'success',
      data: {
        total: 580,
        page: params.page || 1,
        page_size: params.page_size || 20,
        pages: 29,
        list: [
          {
            major_code: '080901',
            major_name: '计算机科学与技术',
            category: '工学',
            demand_level: '高'
          },
          {
            major_code: '080902',
            major_name: '软件工程',
            category: '工学',
            demand_level: '高'
          }
        ]
      }
    };
  }
};

// 学长分享接口
const shareAPI = {
  // 获取分享列表
  getShares: async (params) => {
    await delay(800);
    return {
      code: 200,
      msg: 'success',
      data: {
        total: 256,
        page: params.page || 1,
        page_size: params.page_size || 20,
        pages: 13,
        list: [
          {
            id: 1,
            college_name: '清华大学',
            major_name: '计算机科学与技术',
            share_type: 'experience',
            title: '清华计算机的学习生活分享',
            content: '清华大学计算机专业课程设置很合理，大一大二打基础，大三大四做项目...',
            tags: ['计算机', '清华', '学习经验'],
            view_count: 1250,
            like_count: 89,
            published_at: '2024-05-15',
            is_liked: false
          },
          {
            id: 2,
            college_name: '北京大学',
            major_name: '软件工程',
            share_type: 'advice',
            title: '北大软工的填报建议',
            content: '北大软件工程专业对数学要求比较高，如果数学不是强项建议慎重...',
            tags: ['软件工程', '北大', '填报建议'],
            view_count: 980,
            like_count: 67,
            published_at: '2024-05-10',
            is_liked: true
          }
        ]
      }
    };
  },

  // 创建分享
  createShare: async (shareData) => {
    await delay(1000);
    return {
      code: 200,
      msg: 'success',
      data: {
        share_id: Date.now()
      }
    };
  },

  // 点赞分享
  likeShare: async (shareId, like) => {
    await delay(300);
    return {
      code: 200,
      msg: 'success',
      data: {
        like_count: like ? 88 : 87
      }
    };
  }
};

// 测评规划接口
const assessmentAPI = {
  // 职业测评
  careerAssessment: async (assessmentData) => {
    await delay(2000);
    return {
      code: 200,
      msg: 'success',
      data: {
        assessment_type: assessmentData.assessment_type,
        result: {
          type: 'I',
          dimension: '内向直觉思考',
          description: '你是一个内向、注重直觉和逻辑思维的人'
        },
        score_details: {
          E_I: 30,
          S_N: 70,
          T_F: 80,
          J_P: 60
        },
        recommendations: '根据测评结果，你适合从事需要独立思考和创造力的工作，如科研、编程、设计等领域。建议选择相关专业进行深入学习。',
        created_at: new Date().toISOString()
      }
    };
  },

  // 选科规划查询
  getSubjectPlan: async (params) => {
    await delay(800);
    return {
      code: 200,
      msg: 'success',
      data: {
        subject_combination: '物理+历史',
        recommended_majors: ['计算机科学', '电子信息', '自动化'],
        success_rate: 0.85,
        avg_score_required: 580
      }
    };
  }
};

module.exports = {
  userAPI,
  volunteerAPI,
  collegeAPI,
  shareAPI,
  assessmentAPI
};
