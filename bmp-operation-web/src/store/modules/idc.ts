import {defineStore} from 'pinia'; // 定义容器名
import {locationItem} from 'utils/publicClass.ts';
import {CurrencyType} from '@utils/publicType';
import storeName from 'store/storeName.ts'; // 容器名
import {IdcStoreStateType} from '../typeManagement';

const idcStore = defineStore(storeName.idcData, {
    state: (): IdcStoreStateType => {
        return {
            loading: true,
            idcData: [],
            idcDataNoOpt: []
        }
    },

    actions: {
        async getIdc (apiReq: Function) {
            try {
                const idcApi = await apiReq({isAll: '1'});
                if (idcApi?.data?.result?.idcs?.length) {
                    const {idcs} : {idcs: CurrencyType[]} = idcApi.data.result;
                    const newIdcs = idcs.map((item: CurrencyType) => {
                        return {
                            ...item,
                            newIdcName: this.setName(locationItem.getLocationItem, item)
                        };
                    })
                    this.idcData = newIdcs;
                    return;
                }
                throw new Error('');
            }
            catch {
                this.idcData = [];
            }
            finally {
                this.loading = false;
            }
        },
    
        setName (name: string, item: CurrencyType) {
            const keyName: Map<string, string> = new Map([
                ['zh_CN', item.name],
                ['en_US', item.nameEn]
            ]);
            return keyName.get(name);
        },

        async idcListNoOpt (apiReq: Function) {
            try {
                const idcApi = await apiReq({isAll: '1'});
                if (idcApi?.data?.result?.idcs?.length) {
                    const {idcs}: {idcs: CurrencyType[]} = idcApi.data.result;
                    this.idcDataNoOpt = idcs;
                    return;
                }
                throw new Error('');
            }
            catch {
                this.idcDataNoOpt = [];
            }
        }
    }
});

export default idcStore;