/**
 * 自动部署脚本 - 将构建好的网站部署到GitHub Pages
 * 
 * 这个脚本会执行以下操作：
 * 1. 确保dist目录存在（假设已通过npm run build构建）
 * 2. 创建临时目录用于gh-pages分支
 * 3. 克隆仓库的gh-pages分支（如果存在）或创建一个新的孤立分支
 * 4. 复制dist内容到临时目录
 * 5. 提交并推送到gh-pages分支
 */

const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');
const os = require('os');

// 配置
const DIST_DIR = path.resolve(__dirname, '../dist');
const REPO_URL = 'https://github.com/cyberspacesec/go-iconhash.git';
const BRANCH = 'gh-pages';
const TEMP_DIR = path.join(os.tmpdir(), `gh-pages-${Date.now()}`);

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

// 清理函数
function cleanup() {
  if (fs.existsSync(TEMP_DIR)) {
    log(`清理临时目录: ${TEMP_DIR}`);
    fs.rmSync(TEMP_DIR, { recursive: true, force: true });
  }
}

// 主函数
function deployToGitHubPages() {
  log('开始部署IconHash网站到GitHub Pages分支...', colors.green);

  // 设置错误处理和清理
  process.on('SIGINT', () => {
    log('\n收到中断信号，正在清理...', colors.yellow);
    cleanup();
    process.exit(1);
  });

  // 确认dist目录是否存在
  if (!fs.existsSync(DIST_DIR)) {
    log('错误: dist目录不存在，请先运行 npm run build', colors.red);
    process.exit(1);
  }

  // 创建临时目录
  log(`创建临时目录: ${TEMP_DIR}`);
  fs.mkdirSync(TEMP_DIR, { recursive: true });

  try {
    // 获取仓库的远程URL
    const rootDir = path.resolve(__dirname, '../..');
    let repoUrl;
    
    try {
      // 尝试从当前仓库获取URL
      repoUrl = execSync('git config --get remote.origin.url', { cwd: rootDir }).toString().trim();
      log(`使用当前仓库URL: ${repoUrl}`);
    } catch (error) {
      // 如果失败，使用默认URL
      repoUrl = REPO_URL;
      log(`无法获取仓库URL，使用默认值: ${repoUrl}`);
    }

    // 检查gh-pages分支是否存在
    let branchExists = false;
    try {
      execSync(`git ls-remote --heads ${repoUrl} ${BRANCH}`, { stdio: 'pipe' });
      branchExists = true;
      log(`${BRANCH} 分支已存在`);
    } catch (error) {
      log(`${BRANCH} 分支不存在，将创建新分支`);
    }

    // 克隆仓库或创建新的gh-pages分支
    if (branchExists) {
      // 克隆仅gh-pages分支，深度为1（只获取最近一次提交）
      log(`克隆 ${BRANCH} 分支到临时目录...`);
      if (!executeCommand(`git clone --branch ${BRANCH} --single-branch --depth 1 ${repoUrl} ${TEMP_DIR}`)) {
        throw new Error(`克隆 ${BRANCH} 分支失败`);
      }
    } else {
      // 初始化新仓库
      log('初始化新的Git仓库...');
      if (!executeCommand(`git init ${TEMP_DIR}`)) {
        throw new Error('初始化Git仓库失败');
      }

      // 配置远程仓库
      if (!executeCommand(`git remote add origin ${repoUrl}`, TEMP_DIR)) {
        throw new Error('配置远程仓库失败');
      }

      // 创建一个孤立的分支（没有历史记录）
      if (!executeCommand(`git checkout --orphan ${BRANCH}`, TEMP_DIR)) {
        throw new Error(`创建 ${BRANCH} 分支失败`);
      }
    }

    // 清空临时目录中除了.git之外的所有内容
    log('清空临时目录中的旧内容...');
    const tempFiles = fs.readdirSync(TEMP_DIR);
    for (const file of tempFiles) {
      if (file !== '.git') {
        const filePath = path.join(TEMP_DIR, file);
        if (fs.lstatSync(filePath).isDirectory()) {
          fs.rmSync(filePath, { recursive: true, force: true });
        } else {
          fs.unlinkSync(filePath);
        }
      }
    }

    // 复制dist目录内容到临时目录
    log('复制构建文件到临时目录...');
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
    
    copyFiles(DIST_DIR, TEMP_DIR);

    // 创建必要的GitHub Pages文件
    log('创建.nojekyll文件...');
    fs.writeFileSync(path.join(TEMP_DIR, '.nojekyll'), '');

    // 添加CNAME文件（如果需要自定义域名）
    // fs.writeFileSync(path.join(TEMP_DIR, 'CNAME'), 'your-domain.com');

    // Git操作
    log('配置Git用户信息...');
    // 使用环境变量或默认值配置Git用户信息
    executeCommand('git config user.name "GitHub Pages Bot"', TEMP_DIR);
    executeCommand('git config user.email "bot@example.com"', TEMP_DIR);

    // 添加所有文件
    log('添加文件到Git...');
    if (!executeCommand('git add -A', TEMP_DIR)) {
      throw new Error('添加文件失败');
    }

    // 提交更改
    log('提交更改...');
    if (!executeCommand('git commit -m "更新GitHub Pages网站内容"', TEMP_DIR)) {
      throw new Error('提交更改失败');
    }

    // 强制推送到gh-pages分支
    log(`推送到 ${BRANCH} 分支...`);
    if (!executeCommand(`git push -f origin ${BRANCH}`, TEMP_DIR)) {
      throw new Error(`推送到 ${BRANCH} 分支失败`);
    }

    log('🎉 部署成功完成！', colors.green);
    log(`网站将在几分钟内可访问: https://cyberspacesec.github.io/go-iconhash/`, colors.green);
    log('提示: 请确保在GitHub仓库设置中将GitHub Pages的发布源设置为gh-pages分支', colors.yellow);
  } catch (error) {
    log(`部署失败: ${error.message}`, colors.red);
    process.exit(1);
  } finally {
    // 清理临时目录
    cleanup();
  }
}

// 执行部署
deployToGitHubPages(); 