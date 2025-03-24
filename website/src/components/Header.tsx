import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Menu, Dropdown, Button } from 'antd';
import { GithubOutlined, GlobalOutlined, MenuOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const HeaderContainer = styled.header`
  position: sticky;
  top: 0;
  background-color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  padding: 0.8rem 0;
`;

const NavContainer = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const Logo = styled.a`
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-color);
  span {
    color: var(--primary-color);
  }
`;

const NavLinks = styled.ul`
  display: flex;
  gap: 1.5rem;
  margin: 0;

  @media (max-width: 768px) {
    display: none;
  }
`;

const NavLink = styled.a`
  color: var(--text-color);
  font-weight: 500;
  
  &:hover {
    color: var(--primary-color);
  }
`;

const MobileMenuButton = styled(Button)`
  display: none;
  
  @media (max-width: 768px) {
    display: flex;
    align-items: center;
    justify-content: center;
  }
`;

const RightSection = styled.div`
  display: flex;
  align-items: center;
  gap: 1rem;
`;

const Header: React.FC = () => {
  const { t, i18n } = useTranslation();
  const [menuVisible, setMenuVisible] = useState(false);
  
  const changeLanguage = (lang: string) => {
    i18n.changeLanguage(lang);
  };
  
  const languageMenu = (
    <Menu>
      <Menu.Item key="en" onClick={() => changeLanguage('en')}>
        English
      </Menu.Item>
      <Menu.Item key="zh" onClick={() => changeLanguage('zh')}>
        中文
      </Menu.Item>
    </Menu>
  );
  
  const mobileMenu = (
    <Menu>
      <Menu.Item key="home">
        <a href="#">{t('header.home')}</a>
      </Menu.Item>
      <Menu.Item key="features">
        <a href="#features">{t('header.features')}</a>
      </Menu.Item>
      <Menu.Item key="documentation">
        <a href="#installation">{t('header.documentation')}</a>
      </Menu.Item>
      <Menu.Item key="examples">
        <a href="#examples">{t('header.examples')}</a>
      </Menu.Item>
      <Menu.Item key="api">
        <a href="#api">{t('header.api')}</a>
      </Menu.Item>
      <Menu.Item key="mcp">
        <a href="#mcp">{t('header.mcp')}</a>
      </Menu.Item>
      <Menu.Item key="github">
        <a href="https://github.com/cyberspacesec/go-iconhash" target="_blank" rel="noopener noreferrer">
          {t('header.github')}
        </a>
      </Menu.Item>
    </Menu>
  );

  return (
    <HeaderContainer>
      <div className="container">
        <NavContainer>
          <Logo href="#">
            Icon<span>Hash</span>
          </Logo>
          
          <NavLinks>
            <li><NavLink href="#">{t('header.home')}</NavLink></li>
            <li><NavLink href="#features">{t('header.features')}</NavLink></li>
            <li><NavLink href="#installation">{t('header.documentation')}</NavLink></li>
            <li><NavLink href="#examples">{t('header.examples')}</NavLink></li>
            <li><NavLink href="#api">{t('header.api')}</NavLink></li>
            <li><NavLink href="#mcp">{t('header.mcp')}</NavLink></li>
            <li>
              <NavLink 
                href="https://github.com/cyberspacesec/go-iconhash" 
                target="_blank" 
                rel="noopener noreferrer"
              >
                {t('header.github')}
              </NavLink>
            </li>
          </NavLinks>
          
          <RightSection>
            <Dropdown overlay={languageMenu} trigger={['click']}>
              <Button icon={<GlobalOutlined />} type="text">
                {i18n.language === 'zh' ? '中文' : 'EN'}
              </Button>
            </Dropdown>
            
            <Dropdown 
              overlay={mobileMenu} 
              trigger={['click']} 
              visible={menuVisible}
              onVisibleChange={setMenuVisible}
            >
              <MobileMenuButton icon={<MenuOutlined />} type="text" />
            </Dropdown>
          </RightSection>
        </NavContainer>
      </div>
    </HeaderContainer>
  );
};

export default Header; 