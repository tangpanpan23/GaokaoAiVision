# 灯塔志愿 - 高考志愿填报小程序技术方案

## 🎯 项目概述

"灯塔志愿"是一款基于微信小程序云开发的智能高考志愿填报指导平台，致力于为考生提供公平、透明、个性化的志愿填报服务。

## 🏗️ 技术架构

### 整体架构图

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   小程序前端     │    │    云函数后端     │    │    云数据库      │
│                 │    │                 │    │                 │
│  • 原生小程序   │◄──►│  • Node.js函数   │◄──►│  • MongoDB      │
│  • WXSS样式     │    │  • 业务逻辑处理  │    │  • 文档存储      │
│  • WXML模板     │    │  • AI能力集成    │    │  • 自动索引      │
│  • JS逻辑       │    │  • 数据验证      │    │  • 实时同步      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
       │                       │                       │
       ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   外部AI服务     │    │   数据缓存       │    │   监控告警       │
│                 │    │                 │    │                 │
│  • 通义千问      │    │  • 本地存储      │    │  • 云控制台      │
│  • 文心一言      │    │  • 智能缓存      │    │  • 错误日志      │
│  • GPT系列       │    │  • 离线支持      │    │  • 性能监控      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 📊 数据存储方案

### 数据库设计原则

1. **文档化存储**: 使用MongoDB文档模型，支持灵活的数据结构
2. **分层设计**: 按功能和访问频率分层存储
3. **索引优化**: 为查询频繁的字段创建合适的索引
4. **数据完整性**: 通过应用层保证数据一致性
5. **扩展性**: 支持水平扩展和垂直扩展

### 核心集合设计

#### 1. admission_scores - 录取分数线集合

**数据结构**:
```javascript
{
  _id: "score_2024_10001_080901",  // 唯一标识
  year: 2024,                        // 年份
  province: "北京",                   // 省份
  collegeCode: "10001",               // 学校代码
  collegeName: "清华大学",            // 学校名称
  majorCode: "080901",                // 专业代码
  majorName: "计算机科学与技术",      // 专业名称
  batch: "一本",                      // 录取批次
  scoreType: 2,                       // 分数类型: 1-文科, 2-理科, 3-综合改革
  minScore: 685,                      // 最低分
  minRank: 150,                       // 最低位次
  avgScore: 690,                      // 平均分
  enrollmentCount: 50,                // 录取人数
  dataSource: "教育部官网",           // 数据来源
  dataQuality: 5,                     // 数据质量评分(1-5)
  createdAt: "2024-01-01T00:00:00Z", // 创建时间
  updatedAt: "2024-01-01T00:00:00Z"  // 更新时间
}
```

**索引设计**:
```javascript
// 复合查询索引 - 核心查询
{
  province: 1,
  scoreType: 1,
  minScore: -1,
  collegeName: "text",
  majorName: "text"
}

// 年份索引 - 时间筛选
{
  year: -1
}

// 分数范围索引 - 范围查询
{
  minScore: -1,
  maxScore: -1
}
```

#### 2. colleges - 学校信息集合

**数据结构**:
```javascript
{
  _id: "college_10001",
  collegeCode: "10001",
  collegeName: "清华大学",
  province: "北京",
  city: "北京",
  level: "985",           // 985, 211, 双一流, 普通本科
  type: "综合",           // 综合, 理工, 师范, 农业, etc.
  nature: "公办",         // 公办, 民办
  website: "https://www.tsinghua.edu.cn",
  description: "清华大学是中国著名高等学府...",
  ranking: 1,             // 综合排名
  employmentRate: 98.5,   // 就业率
  createdAt: "2024-01-01T00:00:00Z",
  updatedAt: "2024-01-01T00:00:00Z"
}
```

**索引设计**:
```javascript
// 搜索索引 - 全文搜索
{
  collegeName: "text",
  level: 1,
  province: 1
}

// 排名索引 - 排序查询
{
  ranking: 1
}
```

#### 3. user_profiles - 用户档案集合

