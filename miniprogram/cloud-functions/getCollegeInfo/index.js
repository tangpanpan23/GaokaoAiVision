// 云函数：getCollegeInfo
// 获取学校详细信息

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
  console.log('收到学校信息查询请求:', event)

  try {
    const { collegeCode, collegeName, province, level, page = 1, pageSize = 20 } = event

    // 构建查询条件
    let whereClause = {}

    if (collegeCode) {
      whereClause.collegeCode = collegeCode
    }

    if (collegeName) {
      whereClause.collegeName = db.RegExp({
        regexp: collegeName,
        options: 'i'
      })
    }

    if (province) {
      whereClause.province = province
    }

    if (level) {
      whereClause.level = level
    }

    // 如果没有查询条件，返回热门学校
    if (Object.keys(whereClause).length === 0) {
      return await getHotColleges(page, pageSize)
    }

    // 获取总数
    const countResult = await db.collection('colleges').where(whereClause).count()
    const total = countResult.total

    // 分页查询
    const skip = (parseInt(page) - 1) * parseInt(pageSize)
    const result = await db.collection('colleges')
      .where(whereClause)
      .orderBy('ranking', 'asc')
      .skip(skip)
      .limit(parseInt(pageSize))
      .get()

    // 获取学校的录取数据统计
    const collegesWithStats = await Promise.all(
      result.data.map(async (college) => {
        const stats = await getCollegeAdmissionStats(college.collegeCode)
        return {
          ...college,
          admissionStats: stats,
          // 添加格式化字段
          levelText: getLevelText(college.level),
          typeText: getTypeText(college.type),
          natureText: college.nature === '公办' ? '公办' : '民办'
        }
      })
    )

    return {
      code: 200,
      message: '查询成功',
      data: {
        list: collegesWithStats,
        pagination: {
          total,
          page: parseInt(page),
          pageSize: parseInt(pageSize),
          totalPages: Math.ceil(total / parseInt(pageSize)),
          hasNext: parseInt(page) * parseInt(pageSize) < total,
          hasPrev: parseInt(page) > 1
        }
      }
    }

  } catch (error) {
    console.error('学校信息查询失败:', error)

    return {
      code: 500,
      message: error.message || '服务器内部错误',
      error: process.env.NODE_ENV === 'development' ? error.stack : undefined
    }
  }
}

// 获取热门学校
async function getHotColleges(page = 1, pageSize = 20) {
  const hotColleges = [
    '清华大学', '北京大学', '复旦大学', '上海交通大学',
    '浙江大学', '中国科学技术大学', '南京大学', '中山大学',
    '华中科技大学', '西安交通大学', '哈尔滨工业大学', '同济大学'
  ]

  const skip = (parseInt(page) - 1) * parseInt(pageSize)
  const limit = parseInt(pageSize)
  const collegesToQuery = hotColleges.slice(skip, skip + limit)

  const result = await db.collection('colleges')
    .where({
      collegeName: _.in(collegesToQuery)
    })
    .orderBy('ranking', 'asc')
    .get()

  // 为热门学校添加录取统计
  const collegesWithStats = await Promise.all(
    result.data.map(async (college) => {
      const stats = await getCollegeAdmissionStats(college.collegeCode)
      return {
        ...college,
        admissionStats: stats,
        isHot: true,
        levelText: getLevelText(college.level),
        typeText: getTypeText(college.type),
        natureText: college.nature === '公办' ? '公办' : '民办'
      }
    })
  )

  return {
    code: 200,
    message: '热门学校查询成功',
    data: {
      list: collegesWithStats,
      pagination: {
        total: hotColleges.length,
        page: parseInt(page),
        pageSize: parseInt(pageSize),
        totalPages: Math.ceil(hotColleges.length / parseInt(pageSize)),
        hasNext: (parseInt(page) * parseInt(pageSize)) < hotColleges.length,
        hasPrev: parseInt(page) > 1
      },
      isHotList: true
    }
  }
}

// 获取学校录取数据统计
async function getCollegeAdmissionStats(collegeCode) {
  try {
    // 获取最近3年的录取数据
    const currentYear = new Date().getFullYear()
    const recentYears = [currentYear, currentYear - 1, currentYear - 2]

    const stats = await db.collection('admission_scores')
      .where({
        collegeCode: collegeCode,
        year: _.in(recentYears)
      })
      .get()

    if (stats.data.length === 0) {
      return {
        hasData: false,
        message: '暂无录取数据'
      }
    }

    // 计算统计信息
    const scores = stats.data.map(item => item.minScore).filter(score => score > 0)
    const ranks = stats.data.map(item => item.minRank).filter(rank => rank > 0)

    const avgScore = scores.length > 0 ? Math.round(scores.reduce((a, b) => a + b, 0) / scores.length) : 0
    const minScore = scores.length > 0 ? Math.min(...scores) : 0
    const maxScore = scores.length > 0 ? Math.max(...scores) : 0

    const avgRank = ranks.length > 0 ? Math.round(ranks.reduce((a, b) => a + b, 0) / ranks.length) : 0

    return {
      hasData: true,
      totalRecords: stats.data.length,
      avgScore,
      minScore,
      maxScore,
      avgRank,
      scoreRange: minScore && maxScore ? `${minScore}-${maxScore}` : '暂无',
      dataYears: recentYears,
      // 难度评估
      difficulty: assessDifficulty(avgScore)
    }
  } catch (error) {
    console.error('获取学校录取统计失败:', error)
    return {
      hasData: false,
      message: '统计数据获取失败'
    }
  }
}

// 评估学校难度
function assessDifficulty(avgScore) {
  if (avgScore >= 650) return { level: '极高', color: '#dc3545', description: '顶尖名校，竞争激烈' }
  if (avgScore >= 620) return { level: '很高', color: '#fd7e14', description: '知名高校，竞争较大' }
  if (avgScore >= 580) return { level: '中等', color: '#ffc107', description: '普通一本，竞争适中' }
  if (avgScore >= 540) return { level: '较低', color: '#28a745', description: '二本院校，相对容易' }
  return { level: '偏低', color: '#6c757d', description: '三本及以下，录取较易' }
}

// 获取等级文本
function getLevelText(level) {
  const levelMap = {
    '985': '985工程',
    '211': '211工程',
    '一流': '一流大学',
    '一本': '一本院校',
    '二本': '二本院校',
    '三本': '三本院校',
    '专科': '专科院校'
  }
  return levelMap[level] || level || '普通本科'
}

// 获取类型文本
function getTypeText(type) {
  const typeMap = {
    '综合': '综合大学',
    '理工': '理工院校',
    '师范': '师范院校',
    '农业': '农业院校',
    '林业': '林业院校',
    '医药': '医药院校',
    '军事': '军事院校',
    '财经': '财经院校',
    '政法': '政法院校',
    '体育': '体育院校',
    '艺术': '艺术院校',
    '民族': '民族院校',
    '语言': '语言院校'
  }
  return typeMap[type] || type || '综合'
}
