// 云函数：getUserProfile
// 获取用户高考档案

const cloud = require('wx-server-sdk')

// 初始化云环境
cloud.init({
  env: cloud.DYNAMIC_CURRENT_ENV
})

// 云数据库
const db = cloud.database()

// 主函数
exports.main = async (event, context) => {
  console.log('收到用户档案查询请求:', event)

  // 获取用户OpenID
  const { OPENID } = cloud.getWXContext()

  if (!OPENID) {
    return {
      code: 401,
      message: '用户未授权'
    }
  }

  try {
    // 查询用户档案
    const result = await db.collection('user_profiles')
      .where({ userId: OPENID })
      .get()

    if (result.data.length === 0) {
      return {
        code: 404,
        message: '用户档案不存在',
        data: {
          hasProfile: false,
          needCreate: true
        }
      }
    }

    const profile = result.data[0]

    // 处理档案数据
    const processedProfile = {
      id: profile._id,
      userId: profile.userId,
      graduationYear: profile.graduationYear,
      province: profile.province,
      scoreType: profile.scoreType,
      scoreTypeText: profile.scoreType === 1 ? '文科' : profile.scoreType === 2 ? '理科' : '综合改革',
      totalScore: profile.totalScore,
      rank: profile.rank,
      subjects: profile.subjects,
      interestTags: profile.interestTags || [],
      personalityType: profile.personalityType,
      targetCollege: profile.targetCollege,
      targetMajor: profile.targetMajor,
      createdAt: profile.createdAt,
      updatedAt: profile.updatedAt,

      // 计算属性
      scoreLevel: getScoreLevel(profile.totalScore, profile.scoreType),
      rankLevel: getRankLevel(profile.rank, profile.province),
      interestTagsCount: (profile.interestTags || []).length,
      completeness: calculateProfileCompleteness(profile)
    }

    return {
      code: 200,
      message: '查询成功',
      data: {
        hasProfile: true,
        profile: processedProfile
      }
    }

  } catch (error) {
    console.error('获取用户档案失败:', error)

    return {
      code: 500,
      message: error.message || '服务器内部错误',
      error: process.env.NODE_ENV === 'development' ? error.stack : undefined
    }
  }
}

// 获取分数等级
function getScoreLevel(score, scoreType) {
  if (!score) return '未知'

  // 根据不同分数类型调整基准
  const baseScore = scoreType === 1 ? 550 : scoreType === 2 ? 500 : 520

  if (score >= baseScore + 100) return '优秀'
  if (score >= baseScore + 50) return '良好'
  if (score >= baseScore) return '中等'
  if (score >= baseScore - 50) return '较低'
  return '偏低'
}

// 获取位次等级
function getRankLevel(rank, province) {
  if (!rank || !province) return '未知'

  // 根据省份调整基准（简化版）
  const provinceLevels = {
    '北京': { excellent: 500, good: 2000, average: 5000 },
    '上海': { excellent: 300, good: 1500, average: 4000 },
    '广东': { excellent: 1000, good: 3000, average: 8000 },
    // 默认值
    'default': { excellent: 800, good: 2500, average: 6000 }
  }

  const levels = provinceLevels[province] || provinceLevels.default

  if (rank <= levels.excellent) return '优秀'
  if (rank <= levels.good) return '良好'
  if (rank <= levels.average) return '中等'
  return '偏低'
}

// 计算档案完整度
function calculateProfileCompleteness(profile) {
  const fields = [
    'graduationYear', 'province', 'scoreType', 'totalScore',
    'subjects', 'interestTags', 'personalityType'
  ]

  let completed = 0
  let total = fields.length

  fields.forEach(field => {
    const value = profile[field]
    if (value !== undefined && value !== null && value !== '') {
      if (Array.isArray(value) && value.length > 0) {
        completed++
      } else if (!Array.isArray(value)) {
        completed++
      }
    }
  })

  // rank 是可选字段，不计入完整度
  return Math.round((completed / total) * 100)
}