**数据结构**:
```javascript
{
  _id: "user_o6zAJs-_RmfKkS4UlOsiMDLECbFc",
  userId: "o6zAJs-_RmfKkS4UlOsiMDLECbFc",  // 微信OpenID
  unionId: "union_123456",                    // 微信UnionID
  nickname: "考生小明",                        // 昵称
  avatar: "https://...",                       // 头像
  graduationYear: 2024,                        // 毕业年份
  province: "北京",                             // 省份
  scoreType: 2,                                // 分数类型
  totalScore: 650,                             // 总分
  rank: 5000,                                  // 位次
  subjects: "物理+历史",                       // 选考科目
  interestTags: ["计算机", "人工智能"],        // 兴趣标签
  personalityType: "INTJ",                     // 性格类型
  targetCollege: "清华大学",                    // 目标学校
  targetMajor: "计算机科学与技术",              // 目标专业
  createdAt: "2024-01-01T00:00:00Z",
  updatedAt: "2024-01-01T00:00:00Z"
}
```

**索引设计**:
```javascript
// 用户ID索引 - 主键查询
{
  userId: 1
}

// 省份分数索引 - 推荐查询
{
  province: 1,
  scoreType: 1,
  totalScore: -1
}
```

#### 4. volunteer_suggestions - 志愿推荐记录集合

**数据结构**:
```javascript
{
  _id: "suggestion_1737234567890",
  userId: "o6zAJs-_RmfKkS4UlOsiMDLECbFc",
  inputData: {                           // 输入数据
    province: "北京",
    score: 650,
    rank: 5000,
    subjects: "物理+历史",
    interestTags: ["计算机", "人工智能"]
  },
  result: {                             // AI分析结果
    categories: [...],                  // 推荐分类
    analysisSummary: "...",             // 分析总结
    recommendations: [...]              // 建议列表
  },
  aiModel: "qwen-turbo",               // 使用的AI模型
  responseTime: 2500,                  // 响应时间(ms)
  createdAt: "2024-01-01T00:00:00Z"
}
```

## 🤖 AI能力集成方案

### AI模型选择策略

#### 1. 多模型并存
```javascript
function selectAIModel(userInput, context) {
  const complexity = calculateComplexity(userInput)

  // 高复杂度使用文心一言
  if (complexity > 0.8 && AI_CONFIG.ernie.enabled) {
    return 'ernie'
  }

  // 中等复杂度使用通义千问
  if (AI_CONFIG.qwen.enabled) {
    return 'qwen'
  }

  // 默认使用通义千问
  return 'qwen'
}
```

#### 2. 模型能力对比

| 特性 | 通义千问 (qwen) | 文心一言 (ernie) | GPT-4 |
|------|----------------|------------------|-------|
| 性价比 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ |
| 中文理解 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 推理能力 | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| 教育场景 | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 响应速度 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 成本控制 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐ |

### Prompt工程设计

#### 志愿推荐Prompt模板

```javascript
const VOLUNTEER_ANALYSIS_PROMPT = `
你是一位资深高考志愿规划师，具有15年以上志愿填报指导经验。请根据以下考生信息进行个性化志愿推荐分析。

【考生基本信息】
- 省份: {{province}}
- 分数类型: {{scoreType}} (1=文科, 2=理科, 3=综合改革)
- 高考分数: {{score}}分
- 位次排名: {{rank}}名
- 选考科目: {{subjects}}
- 兴趣方向: {{interests}}

【匹配候选专业】
{{candidates}}

【分析要求】
请从以下维度进行深度分析：

1. **分数竞争力评估**
   - 基于历史录取数据分析分数优势
   - 考虑省份竞争力和批次差异
   - 评估录取成功率

2. **专业兴趣匹配度**
   - 根据兴趣标签匹配适合专业
   - 分析专业就业前景和发展空间
   - 考虑专业学习难度和就业竞争力

3. **学校综合实力**
   - 评估学校学术实力和师资力量
   - 分析地理位置和校园环境
   - 考虑学校特色和优势专业

4. **职业发展建议**
   - 分析专业就业方向和薪资水平
   - 评估职业发展前景和稳定性
   - 建议职业规划路径

【输出格式】
请严格按照以下JSON格式输出：

