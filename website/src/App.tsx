import React from 'react';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/lib/locale/zh_CN';
import enUS from 'antd/lib/locale/en_US';
import { useTranslation } from 'react-i18next';

import GlobalStyles from './styles/GlobalStyles';
import Header from './components/Header';
import Hero from './components/Hero';
import Features from './components/Features';
import Installation from './components/Installation';
import Examples from './components/Examples';
import ApiReference from './components/ApiReference';
import McpSection from './components/McpSection';
import Footer from './components/Footer';

const App: React.FC = () => {
  const { i18n } = useTranslation();
  
  // Set Ant Design's locale based on the current i18n language
  const getAntLocale = () => {
    return i18n.language === 'zh' ? zhCN : enUS;
  };

  return (
    <ConfigProvider locale={getAntLocale()}>
      <GlobalStyles />
      <Header />
      <main>
        <Hero />
        <Features />
        <Installation />
        <Examples />
        <ApiReference />
        <McpSection />
      </main>
      <Footer />
    </ConfigProvider>
  );
};

export default App; 