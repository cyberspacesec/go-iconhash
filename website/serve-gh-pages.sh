#!/bin/bash

# 首先构建并部署到docs目录
echo "Building and deploying to docs/ directory..."
./deploy.sh

if [ $? -ne 0 ]; then
  echo "Deployment failed. Exiting."
  exit 1
fi

# 检测docs目录是否存在
DOCS_DIR="../docs"
if [ ! -d "$DOCS_DIR" ]; then
  echo "Error: docs/ directory not found at $DOCS_DIR"
  exit 1
fi

echo "Starting local server for GitHub Pages preview..."
echo "Open your browser and navigate to http://localhost:8000"

# 切换到docs目录
cd "$DOCS_DIR"

# 检查可用的HTTP服务器并启动本地服务
if command -v python3 &>/dev/null; then
  echo "Using Python 3 HTTP server"
  python3 -m http.server 8000
elif command -v python &>/dev/null; then
  echo "Using Python 2 HTTP server"
  python -m SimpleHTTPServer 8000
elif command -v npx &>/dev/null; then
  echo "Using npx http-server"
  npx http-server -p 8000 --cors
elif command -v php &>/dev/null; then
  echo "Using PHP server"
  php -S localhost:8000
else
  echo "Error: No suitable HTTP server found."
  echo "Please install Python, Node.js with npx, or PHP to serve the website locally."
  exit 1
fi 