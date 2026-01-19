// 云函数：initDatabase
// 初始化数据库和导入基础数据

const cloud = require('wx-server-sdk')

// 初始化云环境
cloud.init({
  env: cloud.DYNAMIC_CURRENT_ENV
})

// 云数据库
const db = cloud.database()
const _ = db.command

// 基础数据
const PROVINCES = [
  '北京', '上海', '天津', '重庆', '河北', '山西', '辽宁', '吉林',
  '黑龙江', '江苏', '浙江', '安徽', '福建', '江西', '山东', '河南',
  '湖北', '湖南', '广东', '广西', '海南', '四川', '贵州', '云南',
  '西藏', '陕西', '甘肃', '青海', '宁夏', '新疆'
]

const COLLEGE_LEVELS = ['985', '211', '一流', '一本', '二本', '三本', '专科']

const SAMPLE_COLLEGES = [
  {
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
    employmentRate: 98.5
  },
  {
    collegeCode: '10002',
    collegeName: '北京大学',
    province: '北京',
    city: '北京',
    level: '985',
    type: '综合',
    nature: '公办',
    website: 'https://www.pku.edu.cn',
    description: '北京大学是中国最古老的大学之一，位于北京市海淀区。',
    ranking: 2,
    employmentRate: 98.2
  },
  {
    collegeCode: '10003',
    collegeName: '复旦大学',
    province: '上海',
    city: '上海',
    level: '985',
    type: '综合',
    nature: '公办',
    website: 'https://www.fudan.edu.cn',
    description: '复旦大学位于上海市，是中国顶尖大学之一。',
    ranking: 3,
    employmentRate: 97.8
  },
  {
    collegeCode: '10004',
    collegeName: '上海交通大学',
    province: '上海',
    city: '上海',
    level: '985',
    type: '综合',
    nature: '公办',
    website: 'https://www.sjtu.edu.cn',
    description: '上海交通大学位于上海市，是中国著名理工科大学。',
    ranking: 4,
    employmentRate: 97.5
  }
]

const SAMPLE_ADMISSION_SCORES = [
  {
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
    dataQuality: 5
  },
  {
    year: 2024,
    province: '北京',
    collegeCode: '10002',
    collegeName: '北京大学',
    majorCode: '080901',
    majorName: '计算机科学与技术',
    batch: '一本',
    scoreType: 2,
    minScore: 680,
    minRank: 200,
    avgScore: 685,
    enrollmentCount: 45,
    dataSource: '北京市教育考试院',
    dataQuality: 5
  },
  {
    year: 2024,
    province: '上海',
    collegeCode: '10003',
    collegeName: '复旦大学',
    majorCode: '080901',
    majorName: '计算机科学与技术',
    batch: '一本',
    scoreType: 2,
    minScore: 645,
    minRank: 300,
    avgScore: 650,
    enrollmentCount: 40,
    dataSource: '上海市教育考试院',
    dataQuality: 5
  }
]

// 主函数
exports.main = async (event, context) => {
  console.log('收到数据库初始化请求:', event)

  try {
    const { action, data } = event

    switch (action) {
      case 'init_colleges':
        return await initColleges()
      case 'init_scores':
        return await initAdmissionScores()
      case 'clear_all':
        return await clearAllData()
      case 'import_batch':
        return await importBatchData(data)
      case 'check_status':
        return await checkDatabaseStatus()
      default:
        return {
          code: 400,
          message: '无效的操作类型'
        }
    }

  } catch (error) {
    console.error('数据库初始化失败:', error)

    return {
      code: 500,
      message: error.message || '服务器内部错误',
      error: process.env.NODE_ENV === 'development' ? error.stack : undefined
    }
  }
}

// 初始化学校数据
async function initColleges() {
  console.log('开始初始化学校数据...')

  const result = await db.collection('colleges').add({
    data: SAMPLE_COLLEGES.map(college => ({
      ...college,
      createdAt: db.serverDate(),
      updatedAt: db.serverDate()
    }))
  })

  return {
    code: 200,
    message: `成功初始化 ${SAMPLE_COLLEGES.length} 所学校数据`,
    data: {
      insertedCount: SAMPLE_COLLEGES.length,
      result
    }
  }
}

