# 微信云开发最佳实践指南

基于微信云开发官方文档的最佳实践和规范。

## 📚 核心能力概览

### 1. 数据库（JSON数据库）

#### 基本概念
- **数据库 (database)**: 一个云环境对应一个数据库
- **集合 (collection)**: 相当于关系型数据库的表
- **记录 (record/doc)**: 相当于关系型数据库的行
- **字段 (field)**: 相当于关系型数据库的列

#### 关键特性
- 每条记录都是 JSON 格式对象
- 自动包含 `_id` 字段（唯一标识）
- 小程序端创建的记录自动包含 `_openid` 字段（创建者标识）
- 管理端创建的记录不包含 `_openid` 字段

#### 最佳实践

**✅ 推荐做法**:
```javascript
// 1. 使用 DYNAMIC_CURRENT_ENV 自动选择环境
cloud.init({
  env: cloud.DYNAMIC_CURRENT_ENV
})

// 2. 小程序端：严格的权限控制
const db = wx.cloud.database()
db.collection('admission_scores')
  .where({
    province: '北京',
    scoreType: 2
  })
  .get()
  .then(res => {
    console.log('查询成功', res.data)
  })

// 3. 云函数端：管理员权限，可操作所有数据
const db = cloud.database()
const _ = db.command
db.collection('admission_scores')
  .where({
    minScore: _.gte(600)
  })
  .get()
```

**❌ 避免做法**:
```javascript
// 不要硬编码环境ID
cloud.init({
  env: 'lighthouse-volunteer-dev'  // ❌ 不推荐
})

// 不要在小程序端操作敏感数据
// 敏感操作应在云函数中进行
```

### 2. 数据模型（高级工具）

#### 核心特点
- **数据校验**: 自动检查数据正确性
- **关联关系**: 简化数据间的关系管理
- **自动生成代码**: 快速生成 CRUD 操作
- **CMS 管理端**: 提供易用的数据管理界面
- **AI 智能分析**: 利用 AI 分析数据
- **MySQL 支持**: 支持复杂查询操作

#### 使用示例

**初始化数据模型 SDK**:
```javascript
const { init } = require('@cloudbase/wx-cloud-client-sdk')

const client = init(wx.cloud)
const models = client.models
```

**数据校验示例**:
```javascript
// 定义模型：admission_score
// 字段：year (数字), province (字符串), minScore (数字)

try {
  await models.admission_score.create({
    data: {
      year: 2024,
      province: "北京",
      minScore: 685,
      // 如果类型错误会自动校验失败
    },
  });
} catch (error) {
  console.error("数据校验失败：", error);
}
```

**关联查询示例**:
```javascript
// 查询学校及其录取分数
const { data } = await models.college.list({
  select: {
    _id: true,
    collegeName: true,
    province: true,
    // 关联查询录取分数
    admissionScores: {
      _id: true,
      year: true,
      minScore: true,
      minRank: true,
    },
  },
  filter: {
    where: {
      level: {
        $eq: '985'
      }
    },
  },
  getCount: true,
});
```

### 3. 存储

#### 基本操作

**上传文件**:
```javascript
// 小程序端上传
wx.chooseImage({
  success: chooseResult => {
    wx.cloud.uploadFile({
      cloudPath: `user-avatars/${Date.now()}.png`,
      filePath: chooseResult.tempFilePaths[0],
      success: res => {
        console.log('上传成功', res.fileID)
      },
      fail: err => {
        console.error('上传失败', err)
      }
    })
  },
})

// 云函数端上传
const cloud = require('wx-server-sdk')
cloud.init()
const result = await cloud.uploadFile({
  cloudPath: 'admin-data/backup.json',
  fileContent: JSON.stringify(data)
})
```

**下载文件**:
```javascript
// 小程序端下载
wx.cloud.downloadFile({
  fileID: 'cloud://xxx.xxx/xxx.png',
  success: res => {
    console.log('下载成功', res.tempFilePath)
  }
})

// 云函数端下载
const result = await cloud.downloadFile({
  fileID: 'cloud://xxx.xxx/xxx.json'
})
const content = result.fileContent.toString()
```

### 4. 云函数

#### 基本结构

