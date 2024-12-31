// 避免文件的频繁变更导致浏览器缓存失效，更好的利用缓存。提升用户体验
module.exports = (config) => {
	config.optimization.runtimeChunk = true;
	return config;
};
