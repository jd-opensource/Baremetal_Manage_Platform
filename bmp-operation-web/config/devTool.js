const env = process.env.VUE_APP_ENV;

module.exports = (config) => {
	config.devtool = env === 'dev' ? 'source-map' : 'cheap-module-souce-map';
	return config;
};