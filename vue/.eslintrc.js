module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/recommended',
    '@vue/airbnb',
  ],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-plusplus': ['error', { allowForLoopAfterthoughts: true }],
    'import/extensions': ['error', 'always', { vue: 'never' }],
    'vue/component-name-in-template-casing': 'error',
    'vue/max-attributes-per-line': ['error', { singleline: 3 }],
    'vue/singleline-html-element-content-newline': 'off',
  },
};
