module.exports = {
  testResultsProcessor: 'jest-sonar-reporter',
  snapshotSerializers: ['<rootDir>/config/jest/styledComponentsSerializer'],
  coverageProvider: 'v8',
  globalSetup: '<rootDir>/config/jest/globalSetup.js',
  setupFiles: ['<rootDir>/config/jest/mocks.js'],
  setupFilesAfterEnv: ['<rootDir>/config/jest/setupTests.js'],
  testMatch: ['<rootDir>/src/**/__tests__/**/*.(j|t)s?(x)', '<rootDir>/src/**/?(*.)(spec|test).(j|t)s?(x)'],
  testPathIgnorePatterns: ['/__fixtures__/', '/__utils__/'],
  testEnvironment: 'jsdom',
  testEnvironmentOptions: {
    url: 'http://localhost/',
  },
  transformIgnorePatterns: ['[/\\\\]node_modules[/\\\\].+\\.(js|jsx|mjs|ts|tsx)$'],
  moduleDirectories: ['<rootDir>/src', 'node_modules'],
  moduleNameMapper: {
    '\\.(css|scss)$': 'identity-obj-proxy',
  },
  moduleFileExtensions: ['web.ts', 'ts', 'web.tsx', 'tsx', 'web.js', 'js', 'web.jsx', 'jsx', 'json', 'node', 'mjs'],
  transform: {
    "^.+\\.(t|j)sx?$": ["@swc/jest", {
      jsc: {
        transform: {
          react: {
            runtime: 'automatic',
          },
        },
        "experimental": {
          "plugins": [["swc_mut_cjs_exports", {}]]
        }
      },
    }],
    '\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$':
      '<rootDir>/config/jest/fileTransform.js',
  },
};