**标准云函数模板**:
```javascript
// index.js
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
  // 获取用户信息（可信的）
  const { OPENID, APPID } = cloud.getWXContext()
  
  // 获取请求参数
  const { a, b } = event
  
  // 业务逻辑
  const sum = a + b
  
  // 返回结果
  return {
    code: 200,
    message: 'success',
    data: {
      OPENID,
      APPID,
      sum
    }
  }
}
```

#### 关键要点

**1. 获取用户信息**:
```javascript
// ✅ 推荐：使用 getWXContext 获取可信的用户信息
const { OPENID, APPID } = cloud.getWXContext()

// ❌ 避免：直接使用 event 中的用户信息（不可信）
const { openid } = event.userInfo  // 可能被伪造
```

**2. 错误处理**:
```javascript
exports.main = async (event, context) => {
  try {
    // 业务逻辑
    return {
      code: 200,
      data: result
    }
  } catch (error) {
    console.error('云函数执行失败:', error)
    return {
      code: 500,
      message: error.message || '服务器内部错误'
    }
  }
}
```

**3. 参数验证**:
```javascript
exports.main = async (event, context) => {
  const { province, score, scoreType } = event
  
  // 参数验证
  if (!province || !score || !scoreType) {
    return {
      code: 400,
      message: '缺少必要参数'
    }
  }
  
  if (score < 0 || score > 750) {
    return {
      code: 400,
      message: '分数范围无效'
    }
  }
  
  // 业务逻辑...
}
```

**4. 请求大小限制**:
- 云函数的请求参数 `data` 有大小限制：**100KB**
- 超过限制需要分片处理或使用存储

### 5. 云调用

#### 使用场景
- 发送模板消息
- 获取小程序码
- 其他小程序开放接口

#### 示例
```javascript
// 在云函数中发送模板消息
const cloud = require('wx-server-sdk')
cloud.init()

exports.main = async (event, context) => {
  const { OPENID } = cloud.getWXContext()
  
  // 使用云调用发送模板消息
  const result = await cloud.openapi.templateMessage.send({
    touser: OPENID,
    templateId: 'xxx',
    page: 'index',
    data: {
      keyword1: { value: '志愿推荐' },
      keyword2: { value: '分析完成' }
    }
  })
  
  return result
}
```

### 6. HTTP API

#### 使用场景
在小程序外访问云开发资源，如：
- Web 端应用
- 管理后台
- 第三方系统集成

#### 认证方式
```javascript
// 使用 access_token 认证
const axios = require('axios')

const response = await axios.post('https://api.weixin.qq.com/cgi-bin/token', {
  grant_type: 'client_credential',
  appid: 'your-appid',
  secret: 'your-secret'
})

const accessToken = response.data.access_token

// 使用 access_token 调用 HTTP API
const dbResult = await axios.post(
  'https://api.weixin.qq.com/tcb/databasequery',
  {
    env: 'your-env-id',
    query: 'db.collection("admission_scores").get()'
  },
  {
    headers: {
      'Authorization': `Bearer ${accessToken}`
    }
  }
)
```

## 🎯 项目最佳实践

### 1. 环境管理

**✅ 推荐**:
```javascript
// 使用 DYNAMIC_CURRENT_ENV 自动选择环境
cloud.init({
  env: cloud.DYNAMIC_CURRENT_ENV
})
```

**环境配置**:
- **开发环境**: lighthouse-volunteer-dev
- **生产环境**: lighthouse-volunteer-prod
- **测试环境**: lighthouse-volunteer-test

### 2. 数据库设计

**集合命名规范**:
- 使用小写字母和下划线
- 集合名要清晰表达用途
- 示例: `admission_scores`, `user_profiles`, `volunteer_suggestions`

**字段命名规范**:
- 使用驼峰命名法
- 保持字段名简洁明了
- 示例: `collegeName`, `minScore`, `dataSource`

**索引设计**:
```javascript
// 为查询频繁的字段创建索引
// 在云控制台或使用代码创建

// 复合索引示例
{
  province: 1,
  scoreType: 1,
  minScore: -1
}

// 文本索引示例
{
  collegeName: "text",
  majorName: "text"
}
```

### 3. 云函数组织

