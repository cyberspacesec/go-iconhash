/**
 * 自动部署脚本 - 将构建好的网站部署到GitHub Pages
 * 
 * 这个脚本会执行以下操作：
 * 1. 确保dist目录存在（假设已通过npm run build构建）
 * 2. 创建或清空项目根目录下的docs目录
 * 3. 复制dist内容到docs目录
 * 4. 创建.nojekyll文件防止GitHub用Jekyll处理
 * 5. 添加到git并提交
 * 6. 推送到GitHub
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

// 配置
const DIST_DIR = path.resolve(__dirname, '../dist');
const DOCS_DIR = path.resolve(__dirname, '../../docs');

// 彩色输出
const colors = {
  reset: '\x1b[0m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  blue: '\x1b[34m',
  red: '\x1b[31m'
};

// 辅助函数 - 日志
function log(message, color = colors.blue) {
  console.log(`${color}${message}${colors.reset}`);
}

// 辅助函数 - 执行命令
function executeCommand(command, cwd = process.cwd()) {
  try {
    log(`执行命令: ${command}`, colors.yellow);
    execSync(command, { stdio: 'inherit', cwd });
    return true;
  } catch (error) {
    log(`命令执行失败: ${error}`, colors.red);
    return false;
  }
}

// 主函数
function deployToGitHubPages() {
  log('开始部署IconHash网站到GitHub Pages...', colors.green);

  // 确认dist目录是否存在
  if (!fs.existsSync(DIST_DIR)) {
    log('错误: dist目录不存在，请先运行 npm run build', colors.red);
    process.exit(1);
  }

  // 创建或清空docs目录
  if (!fs.existsSync(DOCS_DIR)) {
    log('创建docs目录...');
    fs.mkdirSync(DOCS_DIR, { recursive: true });
  } else {
    log('清空docs目录...');
    const files = fs.readdirSync(DOCS_DIR);
    for (const file of files) {
      if (file !== '.git') { // 保留.git目录
        const filePath = path.join(DOCS_DIR, file);
        if (fs.lstatSync(filePath).isDirectory()) {
          fs.rmSync(filePath, { recursive: true, force: true });
        } else {
          fs.unlinkSync(filePath);
        }
      }
    }
  }

  // 复制构建文件到docs目录
  log('复制构建文件到docs目录...');
  const copyFiles = (src, dest) => {
    const entries = fs.readdirSync(src, { withFileTypes: true });
    
    for (const entry of entries) {
      const srcPath = path.join(src, entry.name);
      const destPath = path.join(dest, entry.name);
      
      if (entry.isDirectory()) {
        fs.mkdirSync(destPath, { recursive: true });
        copyFiles(srcPath, destPath);
      } else {
        fs.copyFileSync(srcPath, destPath);
      }
    }
  };
  
  copyFiles(DIST_DIR, DOCS_DIR);

  // 创建.nojekyll文件
  log('创建.nojekyll文件...');
  fs.writeFileSync(path.join(DOCS_DIR, '.nojekyll'), '');

  // Git操作
  log('将更改提交到Git...');
  const rootDir = path.resolve(__dirname, '../..');
  
  // 检查Git状态
  if (executeCommand('git status', rootDir)) {
    // 添加文件到Git
    if (executeCommand('git add docs/', rootDir)) {
      // 提交更改
      if (executeCommand('git commit -m "更新GitHub Pages网站内容"', rootDir)) {
        // 推送到GitHub
        if (executeCommand('git push origin main', rootDir)) {
          log('🎉 部署成功完成！', colors.green);
          log('网站将在几分钟内可访问: https://cyberspacesec.github.io/go-iconhash/', colors.green);
        } else {
          log('推送到GitHub失败', colors.red);
        }
      } else {
        log('提交更改失败', colors.red);
      }
    } else {
      log('添加文件到Git失败', colors.red);
    }
  } else {
    log('检查Git状态失败', colors.red);
  }
}

// 执行部署
deployToGitHubPages(); 