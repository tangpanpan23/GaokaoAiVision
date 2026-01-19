# 微信小程序云函数说明

## 云函数架构

```
小程序前端
    ↓ (wx.cloud.callFunction)
云函数 (Node.js)
    ↓ (数据库查询/外部API调用)
云数据库 / 外部服务
```

## 已实现的云函数

### 1. getAdmissionScores - 获取录取分数线
- **功能**: 查询历年录取分数线数据
- **参数**:
  ```javascript
  {
    province: "北京",           // 省份
    scoreType: 2,             // 分数类型: 1-文科, 2-理科, 3-综合改革
    minScore: 600,            // 最低分数
    maxScore: 700,            // 最高分数
    collegeName: "清华大学",    // 学校名称关键词
    majorName: "计算机",       // 专业名称关键词
    year: 2024,               // 年份
    page: 1,                  // 页码
    pageSize: 20              // 每页数量
  }
  ```
- **返回**:
  ```javascript
  {
    code: 200,
    data: {
      list: [...],     // 分数线数据
      total: 100,      // 总数量
      page: 1,         // 当前页
      pageSize: 20     // 每页数量
    }
  }
  ```

### 2. generateVolunteerSuggestion - 生成志愿推荐
- **功能**: AI分析生成志愿推荐
- **参数**:
  ```javascript
  {
    province: "北京",
    scoreType: 2,
    score: 650,
    rank: 5000,
    subjects: "物理+历史",
    interestTags: ["计算机", "人工智能"]
  }
  ```
- **返回**:
  ```javascript
  {
    code: 200,
    data: {
      categories: [
        {
          category: "冲",
          reason: "分数有竞争力",
          colleges: [...]
        }
      ],
      analysisSummary: "分析总结",
      recommendations: ["建议1", "建议2"]
    }
  }
  ```

### 3. getCollegeInfo - 获取学校信息
- **功能**: 查询学校详细信息
- **参数**:
  ```javascript
  {
    collegeCode: "10001"
  }
  ```

### 4. getMajorInfo - 获取专业信息
- **功能**: 查询专业详细信息
- **参数**:
  ```javascript
  {
    majorCode: "080901"
  }
  ```

### 5. saveUserProfile - 保存用户档案
- **功能**: 保存用户高考信息
- **参数**:
  ```javascript
  {
    graduationYear: 2024,
    province: "北京",
    scoreType: 2,
    totalScore: 650,
    rank: 5000,
    subjects: "物理+历史",
    interestTags: ["计算机", "金融"],
    personalityType: "INTJ"
  }
  ```

## 云数据库设计

### admission_scores 集合
```javascript
{
  _id: "score_2024_10001_080901",
  year: 2024,
  province: "北京",
  collegeCode: "10001",
  collegeName: "清华大学",
  majorCode: "080901",
  majorName: "计算机科学与技术",
  batch: "一本",
  scoreType: 2,           // 1-文科, 2-理科, 3-综合改革
  minScore: 685,
  minRank: 150,
  avgScore: 690,
  enrollmentCount: 50,
  dataSource: "教育部官网",
  dataQuality: 5,         // 1-5 数据质量评分
  createdAt: "2024-01-01T00:00:00Z",
  updatedAt: "2024-01-01T00:00:00Z"
}
```

