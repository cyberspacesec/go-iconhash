import React from 'react';
import { useTranslation } from 'react-i18next';
import { Row, Col, Card } from 'antd';
import {
  FileAddOutlined,
  FileExcelOutlined,
  ApiOutlined,
  RobotOutlined,
  CodeOutlined,
  DesktopOutlined
} from '@ant-design/icons';
import styled from 'styled-components';

const FeaturesSection = styled.section`
  background-color: var(--bg-light);
  padding: 5rem 0;
`;

const SectionTitle = styled.h2`
  font-size: 2.5rem;
  font-weight: 700;
  text-align: center;
  margin-bottom: 3rem;
  color: var(--text-color);
`;

const FeatureCard = styled(Card)`
  height: 100%;
  border-radius: var(--border-radius);
  box-shadow: var(--shadow);
  transition: var(--transition);
  border: none;
  
  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  }
  
  .ant-card-body {
    height: 100%;
    display: flex;
    flex-direction: column;
  }
`;

const IconWrapper = styled.div`
  font-size: 2.5rem;
  color: var(--primary-color);
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const FeatureTitle = styled.h3`
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  text-align: center;
`;

const FeatureDescription = styled.p`
  color: var(--text-light);
  text-align: center;
  flex-grow: 1;
`;

interface FeatureItemProps {
  icon: React.ReactNode;
  title: string;
  description: string;
}

const FeatureItem: React.FC<FeatureItemProps> = ({ icon, title, description }) => {
  return (
    <FeatureCard>
      <IconWrapper>{icon}</IconWrapper>
      <FeatureTitle>{title}</FeatureTitle>
      <FeatureDescription>{description}</FeatureDescription>
    </FeatureCard>
  );
};

const Features: React.FC = () => {
  const { t } = useTranslation();
  
  const features = [
    {
      key: 'inputs',
      icon: <FileAddOutlined />,
      title: t('features.multipleInputs.title'),
      description: t('features.multipleInputs.description')
    },
    {
      key: 'outputs',
      icon: <FileExcelOutlined />,
      title: t('features.multipleOutputs.title'),
      description: t('features.multipleOutputs.description')
    },
    {
      key: 'api',
      icon: <ApiOutlined />,
      title: t('features.apiServer.title'),
      description: t('features.apiServer.description')
    },
    {
      key: 'mcp',
      icon: <RobotOutlined />,
      title: t('features.mcpSupport.title'),
      description: t('features.mcpSupport.description')
    },
    {
      key: 'docker',
      icon: <DesktopOutlined />,
      title: t('features.docker.title'),
      description: t('features.docker.description')
    },
    {
      key: 'cli',
      icon: <CodeOutlined />,
      title: t('features.cli.title'),
      description: t('features.cli.description')
    }
  ];

  return (
    <FeaturesSection id="features">
      <div className="container">
        <SectionTitle>{t('features.title')}</SectionTitle>
        
        <Row gutter={[24, 24]}>
          {features.map(feature => (
            <Col xs={24} sm={12} lg={8} key={feature.key}>
              <FeatureItem
                icon={feature.icon}
                title={feature.title}
                description={feature.description}
              />
            </Col>
          ))}
        </Row>
      </div>
    </FeaturesSection>
  );
};

export default Features; 