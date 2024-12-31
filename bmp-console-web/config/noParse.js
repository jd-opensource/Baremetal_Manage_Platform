module.exports = (config) => {
    // 减少不必要的解析，让 Webpack 忽略对部分没采用模块化的文件的递归解析处理，这样做的好处是能提高构建性能
	config.module.noParse = '/^(vue|vue-router|vue-i18n)$/';
	return config
};
