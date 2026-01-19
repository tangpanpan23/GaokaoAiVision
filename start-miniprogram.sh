#!/bin/bash

# 灯塔志愿小程序启动脚本

echo "🚀 灯塔志愿微信小程序"
echo "========================"
echo ""

# 检查微信开发者工具是否安装
check_wechat_devtools() {
    echo "📋 检查微信开发者工具..."
    if command -v "微信开发者工具" &> /dev/null; then
        echo "✅ 微信开发者工具已安装"
        return 0
    else
        echo "⚠️  未检测到微信开发者工具"
        echo "请从 https://developers.weixin.qq.com/miniprogram/dev/devtools/download.html 下载安装"
        return 1
    fi
}

# 检查项目文件完整性
check_project_files() {
    echo "📋 检查项目文件..."

    local required_files=(
        "miniprogram/app.js"
        "miniprogram/app.json"
        "miniprogram/app.wxss"
        "miniprogram/pages/home/index.js"
        "miniprogram/pages/home/index.wxml"
        "miniprogram/pages/home/index.wxss"
        "miniprogram/pages/volunteer-suggestion/index.js"
        "miniprogram/pages/volunteer-suggestion/index.wxml"
        "miniprogram/pages/volunteer-suggestion/index.wxss"
        "miniprogram/utils/api.js"
        "miniprogram/utils/util.js"
    )

    local missing_files=()

    for file in "${required_files[@]}"; do
        if [[ ! -f "$file" ]]; then
            missing_files+=("$file")
        fi
    done

    if [[ ${#missing_files[@]} -eq 0 ]]; then
        echo "✅ 项目文件完整"
        return 0
    else
        echo "❌ 缺少以下文件:"
        for file in "${missing_files[@]}"; do
            echo "   - $file"
        done
        return 1
    fi
}

# 显示项目信息
show_project_info() {
    echo ""
    echo "📱 项目信息"
    echo "-----------"
    echo "项目名称: 灯塔志愿"
    echo "项目类型: 微信小程序"
    echo "项目路径: $(pwd)/miniprogram"
    echo "开发模式: 模拟数据模式"
    echo ""

    echo "🎯 核心功能"
    echo "-----------"
    echo "✅ 志愿智能推荐 (AI分析)"
    echo "✅ 分数查询 (历年数据)"
    echo "✅ 学长分享 (真实经历)"
    echo "✅ 职业测评 (性格分析)"
    echo "✅ 个人中心 (档案管理)"
    echo ""

    echo "🔧 技术栈"
    echo "---------"
    echo "前端框架: 原生微信小程序"
    echo "数据交互: RESTful API (模拟)"
    echo "UI设计: Material Design"
    echo "状态管理: 页面级管理"
    echo ""
}

# 显示使用说明
show_usage_guide() {
    echo "📚 使用指南"
    echo "-----------"
    echo ""
    echo "1. 打开微信开发者工具"
    echo "2. 点击 '导入项目'"
    echo "3. 项目路径选择: $(pwd)/miniprogram"
    echo "4. AppID: 使用测试号或你的小程序AppID"
    echo "5. 点击 '确定' 导入项目"
    echo "6. 点击 '编译' 运行小程序"
    echo ""
    echo "🎮 操作说明"
    echo "-----------"
    echo "1. 在首页点击各项功能进入对应页面"
    echo "2. 志愿推荐页: 填写信息，AI生成推荐方案"
    echo "3. 所有功能都支持离线演示（模拟数据）"
    echo "4. 支持微信授权登录（模拟）"
    echo ""
    echo "🔍 调试技巧"
    echo "-----------"
    echo "1. 控制台: 查看日志和错误信息"
    echo "2. Network: 监控API请求（模拟）"
    echo "3. Storage: 检查本地存储数据"
    echo "4. WXML: 实时查看页面结构"
    echo ""
}

# 主函数
main() {
    echo "启动检查中..."
    echo ""

    # 检查项目文件
    if ! check_project_files; then
        echo ""
        echo "❌ 项目文件不完整，请检查项目结构"
        exit 1
    fi

    # 显示项目信息
    show_project_info

    # 显示使用指南
    show_usage_guide

    echo "🎉 项目检查完成！"
    echo ""
    echo "现在你可以按照上述步骤打开微信开发者工具并导入项目了。"
    echo ""
    echo "如有问题，请查看 MINIPROGRAM_README.md 文件获取详细文档。"
    echo ""
    echo "祝你开发顺利！🚀"
}

# 运行主函数
main "$@"
