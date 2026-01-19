// 云函数：generateVolunteerSuggestion
// 生成志愿推荐分析

const cloud = require('wx-server-sdk')
const axios = require('axios')

// 初始化云环境
cloud.init({
  env: cloud.DYNAMIC_CURRENT_ENV
})

// 云数据库
const db = cloud.database()
const _ = db.command

// AI API配置
const AI_CONFIG = {
  qwen: {
    apiKey: process.env.QWEN_API_KEY,
    baseURL: 'https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation',
    model: 'qwen-turbo'
  },
  ernie: {
    apiKey: process.env.ERNIE_API_KEY,
    secretKey: process.env.ERNIE_SECRET_KEY,
    accessToken: null,
    baseURL: 'https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop',
    model: 'chat/completions'
  }
}

// 获取AI访问令牌
async function getErnieAccessToken() {
  if (AI_CONFIG.ernie.accessToken) {
    return AI_CONFIG.ernie.accessToken
  }

  try {
    const response = await axios.get('https://aip.baidubce.com/oauth/2.0/token', {
      params: {
        grant_type: 'client_credentials',
        client_id: AI_CONFIG.ernie.apiKey,
        client_secret: AI_CONFIG.ernie.secretKey
      }
    })

    AI_CONFIG.ernie.accessToken = response.data.access_token
    return AI_CONFIG.ernie.accessToken
  } catch (error) {
    console.error('获取文心一言访问令牌失败:', error)
    throw error
  }
}

// 选择最合适的AI模型
function selectAIModel(userInput, context) {
  // 简单策略：根据输入复杂度选择模型
  const complexity = calculateComplexity(userInput)

  if (complexity > 0.8 && AI_CONFIG.ernie.apiKey) {
    return 'ernie'  // 复杂分析用文心一言
  } else if (AI_CONFIG.qwen.apiKey) {
    return 'qwen'   // 默认用通义千问
  } else {
    throw new Error('没有可用的AI模型')
  }
}

// 计算输入复杂度
function calculateComplexity(input) {
  let complexity = 0

  // 分数越高复杂度越大
  if (input.score > 680) complexity += 0.3
  else if (input.score > 650) complexity += 0.2
  else if (input.score > 600) complexity += 0.1

  // 有位次信息复杂度增加
  if (input.rank > 0) complexity += 0.2

  // 兴趣标签越多复杂度越大
  if (input.interestTags && input.interestTags.length > 2) complexity += 0.2

  return Math.min(complexity, 1)
}

// 调用通义千问API
async function callQwen(prompt) {
  const response = await axios.post(AI_CONFIG.qwen.baseURL, {
    model: AI_CONFIG.qwen.model,
    input: {
      messages: [
        {
          role: 'system',
          content: '你是一位资深高考志愿规划师，具有丰富的高考志愿填报经验。请基于数据和事实给出专业、客观的建议。'
        },
        {
          role: 'user',
          content: prompt
        }
      ]
    },
    parameters: {
      temperature: 0.7,
      max_tokens: 2000
    }
  }, {
    headers: {
      'Authorization': `Bearer ${AI_CONFIG.qwen.apiKey}`,
      'Content-Type': 'application/json'
    }
  })

  return response.data.output.choices[0].message.content
}

// 调用文心一言API
async function callErnie(prompt) {
  const accessToken = await getErnieAccessToken()

  const response = await axios.post(`${AI_CONFIG.ernie.baseURL}/chat/completions`, {
    model: 'ernie-4.0-8k',
    messages: [
      {
        role: 'user',
        content: prompt
      }
    ],
    temperature: 0.7,
    max_tokens: 2000
  }, {
    headers: {
      'Authorization': `Bearer ${accessToken}`,
      'Content-Type': 'application/json'
    }
  })

  return response.data.result
}

