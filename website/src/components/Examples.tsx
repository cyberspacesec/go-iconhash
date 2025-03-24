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
iconhash url https://www.example.com

# With debug output
iconhash url -d https://www.example.com

# With Shodan format
iconhash url https://www.example.com --shodan

# With URL flag explicitly
iconhash url -u https://www.example.com/favicon.ico`}
              </pre>
            </CodeBlock>
          </TabPane>
          
          <TabPane tab={t('examples.hashFromFile.title')} key="file">
            <ExampleDescription>{t('examples.hashFromFile.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Basic usage
iconhash file favicon.ico

# With explicit file flag
iconhash file -f favicon.ico

# With uint32 format
iconhash file favicon.ico --uint32

# With shorthand syntax
iconhash -- favicon.ico`}
              </pre>
            </CodeBlock>
          </TabPane>
          
          <TabPane tab={t('examples.hashFromBase64.title')} key="base64">
            <ExampleDescription>{t('examples.hashFromBase64.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Basic usage
iconhash base64 encoded.txt

# With explicit base64 flag
iconhash base64 -b encoded.txt

# With FOFA format output
iconhash base64 encoded.txt --fofa

# With shorthand syntax
iconhash -b encoded.txt`}
              </pre>
            </CodeBlock>
          </TabPane>
          
          <TabPane tab={t('examples.apiServer.title')} key="server">
            <ExampleDescription>{t('examples.apiServer.description')}</ExampleDescription>
            <CodeBlock>
              <pre>
{`# Start on default port (8000)
iconhash server

# Start on custom port with debug output
iconhash server -p 3000 -d

# Start with authentication token
iconhash server --auth-token "your-secret-token"

# Bind to specific host with custom timeouts
iconhash server --host 0.0.0.0 --read-timeout 60 --write-timeout 60`}
              </pre>
            </CodeBlock>
          </TabPane>
        </StyledTabs>
      </div>
    </ExamplesSection>
  );
};

export default Examples; 