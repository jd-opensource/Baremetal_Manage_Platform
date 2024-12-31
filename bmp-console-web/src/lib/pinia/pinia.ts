import {createPinia, Pinia} from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const pinia: Pinia = createPinia();
// 数据持久化
pinia.use(piniaPluginPersistedstate)
export default pinia;

