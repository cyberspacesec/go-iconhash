import { createGlobalStyle } from 'styled-components';

const GlobalStyles = createGlobalStyle`
  :root {
    --primary-color: #4299e1;
    --primary-dark: #3182ce;
    --secondary-color: #2d3748;
    --text-color: #1a202c;
    --text-light: #718096;
    --bg-color: #ffffff;
    --bg-light: #f7fafc;
    --bg-dark: #1a202c;
    --border-color: #e2e8f0;
    --code-bg: #f8f9fa;
    --code-color: #2d3748;
    --error-color: #e53e3e;
    --success-color: #38a169;
    --warning-color: #f6ad55;
    --border-radius: 6px;
    --shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    --transition: all 0.3s ease;
  }

  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  html {
    scroll-behavior: smooth;
  }

  body {
    font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    color: var(--text-color);
    background-color: var(--bg-color);
    line-height: 1.5;
    overflow-x: hidden;
  }

  pre, code {
    font-family: 'Fira Code', monospace;
  }

  a {
    color: var(--primary-color);
    text-decoration: none;
    transition: var(--transition);
  }

  a:hover {
    color: var(--primary-dark);
  }

  img {
    max-width: 100%;
    height: auto;
  }

  button {
    cursor: pointer;
    font-family: inherit;
  }

  ul {
    list-style: none;
  }

  .container {
    width: 100%;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 1.5rem;
  }

  section {
    padding: 5rem 0;
  }

  @media (max-width: 768px) {
    section {
      padding: 3rem 0;
    }
  }
`;

export default GlobalStyles; 