/**
 * @file
 * @author
*/

import {imagesDetailAPI} from 'api/image/request.ts';
import {CurrencyType} from '@utils/publicType';
import store from 'store/index.ts';
import {msgTip} from 'utils/index.ts';
import {language, CurrentInstance} from 'utils/publicClass.ts';
import router from 'router/index.ts';
import AllStaticData from 'staticData/staticData.ts';

type BaseInfoData = {
    label: string;
    info: string;
    class: string;
    hasEdit: boolean;
    formItemClass?: string[];
}

class BaseInfo {
    infoLoading: Ref<boolean> = ref<boolean>(true);
    deviceTypeNum: number = 0;
    detail: any = {};
    systemPartitionInfo = store.sysPartitionInfo;
    // 复杂数据类型
    reactiveArr: any = reactive<any>({
        deviceTypeNum: 0,
        detail: {},
        description: ''
    });
    query: any;
    instanceProxy = new CurrentInstance().proxy;
    baseInfoData: BaseInfoData[] = [
        {
            label: language.t('imageDetail.label.imageName'),
            info: 'imageName',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.imageId'),
            info: 'imageId',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.imageType'),
            info: 'sourceName',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.architecture'),
            info: 'architecture',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.operateSysPlatform'),
            info: 'osType',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.operateSysVersion'),
            info: 'osVersion',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.format'),
            info: 'format',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.bootMode'),
            info: 'bootMode',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.sysPartitionModule'),
            info: 'newSystemPartition',
            formItemClass: ['sys-partition-module', ''],
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.numberOfAssociatedModels'),
            info: 'deviceTypeNum',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.createTime'),
            info: 'createdTime',
            class: 'set-wrap info-name',
            hasEdit: false
        },
        {
            label: language.t('imageDetail.label.imageDesc'),
            info: 'description',
            class: 'set-wrap info-name',
            hasEdit: true
        }
    ]

    constructor (query: any) {
        this.query = query;
    };

    getImageDetail = () => {
        this.infoLoading.value = true;
        imagesDetailAPI(
            {
                imageId: this.query?.imageId
            }
        )
        .then(({data} : {data: {result: CurrencyType}}) => {
            if (data?.result && Object.keys(data.result).length) {
                const newSystemPartition = this.systemPartitionInfo.systemPartitionData(data.result.systemPartition, false);
                this.reactiveArr.detail = data.result;
                this.reactiveArr.deviceTypeNum = data.result.deviceTypeNum;
                this.reactiveArr.description = data.result.description;
                this.reactiveArr.detail.newSystemPartition = newSystemPartition;
                return;
            }
            return Promise.reject('');
        })
        .catch(({message} : {message: string}) => {
            const path: string = this.instanceProxy.$defInfo.routerPath('imageList');
            this.reactiveArr.detail = {};
            if (AllStaticData.noFoundData.includes(message)) {
                router.push(path);
                return;
            }
            message && msgTip.error(message);
        })
        .finally(() => {
            this.infoLoading.value = false;
        });
    };
};

export default BaseInfo;