// 生成志愿推荐Prompt
function generateVolunteerPrompt(userInput, candidates) {
  return `
你是一位资深高考志愿规划师。请根据以下考生信息和候选学校数据，生成详细的志愿推荐分析报告。

【考生信息】
- 省份: ${userInput.province}
- 分数类型: ${userInput.scoreType === 1 ? '文科' : userInput.scoreType === 2 ? '理科' : '综合改革'}
- 高考分数: ${userInput.score}分
${userInput.rank > 0 ? `- 位次: ${userInput.rank}名` : ''}
- 选考科目: ${userInput.subjects || '未指定'}
- 兴趣方向: ${userInput.interestTags ? userInput.interestTags.join('、') : '未指定'}

【候选学校数据】
${candidates.map((item, index) => `
${index + 1}. ${item.collegeName} ${item.majorName ? `(${item.majorName})` : ''}
   - 录取批次: ${item.batch}
   - 最低分数: ${item.minScore}分
   ${item.minRank ? `- 最低位次: ${item.minRank}名` : ''}
   - 年份: ${item.year}年
`).join('')}

【分析要求】
请从"冲刺志愿"、"稳妥志愿"、"保底志愿"三个层次进行分析：

1. **冲刺志愿**: 比考生分数高10-20分的学校，冲一冲有希望
2. **稳妥志愿**: 与考生分数相近的学校，录取把握较大
3. **保底志愿**: 比考生分数低10-20分的学校，确保录取

对于每个层次，请：
- 列出3-5所推荐学校及其理由
- 分析录取概率和风险
- 考虑专业与兴趣的匹配度
- 评估学校综合实力和就业前景

【输出格式】
请严格按照以下JSON格式输出：

{
  "categories": [
    {
      "category": "冲",
      "reason": "分数略有竞争力，值得冲刺",
      "colleges": [
        {
          "collegeCode": "10001",
          "collegeName": "清华大学",
          "majorName": "计算机科学与技术",
          "minScore": 685,
          "minRank": 150,
          "matchingScore": 0.7,
          "advantages": "综合实力强，就业前景好",
          "considerations": "竞争激烈，需要优秀表现"
        }
      ]
    },
    {
      "category": "稳",
      "reason": "分数匹配度高，录取把握大",
      "colleges": [...]
    },
    {
      "category": "保",
      "reason": "分数优势明显，确保录取",
      "colleges": [...]
    }
  ],
  "analysisSummary": "总体分析总结...",
  "recommendations": [
    "建议1: 具体建议内容",
    "建议2: 具体建议内容",
    "建议3: 具体建议内容"
  ]
}

请确保输出的是有效的JSON格式，不要包含其他文本。
`
}

// 获取候选学校数据
async function getCandidateColleges(userInput) {
  const { province, scoreType, score, subjects } = userInput

  // 查询分数相近的学校（上下30分范围）
  const minScore = Math.max(0, score - 30)
  const maxScore = score + 30

  const result = await db.collection('admission_scores')
    .where({
      province: province,
      scoreType: scoreType,
      minScore: _.gte(minScore).and(_.lte(maxScore))
    })
    .orderBy('minScore', 'desc')
    .limit(50)
    .get()

  return result.data || []
}

