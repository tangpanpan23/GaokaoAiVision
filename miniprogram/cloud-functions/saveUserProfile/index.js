// 云函数：saveUserProfile
// 保存用户高考档案

const cloud = require('wx-server-sdk')

// 初始化云环境
cloud.init({
  env: cloud.DYNAMIC_CURRENT_ENV
})

// 云数据库
const db = cloud.database()
const _ = db.command

// 主函数
exports.main = async (event, context) => {
  console.log('收到用户档案保存请求:', event)

  // 获取用户OpenID
  const { OPENID } = cloud.getWXContext()

  if (!OPENID) {
    return {
      code: 401,
      message: '用户未授权'
    }
  }

  try {
    const {
      graduationYear,
      province,
      scoreType,
      totalScore,
      rank,
      subjects,
      interestTags,
      personalityType,
      targetCollege,
      targetMajor
    } = event

    // 数据验证
    const validation = validateUserProfile(event)
    if (!validation.valid) {
      return {
        code: 400,
        message: validation.message
      }
    }

    // 准备档案数据
    const profileData = {
      userId: OPENID,
      graduationYear: parseInt(graduationYear),
      province,
      scoreType: parseInt(scoreType),
      totalScore: parseInt(totalScore),
      rank: rank ? parseInt(rank) : null,
      subjects: subjects || '',
      interestTags: Array.isArray(interestTags) ? interestTags : [],
      personalityType: personalityType || '',
      targetCollege: targetCollege || '',
      targetMajor: targetMajor || '',
      updatedAt: db.serverDate(),
      createdAt: db.serverDate()
    }

    // 检查是否已存在档案
    const existingProfile = await db.collection('user_profiles')
      .where({ userId: OPENID })
      .get()

    let result
    if (existingProfile.data.length > 0) {
      // 更新现有档案
      result = await db.collection('user_profiles')
        .doc(existingProfile.data[0]._id)
        .update({
          data: {
            ...profileData,
            createdAt: existingProfile.data[0].createdAt // 保留创建时间
          }
        })
    } else {
      // 创建新档案
      result = await db.collection('user_profiles').add({
        data: profileData
      })
    }

    // 获取完整的档案信息用于返回
    const updatedProfile = await db.collection('user_profiles')
      .where({ userId: OPENID })
      .get()

    return {
      code: 200,
      message: existingProfile.data.length > 0 ? '档案更新成功' : '档案创建成功',
      data: {
        profile: updatedProfile.data[0],
        isNew: existingProfile.data.length === 0
      }
    }

  } catch (error) {
    console.error('保存用户档案失败:', error)

    return {
      code: 500,
      message: error.message || '服务器内部错误',
      error: process.env.NODE_ENV === 'development' ? error.stack : undefined
    }
  }
}

// 数据验证函数
function validateUserProfile(data) {
  const { graduationYear, province, scoreType, totalScore } = data

  // 必填字段验证
  if (!graduationYear) {
    return { valid: false, message: '请选择毕业年份' }
  }

  if (!province) {
    return { valid: false, message: '请选择高考省份' }
  }

  if (!scoreType) {
    return { valid: false, message: '请选择分数类型' }
  }

  if (totalScore === undefined || totalScore === null) {
    return { valid: false, message: '请输入高考总分' }
  }

  // 数据范围验证
  const currentYear = new Date().getFullYear()
  if (graduationYear < currentYear - 10 || graduationYear > currentYear + 5) {
    return { valid: false, message: '毕业年份范围无效' }
  }

  if (scoreType < 1 || scoreType > 3) {
    return { valid: false, message: '分数类型无效' }
  }

  if (totalScore < 0 || totalScore > 750) {
    return { valid: false, message: '分数范围无效，应在0-750之间' }
  }

  // 可选字段验证
  if (data.rank !== undefined && data.rank !== null && data.rank < 0) {
    return { valid: false, message: '位次不能为负数' }
  }

  if (data.interestTags && !Array.isArray(data.interestTags)) {
    return { valid: false, message: '兴趣标签格式错误' }
  }

  if (data.interestTags && data.interestTags.length > 10) {
    return { valid: false, message: '兴趣标签最多选择10个' }
  }

  return { valid: true }
}
