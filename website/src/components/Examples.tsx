import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Tabs, Typography } from 'antd';
import styled from 'styled-components';

const { TabPane } = Tabs;
const { Paragraph } = Typography;

const ExamplesSection = styled.section`
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

const ExampleDescription = styled(Paragraph)`
  font-size: 1.1rem;
  margin-bottom: 1.5rem;
`;

const CodeBlock = styled.div`
  background-color: var(--code-bg);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow);
  overflow: hidden;
  
  pre {
    padding: 1.5rem;
    margin: 0;
    overflow-x: auto;
    font-family: 'Fira Code', monospace;
    color: var(--code-color);
  }
`;

const StyledTabs = styled(Tabs)`
  .ant-tabs-nav {
    margin-bottom: 2rem;
  }
  
  .ant-tabs-tab {
    padding: 0.75rem 1.5rem;
    font-size: 1rem;
    font-weight: 500;
  }
  
  .ant-tabs-tab-active {
    font-weight: 600;
  }
`;

const Examples: React.FC = () => {
  const { t } = useTranslation();
  
  return (
    <ExamplesSection id="examples">
      <div className="container">
        <SectionTitle>{t('examples.title')}</SectionTitle>
        
        <StyledTabs defaultActiveKey="url">
          <TabPane tab={t('examples.hashFromUrl.title')} key="url">
            <ExampleDescription>{t('examples.hashFromUrl.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Basic usage
iconhash https://www.example.com/favicon.ico

# With debug output
iconhash -d -u https://www.example.com/favicon.ico

# With Shodan format
iconhash -u https://www.example.com/favicon.ico --shodan --fofa=false`}
              </pre>
            </CodeBlock>
          </TabPane>
          
          <TabPane tab={t('examples.hashFromFile.title')} key="file">
            <ExampleDescription>{t('examples.hashFromFile.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Basic usage
iconhash favicon.ico

# With explicit file flag
iconhash -f favicon.ico

# With multiple output formats
iconhash -f favicon.ico --shodan --fofa`}
              </pre>
            </CodeBlock>
          </TabPane>
          
          <TabPane tab={t('examples.hashFromBase64.title')} key="base64">
            <ExampleDescription>{t('examples.hashFromBase64.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Basic usage
iconhash -b64 encoded.txt

# With custom timeout
iconhash -b64 encoded.txt -t 30`}
              </pre>
            </CodeBlock>
          </TabPane>
          
          <TabPane tab={t('examples.apiServer.title')} key="server">
            <ExampleDescription>{t('examples.apiServer.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Start on default port (8080)
iconhash server

# Start on custom port with debug output
iconhash server -p 3000 -d

# Start with authentication
iconhash server -a "your-secret-token"`}
              </pre>
            </CodeBlock>
          </TabPane>
        </StyledTabs>
      </div>
    </ExamplesSection>
  );
};

export default Examples; 