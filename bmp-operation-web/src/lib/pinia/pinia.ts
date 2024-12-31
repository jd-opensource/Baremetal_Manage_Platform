import {createPinia, Pinia} from 'pinia';

import piniaPluginPersist from 'pinia-plugin-persist';

const pinia: Pinia = createPinia();

pinia.use(piniaPluginPersist);

export default pinia;