{
  "categories": [
    {
      "category": "冲刺志愿",
      "reason": "分数有明显竞争力，值得冲击理想学校",
      "colleges": [
        {
          "collegeCode": "10001",
          "collegeName": "清华大学",
          "majorName": "计算机科学与技术",
          "minScore": 685,
          "minRank": 150,
          "matchingScore": 0.85,
          "advantages": "顶尖计算机专业，师资力量雄厚，就业前景极佳",
          "considerations": "竞争极为激烈，需要全力以赴",
          "recommendationLevel": "A+"
        }
      ]
    }
  ],
  "analysisSummary": "综合分析：你的分数在{{province}}省具有较强竞争力，建议重点关注计算机、人工智能等热门专业...",
  "recommendations": [
    "优先考虑985/211高校的王牌专业",
    "结合个人兴趣选择专业方向，避免盲目跟风",
    "关注学校就业数据和专业发展前景",
    "合理安排冲稳保志愿梯度，确保录取率",
    "建议咨询专业老师获取针对性指导"
  ],
  "riskAssessment": {
    "overallRisk": "中等",
    "riskFactors": ["专业竞争激烈", "就业压力较大"],
    "mitigationStrategies": ["提前实习积累经验", "保持学习能力"]
  }
}

【重要提醒】
- 所有分析仅供参考，最终决策需结合个人实际情况
- 志愿填报政策可能发生变化，请以官方信息为准
- 建议在专业人士指导下完成志愿填报
- 保持积极心态，相信选择比努力更重要
`
```

### AI调用优化策略

#### 1. 缓存机制
```javascript
class AISuggestionCache {
  constructor() {
    this.cache = new Map()
    this.maxAge = 24 * 60 * 60 * 1000  // 24小时
  }

  // 生成缓存键
  generateKey(params) {
    const keyData = {
      province: params.province,
      score: Math.floor(params.score / 5) * 5,  // 分数按5分区间缓存
      scoreType: params.scoreType,
      subjects: params.subjects,
      interests: params.interestTags?.sort().join(',')
    }
    return JSON.stringify(keyData)
  }

  // 获取缓存
  get(params) {
    const key = this.generateKey(params)
    const cached = this.cache.get(key)

    if (cached && Date.now() - cached.timestamp < this.maxAge) {
      return cached.data
    }

    return null
  }

  // 设置缓存
  set(params, data) {
    const key = this.generateKey(params)
    this.cache.set(key, {
      data,
      timestamp: Date.now()
    })

    // 限制缓存大小
    if (this.cache.size > 1000) {
      const oldestKey = this.cache.keys().next().value
      this.cache.delete(oldestKey)
    }
  }
}
```

#### 2. 并发控制
```javascript
class AIRequestLimiter {
  constructor(maxConcurrent = 3, interval = 1000) {
    this.maxConcurrent = maxConcurrent
    this.interval = interval
    this.running = 0
    this.queue = []
  }

  async call(apiCall) {
    return new Promise((resolve, reject) => {
      this.queue.push({ apiCall, resolve, reject })
      this.processQueue()
    })
  }

  async processQueue() {
    if (this.running >= this.maxConcurrent || this.queue.length === 0) {
      return
    }

    this.running++
    const { apiCall, resolve, reject } = this.queue.shift()

    try {
      const result = await apiCall()
      resolve(result)
    } catch (error) {
      reject(error)
    } finally {
      this.running--
      // 延迟处理下一个请求
      setTimeout(() => this.processQueue(), this.interval)
    }
  }
}
```

## 🔒 安全与性能优化

### 数据安全设计

#### 1. 用户隐私保护
- 微信授权获取最小必要信息
- 用户数据加密存储
- 敏感信息脱敏处理
- 符合GDPR和国内隐私保护要求

#### 2. API安全防护
- 请求频率限制 (Rate Limiting)
- 参数验证和过滤
- 错误信息脱敏
- 异常监控和告警

#### 3. 数据传输安全
- HTTPS强制传输
- 请求签名验证
- 敏感数据加密

### 性能优化策略

#### 1. 数据库优化
- **索引策略**: 为查询频繁字段创建复合索引
- **查询优化**: 使用分页查询，避免全表扫描
- **缓存机制**: 热点数据缓存，减少数据库压力
- **读写分离**: 考虑读写分离架构

