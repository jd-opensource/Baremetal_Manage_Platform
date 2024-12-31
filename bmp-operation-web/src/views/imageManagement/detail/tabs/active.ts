import {language} from 'utils/publicClass.ts';
import {CurrencyType6, CurrencyType} from '@utils/publicType';
import router from 'router/index.ts';
import store from 'store/index.ts';

interface DataType {
    activeName: Ref<string>;
    tabsPangeData: CurrencyType6[];
}

const tabsOperate = (associatedModel: any, query: any, baseInfo: any) => {
    const data: DataType = {
        // 标签页切换的model
        activeName: ref<string>(''),
        tabsPangeData: [ // 标签页切换信息
            {
                label: language.t('imageDetail.tabs.basicInfo'),
                name: 'basicInfo'
            },
            {
                label: language.t('imageDetail.tabs.model'),
                name: 'model'
            }
        ]
    };

    onMounted(() => {
        if (query?.deviceTypeNum?.length) {
            data.activeName.value = 'model';
            return;
        }
        data.activeName.value = 'basicInfo';
    })

    watch(() => data.activeName.value, (newValue: string) => {
        const requestData = [
            [
                (newValue: string) => Object.is(newValue, 'basicInfo'),
                () => baseInfo.getImageDetail()
            ],
            [
                (newValue: string) => Object.is(newValue, 'model'),
                () => associatedModel.getTable()
            ]
        ];
        for(const key of requestData) {
            if (key[0](newValue)) {
                key[1](newValue);
                break;
            }
        }
    });

    /**
     * 标签页切换
     * @param {string} value 当前标签页点击切换的名称
     * @return {string} activeName.value 当前标签页切换的数据
    */
    const tabsChange = (value: string): void => {
        data.activeName.value = value;
        const path: string = router.currentRoute.value.path;
        const query: CurrencyType = router.currentRoute.value.query;
        let obj: CurrencyType = {};
        let activeQuery: string = '';
        for (let index in query) {
            obj = {...obj, [index]: query[index], deviceTypeNum: query?.deviceTypeNum || '0'}
            if (value === 'basicInfo') {
                obj.deviceTypeNum = '';
                store.filterEmpty.deleteEmtpyData(obj);
            }
            else {
                obj.deviceTypeNum = query?.deviceTypeNum || '0';
            }
        }
        for (let key in obj) {
            activeQuery += `${[key]}=${obj[key]}&`
        }
        activeQuery = activeQuery.slice(0, -1);
        router.push(`${path}?${activeQuery}`);
    };

    return {
        ...data,
        tabsChange
    }
};

export default tabsOperate;