### colleges 集合
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
  ranking: 1,
  employmentRate: 98.5,
  createdAt: "2024-01-01T00:00:00Z"
}
```

### majors 集合
```javascript
{
  _id: "major_080901",
  majorCode: "080901",
  majorName: "计算机科学与技术",
  category: "工学",
  subcategory: "计算机类",
  degree: "学士",
  duration: 4,            // 学制(年)
  description: "计算机科学与技术专业...",
  employmentDirection: "软件开发、系统分析...",
  salaryRange: "10k-30k",
  demandLevel: "高",      // 高, 中, 低
  createdAt: "2024-01-01T00:00:00Z"
}
```

### user_profiles 集合
```javascript
{
  _id: "user_{openid}",
  openId: "o6zAJs-_RmfKkS4UlOsiMDLECbFc",
  unionId: "union_123456",
  nickname: "考生小明",
  avatar: "https://...",
  graduationYear: 2024,
  province: "北京",
  scoreType: 2,
  totalScore: 650,
  rank: 5000,
  subjects: "物理+历史",
  interestTags: ["计算机", "人工智能", "大数据"],
  personalityType: "INTJ",
  targetCollege: "清华大学",
  targetMajor: "计算机科学与技术",
  createdAt: "2024-01-01T00:00:00Z",
  updatedAt: "2024-01-01T00:00:00Z"
}
```

### volunteer_suggestions 集合
```javascript
{
  _id: "suggestion_{timestamp}",
  userId: "user_{openid}",
  inputData: {
    province: "北京",
    score: 650,
    // ... 其他输入数据
  },
  result: {
    categories: [...],
    analysisSummary: "...",
    recommendations: [...]
  },
  aiModel: "qwen-turbo",
  createdAt: "2024-01-01T00:00:00Z"
}
```

## AI能力集成

### 支持的AI模型

1. **通义千问 (qwen)**
   - 优势: 中文理解优秀，性价比高
   - 适用: 志愿分析、政策咨询
   - API: https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation

2. **文心一言 (ernie)**
   - 优势: 百度生态，中文优化
   - 适用: 教育场景，专业分析
   - API: https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop

3. **GPT系列**
   - 优势: 全球领先，推理能力强
   - 适用: 复杂分析，多维度评估
   - API: https://api.openai.com/v1/chat/completions

### AI调用策略

#### 1. 模型选择策略
```javascript
function selectAIModel(userInput, context) {
  // 根据输入复杂度选择模型
  if (context.complexity > 0.8) {
    return 'gpt-4';        // 复杂分析用GPT-4
  } else if (context.isEducation) {
    return 'ernie';        // 教育场景用文心一言
  } else {
    return 'qwen';         // 默认用通义千问
  }
}
```

#### 2. Prompt工程设计
```javascript
const VOLUNTEER_ANALYSIS_PROMPT = `
你是一位资深高考志愿规划师。请根据以下信息，为考生生成志愿推荐分析：

考生信息:
- 省份: {{province}}
- 分数类型: {{scoreType}}
- 分数: {{score}}
- 位次: {{rank}}
- 选考科目: {{subjects}}
- 兴趣方向: {{interests}}

匹配的候选专业列表:
{{candidates}}

请从"冲刺、稳妥、保底"三个层次分析，重点考虑:
1. 分数竞争力分析
2. 专业兴趣匹配度
3. 就业前景评估
4. 学校综合实力

请提供详细的分析报告和具体的志愿建议。
`;