#### 2. 前端优化
- **数据分页**: 大数据量分页加载
- **懒加载**: 图片和内容按需加载
- **缓存策略**: 本地存储缓存用户数据
- **预加载**: 预测用户行为预加载数据

#### 3. AI服务优化
- **模型选择**: 根据复杂度选择合适的AI模型
- **结果缓存**: 相似查询结果缓存
- **并发控制**: 限制同时进行的AI请求数量
- **降级策略**: AI服务异常时的降级处理

## 📈 监控与运维

### 监控指标

#### 1. 业务指标
- DAU (日活跃用户数)
- 志愿推荐成功率
- 用户留存率
- 功能使用统计

#### 2. 技术指标
- API响应时间
- 数据库查询性能
- 云函数执行时间
- 错误率统计

#### 3. AI服务指标
- AI接口调用成功率
- 平均响应时间
- 模型切换统计
- 缓存命中率

### 告警机制

#### 1. 错误告警
- API错误率超过阈值
- 数据库连接异常
- AI服务调用失败
- 用户反馈异常

#### 2. 性能告警
- 响应时间过长
- 资源使用过高
- 并发请求过多

### 数据分析

#### 1. 用户行为分析
- 功能使用偏好
- 志愿填报时间分布
- 地域分布统计
- 专业选择趋势

#### 2. 系统性能分析
- 峰值使用时间
- 数据库查询热点
- 缓存效果评估
- 成本效益分析

## 🚀 扩展性设计

### 水平扩展

#### 1. 云函数扩展
- 多实例部署
- 负载均衡
- 自动扩缩容

#### 2. 数据库扩展
- 分片存储
- 读写分离
- 多地域部署

#### 3. AI服务扩展
- 多模型支持
- 服务降级
- 第三方AI集成

### 功能扩展

#### 1. 高级功能
- 实时志愿填报指导
- 专业导师一对一咨询
- 校友就业数据分析
- 学校开放日信息

#### 2. 数据源扩展
- 多渠道数据采集
- 用户生成内容
- 第三方数据合作
- 实时数据更新

#### 3. AI能力扩展
- 多轮对话咨询
- 个性化学习推荐
- 职业规划建议
- 心理辅导支持

## 💰 成本控制

### 资源成本估算

| 资源类型 | 预估月成本 | 说明 |
|----------|-----------|------|
| 云数据库 | ¥50-200 | 按读写次数收费 |
| 云函数 | ¥20-100 | 按调用次数收费 |
| 云存储 | ¥10-50 | 按存储量收费 |
| AI API | ¥100-500 | 通义千问+文心一言 |
| CDN | ¥20-100 | 静态资源加速 |

**总计**: ¥200-950/月

### 成本优化策略

#### 1. 技术优化
- 合理使用缓存减少API调用
- 优化数据库查询减少资源消耗
- 使用压缩算法减少数据传输

#### 2. 业务优化
- 合理定价策略平衡收入和成本
- 用户分层服务降低高成本用户比例
- 批量处理减少单次请求成本

#### 3. 运营优化
- 精准营销提高用户质量
- 数据分析指导产品优化
- 定期review和调整资源配置

## 🎯 技术选型理由

### 微信小程序云开发
- **快速开发**: 无需服务器运维，一站式开发
- **弹性伸缩**: 按需付费，避免资源浪费
- **微信生态**: 天然的用户流量入口
- **数据安全**: 符合微信安全标准

### MongoDB云数据库
- **文档模型**: 灵活的数据结构设计
- **自动索引**: 智能索引推荐和优化
- **高可用**: 多副本集确保数据安全
- **易扩展**: 支持水平扩展和分片

### 多AI模型集成
- **能力互补**: 不同模型优势互补
- **成本平衡**: 根据使用场景选择性价比最优模型
- **服务稳定**: 单模型故障时自动切换
- **持续优化**: 定期评估和更新模型选择策略

---

**技术方案版本**: v1.0
**最后更新时间**: 2026-01-19
**文档作者**: AI Assistant
