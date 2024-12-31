/**
 * @file
 * @author
*/

import {CurrencyType} from '@utils/publicType';
import {getApiKeyAPI} from 'api/userCenter/request.ts';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';
import UserCenterDataStatic from 'staticData/userCenter/index.ts';

interface ReactiveArrType {
    keyNameData: string[];
    tableData: CurrencyType[];
}

class ApiKey {
    mitt = new CurrentInstance();
    isBtnDisabled: Ref<boolean> = ref<boolean>(true);
    isLoading: Ref<boolean> = ref<boolean>(true);
    reactiveArr: ReactiveArrType = reactive<ReactiveArrType>({
        keyNameData: [],
        tableData: []
    });

    constructor() {
        this.getApiKey();
    }

    getApiKey = () => {
        this.isLoading.value = true;
        getApiKeyAPI(
            {
                isAll: '1'
            }
        )
        .then(({data}: {data: {result: {apikeys: CurrencyType[]; totalCount: number}}}) => {
            if (data?.result?.apikeys?.length) {
                this.reactiveArr.tableData = data.result.apikeys.map((item: CurrencyType, index: number) => {
                    UserCenterDataStatic.apiKeyTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                    return item;
                });
                paginationOperate.total.value = data.result.totalCount;
                this.reactiveArr.keyNameData = data.result.apikeys.map((item: CurrencyType) => item.name);
                this.isBtnDisabled.value = data.result.totalCount >= 2;
                return;
            }
            return Promise.reject();
        })
        .catch(() => {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            this.isBtnDisabled.value = false;
        })
        .finally(() => {
            this.isLoading.value = false;
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('api-key-data', this.reactiveArr.tableData);
        })
    };
};

export default ApiKey;
