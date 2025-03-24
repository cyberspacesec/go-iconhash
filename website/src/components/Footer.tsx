import React from 'react';
import { useTranslation } from 'react-i18next';
import { Radio, Space } from 'antd';
import { GithubOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const FooterContainer = styled.footer`
  background-color: var(--bg-dark);
  color: white;
  padding: 3rem 0;
`;

const FooterContent = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const FooterLogo = styled.div`
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
  
  span {
    color: var(--primary-color);
  }
`;

const FooterLinks = styled.div`
  display: flex;
  gap: 1.5rem;
  margin-bottom: 1.5rem;
  
  @media (max-width: 576px) {
    flex-direction: column;
    align-items: center;
    gap: 0.75rem;
  }
`;

const FooterLink = styled.a`
  color: white;
  font-weight: 500;
  opacity: 0.8;
  transition: var(--transition);
  
  &:hover {
    opacity: 1;
    color: var(--primary-color);
  }
`;

const SocialLinks = styled.div`
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
`;

const SocialLink = styled.a`
  color: white;
  font-size: 1.5rem;
  opacity: 0.8;
  transition: var(--transition);
  
  &:hover {
    opacity: 1;
    color: var(--primary-color);
  }
`;

const Copyright = styled.div`
  opacity: 0.6;
  text-align: center;
  margin-bottom: 1rem;
`;

const License = styled.div`
  opacity: 0.6;
  text-align: center;
`;

const LanguageSelector = styled.div`
  margin-top: 1.5rem;
`;

const Footer: React.FC = () => {
  const { t, i18n } = useTranslation();
  
  const handleLanguageChange = (e: any) => {
    i18n.changeLanguage(e.target.value);
  };

  return (
    <FooterContainer>
      <div className="container">
        <FooterContent>
          <FooterLogo>
            Icon<span>Hash</span>
          </FooterLogo>
          
          <FooterLinks>
            <FooterLink href="#">{t('header.home')}</FooterLink>
            <FooterLink href="#features">{t('header.features')}</FooterLink>
            <FooterLink href="#installation">{t('header.documentation')}</FooterLink>
            <FooterLink href="#examples">{t('header.examples')}</FooterLink>
            <FooterLink href="#api">{t('header.api')}</FooterLink>
            <FooterLink href="#mcp">{t('header.mcp')}</FooterLink>
          </FooterLinks>
          
          <SocialLinks>
            <SocialLink 
              href="https://github.com/cyberspacesec/go-iconhash" 
              target="_blank" 
              rel="noopener noreferrer"
            >
              <GithubOutlined />
            </SocialLink>
          </SocialLinks>
          
          <Copyright>{t('footer.copyright')}</Copyright>
          <License>{t('footer.license')}</License>
          
          <LanguageSelector>
            <Space>
              <span>{t('footer.language')}:</span>
              <Radio.Group 
                value={i18n.language} 
                onChange={handleLanguageChange}
                optionType="button"
                buttonStyle="solid"
              >
                <Radio.Button value="en">English</Radio.Button>
                <Radio.Button value="zh">中文</Radio.Button>
              </Radio.Group>
            </Space>
          </LanguageSelector>
        </FooterContent>
      </div>
    </FooterContainer>
  );
};

export default Footer; 