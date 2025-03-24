import React from 'react';
import { useTranslation } from 'react-i18next';
import { Button, Space } from 'antd';
import { DownloadOutlined, PlayCircleOutlined, GithubOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const HeroSection = styled.section`
  background: linear-gradient(to bottom, var(--bg-light), white);
  min-height: 70vh;
  display: flex;
  align-items: center;
  text-align: center;
  padding: 6rem 0;
`;

const HeroTitle = styled.h1`
  font-size: 3.5rem;
  font-weight: 800;
  margin-bottom: 1rem;
  color: var(--text-color);
  
  @media (max-width: 768px) {
    font-size: 2.5rem;
  }
`;

const HeroSubtitle = styled.h2`
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  color: var(--text-color);
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
  
  @media (max-width: 768px) {
    font-size: 1.2rem;
  }
`;

const HeroDescription = styled.p`
  font-size: 1.2rem;
  margin-bottom: 2rem;
  color: var(--text-light);
  max-width: 700px;
  margin-left: auto;
  margin-right: auto;
  
  @media (max-width: 768px) {
    font-size: 1rem;
  }
`;

const ButtonGroup = styled(Space)`
  margin-top: 1rem;
`;

const StyledButton = styled(Button)`
  height: 42px;
  font-size: 1rem;
`;

const Hero: React.FC = () => {
  const { t } = useTranslation();
  
  return (
    <HeroSection>
      <div className="container">
        <HeroTitle>{t('hero.title')}</HeroTitle>
        <HeroSubtitle>{t('hero.subtitle')}</HeroSubtitle>
        <HeroDescription>{t('hero.description')}</HeroDescription>
        
        <ButtonGroup size="large">
          <StyledButton 
            type="primary" 
            icon={<DownloadOutlined />} 
            size="large"
            href="https://github.com/cyberspacesec/go-iconhash/releases" 
            target="_blank"
          >
            {t('hero.download')}
          </StyledButton>
          
          <StyledButton 
            icon={<PlayCircleOutlined />} 
            href="#examples"
          >
            {t('hero.tryOnline')}
          </StyledButton>
          
          <StyledButton 
            icon={<GithubOutlined />} 
            href="https://github.com/cyberspacesec/go-iconhash" 
            target="_blank"
          >
            {t('hero.viewOnGitHub')}
          </StyledButton>
        </ButtonGroup>
      </div>
    </HeroSection>
  );
};

export default Hero; 