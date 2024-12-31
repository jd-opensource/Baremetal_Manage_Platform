module.exports = {
    devServer: {
        open: true, // 是否直接打开浏览器
        host: 'localhost', // 本地服务器
        port: 8089, // 端口号
        proxy: { // 跨域
            '/console-web': {
                target: 'http://bmp-console.bmp.local',
                changeOrigin: true, // 是否需要跨域
            },
            '/oob-alert': {
                target:'http://bmp-oob-alert.bmp.local',
                changeOrigin: true, // 是否需要跨域
                pathRewrite: {
                    '^/oob-alert': ''   //需要rewrite的,
                }
            }
        },
        sockHost: 'localhost',
        disableHostCheck: true,
        progress: false
    }
};
