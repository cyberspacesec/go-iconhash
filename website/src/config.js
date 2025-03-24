/**
 * 全局配置
 */

// 在生产环境中使用GitHub Pages路径，在开发环境中使用根路径
const BASE_URL = process.env.NODE_ENV === 'production' ? '/go-iconhash' : '';

export default {
  BASE_URL
}; 