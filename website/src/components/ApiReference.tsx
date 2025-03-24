import React from 'react';
import { useTranslation } from 'react-i18next';
import { Typography, Table, Space, Card, Divider } from 'antd';
import styled from 'styled-components';

const { Title, Paragraph, Text } = Typography;

const ApiSection = styled.section`
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

const SubSection = styled.div`
  margin-bottom: 3rem;
`;

const SubTitle = styled.h3`
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  color: var(--text-color);
`;

const StyledTable = styled(Table)`
  box-shadow: var(--shadow);
  border-radius: var(--border-radius);
  overflow: hidden;
  
  .ant-table-thead > tr > th {
    background-color: var(--bg-light);
    font-weight: 600;
  }
`;

const CodeBlock = styled.pre`
  background-color: var(--code-bg);
  padding: 1rem;
  border-radius: var(--border-radius);
  overflow-x: auto;
  font-family: 'Fira Code', monospace;
  color: var(--code-color);
`;

const ApiReference: React.FC = () => {
  const { t } = useTranslation();
  
  const endpointColumns = [
    {
      title: t('api.endpoints'),
      dataIndex: 'endpoint',
      key: 'endpoint',
    },
    {
      title: t('api.method'),
      dataIndex: 'method',
      key: 'method',
    },
    {
      title: t('api.description'),
      dataIndex: 'description',
      key: 'description',
    },
  ];
  
  const endpointData = [
    {
      key: '1',
      endpoint: '/health',
      method: 'GET',
      description: t('api.healthDesc'),
    },
    {
      key: '2',
      endpoint: '/hash/url',
      method: 'GET, POST',
      description: t('api.urlDesc'),
    },
    {
      key: '3',
      endpoint: '/hash/file',
      method: 'POST',
      description: t('api.fileDesc'),
    },
    {
      key: '4',
      endpoint: '/hash/base64',
      method: 'POST',
      description: t('api.base64Desc'),
    },
    {
      key: '5',
      endpoint: '/mcp',
      method: 'POST',
      description: t('api.mcpDesc'),
    },
  ];

  return (
    <ApiSection id="api">
      <div className="container">
        <SectionTitle>{t('api.title')}</SectionTitle>
        <SectionDescription>{t('api.description')}</SectionDescription>
        
        <SubSection>
          <SubTitle>{t('api.endpoints')}</SubTitle>
          <StyledTable 
            columns={endpointColumns} 
            dataSource={endpointData} 
            pagination={false}
          />
        </SubSection>
        
        <SubSection>
          <SubTitle>{t('api.authentication')}</SubTitle>
          <Paragraph>
            {t('api.authDescription')}
          </Paragraph>
          <Space direction="vertical" style={{ width: '100%' }}>
            <Text code>Authorization: Bearer your-token</Text>
            <Text code>?token=your-token</Text>
          </Space>
        </SubSection>
        
        <SubSection>
          <SubTitle>{t('api.examples')}</SubTitle>
          <Card title={t('api.urlExample')} bordered={false}>
            <CodeBlock>
{`# Get hash from URL
curl -X GET "http://localhost:8080/hash/url?url=https://example.com/favicon.ico"

# With authentication
curl -X GET "http://localhost:8080/hash/url?url=https://example.com/favicon.ico" \\
  -H "Authorization: Bearer your-token"`}
            </CodeBlock>
          </Card>
          <Divider />
          <Card title={t('api.fileExample')} bordered={false}>
            <CodeBlock>
{`# Upload a file for hashing
curl -X POST "http://localhost:8080/hash/file" \\
  -F "file=@/path/to/favicon.ico" \\
  -H "Content-Type: multipart/form-data"`}
            </CodeBlock>
          </Card>
        </SubSection>
      </div>
    </ApiSection>
  );
};

export default ApiReference; 