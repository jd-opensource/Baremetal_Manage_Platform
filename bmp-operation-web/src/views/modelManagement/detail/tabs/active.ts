import {language} from 'utils/publicClass.ts';
import {CurrencyType, CurrencyType6} from '@utils/publicType';
import router from 'router/index.ts';

interface DataType {
    activeName: Ref<string>;
    tabsPangeData: CurrencyType6[];
}

const tabsOperate = (tableList: any, dataInit: any, baseInfo: any) => {
    const data: DataType = {
        // 标签页切换的model
        activeName: ref<string>('essentialInformation'),
        tabsPangeData: [ // 标签页切换信息
            {
                label: language.t('modelDetail.tabs.basicInfo.name'),
                name: 'essentialInformation'
            },
            {
                label: language.t('modelDetail.tabs.relationImage.name'),
                name: 'associatedMirror'
            }
        ]
    };

    onMounted(() => {
        if (dataInit.routerType === 'image') {
            data.activeName.value = 'associatedMirror';
            tableList.init();
            return;
        }
        baseInfo.getModelDetail();
    });

    watch(() => data.activeName.value, (): void => {
        if (!tableList.reactiveArr.tableData.length && dataInit.routerType !== 'image') {
            tableList.init();
            dataInit.ossStore?.getOssData();
        }
    }, {deep: true});

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
            obj = Object.assign(obj, {[index]: query[index]}, {type: value === 'essentialInformation' ? 'modelDetail' : 'image'})
            activeQuery += `${[index]}=${obj[index]}&`
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