// 初始化录取分数数据
async function initAdmissionScores() {
  console.log('开始初始化录取分数数据...')

  const result = await db.collection('admission_scores').add({
    data: SAMPLE_ADMISSION_SCORES.map(score => ({
      ...score,
      createdAt: db.serverDate(),
      updatedAt: db.serverDate()
    }))
  })

  return {
    code: 200,
    message: `成功初始化 ${SAMPLE_ADMISSION_SCORES.length} 条录取分数数据`,
    data: {
      insertedCount: SAMPLE_ADMISSION_SCORES.length,
      result
    }
  }
}

// 清空所有数据
async function clearAllData() {
  console.log('开始清空数据库...')

  const collections = ['colleges', 'admission_scores', 'user_profiles', 'volunteer_suggestions']

  const results = {}

  for (const collection of collections) {
    try {
      const result = await db.collection(collection).where({}).remove()
      results[collection] = {
        success: true,
        deletedCount: result.stats.removed
      }
    } catch (error) {
      results[collection] = {
        success: false,
        error: error.message
      }
    }
  }

  return {
    code: 200,
    message: '数据清空完成',
    data: results
  }
}

// 批量导入数据
async function importBatchData(data) {
  if (!data || !Array.isArray(data)) {
    return {
      code: 400,
      message: '导入数据格式错误'
    }
  }

  const { collection, items } = data

  if (!collection || !['colleges', 'admission_scores'].includes(collection)) {
    return {
      code: 400,
      message: '无效的集合名称'
    }
  }

  console.log(`开始批量导入 ${items.length} 条数据到 ${collection}...`)

  // 分批导入，每批最多500条
  const batchSize = 500
  let successCount = 0
  let failCount = 0
  const errors = []

  for (let i = 0; i < items.length; i += batchSize) {
    const batch = items.slice(i, i + batchSize)

    try {
      const batchData = batch.map(item => ({
        ...item,
        createdAt: db.serverDate(),
        updatedAt: db.serverDate()
      }))

      await db.collection(collection).add({
        data: batchData
      })

      successCount += batch.length
    } catch (error) {
      failCount += batch.length
      errors.push({
        batch: Math.floor(i / batchSize) + 1,
        error: error.message,
        itemCount: batch.length
      })
    }
  }

  return {
    code: 200,
    message: `批量导入完成：成功 ${successCount} 条，失败 ${failCount} 条`,
    data: {
      totalCount: items.length,
      successCount,
      failCount,
      errors
    }
  }
}

// 检查数据库状态
async function checkDatabaseStatus() {
  console.log('检查数据库状态...')

  const collections = ['colleges', 'admission_scores', 'user_profiles', 'volunteer_suggestions']
  const status = {}

  for (const collection of collections) {
    try {
      const result = await db.collection(collection).count()
      status[collection] = {
        exists: true,
        count: result.total
      }
    } catch (error) {
      status[collection] = {
        exists: false,
        error: error.message
      }
    }
  }

  // 检查索引
  const indexes = await checkIndexes()

  return {
    code: 200,
    message: '数据库状态检查完成',
    data: {
      collections: status,
      indexes,
      recommendations: generateRecommendations(status)
    }
  }
}

// 检查索引状态
async function checkIndexes() {
  // 这里可以检查重要的索引是否存在
  // 云数据库的索引需要在控制台手动创建
  return {
    admission_scores_query: {
      exists: false, // 需要手动创建
      required: true,
      fields: ['province', 'scoreType', 'minScore', 'collegeName', 'majorName']
    },
    colleges_search: {
      exists: false, // 需要手动创建
      required: true,
      fields: ['collegeName', 'level', 'province']
    }
  }
}

// 生成推荐操作
function generateRecommendations(status) {
  const recommendations = []

  if (status.colleges?.count === 0) {
    recommendations.push('需要初始化学校数据')
  }

  if (status.admission_scores?.count === 0) {
    recommendations.push('需要初始化录取分数数据')
  }

  if (!status.colleges?.exists) {
    recommendations.push('需要创建 colleges 集合')
  }

  if (!status.admission_scores?.exists) {
    recommendations.push('需要创建 admission_scores 集合')
  }

  return recommendations
}
