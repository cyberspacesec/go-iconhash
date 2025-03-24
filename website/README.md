# IconHash Website (React Version)

This is the React-based website for the IconHash tool, a powerful favicon hash calculator for cybersecurity reconnaissance.

## Technologies Used

- React 18
- TypeScript
- Ant Design UI framework
- styled-components for styling
- i18next for internationalization support

## Features

- Responsive design that works on desktop, tablet, and mobile devices
- Multilingual support (English and Chinese)
- Modern and clean user interface with Ant Design components
- Interactive examples and documentation

## Development

### Prerequisites

- Node.js (v14 or newer)
- npm or yarn

### Local Development

To run the website locally:

1. Clone the repository:
   ```
   git clone https://github.com/cyberspacesec/go-iconhash.git
   cd go-iconhash/website
   ```

2. Install dependencies:
   ```
   npm install
   # or
   yarn
   ```

3. Start the development server:
   ```
   npm start
   # or
   yarn start
   ```

4. Visit `http://localhost:3001` in your browser

### Building for Production

To create a production build:

```
npm run build
# or
yarn build
```

The built files will be located in the `dist` directory.

### GitHub Pages Deployment

The website is configured to be deployed to GitHub Pages. To deploy the website:

```
npm run deploy
# or
./deploy.sh
```

This script will:
1. Build the website
2. Copy the built files to the `../docs` directory
3. Commit the changes to git
4. Push the changes to GitHub

After the deployment is complete, the website will be available at:
https://cyberspacesec.github.io/go-iconhash/

Note: GitHub may take a few minutes to process the changes.

### Local Development Server

If you want to test the website locally without deploying to GitHub Pages:

```
npm run serve-local
# or
./serve-gh-pages.sh
```

This will build the project, copy the files to the `../docs` directory, and start a local server to preview the website. Visit `http://localhost:8000` in your browser to see it.

## Language Support

The website supports multiple languages through i18next:

- English: `src/locales/en/translation.json`
- Chinese: `src/locales/zh/translation.json`

To add support for a new language:

1. Create a new JSON file in the `src/locales/{lang}` directory (e.g., `src/locales/fr/translation.json`)
2. Copy the structure from an existing language file
3. Translate all text values
4. Add the language import and configuration in `src/i18n.ts`

## Project Structure

- `src/components/` - React components
- `src/locales/` - Translation files
- `src/styles/` - Global styles
- `src/assets/` - Static assets like images
- `src/i18n.ts` - i18next configuration
- `src/App.tsx` - Main application component
- `src/index.tsx` - Application entry point

## License

This website is released under the [MIT License](../LICENSE). 