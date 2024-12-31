// 移除console
const TerserPlugin = require('terser-webpack-plugin');
module.exports = () => {
    if (['prod', 'pre'].includes(process.env.VUE_APP_ENV)) {
        new TerserPlugin({
            parallel: true,
            terserOptions: {
                ecma: undefined,
                warnings: false,
                parse: {},
                compress: {
                    drop_console: true,
                    drop_debugger: false,
                    pure_funcs: ['console.log'], // 移除console
                },
            }
        })
    }
}
