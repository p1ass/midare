module.exports = {
  plugins: ['prettier'],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:import/errors',
    'plugin:import/warnings',
    'plugin:import/typescript',
    'plugin:@next/next/recommended',
    'prettier',
  ],
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly',
  },
  rules: {
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    'import/order': [
      'error',
      {
        pathGroups: [
          {
            pattern: '@/**',
            group: 'parent',
            position: 'after',
          },
        ],
        'newlines-between': 'always',
      },
    ],
    'import/no-default-export': 'warn',
    'prettier/prettier': [
      'error',
      {
        singleQuote: true,
        printWidth: 100,
        tabWidth: 2,
        semi: false,
      },
    ],
    semi: ['error', 'never'],
    indent: ['error', 2],
  },
  // Next.js向けのページコンポーネントはdefault exportしか使えないなので除外
  overrides: [
    {
      files: ['src/pages/**/*.tsx'],
      rules: {
        'import/no-default-export': 'off',
      },
    },
  ],
  settings: {
    react: {
      version: 'detect',
    },
    'import/resolver': {
      typescript: {
        project: '.',
      },
    },
  },
}
