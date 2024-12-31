// 移除console
const TerserPlugin = require('terser-webpack-plugin');

module.exports = (config) => {
    if (process.env.VUE_APP_ENV === 'prod') {
        config.optimization.minimizer([
            new TerserPlugin({
                test: /\.js(\?.*)?$/i,
                terserOptions: {
                    compress: {
                        drop_console: true,
                        pure_funcs: ['console.log']
                    }
                }
            })
        ])
    }
    else {
        // 在其他环境期间禁用优化以加快速度
        config.optimization.minimize(false);
    }
    return config;
}
