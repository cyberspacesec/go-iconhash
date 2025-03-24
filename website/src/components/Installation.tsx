import React from 'react';
import { useTranslation } from 'react-i18next';
import { Row, Col, Card, Typography } from 'antd';
import styled from 'styled-components';

const { Title, Paragraph, Text } = Typography;

const InstallationSection = styled.section`
  padding: 5rem 0;
`;

const SectionTitle = styled.h2`
  font-size: 2.5rem;
  font-weight: 700;
  text-align: center;
  margin-bottom: 3rem;
  color: var(--text-color);
`;

const InstallCard = styled(Card)`
  height: 100%;
  border-radius: var(--border-radius);
  box-shadow: var(--shadow);
  
  .ant-card-head {
    background-color: var(--bg-light);
    border-top-left-radius: var(--border-radius);
    border-top-right-radius: var(--border-radius);
  }
  
  .ant-card-head-title {
    font-weight: 600;
  }
  
  pre {
    background-color: var(--code-bg);
    padding: 1rem;
    border-radius: var(--border-radius);
    overflow-x: auto;
    font-family: 'Fira Code', monospace;
    color: var(--code-color);
  }
`;

const Installation: React.FC = () => {
  const { t } = useTranslation();
  
  return (
    <InstallationSection id="installation">
      <div className="container">
        <SectionTitle>{t('installation.title')}</SectionTitle>
        
        <Row gutter={[24, 24]}>
          <Col xs={24} md={8}>
            <InstallCard 
              title={t('installation.fromSource.title')}
              bordered={false}
            >
              <Paragraph>{t('installation.fromSource.description')}</Paragraph>
              <pre>
{`# Clone the repository
git clone https://github.com/cyberspacesec/go-iconhash.git
cd go-iconhash

# Build the binary
make build`}
              </pre>
            </InstallCard>
          </Col>
          
          <Col xs={24} md={8}>
            <InstallCard 
              title={t('installation.usingGo.title')}
              bordered={false}
            >
              <Paragraph>{t('installation.usingGo.description')}</Paragraph>
              <pre>
{`# Install latest version
go install github.com/cyberspacesec/go-iconhash/cmd/iconhash@latest`}
              </pre>
            </InstallCard>
          </Col>
          
          <Col xs={24} md={8}>
            <InstallCard 
              title={t('installation.usingDocker.title')}
              bordered={false}
            >
              <Paragraph>{t('installation.usingDocker.description')}</Paragraph>
              <pre>
{`# Pull the image
docker pull cyberspacesec/iconhash:latest

# Run with Docker
docker run --rm cyberspacesec/iconhash:latest -u https://example.com/favicon.ico`}
              </pre>
            </InstallCard>
          </Col>
        </Row>
      </div>
    </InstallationSection>
  );
};

export default Installation; 