module.exports = {
    env: {
      browser: true,
      es6: true,
      node: true
    },
    plugins: ['react','react-hooks', '@typescript-eslint', 'prettier'],
    extends: [
      'eslint:recommended',
      'plugin:@typescript-eslint/eslint-recommended',
      'react-app',
      'plugin:react-hooks/recommended',
      'plugin:prettier/recommended',
      'prettier/@typescript-eslint',
      'prettier/react'
    ],
    globals: {
      Atomics: 'readonly',
      SharedArrayBuffer: 'readonly'
    },
    parser: '@typescript-eslint/parser',
    parserOptions: {
      ecmaFeatures: {
        jsx: true
      },
      ecmaVersion: 2018,
      sourceType: 'module'
    },
    rules: {
      'prettier/prettier': ['error', { 
        'singleQuote': true ,
        'printWidth': 100,
        'tabWidth': 2,
        'semi': false
      }],
      semi: ['error', 'never'],
      indent: ['error', 2]
    }
  }
