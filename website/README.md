# IconHash Website (React Version)

This is the React-based website for the IconHash tool, a powerful favicon hash calculator for cybersecurity reconnaissance.

## Technologies Used

- React 18
- TypeScript
- Ant Design UI framework
- styled-components for styling
- i18next for internationalization support

## Features

- Responsive design that works on desktop, tablet, and mobile devices
- Multilingual support (English and Chinese)
- Modern and clean user interface with Ant Design components
- Interactive examples and documentation

## Development

### Prerequisites

- Node.js (v14 or newer)
- npm or yarn

### Local Development

To run the website locally:

1. Clone the repository:
   ```
   git clone https://github.com/cyberspacesec/go-iconhash.git
   cd go-iconhash/website
   ```

2. Install dependencies:
   ```
   npm install
   # or
   yarn
   ```

3. Start the development server:
   ```
   npm start
   # or
   yarn start
   ```

4. Visit `http://localhost:3001` in your browser

### Building for Production

To create a production build:

```
npm run build
# or
yarn build
```

The built files will be located in the `dist` directory.

### GitHub Pages Deployment

网站配置为部署到GitHub Pages专用的gh-pages分支。要部署网站：

```
# 使用Node.js脚本部署（推荐）
npm run deploy-gh-pages

# 或使用Bash脚本部署（旧方式，部署到docs目录）
npm run deploy
# 或直接运行
./deploy.sh
```

新的部署过程会执行以下操作:
1. 构建网站
2. 创建临时目录
3. 克隆或创建gh-pages分支
4. 将构建文件复制到临时目录
5. 提交并推送到gh-pages分支
6. 清理临时目录

> **重要**: 在GitHub仓库设置中，确保将GitHub Pages的发布源设置为`gh-pages`分支，而不是`main`分支的`/docs`文件夹。

部署完成后，网站将在以下地址可访问:
https://cyberspacesec.github.io/go-iconhash/

注意: GitHub可能需要几分钟时间来处理更改。

### Local Development Server

If you want to test the website locally without deploying to GitHub Pages:

```
npm run serve-local
# or
./serve-gh-pages.sh
```

This will build the project, copy the files to the `../docs` directory, and start a local server to preview the website. Visit `http://localhost:8000` in your browser to see it.

## Language Support

The website supports multiple languages through i18next:

- English: `src/locales/en/translation.json`
- Chinese: `src/locales/zh/translation.json`

To add support for a new language:

1. Create a new JSON file in the `src/locales/{lang}` directory (e.g., `src/locales/fr/translation.json`)
2. Copy the structure from an existing language file
3. Translate all text values
4. Add the language import and configuration in `src/i18n.ts`

## Project Structure

- `src/components/` - React components
  - `ApiReference.tsx` - API文档组件
  - `Examples.tsx` - 使用示例组件
  - `Features.tsx` - 功能特性组件
  - `Footer.tsx` - 页脚组件
  - `Header.tsx` - 页眉组件
  - `Hero.tsx` - 主页英雄区域组件
  - `Installation.tsx` - 安装指南组件
  - `McpSection.tsx` - Model Context Protocol部分组件
- `src/locales/` - Translation files
  - `en/translation.json` - 英文翻译文件
  - `zh/translation.json` - 中文翻译文件
- `src/styles/` - Global styles
- `src/assets/` - Static assets like images
- `src/i18n.ts` - i18next configuration
- `src/App.tsx` - Main application component
- `src/index.tsx` - Application entry point

## API Endpoints

The website documentation covers the following API endpoints:

- `/health` - Health check endpoint (no authentication required)
- `/hash/url` - Calculate hash from a URL
- `/hash/file` - Calculate hash from an uploaded file
- `/hash/base64` - Calculate hash from base64 encoded data
- `/mcp` - Model Context Protocol endpoint

## Development Notes

After making changes to the IconHash tool's structure or functionality, please ensure that the website documentation is updated accordingly. The website serves as the primary documentation source for users, so it's important to keep it in sync with the actual code behavior.

## License

This website is released under the [MIT License](../LICENSE). 