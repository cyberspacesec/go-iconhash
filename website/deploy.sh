#!/bin/bash

# 确保任何错误发生时脚本会退出
set -e

echo "开始部署IconHash网站到GitHub Pages..."

# 构建项目
echo "1. 构建React网站..."
npm run build

# 创建或清空docs目录
DOCS_DIR="../docs"

if [ ! -d "$DOCS_DIR" ]; then
  echo "2. 创建docs目录..."
  mkdir -p "$DOCS_DIR"
else
  echo "2. 清空docs目录..."
  rm -rf "$DOCS_DIR"/*
fi

# 复制构建文件到docs目录
echo "3. 复制构建文件到docs目录..."
cp -r dist/* "$DOCS_DIR"

# 创建.nojekyll文件确保GitHub Pages不用Jekyll处理
touch "$DOCS_DIR/.nojekyll"

# 提交更改
echo "4. 提交更改到Git仓库..."
cd ..
git add docs/
git commit -m "更新GitHub Pages网站"

# 推送到GitHub
echo "5. 推送更改到GitHub..."
git push origin main

echo "部署完成！网站应该很快就能在 https://cyberspacesec.github.io/go-iconhash/ 访问到"
echo "注意：GitHub可能需要几分钟时间来处理更改" 