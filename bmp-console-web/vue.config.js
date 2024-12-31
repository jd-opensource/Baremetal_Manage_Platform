// 页面路径
const setAlias = require('./config/setAlias');
// 跨域配置
const devServer = require('./config/devServer');
// 避免文件的频繁变更导致浏览器缓存失效，更好的利用缓存。提升用户体验
const runtimeChunk = require('./config/runtimeChunk');
// 可以看到启动时的消耗模块，方便后续做优化
const speedMeasure = require('./config/speedMeasure');
// prod环境移除console
const terserPlugin = require('./config/terserPlugin');
// 减少不必要的解析
const noParse  = require('./config/noParse');

module.exports = {
    productionSourceMap: false, // 是否在构建生产包时生成 sourceMap 文件，false将提高构建速度
    // 修改默认的webpack的配置
    configureWebpack: config => {
        runtimeChunk(config);
        speedMeasure(config);
        noParse(config);
    },
    // 调整内部的webpack配置
    chainWebpack: config => {
        // 配置alias别名
        setAlias(config);
        // 删除懒加载模块的prefetch，降低带宽压力
        config.plugins.delete('prefetch');
        terserPlugin(config);
    },
    ...devServer
};
