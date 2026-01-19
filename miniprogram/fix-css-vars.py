#!/usr/bin/env python3
"""
批量替换WXSS文件中的CSS变量为实际颜色值
"""
import os
import re
import glob

# CSS变量映射表
CSS_VARS = {
    '--primary-color': '#007bff',
    '--primary-dark': '#0056b3',
    '--success-color': '#28a745',
    '--warning-color': '#ffc107',
    '--danger-color': '#dc3545',
    '--info-color': '#17a2b8',
    '--text-primary': '#333333',
    '--text-secondary': '#666666',
    '--text-muted': '#999999',
    '--text-white': '#ffffff',
    '--bg-white': '#ffffff',
    '--bg-light': '#f8f9fa',
    '--bg-gray': '#e9ecef',
    '--border-color': '#e9ecef',
    '--shadow-color': 'rgba(0, 0, 0, 0.06)',
}

def replace_css_vars(content):
    """替换CSS变量为实际值"""
    for var_name, var_value in CSS_VARS.items():
        pattern = f'var\\({re.escape(var_name)}\\)'
        content = re.sub(pattern, var_value, content)
    return content

def process_file(filepath):
    """处理单个文件"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        new_content = replace_css_vars(content)
        
        if content != new_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(new_content)
            print(f'✅ 已处理: {filepath}')
            return True
        else:
            print(f'⏭️  跳过: {filepath} (无CSS变量)')
            return False
    except Exception as e:
        print(f'❌ 错误: {filepath} - {e}')
        return False

def main():
    """主函数"""
    print('🚀 开始批量替换CSS变量...\n')
    
    # 查找所有WXSS文件
    wxss_files = glob.glob('**/*.wxss', recursive=True)
    
    if not wxss_files:
        print('❌ 未找到WXSS文件')
        return
    
    print(f'📋 找到 {len(wxss_files)} 个WXSS文件\n')
    
    processed = 0
    for filepath in wxss_files:
        if process_file(filepath):
            processed += 1
    
    print(f'\n✅ 处理完成！共处理 {processed} 个文件')

if __name__ == '__main__':
    main()
