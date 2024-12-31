module.exports = {
    devServer: {
        overlay: { // 让浏览器 overlay 同时显示警告和错误
            warnings: true,
            errors: true
        },
        open: true, // 是否直接打开浏览器
        host: 'localhost', // 本地服务器
        port: 8089, // 端口号
        hotOnly: true, // 热更新，在某些模块不支持热更新的情况下，在控制台输出热更新失败
        proxy: { // 跨域
            '/operation-web': {
                target: 'http://bmp-operation.bmp.local',
                // ws: true,
                changeOrigin: true // 是否需要跨域
            },
            '/template-upload': {
                target: 'http://bmp-operation.bmp.local',
                changeOrigin: true, // 是否需要跨域
                // ws: true,
                pathRewrite: {
                    '^/template-upload': ''
                }
            },
            '/oob-alert': {
                target: 'http://bmp-oob-alert.bmp.local',
                changeOrigin: true, // 是否需要跨域
                pathRewrite: {
                    '^/oob-alert': ''
                }
            }
        },
        sockHost: 'localhost',
        disableHostCheck: true, // 禁止主机检查
        progress: false
    }
};