**函数命名规范**:
- 使用动词开头
- 清晰表达函数功能
- 示例: `getAdmissionScores`, `generateVolunteerSuggestion`, `saveUserProfile`

**函数结构**:
```
cloud-functions/
├── generateVolunteerSuggestion/
│   ├── index.js
│   ├── package.json
│   └── config.json
├── getAdmissionScores/
│   ├── index.js
│   └── package.json
└── saveUserProfile/
    ├── index.js
    └── package.json
```

### 4. 错误处理

**统一错误格式**:
```javascript
// 成功响应
{
  code: 200,
  message: 'success',
  data: { ... }
}

// 错误响应
{
  code: 400,  // 400: 参数错误, 401: 未授权, 404: 未找到, 500: 服务器错误
  message: '错误描述',
  error: process.env.NODE_ENV === 'development' ? error.stack : undefined
}
```

**错误日志记录**:
```javascript
// 记录错误到数据库
try {
  // 业务逻辑
} catch (error) {
  console.error('业务错误:', error)
  
  // 记录错误日志
  await db.collection('error_logs').add({
    data: {
      functionName: 'generateVolunteerSuggestion',
      error: error.message,
      stack: error.stack,
      userId: OPENID,
      timestamp: db.serverDate()
    }
  })
  
  throw error
}
```

### 5. 性能优化

**数据库查询优化**:
```javascript
// ✅ 使用索引字段查询
db.collection('admission_scores')
  .where({
    province: '北京',      // 有索引
    scoreType: 2,         // 有索引
    minScore: _.gte(600)  // 有索引
  })
  .get()

// ✅ 使用分页查询
db.collection('admission_scores')
  .where({ ... })
  .skip((page - 1) * pageSize)
  .limit(pageSize)
  .get()

// ❌ 避免全表扫描
db.collection('admission_scores')
  .get()  // 会查询所有数据
```

**云函数优化**:
```javascript
// ✅ 使用 Promise.all 并行处理
const [scores, colleges, majors] = await Promise.all([
  db.collection('admission_scores').get(),
  db.collection('colleges').get(),
  db.collection('majors').get()
])

// ✅ 使用缓存减少重复计算
const cacheKey = `suggestion_${province}_${score}_${scoreType}`
const cached = await getCache(cacheKey)
if (cached) {
  return cached
}
```

### 6. 安全实践

**权限控制**:
```javascript
// 小程序端：只能操作自己的数据
db.collection('user_profiles')
  .where({
    _openid: '{openid}'  // 自动匹配当前用户
  })
  .get()

// 云函数端：管理员权限，可操作所有数据
// 但需要验证用户身份
const { OPENID } = cloud.getWXContext()
if (!OPENID) {
  return { code: 401, message: '未授权' }
}
```

**数据验证**:
```javascript
// 验证用户输入
function validateInput(data) {
  if (!data.province || typeof data.province !== 'string') {
    throw new Error('省份参数无效')
  }
  
  if (!data.score || data.score < 0 || data.score > 750) {
    throw new Error('分数参数无效')
  }
  
  return true
}
```

**敏感数据处理**:
```javascript
// 敏感数据在云函数中处理
// 不要在小程序端直接操作敏感数据
exports.main = async (event, context) => {
  const { OPENID } = cloud.getWXContext()
  
  // 在云函数中处理敏感操作
  const result = await processSensitiveData(OPENID)
  
  return result
}
```

## 📊 监控和调试

### 1. 云控制台监控
- 实时查看云函数调用次数和耗时
- 监控数据库读写统计
- 查看存储使用情况
- 分析错误日志

### 2. 本地调试
```javascript
// 云函数本地调试
// 在微信开发者工具中：
// 1. 右键云函数文件夹
// 2. 选择"本地调试"
// 3. 设置断点和测试数据
```

### 3. 日志记录
```javascript
// 使用 console.log 记录日志
console.log('用户请求:', event)
console.error('错误信息:', error)

// 日志会自动记录到云控制台
```

## 🔗 相关资源

- [微信云开发官方文档](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/basis/getting-started.html)
- [数据库指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/database.html)
- [云函数指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/functions.html)
- [存储指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/storage.html)
- [数据模型指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/datamodel.html)

---

**文档版本**: v1.0
**最后更新**: 2026-01-19
