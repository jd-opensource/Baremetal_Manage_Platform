import {defineStore} from 'pinia'; // 定义容器名
import {FaultSpecStateType} from '../typeManagement';
import {faultSpecAPI} from 'api/surveillance/faultRule/request.ts';
import storeName from 'store/storeName.ts'; // 容器名
import SurveillanceStaticeData from 'staticData/surveillance/index.ts';

const faultSpecStore = defineStore(storeName.faultSpecData, {
    state: (): FaultSpecStateType => {
        return {
            loading: true,
            faultSpecData: []
        }
    },

    actions: {

        getFaultSpec () {
            faultSpecAPI({})
            .then(({data}: {data: {result: string[]}}) => {
                if (data?.result?.length) {
                    this.faultSpecData = data.result;
                    return;
                }
                return Promise.reject();
            })
            .catch(() => {
                this.faultSpecData = SurveillanceStaticeData.faultyAccessories;
            })
            .finally(() => this.loading = false);
        }
    }
});

export default faultSpecStore;