const POLICY_CONSULTATION_PROMPT = `
你是一位高考政策专家。请回答考生的志愿填报问题:

问题: {{query}}
考生背景: {{background}}

请提供准确、实用的建议，注意:
1. 基于最新政策法规
2. 结合考生实际情况
3. 突出关键决策点
4. 保持客观中立的态度
`;
```

#### 3. 结果缓存策略
```javascript
// 缓存相似查询的结果
async function getCachedSuggestion(cacheKey, params) {
  const cache = wx.getStorageSync('ai_cache') || {};

  // 检查缓存是否有效 (24小时内)
  if (cache[cacheKey] && Date.now() - cache[cacheKey].timestamp < 24 * 60 * 60 * 1000) {
    return cache[cacheKey].result;
  }

  // 缓存失效，重新调用AI
  const result = await callAI(params);
  cache[cacheKey] = {
    result,
    timestamp: Date.now()
  };
  wx.setStorageSync('ai_cache', cache);

  return result;
}
```

## 性能优化

### 1. 数据分页加载
```javascript
// 分页查询分数线数据
async function getAdmissionScoresPaged(params) {
  const { page = 1, pageSize = 20 } = params;

  return await wx.cloud.database()
    .collection('admission_scores')
    .where(buildQuery(params))
    .orderBy('minScore', 'desc')
    .skip((page - 1) * pageSize)
    .limit(pageSize)
    .get();
}
```

### 2. 索引优化
```javascript
// 云数据库索引配置
const indexes = [
  {
    name: 'admission_scores_query',
    keys: [
      { province: 1 },
      { scoreType: 1 },
      { minScore: -1 },
      { collegeName: 'text' },
      { majorName: 'text' }
    ]
  },
  {
    name: 'colleges_search',
    keys: [
      { collegeName: 'text' },
      { level: 1 },
      { province: 1 }
    ]
  }
];
```

### 3. 数据预加载
```javascript
// 预加载热门学校数据
async function preloadHotData() {
  const hotColleges = ['清华大学', '北京大学', '上海交通大学', '浙江大学'];

  for (const college of hotColleges) {
    const data = await wx.cloud.database()
      .collection('admission_scores')
      .where({ collegeName: college })
      .get();

    // 缓存到本地存储
    wx.setStorageSync(`hot_${college}`, data);
  }
}
```

## 安全考虑

### 1. 数据验证
```javascript
function validateUserInput(data) {
  // 验证分数范围
  if (data.score < 0 || data.score > 750) {
    throw new Error('分数范围无效');
  }

  // 验证省份
  const validProvinces = ['北京', '上海', '广东', /* ... */];
  if (!validProvinces.includes(data.province)) {
    throw new Error('省份信息无效');
  }

  // 验证位次
  if (data.rank < 0) {
    throw new Error('位次不能为负数');
  }
}
```

### 2. API限流
```javascript
class RateLimiter {
  constructor(maxRequests = 10, windowMs = 60000) {
    this.requests = [];
    this.maxRequests = maxRequests;
    this.windowMs = windowMs;
  }

  canMakeRequest() {
    const now = Date.now();
    this.requests = this.requests.filter(time => now - time < this.windowMs);

    if (this.requests.length >= this.maxRequests) {
      return false;
    }

    this.requests.push(now);
    return true;
  }
}
```

### 3. 错误处理
```javascript
async function safeAPICall(apiCall, fallback = null) {
  try {
    return await apiCall();
  } catch (error) {
    console.error('API调用失败:', error);

    // 记录错误日志
    await wx.cloud.callFunction({
      name: 'logError',
      data: {
        error: error.message,
        stack: error.stack,
        context: 'api_call'
      }
    });

    // 返回后备数据
    return fallback;
  }
}
```

## 部署和维护

### 1. 环境配置
```javascript
// 开发环境
const DEV_CONFIG = {
  envId: 'lighthouse-volunteer-dev',
  aiProviders: ['qwen'],
  enableDebug: true
};

// 生产环境
const PROD_CONFIG = {
  envId: 'lighthouse-volunteer-prod',
  aiProviders: ['qwen', 'ernie', 'gpt'],
  enableDebug: false
};
```

### 2. 监控和告警
```javascript
// 性能监控
async function monitorPerformance() {
  const performance = wx.getPerformance();

  // 监控关键指标
  const metrics = {
    firstPaint: performance.getEntriesByType('paint')[0],
    domContentLoaded: performance.getEntriesByType('navigation')[0],
    loadComplete: performance.getEntriesByType('navigation')[0]
  };

  // 上报到监控系统
  await wx.cloud.callFunction({
    name: 'reportMetrics',
    data: metrics
  });
}

// 错误监控
wx.onError && wx.onError((error) => {
  wx.cloud.callFunction({
    name: 'reportError',
    data: {
      message: error.message,
      stack: error.stack,
      userAgent: wx.getSystemInfoSync()
    }
  });
});
```

这个方案提供了完整的高考志愿填报小程序云开发架构，包括数据存储、AI集成、性能优化和安全考虑。
