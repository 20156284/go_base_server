// .eslintrc.js
module.exports = {
    env: {
        browser: true,
        es2021: true,
        node: true
    },
    extends: [
        'eslint:recommended',
        'plugin:vue/essential',
        'plugin:@typescript-eslint/recommended',
        'plugin:prettier/recommended',
        'prettier/@typescript-eslint'
    ],
    parserOptions: {
        ecmaVersion: 2021,
        parser: '@typescript-eslint/parser',
        sourceType: 'module',
        ecmaFeatures: {
             tsx: true, // Allows for the parsing of JSX
            // jsx: true,
        },
    },
    plugins: ['vue', '@typescript-eslint'],
    rules: {}
}