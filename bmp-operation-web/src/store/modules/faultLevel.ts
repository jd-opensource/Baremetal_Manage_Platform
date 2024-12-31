import {defineStore} from 'pinia'; // 定义容器名
import {FaultLevelStateType} from '../typeManagement';
import {faultLevelAPI} from 'api/surveillance/faultRule/request.ts';
import storeName from 'store/storeName.ts'; // 容器名
import SurveillanceStaticeData from 'staticData/surveillance/index.ts';

const faultLevelStore = defineStore(storeName.faultLevelData, {
    state: (): FaultLevelStateType => {
        return {
            loading: true,
            faultLevelData: []
        }
    },

    actions: {

        getFaultLevel () {
            faultLevelAPI({})
            .then(({data}: {data: {result: string[]}}) => {
                if (data?.result?.length) {
                    this.faultLevelData = data.result;
                    return;
                }
                return Promise.reject();
            })
            .catch(() => {
                this.faultLevelData = SurveillanceStaticeData.faultLevel;
            })
            .finally(() => this.loading = false);
        }
    }
});

export default faultLevelStore;
