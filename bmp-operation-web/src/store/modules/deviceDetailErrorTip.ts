import {defineStore} from 'pinia'; // 定义容器名
// import {faultLevelAPI} from 'api/surveillance/faultRule/request.ts';
import storeName from 'store/storeName.ts'; // 容器名
// import SurveillanceStaticeData from 'staticData/surveillance/index.ts';

const deviceDetailErrorTipStore = defineStore(storeName.deviceDetailErrorTipOpt, {
    state: () => {
        return {
            errorStatus: false
        }
    },
    // 类似computed
    getters: {
        // 获取raids
        getErrorStatus (state) {
            return state.errorStatus;
        }
    },

    // actions: {

    //     setErrorStatus() {
    //         this.errorStatus = !this.errorStatus;
    //     }
    // }
});

export default deviceDetailErrorTipStore;
