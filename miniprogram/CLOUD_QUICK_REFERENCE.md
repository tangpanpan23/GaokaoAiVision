# 微信云开发快速参考指南

基于官方文档的核心概念和API速查表。

## 🗄️ 数据库操作

### 小程序端

```javascript
// 初始化
const db = wx.cloud.database()

// 查询
db.collection('admission_scores')
  .where({
    province: '北京',
    scoreType: 2
  })
  .get()
  .then(res => console.log(res.data))

// 添加
db.collection('user_profiles').add({
  data: {
    province: '北京',
    score: 650
  }
})

// 更新
db.collection('user_profiles')
  .doc('doc-id')
  .update({
    data: {
      score: 660
    }
  })

// 删除
db.collection('user_profiles')
  .doc('doc-id')
  .remove()
```

### 云函数端

```javascript
const cloud = require('wx-server-sdk')
cloud.init({ env: cloud.DYNAMIC_CURRENT_ENV })
const db = cloud.database()
const _ = db.command

// 查询（管理员权限）
db.collection('admission_scores')
  .where({
    minScore: _.gte(600)
  })
  .get()

// 批量操作
const _ = db.command
db.collection('admission_scores')
  .where({
    year: _.in([2022, 2023, 2024])
  })
  .update({
    data: {
      dataQuality: 5
    }
  })
```

## 📦 存储操作

### 上传文件

```javascript
// 小程序端
wx.chooseImage({
  success: res => {
    wx.cloud.uploadFile({
      cloudPath: `avatars/${Date.now()}.png`,
      filePath: res.tempFilePaths[0],
      success: res => console.log(res.fileID)
    })
  }
})

// 云函数端
const cloud = require('wx-server-sdk')
const result = await cloud.uploadFile({
  cloudPath: 'data/backup.json',
  fileContent: Buffer.from(JSON.stringify(data))
})
```

### 下载文件

```javascript
// 小程序端
wx.cloud.downloadFile({
  fileID: 'cloud://xxx.xxx/xxx.png',
  success: res => console.log(res.tempFilePath)
})

// 云函数端
const result = await cloud.downloadFile({
  fileID: 'cloud://xxx.xxx/xxx.json'
})
```

## ⚡ 云函数

### 基本结构

```javascript
const cloud = require('wx-server-sdk')
cloud.init({ env: cloud.DYNAMIC_CURRENT_ENV })

exports.main = async (event, context) => {
  // 获取用户信息（可信的）
  const { OPENID, APPID } = cloud.getWXContext()
  
  // 获取参数
  const { a, b } = event
  
  // 业务逻辑
  const result = a + b
  
  // 返回结果
  return {
    code: 200,
    data: { result }
  }
}
```

### 调用云函数

```javascript
// 小程序端
wx.cloud.callFunction({
  name: 'add',
  data: { a: 12, b: 19 }
}).then(res => console.log(res.result))

// Promise 方式
const result = await wx.cloud.callFunction({
  name: 'add',
  data: { a: 12, b: 19 }
})
```

## 🔍 查询操作符

```javascript
const _ = db.command

// 比较操作符
_.eq(value)      // 等于
_.neq(value)     // 不等于
_.gt(value)      // 大于
_.gte(value)     // 大于等于
_.lt(value)      // 小于
_.lte(value)     // 小于等于
_.in(array)      // 在数组中
_.nin(array)     // 不在数组中

// 逻辑操作符
_.and([...])     // 且
_.or([...])      // 或
_.not(...)       // 非

// 示例
db.collection('admission_scores')
  .where({
    minScore: _.gte(600).and(_.lte(700)),
    province: _.in(['北京', '上海', '广东'])
  })
  .get()
```

## 📊 数据模型（高级）

```javascript
const { init } = require('@cloudbase/wx-cloud-client-sdk')

const client = init(wx.cloud)
const models = client.models

// 创建（带数据校验）
await models.admission_score.create({
  data: {
    year: 2024,
    province: "北京",
    minScore: 685
  }
})

// 查询（关联查询）
const { data } = await models.college.list({
  select: {
    _id: true,
    collegeName: true,
    admissionScores: {
      _id: true,
      year: true,
      minScore: true
    }
  },
  filter: {
    where: {
      level: { $eq: '985' }
    }
  },
  getCount: true
})
```

## 🔐 权限和安全

### 小程序端权限
- 只能操作自己的数据（通过 `_openid` 自动匹配）
- 受数据库权限规则限制

### 云函数端权限
- 管理员权限，可操作所有数据
- 使用 `getWXContext()` 获取可信的用户信息

```javascript
// ✅ 推荐：使用 getWXContext
const { OPENID } = cloud.getWXContext()

// ❌ 避免：直接使用 event 中的用户信息
const { openid } = event.userInfo  // 可能被伪造
```

## 📏 限制说明

| 项目 | 限制 |
|------|------|
| 云函数请求参数 | 100KB |
| 数据库单次查询 | 20条（默认） |
| 数据库单次更新 | 500条 |
| 文件上传大小 | 10MB（小程序端） |
| 云函数执行时间 | 60秒（默认） |

## 🎯 最佳实践

### 1. 环境管理
```javascript
// ✅ 使用 DYNAMIC_CURRENT_ENV
cloud.init({ env: cloud.DYNAMIC_CURRENT_ENV })

// ❌ 不要硬编码环境ID
cloud.init({ env: 'lighthouse-volunteer-dev' })
```

### 2. 错误处理
```javascript
try {
  const result = await db.collection('xxx').get()
  return { code: 200, data: result }
} catch (error) {
  console.error('错误:', error)
  return { code: 500, message: error.message }
}
```

### 3. 性能优化
```javascript
// ✅ 使用索引字段查询
db.collection('admission_scores')
  .where({ province: '北京' })  // 有索引
  .get()

// ✅ 使用分页
.skip((page - 1) * pageSize)
.limit(pageSize)

// ✅ 并行查询
const [scores, colleges] = await Promise.all([
  db.collection('admission_scores').get(),
  db.collection('colleges').get()
])
```

### 4. 数据验证
```javascript
// 参数验证
if (!province || !score) {
  return { code: 400, message: '缺少必要参数' }
}

if (score < 0 || score > 750) {
  return { code: 400, message: '分数范围无效' }
}
```

## 🔗 官方文档链接

- [数据库指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/database.html)
- [云函数指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/functions.html)
- [存储指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/storage.html)
- [数据模型指引](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/guide/datamodel.html)
- [HTTP API文档](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference/http-api/)

---

**快速参考版本**: v1.0
**最后更新**: 2026-01-19
