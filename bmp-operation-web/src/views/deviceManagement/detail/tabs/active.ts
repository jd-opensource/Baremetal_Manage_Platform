/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';
import {CurrencyType6, CurrencyType} from '@utils/publicType';
import router from 'router/index.ts';

interface DataType {
    code: Ref<number>;
    activeName: Ref<string>;
    tabsPangeOneData: {label: string; name: string}[];
    tabsPangeTwoData: {label: string; name: string}[];
    tabsPangeThreeData: {label: string; name: string}[];
    tabsPangeData: CurrencyType6[];
}

const tabsOperate = (deviceDetail: {desrcibeAgentStatus(): void; initData(): void;}, query: {type: string}, diskTableOpt: {init(): void;}) => {
    const data: DataType = {
        code: ref<number>(0),
        // 标签页切换的model
        activeName: ref<string>('basicInfo'),
        tabsPangeOneData: [
            {
                label: language.t('deviceDetail.tabs.baseInfo'),
                name: 'basicInfo'
            },
            {
                label: language.t('deviceDetail.tabs.diskDetail'),
                name: 'diskDetail'
            },
            {
                label: language.t('deviceDetail.tabs.operatLog'),
                name: 'operatLog'
            }
        ],
        tabsPangeTwoData: [
            {
                label: language.t('deviceDetail.tabs.baseInfo'),
                name: 'basicInfo'
            },
            {
                label: language.t('deviceDetail.tabs.diskDetail'),
                name: 'diskDetail'
            },
            {
                label: language.t('deviceDetail.tabs.performanceMonitoring'),
                name: 'performanceMonitoring'
            },
            {
                label: language.t('deviceDetail.tabs.operatLog'),
                name: 'operatLog'
            }
        ],
        tabsPangeThreeData: [
            {
                label: language.t('deviceDetail.tabs.baseInfo'),
                name: 'basicInfo'
            },
            {
                label: language.t('deviceDetail.tabs.diskDetail'),
                name: 'diskDetail'
            },
            {
                label: language.t('deviceDetail.tabs.hardwareMonitoring'),
                name: 'hardwareMonitoring'
            },
            {
                label: language.t('deviceDetail.tabs.operatLog'),
                name: 'operatLog'
            }
        ],
        tabsPangeData: [
            {
                label: language.t('deviceDetail.tabs.baseInfo'),
                name: 'basicInfo'
            },
            {
                label:language.t('deviceDetail.tabs.diskDetail'),
                name: 'diskDetail'
            },
            {
                label: language.t('deviceDetail.tabs.hardwareMonitoring'),
                name: 'hardwareMonitoring'
            },
            {
                label: language.t('deviceDetail.tabs.performanceMonitoring'),
                name: 'performanceMonitoring'
            },
            {
                label: language.t('deviceDetail.tabs.operatLog'),
                name: 'operatLog'
            }
        ]
    };

    onMounted(() => {
        data.activeName.value = query.type;
        if (query?.type === 'basicInfo') {
            deviceDetail.initData();
            deviceDetail.desrcibeAgentStatus();
            return;
        }
    })

    watch(() => data.activeName.value, (newValue: string) => {
        const requestData = [
            [
                (newValue: string) => !newValue.localeCompare('basicInfo'),
                () => {
                    deviceDetail.initData();
                    deviceDetail.desrcibeAgentStatus();
                }
            ],
            [
                (newValue: string) => !newValue.localeCompare('diskDetail'),
                () => {
                    deviceDetail.initData();
                    diskTableOpt.init();
                }
            ]
        ];
        for(const key of requestData) {
            if (key[0](newValue)) {
                key[1](newValue);
                break;
            }
        }
    });

    const setData = (status: string, monitorStatus: string) => {
        if (Object.is(status, 'error')) {
            if (Object.is(monitorStatus, 'error')) return data.tabsPangeOneData;
            if (!Object.is(monitorStatus, 'error')) return data.tabsPangeTwoData;
        }
        else {
            if (Object.is(monitorStatus, 'error')) return data.tabsPangeThreeData;
            return data.tabsPangeData;
        }
    }

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
            obj = {...obj, [index]: query[index], type: data.activeName.value}
            obj.type = data.activeName.value;
        }
        for (let key in obj) {
            activeQuery += `${[key]}=${obj[key]}&`
        }
        activeQuery = activeQuery.slice(0, -1);
        router.push(`${path}?${activeQuery}`);
        
    };

    return {
        ...data,
        setData,
        // consolelogs,
        tabsChange
    }
};

export default tabsOperate;