// 处理志愿推荐
async function processVolunteerSuggestion(userInput) {
  try {
    // 1. 获取候选学校数据
    const candidates = await getCandidateColleges(userInput)

    if (candidates.length === 0) {
      throw new Error('没有找到匹配的学校数据，请检查输入信息')
    }

    // 2. 选择AI模型
    const context = {
      complexity: calculateComplexity(userInput),
      isEducation: true
    }
    const selectedModel = selectAIModel(userInput, context)

    // 3. 生成AI分析Prompt
    const prompt = generateVolunteerPrompt(userInput, candidates)

    // 4. 调用AI接口
    let aiResponse
    if (selectedModel === 'qwen') {
      aiResponse = await callQwen(prompt)
    } else if (selectedModel === 'ernie') {
      aiResponse = await callErnie(prompt)
    } else {
      throw new Error('不支持的AI模型')
    }

    // 5. 解析AI响应
    let analysisResult
    try {
      // 尝试提取JSON内容
      const jsonMatch = aiResponse.match(/\{[\s\S]*\}/)
      if (jsonMatch) {
        analysisResult = JSON.parse(jsonMatch[0])
      } else {
        throw new Error('AI响应格式错误')
      }
    } catch (parseError) {
      console.error('解析AI响应失败:', parseError)
      // 如果解析失败，返回默认结构
      analysisResult = {
        categories: [
          {
            category: '稳',
            reason: '基于分数分析的推荐',
            colleges: candidates.slice(0, 5).map(item => ({
              collegeCode: item.collegeCode,
              collegeName: item.collegeName,
              majorName: item.majorName,
              minScore: item.minScore,
              minRank: item.minRank,
              matchingScore: 0.8,
              advantages: '综合实力良好',
              considerations: '建议认真考虑'
            }))
          }
        ],
        analysisSummary: 'AI分析服务暂时不可用，返回基础推荐',
        recommendations: [
          '建议结合个人兴趣选择专业',
          '关注学校就业前景和专业实力',
          '合理安排冲稳保志愿梯度'
        ]
      }
    }

    // 6. 后处理结果
    const processedResult = processAnalysisResult(analysisResult, candidates)

    return {
      code: 200,
      data: processedResult,
      model: selectedModel
    }

  } catch (error) {
    console.error('志愿推荐处理失败:', error)
    throw error
  }
}

// 后处理分析结果
function processAnalysisResult(result, candidates) {
  // 为每个分类添加颜色标识
  const categoryColors = {
    '冲': '#dc3545',
    '稳': '#28a745',
    '保': '#ffc107'
  }

  if (result.categories) {
    result.categories.forEach(category => {
      category.categoryColor = categoryColors[category.category] || '#666666'

      // 为每个学校添加详细信息
      if (category.colleges) {
        category.colleges.forEach(college => {
          // 从候选数据中查找详细信息
          const candidateInfo = candidates.find(c =>
            c.collegeCode === college.collegeCode &&
            c.majorName === college.majorName
          )

          if (candidateInfo) {
            college.batch = candidateInfo.batch
            college.year = candidateInfo.year
            college.matchingScorePercent = Math.round((college.matchingScore || 0) * 100)
          }
        })
      }
    })
  }

  return result
}

// 主函数
exports.main = async (event, context) => {
  console.log('收到志愿推荐请求:', event)

  try {
    // 参数验证
    const { province, scoreType, score, rank, subjects, interestTags } = event

    if (!province || !scoreType || !score) {
      return {
        code: 400,
        message: '缺少必要参数：province, scoreType, score'
      }
    }

    if (score < 0 || score > 750) {
      return {
        code: 400,
        message: '分数范围无效，应在0-750之间'
      }
    }

    // 处理用户输入
    const userInput = {
      province,
      scoreType: parseInt(scoreType),
      score: parseInt(score),
      rank: parseInt(rank) || 0,
      subjects: subjects || '',
      interestTags: Array.isArray(interestTags) ? interestTags : []
    }

    // 生成志愿推荐
    const result = await processVolunteerSuggestion(userInput)

    // 保存推荐记录（可选）
    try {
      await db.collection('volunteer_suggestions').add({
        data: {
          userId: event.userInfo?.openId || 'anonymous',
          inputData: userInput,
          result: result.data,
          aiModel: result.model,
          createdAt: db.serverDate()
        }
      })
    } catch (saveError) {
      console.error('保存推荐记录失败:', saveError)
      // 不影响主要功能
    }

    return result

  } catch (error) {
    console.error('云函数执行失败:', error)

    return {
      code: 500,
      message: error.message || '服务器内部错误',
      error: process.env.NODE_ENV === 'development' ? error.stack : undefined
    }
  }
}
