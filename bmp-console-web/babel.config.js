module.exports = {
    presets: [
        '@vue/cli-plugin-babel/preset'
    ],
    plugins: [
        // 解析es6可选链式语法
        ['@babel/plugin-proposal-optional-chaining']
    ]
}
