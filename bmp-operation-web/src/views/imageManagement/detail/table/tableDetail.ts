/**
 * @file
 * @author
*/

import {imageDeviceTypesAPI} from 'api/image/request.ts';
import {paginationOperate, CurrentInstance} from 'utils/publicClass.ts';
import {msgTip} from 'utils/index.ts';
import {CurrencyType} from '@utils/publicType';
import router from 'router/index.ts';
import ImageStaticData from 'staticData/image/index.ts';
import AllStaticData from 'staticData/staticData.ts';

class AssociatedModel {
    isLoading: Ref<boolean> = ref<boolean>(true);
    // 复杂数据类型
    reactiveArr: any = reactive<any>({
        tableData: [] // 表格关联机型数据
    });
    query: any;
    baseInfo: any;
    instanceProxy = new CurrentInstance().proxy;

    constructor (query: any, baseInfo: any) {
        this.query = query;
        this.baseInfo = baseInfo;
    };

    getTable = () => {
        this.isLoading.value = true;
        imageDeviceTypesAPI(
            {
                imageId: this.query?.imageId,
                architecture: this.query?.architecture,
                isBind: '1',
                pageNumber: paginationOperate.pageNumber.value,
                pageSize: paginationOperate.pageSize.value
            }
        )
        .then(({data}: {data: {result: {totalCount: number; deviceTypes: CurrencyType[]}}}) => {
            if (data?.result?.deviceTypes?.length) {
                data.result.deviceTypes.forEach((item: CurrencyType, index: number) => {
                    ImageStaticData.imageDetailTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                })
                this.reactiveArr.tableData = data.result.deviceTypes;
                paginationOperate.total.value = data.result.totalCount;
                return;
            }
            
            return Promise.reject('')
        })
        .catch(({message} : {message: string}) => {
            paginationOperate.total.value = 0;
            this.reactiveArr.tableData = [];
            message && msgTip.error(message);
            const path: string = this.instanceProxy.$defInfo.routerPath('imageList');
            if (AllStaticData.noFoundData.includes(message)) {
                router.push(path);
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            this.baseInfo.getImageDetail();
        });
    };
};

export default AssociatedModel;
