# 微信小程序云开发部署指南

## 📋 部署概述

"灯塔志愿"高考志愿填报小程序采用微信小程序云开发架构，包括：

- **前端**: 原生微信小程序 (已实现)
- **后端**: 云函数 (Node.js)
- **数据库**: 云数据库 (MongoDB)
- **AI能力**: 集成通义千问、文心一言等AI模型

## 🚀 部署步骤

### 1. 准备工作

#### 1.1 注册微信小程序
1. 访问 [微信公众平台](https://mp.weixin.qq.com)
2. 注册小程序账号
3. 获取 AppID (例如: `wxacc380c307b5192a`)

#### 1.2 开通云开发
1. 登录 [微信开发者工具](https://developers.weixin.qq.com/miniprogram/dev/devtools/download.html)
2. 创建新项目，选择"小程序·云开发"
3. 选择云环境 (推荐创建新环境)

### 2. 云环境配置

#### 2.1 创建云环境
```bash
# 在微信开发者工具中：
1. 点击"云开发"按钮
2. 创建新环境 (如: lighthouse-volunteer-dev)
3. 选择环境规格 (推荐: 基础版 2GB)
4. 开启数据库、存储、云函数
```

#### 2.2 配置环境变量
在云控制台中设置环境变量：

**通义千问 API 配置:**
- `QWEN_API_KEY`: 从阿里云控制台获取
- `QWEN_BASE_URL`: https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation

**文心一言 API 配置:**
- `ERNIE_API_KEY`: 从百度智能云获取
- `ERNIE_SECRET_KEY`: 从百度智能云获取

### 3. 数据库设计

#### 3.1 创建集合
在云控制台数据库管理中创建以下集合：

```javascript
// admission_scores - 录取分数线集合
{
  _id: "score_2024_10001_080901",
  year: 2024,
  province: "北京",
  collegeCode: "10001",
  collegeName: "清华大学",
  majorCode: "080901",
  majorName: "计算机科学与技术",
  batch: "一本",
  scoreType: 2,
  minScore: 685,
  minRank: 150,
  avgScore: 690,
  enrollmentCount: 50,
  dataSource: "教育部官网",
  dataQuality: 5,
  createdAt: "2024-01-01T00:00:00Z",
  updatedAt: "2024-01-01T00:00:00Z"
}

// colleges - 学校信息集合
{
  _id: "college_10001",
  collegeCode: "10001",
  collegeName: "清华大学",
  province: "北京",
  level: "985",
  type: "综合",
  nature: "公办",
  ranking: 1,
  employmentRate: 98.5,
  description: "清华大学是中国著名高等学府...",
  website: "https://www.tsinghua.edu.cn",
  createdAt: "2024-01-01T00:00:00Z"
}

// user_profiles - 用户档案集合
{
  _id: "user_{openid}",
  userId: "o6zAJs-_RmfKkS4UlOsiMDLECbFc",
  graduationYear: 2024,
  province: "北京",
  scoreType: 2,
  totalScore: 650,
  rank: 5000,
  subjects: "物理+历史",
  interestTags: ["计算机", "人工智能"],
  personalityType: "INTJ",
  createdAt: "2024-01-01T00:00:00Z",
  updatedAt: "2024-01-01T00:00:00Z"
}
```

#### 3.2 创建索引
为提升查询性能，创建以下索引：

**admission_scores 集合索引:**
```javascript
// 复合查询索引
{
  province: 1,
  scoreType: 1,
  minScore: -1,
  collegeName: "text",
  majorName: "text"
}

// 年份索引
{
  year: -1
}

// 分数范围索引
{
  minScore: -1
}
```

**colleges 集合索引:**
```javascript
// 搜索索引
{
  collegeName: "text",
  level: 1,
  province: 1
}

// 排名索引
{
  ranking: 1
}
```

### 4. 云函数部署

#### 4.1 上传云函数
```bash
# 在微信开发者工具中：
1. 右键 cloud-functions 文件夹
2. 选择"上传并部署：云端安装依赖"
3. 等待部署完成
```

#### 4.2 验证云函数
在云控制台中测试云函数：

```javascript
// 测试 getAdmissionScores
wx.cloud.callFunction({
  name: 'getAdmissionScores',
  data: {
    province: '北京',
    scoreType: 2,
    page: 1,
    pageSize: 10
  }
}).then(res => {
  console.log('测试结果:', res)
})
```

### 5. 数据导入

#### 5.1 初始化基础数据
使用 `initDatabase` 云函数导入基础数据：

```javascript
// 初始化学校数据
wx.cloud.callFunction({
  name: 'initDatabase',
  data: {
    action: 'init_colleges'
  }
})

// 初始化分数数据
wx.cloud.callFunction({
  name: 'initDatabase',
  data: {
    action: 'init_scores'
  }
})
```

#### 5.2 批量导入高考数据
```javascript
// 准备分数线数据
const scoreData = [
  {
    year: 2024,
    province: '北京',
    collegeCode: '10001',
    collegeName: '清华大学',
    majorName: '计算机科学与技术',
    scoreType: 2,
    minScore: 685,
    minRank: 150
  }
  // ... 更多数据
]

// 批量导入
wx.cloud.callFunction({
  name: 'initDatabase',
  data: {
    action: 'import_batch',
    data: {
      collection: 'admission_scores',
      items: scoreData
    }
  }
})
```

### 6. AI能力配置

#### 6.1 通义千问配置
1. 访问 [阿里云控制台](https://dashscope.aliyuncs.com/)
2. 开通 DashScope 服务
3. 创建 API Key
4. 在云环境变量中配置 `QWEN_API_KEY`

#### 6.2 文心一言配置
1. 访问 [百度智能云](https://console.bce.baidu.com/)
2. 开通文心一言服务
3. 创建应用获取 API Key 和 Secret Key
4. 在云环境变量中配置 `ERNIE_API_KEY` 和 `ERNIE_SECRET_KEY`

#### 6.3 AI调用测试
```javascript
// 测试AI志愿推荐
wx.cloud.callFunction({
  name: 'generateVolunteerSuggestion',
  data: {
    province: '北京',
    scoreType: 2,
    score: 650,
    rank: 5000,
    subjects: '物理+历史',
    interestTags: ['计算机', '人工智能']
  }
}).then(res => {
  console.log('AI推荐结果:', res)
})
```

### 7. 前端配置

#### 7.1 更新配置
在 `miniprogram/project.config.json` 中：

```json
{
  "appid": "你的小程序AppID",
  "cloudfunctionRoot": "./cloud-functions/",
  "setting": {
    // ... 其他配置
  }
}
```

#### 7.2 云环境初始化
在 `app.js` 中添加云环境初始化：

```javascript
// app.js
App({
  onLaunch: function() {
    // 初始化云开发环境
    wx.cloud.init({
      env: 'lighthouse-volunteer-dev',  // 云环境ID
      traceUser: true
    })

    // 其他初始化代码...
  },

  // 全局方法
  getUserInfo: function() {
    return this.globalData.userInfo
  }
})
```

### 8. 测试验证

#### 8.1 功能测试
1. **用户登录**: 测试微信授权登录
2. **档案管理**: 测试用户档案的创建和更新
3. **分数查询**: 测试录取分数线查询功能
4. **志愿推荐**: 测试AI志愿推荐功能
5. **学校信息**: 测试学校信息查询

#### 8.2 性能测试
1. **响应时间**: 确保API响应时间 < 3秒
2. **并发处理**: 测试同时请求处理能力
3. **数据查询**: 验证复杂查询的性能

#### 8.3 错误处理测试
1. **网络异常**: 测试网络断开情况
2. **API限流**: 测试高频请求处理
3. **数据验证**: 测试非法数据处理

### 9. 监控和维护

#### 9.1 云控制台监控
- 实时监控云函数调用次数和耗时
- 查看数据库读写统计
- 监控存储使用情况

#### 9.2 错误日志
```javascript
// 在云函数中记录错误
try {
  // 业务逻辑
} catch (error) {
  console.error('业务错误:', error)

  // 记录到数据库
  await wx.cloud.database()
    .collection('error_logs')
    .add({
      data: {
        functionName: 'functionName',
        error: error.message,
        stack: error.stack,
        userId: event.userInfo?.openId,
        timestamp: new Date()
      }
    })
}
```

#### 9.3 定期维护
- **数据更新**: 每月更新高考录取数据
- **AI模型**: 定期检查和更新AI模型配置
- **性能优化**: 定期检查和优化数据库索引
- **安全更新**: 及时更新API密钥和安全配置

### 10. 成本控制

#### 10.1 资源使用估算
- **云数据库**: 按读写次数收费，预计每月 ¥50-200
- **云函数**: 按调用次数收费，预计每月 ¥20-100
- **云存储**: 按存储量收费，预计每月 ¥10-50
- **AI API**: 按调用次数收费，预计每月 ¥100-500

#### 10.2 成本优化策略
1. **数据缓存**: 使用小程序本地存储减少API调用
2. **批量查询**: 合理使用分页和批量查询
3. **智能缓存**: 缓存热点数据，减少数据库查询
4. **限流控制**: 避免恶意高频请求

## 📚 常见问题

### Q: 云函数部署失败？
A: 检查 package.json 中的依赖版本，确保 Node.js 版本兼容。

### Q: 数据库查询超时？
A: 检查索引配置，优化查询条件，使用分页查询。

### Q: AI接口调用失败？
A: 检查API密钥配置，确认账户余额，查看API限流限制。

### Q: 数据导入缓慢？
A: 使用批量导入，分批次上传，避免单次上传过多数据。

## 🔗 相关链接

- [微信小程序云开发文档](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/basis/getting-started.html)
- [云数据库使用指南](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/database.html)
- [云函数开发指南](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/functions.html)
- [阿里云通义千问](https://dashscope.aliyuncs.com/)
- [百度文心一言](https://yiyan.baidu.com/)

---

**部署完成时间**: 2026-01-19
**文档版本**: v1.0
