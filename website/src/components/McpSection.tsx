import React from 'react';
import { useTranslation } from 'react-i18next';
import { Typography, Card } from 'antd';
import { RobotOutlined } from '@ant-design/icons';
import styled from 'styled-components';

const { Paragraph } = Typography;

const McpSectionContainer = styled.section`
  background-color: var(--bg-light);
  padding: 5rem 0;
`;

const SectionTitle = styled.h2`
  font-size: 2.5rem;
  font-weight: 700;
  text-align: center;
  margin-bottom: 1.5rem;
  color: var(--text-color);
`;

const SectionDescription = styled(Paragraph)`
  text-align: center;
  font-size: 1.1rem;
  max-width: 800px;
  margin: 0 auto 3rem;
  color: var(--text-light);
`;

const CodeBlock = styled.pre`
  background-color: var(--code-bg);
  padding: 1rem;
  border-radius: var(--border-radius);
  overflow-x: auto;
  font-family: 'Fira Code', monospace;
  color: var(--code-color);
`;

const IconContainer = styled.div`
  font-size: 4rem;
  color: var(--primary-color);
  text-align: center;
  margin-bottom: 2rem;
`;

const McpSection: React.FC = () => {
  const { t } = useTranslation();
  
  return (
    <McpSectionContainer id="mcp">
      <div className="container">
        <IconContainer>
          <RobotOutlined />
        </IconContainer>
        <SectionTitle>{t('mcp.title')}</SectionTitle>
        <SectionDescription>{t('mcp.description')}</SectionDescription>
        
        <Card title={t('mcp.example')} bordered={false}>
          <CodeBlock>
{`POST /mcp HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "version": "0.1",
  "context_id": "12345",
  "message": "Calculate the favicon hash for example.com",
  "references": [
    {
      "id": "ref1",
      "url": "https://example.com/favicon.ico"
    }
  ]
}`}
          </CodeBlock>
        </Card>
      </div>
    </McpSectionContainer>
  );
};

export default McpSection; 