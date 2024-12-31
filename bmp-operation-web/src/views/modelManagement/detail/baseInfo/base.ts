/**
 * @file
 * @author
*/

import {msgTip} from 'utils/index.ts';
import {CurrencyType} from '@utils/publicType';
import {modelDetailAPI} from 'api/model/request.ts';
import {CurrentInstance, language} from 'utils/publicClass.ts';
import ModelStaticData from 'staticData/model/index.ts'
import AllStaticData from 'staticData/staticData.ts';

class BasicInfo {
    basicInfoLoading: Ref<boolean> = ref<boolean>(true);
    // 是否请求机型详情-默认true
    isRequestmodelDetail: Ref<boolean> = ref<boolean>(true);
    reactiveArr: any = reactive<any>({
        // tableRef: {},
        // filterParams: {},
        type1: 0,
        type2: 0,
        detail: {},
        formData: {}
    });
    dataInit: any;
    router: any;
    tableList?: any;
    instanceProxy = new CurrentInstance().proxy;

    constructor (dataInit: any, router: any, tableList?: any) {
        this.dataInit = dataInit;
        this.router = router;
        this.tableList = tableList;
    };

    getModelDetail = (): void => {
        this.basicInfoLoading.value = true;
        modelDetailAPI(
            {
                id: this.dataInit.deviceTypeId
            }
        ).then(({data}: {data: {result: CurrencyType}}) => {
            if (data?.result && Object.keys(data.result)?.length) {
                this.reactiveArr.detail = data.result;
                this.reactiveArr.formData.value = this.reactiveArr.detail;
                this.reactiveArr.type1 = data.result.instanceCount;
                this.reactiveArr.type2 = data.result.deviceCount;
                this.isRequestmodelDetail.value = false;
                this.basicInfoLoading.value = false;
                return;
            }
            this.basicInfoLoading.value = false;
        }).catch(({message} : {message: string;}) => {
            const path: string = this.instanceProxy.$defInfo.routerPath('modelList');
            if (AllStaticData.noFoundData.includes(message)) {
                this.router.push(path);
                return;
            }
            this.basicInfoLoading.value = false;
            // !this.tableList.errorTip && 
            msgTip.error(message);
            this.isRequestmodelDetail.value = true;
        }).finally(() => {
            if (this.dataInit.routerType === 'image') {
                if (this.dataInit.ossStore?.source?.length) return;
                this.dataInit.ossStore?.getOssData();
            }
        });
    };

    setDiskInterfaceType = (item: string) => {
        return item === 'notLimited' ? language.t('modelForm.unrestricted') : item;
    }

    setRaidCan = (raidCan: string) => {
        return ModelStaticData.singleRaid0.includes(raidCan) ? language.t('modelForm.raid') : raidCan
    }
};

export default BasicInfo;
