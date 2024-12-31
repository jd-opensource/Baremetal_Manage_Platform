module.exports = {
    root: true,
    env: {
        'node': true
    },
    // eslint关闭
    rules: {
        'no-unused-vars': 'off',
        'linebreak-style': ['error', 'windows'],
        'no-console': 'off', // 禁用console
        'no-debugger': 'off', // 禁用debugger
        'space-before-function-paren': 'off', // 强制在function的左括号之前使用一致的空格
        'no-underscore-dangle': 'off', // 禁止标识符中有悬空下划线
        'no-param-reassign': 'off', // 禁止对函数参数再赋值
        'func-names': 'off', // 禁止命名的 function 表达式
        'no-bitwise': 'off', // 禁止使用按位操作符
        'prefer-rest-params': 'off', // 建议使用剩余参数代替 arguments
        'no-trailing-spaces': 'off', // 禁用行尾空格
        'comma-dangle': 'off', // 禁止使用拖尾逗号
        'quote-props': 'off', // 要求对象字面量属性名称使用引号
        'consistent-return': 'off', // 要求return语句要么总是指定返回的值,要么不指定
        'no-plusplus': 'off', // 禁用一元操作符++和--
        'prefer-spread': 'warn', // 要求使用扩展运算符而非.apply()
        'semi': 'warn', // 要求使用分号代替ASI
        'indent': ['warn', 'tab'], // 强制使用一致的缩进
        'no-tabs': 'warn', // 禁用tab风格字符
        'no-unused-vars': 'warn', // 禁止出现未使用过的变量
        'quotes': 'warn', // 强制使用一致的反勾号、双引号或单引号
        'no-void': 'off', // 禁止使用void操作符
        'no-nested-ternary': 'off', // 禁止使用嵌套的三元表达式
        'import/no-unresolved': 'off',
        'no-return-assign': 'warn', // 禁止在return语句中使用赋值语句
        'linebreak-style': 'off', // 强制使用一致的换行符风格
        'prefer-destructuring': 'off', // 优先使用数组和对象解构
        'no-restricted-syntax': 'warn', // 禁止使用特定的语法
        'arrow-parens': 'off', // 要求箭头函数的参数使用圆括号
        'import/extensions': 'off',
        'import/named': 'off',
        'import/no-deprecated': 'off',
        'import/no-extraneous-dependencies': 'off',
        'fecs-no-require': 'off',
        'fecs-valid-jsdoc': 'off'
    },
    parserOptions: {
        'parser': 'babel-eslint'
    },
    extends: [
        'plugin:@typescript-eslint/no-unused-vars'
    ]
};
