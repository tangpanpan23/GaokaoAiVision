// 云函数：getAdmissionScores
// 获取录取分数线数据

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
  console.log('收到录取分数查询请求:', event)

  try {
    const {
      province,
      scoreType,
      minScore,
      maxScore,
      collegeName,
      majorName,
      year,
      batch,
      page = 1,
      pageSize = 20
    } = event

    // 构建查询条件
    let whereClause = {}

    // 必填条件
    if (province) {
      whereClause.province = province
    }

    if (scoreType) {
      whereClause.scoreType = parseInt(scoreType)
    }

    // 分数范围查询
    if (minScore !== undefined || maxScore !== undefined) {
      whereClause.minScore = {}
      if (minScore !== undefined) {
        whereClause.minScore = _.gte(parseInt(minScore))
      }
      if (maxScore !== undefined) {
        whereClause.minScore = whereClause.minScore && Object.keys(whereClause.minScore).length > 0
          ? _.and(whereClause.minScore, _.lte(parseInt(maxScore)))
          : _.lte(parseInt(maxScore))
      }
    }

    // 关键词搜索
    if (collegeName) {
      whereClause.collegeName = db.RegExp({
        regexp: collegeName,
        options: 'i'  // 不区分大小写
      })
    }

    if (majorName) {
      whereClause.majorName = db.RegExp({
        regexp: majorName,
        options: 'i'
      })
    }

    // 年份筛选
    if (year) {
      whereClause.year = parseInt(year)
    }

    // 批次筛选
    if (batch) {
      whereClause.batch = batch
    }

    // 构建查询
    let query = db.collection('admission_scores').where(whereClause)

    // 获取总数
    const countResult = await query.count()
    const total = countResult.total

    // 分页查询
    const skip = (parseInt(page) - 1) * parseInt(pageSize)
    const result = await query
      .orderBy('year', 'desc')
      .orderBy('minScore', 'desc')
      .skip(skip)
      .limit(parseInt(pageSize))
      .get()

    // 处理结果数据
    const processedData = result.data.map(item => ({
      id: item._id,
      year: item.year,
      province: item.province,
      collegeCode: item.collegeCode,
      collegeName: item.collegeName,
      majorCode: item.majorCode,
      majorName: item.majorName,
      batch: item.batch,
      scoreType: item.scoreType,
      minScore: item.minScore,
      minRank: item.minRank,
      avgScore: item.avgScore,
      enrollmentCount: item.enrollmentCount,
      dataSource: item.dataSource,
      dataQuality: item.dataQuality,
      // 添加格式化字段
      scoreTypeText: item.scoreType === 1 ? '文科' : item.scoreType === 2 ? '理科' : '综合改革',
      minRankFormatted: item.minRank ? `位次${item.minRank}` : '暂无',
      dataQualityStars: '⭐'.repeat(item.dataQuality || 3)
    }))

    // 计算分页信息
    const totalPages = Math.ceil(total / parseInt(pageSize))
    const hasNext = parseInt(page) < totalPages
    const hasPrev = parseInt(page) > 1

    return {
      code: 200,
      message: '查询成功',
      data: {
        list: processedData,
        pagination: {
          total,
          page: parseInt(page),
          pageSize: parseInt(pageSize),
          totalPages,
          hasNext,
          hasPrev
        },
        // 搜索条件摘要
        searchSummary: {
          province,
          scoreType,
          scoreRange: minScore || maxScore ? `${minScore || 0}-${maxScore || 750}` : null,
          collegeKeyword: collegeName,
          majorKeyword: majorName,
          year,
          batch
        }
      }
    }

  } catch (error) {
    console.error('录取分数查询失败:', error)

    return {
      code: 500,
      message: error.message || '服务器内部错误',
      error: process.env.NODE_ENV === 'development' ? error.stack : undefined
    }
  }
}
