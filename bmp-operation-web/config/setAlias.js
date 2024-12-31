const path = require('path');
const resolve = (dir) => path.join(__dirname, '..', dir);

// 设置alias别名、配置简化路径
module.exports = config => {
    config.resolve.alias
        .set('@', resolve('src'))
        .set('api', resolve('src/api'))
        .set('@utils', resolve('src/utils/'))
        .set('assets', resolve('src/assets'))
        .set('utils', resolve('src/utils'))
        .set('mock', resolve('src/mock'))
        .set('staticData', resolve('src/staticData'))
        .set('components', resolve('src/components'))
        .set('hooks', resolve('src/hooks'))
        .set('lib', resolve('src/lib'))
        .set('plugin', resolve('src/plugin'))
        .set('request', resolve('src/request'))
        .set('router', resolve('src/router'))
        .set('store', resolve('src/store'))
        .set('views', resolve('src/views'))
        .set('vue-i18n', 'vue-i18n/dist/vue-i18n.cjs.js') // 解决控制台警告
    return config;
};
