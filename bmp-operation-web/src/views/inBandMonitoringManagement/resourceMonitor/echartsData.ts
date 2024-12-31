/**
 * @file
 * @author
*/

import RuleFormOpt from './utils/ruleForm';
import publicConfigOpt from './utils/publicConfig';
import InitStaticData from './utils/initStaticData';
import {language} from 'utils/publicClass.ts';
import {deepCopy, intervalTime, getMonthHoursMinutes} from 'utils/index.ts';
import {publicOption, seriesDataObj, echartsOpt} from './options';
import {monitorDataAPI} from 'api/inBandMonitoring/resourceMonitor/request.ts';
import {NewDataType, NewTimeType, HourRangesType, ReactiveArrType, EchartsDataObjType} from './type';
import * as echarts from 'echarts';

class EchartsOpt extends RuleFormOpt {
    errorCode: Ref<number> = ref<number>(0);
    pageLoading: Ref<boolean> = ref<boolean>(false);
    cycle: Ref<string> = ref<string>('1min');
    readioSearch: Ref<string> = ref<string>(language.t('monitorEcharts.radio.instance'));
    hours: Ref<string> = ref<string>(language.t('monitorEcharts.tabs.h1'));
    value1: Ref<string> = ref<string>('');
    selectValForm = reactive({
        disk: [],
        mountpoint: [],
        nic: []
    });
    diskVal: Ref<string> = ref<string>('');
    mountpointVal: Ref<string> = ref<string>('');
    nicVal: Ref<string> = ref<string>('');
    reactiveArr: ReactiveArrType = reactive<ReactiveArrType>({
        aggregationMethod: publicConfigOpt.aggregationMethod,
        isLoadingObj: {},
        copyData: [],
        isShowObj: {},
        echartsParamsData: InitStaticData.echartsParamsData,
        // @ts-ignore
        echartsData: InitStaticData.echartsData
    });
    str = `
        ${language.t('table.empty')}
    `;
    props;

    constructor(props: any) {
        super();
        const that = this;
        for (let index in this.ruleForm) {
            (this.reactiveArr.isLoadingObj[`${index}` as keyof typeof that.reactiveArr.isLoadingObj] as boolean) = false;
            (this.reactiveArr.isShowObj[`${index}` as keyof typeof that.reactiveArr.isShowObj] as boolean) = false;
        }
        this.props = props;
        this.reactiveArr.copyData = deepCopy(this.reactiveArr.echartsData);
        // this.initData();
        this.initEchartsParamsData();
        this.watchHours();
        this.watchValue1();
        this.watchEchartsShow();
    }

    watchEchartsShow = () => {
        watch(() => this.reactiveArr.isShowObj, (newValue) => {
            const that = this;
            for (let index in newValue) {
                if (!this.reactiveArr.isShowObj[`${index}` as keyof typeof that.reactiveArr.isShowObj]) {
                    // @ts-ignore
                    for (const key of this.reactiveArr.echartsData) {
                        if (Object.is(key.typeVal, index)) key.aggregationMethod = 'avg';
                    }
                }
            }
        }, {deep: true})
    }

    watchValue1 = () => {
        watch(() => this.value1, (newValue) => {
            this.value1Val(newValue.value);
            if (newValue.value?.length) {
                if (this.hours.value?.length) {
                    this.hours.value = ''
                }
                this.refresh();
            }
            
        }, {deep: true})
    }

    value1Val = (value: string) => {
        if (value !== null) {
            // intervalTime 计算两个时间之间的时间差 多少天时分秒
            const val = intervalTime(value[0], value[1]);
            const objData = [
                [
                    () => !val.days,
                    () => {
                        const hourRanges: HourRangesType = {
                            '0': '1min',
                            '6': '5min', 
                            '12': '10min',
                            '24': '1min'
                        };
                        // @ts-ignore
                        const cycleValue = hourRanges[Math.floor(val.hours / 6) * 6] || '';
                        if (cycleValue?.length) {
                            this.cycle.value = cycleValue;
                        }
                    }
                ],
                [
                    () => val.days,
                    () => {
                        this.cycle.value = val.days >= 14 ? '180min' : val.days >= 7 ? '120min' : val.days >= 3 ? '60min' : '20min';
                    }
                ]
            ];
            for (const key of objData) {
                if (key[0]()!) {
                    key[1]();
                    break;
                }
            }
            return;
        }
        this.hours.value = language.t('monitorEcharts.tabs.h1');
        this.cycle.value = '1min';
    }

    isLoading = (type: string) => {
        const that = this;
        return this.reactiveArr.isLoadingObj[`${type}` as keyof typeof that.reactiveArr.isLoadingObj];
    }

    watchHours = () => {
        watch(() => this.hours, (newValue: {value: string}) => {
            const val: {value: string} = InitStaticData.radioData.filter((item: {label: string}) => item.label === newValue.value)[0];
            if (newValue.value.length) {
                if (this.value1.value?.length) {
                    this.value1.value = ''
                }
                this.cycle.value = val.value;
                this.refresh();
            }
            else {
                if (!Object.is(val, void 0)) {
                    this.cycle.value = val.value;
                    this.refresh();
                }
            }
        }, {
            deep: true
        })
    }

    searchInfo = (ruleForm: any, _: boolean = true) => {
        const data = this.reactiveArr.echartsData;
        const objData = ruleForm.monitoringIndicators;
        const isShowObj = this.reactiveArr.isShowObj;
        ruleForm.keyData = [];
        if (Array.isArray(data)) {
            for (const key of data) {
                for (const item of objData) {
                    if (Array.isArray(key.metricName) && item.includes(key.metricName.join(','))) ruleForm.keyData.push(key);
                    else if (Object.is(item, key.metricName)) ruleForm.keyData.push(key);
                }
            }
        }
        for (const listItem of ruleForm.keyData) {
            for (let index in isShowObj) {
                const status: boolean = Object.is(index, listItem.typeVal);
                (this.reactiveArr.isShowObj[`${index}` as keyof typeof isShowObj] as boolean) = status;
            }
        }
        for (const key of ruleForm.keyData) {
            this.getEcharts(key)
        }
    }

    clearSearchInfo (ruleFormOperate: {ruleForm: {keyData: []}}, ruleForm: {monitoringIndicators: []}) {
        const isShowObj = this.reactiveArr.isShowObj;
        // this.pageLoading.value = false;
        ruleFormOperate.ruleForm.keyData = [];
        ruleForm.monitoringIndicators = [];
        this.reactiveArr.echartsData = this.reactiveArr.copyData;
        for (let index in this.reactiveArr.isShowObj) {
            (this.reactiveArr.isShowObj[`${index}` as keyof typeof isShowObj] as boolean) = false;
        }
    }

    initEchartsParamsData = () => {
        const that = this;
        for (const key of this.reactiveArr.echartsParamsData) {
            key.model = this.ruleForm[`${key.type}` as keyof typeof that.ruleForm];
            key.bigClickVal.model = this.ruleForm[`${key.type}` as keyof typeof that.ruleForm];
        }
    }

    initData = () => {
        if (Array.isArray(this.reactiveArr.echartsData)) {
            for (const key of this.props.ruleForm.keyData) {
                // 循环体
                this.getEcharts(key);
            }
        }
        else {
            throw new Error('数据格式出错，不是一个数组')
        }
    }

    refresh = () => {
        const value = this.readioSearch.value;
        const objData = [
            [
                () => Object.is(value, language.t('monitorEcharts.radio.instance')),
                () => this.initData()
            ]
        ];
        for (const key of objData) {
            if (key[0]()!) {
                return key[1]();
            }
        }
    }

    getTagName = (value: string) => {
        const tagName = new Map([
            [language.t('monitorEcharts.radio.disk'), 'disk'],
            [language.t('monitorEcharts.radio.diskPartition'), 'mountpoint'],
            [language.t('monitorEcharts.radio.netWorkInterface'), 'nic']
        ]);
        return tagName.get(value);
    }

    disabledDate = (time: any) => {
        return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 30;
    }

    showStatus = () => {
        const status = this.reactiveArr.echartsData.some((item) => item.data.length !== '');
        return status;
    }

    echartShowStatus = (type: string) => {
        const that = this;
        // && !this.defaultShowLoading.value;
        return this.reactiveArr.isShowObj[`${type}` as keyof typeof that.reactiveArr.isShowObj] && !this.pageLoading.value;
    }

    getEcharts = async <T extends EchartsDataObjType>(key: T) => {
        const {idcId, instanceId, userName} = this.props.ruleForm;
        this.setLoading(key.typeVal, true)
        this.setShow(key.typeVal, true);
        // this.pageLoading.value = true;
        const param = {
            metricName: typeof key.metricName === 'string' ? key.metricName : key.metricName.join(','),
            instanceId,
            idcId,
            userName,
            // 周期
            timeInterval: Number(this.cycle.value.split('min')[0]) * 60,
        }
        const otherParams = !this.value1.value ? {lastManyTime: Number(this.setTabsParams())} : this.setDateParams();
        try {
            const allParams = {
                ...param,
                ...otherParams,
                downSampleType: key.aggregationMethod
            }
            const res = await monitorDataAPI(allParams);
            if (res?.data?.result?.length) {
                if (!['main6', 'main7', 'main8', 'main9'].includes(key.id)) {
                    const [data] = res.data.result;
                    key.data = data.data.map((item: {value: string}) => parseFloat(Number(item.value).toFixed(2)));
                    key.timeData = data.data.map((item: {timestamp: number}) => getMonthHoursMinutes(item.timestamp));
                    const initTimeData = data.data.map((item: {timestamp: number}) => (item.timestamp));
                    // @ts-ignore
                    this.setData(key, initTimeData);
                }
                else {
                    const newData: NewDataType[] = Array.from(res.data.result, (item: {data: {map: NewDataType['map']}}) => {
                        return item.data.map((ite: {value: string}) => {
                            const value = parseFloat(ite.value);
                            return isNaN(value) ? NaN : value;
                        });
                    });
                    const newTimes: NewTimeType[] = Array.from(res.data.result, (item: {data: {map: NewTimeType['map']}}) => {
                        return item.data.map((ite: {timestamp: number}) => getMonthHoursMinutes(ite.timestamp));
                    });
                    const newArrData = newData.map((item) => {
                        return item.map((t) => parseFloat(Number(t).toFixed(2)));
                    });
                    const initTimes: NewTimeType[] = Array.from(res.data.result, (item: {data: {map: NewTimeType['map']}}) => {
                        return item.data.map((ite: {timestamp: number}) => (ite.timestamp));
                    });
                    key.data = newArrData;
                    key.timeData = [newTimes[0]];
                    // @ts-ignore
                    this.setData(key, initTimes[0]);
                }
                return;
            }
            throw new Error('');
        }
        catch {
            key.data = [];
            key.timeData = [];
            this.reactiveArr.echartsData = this.reactiveArr.copyData;
            // this.props.ruleForm.keyData = [];
            // @ts-ignore
            this.setData(key, []);
            // this.setShow(key.typeVal, false);
        }
        finally {
            this.pageLoading.value = false;
            // this.defaultShowLoading.value = false;
            this.setLoading(key.typeVal, false);
        }
    }

    setLoading = (type: string, bol: boolean) => {
        const that = this;
        for (let index in this.reactiveArr.isLoadingObj) {
            if (index === type) {
                (this.reactiveArr.isLoadingObj[`${index}` as keyof typeof that.reactiveArr.isLoadingObj] as boolean) = bol;
            }
        }
    }

    setShow = (type: string, bol: boolean) => {
        const that = this;
        for (let index in this.reactiveArr.isShowObj) {
            if (index === type) {
                (this.reactiveArr.isShowObj[`${index}` as keyof typeof that.reactiveArr.isShowObj] as boolean) = bol;
            }
        }
    }

    setDateParams = () => {
        return {
            startTime: new Date(this.value1.value[0]).getTime() / 1000,
            endTime: new Date(this.value1.value[1]).getTime() / 1000
        };
    }

    setTabsParams = () => {
        const obj = new Map([
            [language.t('monitorEcharts.tabs.h1'), '1'],
            [language.t('monitorEcharts.tabs.h6'), '6'],
            [language.t('monitorEcharts.tabs.h12'), '12'],
            [language.t('monitorEcharts.tabs.d1'), '24'],
            [language.t('monitorEcharts.tabs.d3'), '72'],
            [language.t('monitorEcharts.tabs.d7'), '168'],
            [language.t('monitorEcharts.tabs.d14'), '336']
        ])
        return obj.get(this.hours.value);
    }

    setData = <T extends string, K extends keyof T>(key: {id: T, data: number[], type: K, timeData: number[], unit: K, hoverUnit: T}, data: number[] | NewTimeType) => {
        nextTick(() => {
            if (document.getElementById(key.id)) {
                this.reqEcharts(key.id, key.data, key.type, key.timeData, key.unit, key.hoverUnit, data)
            }
        })
    }

    reqEcharts = (id: string, ...args: any) => {
        const newData = !Array.isArray(args[1]) ? [
            seriesDataObj(args[1], args[0], publicConfigOpt.singleColor)
        ] : [
            ...this.setEcharts(args[1], args[0])
        ];
        const myChart = echarts?.getInstanceByDom(document.getElementById(id) as HTMLElement);
        publicOption([args[1]], args[2], args[3], newData, args[0], args[4], id, args[5]).then((res: any) => {
            echartsOpt(res, id, myChart);
        })
        .catch((err) => {
            throw new Error(err)
        })
    }

    setEcharts = (data: string[], args: number[][]) => {
        let newData = [];
        for (let i = 0; i < data.length; i++) {
            newData.push(seriesDataObj(data[i], args[i], publicConfigOpt.moreColor[i]));
        }
        return newData;
    }

    selectChange = (val: string, value: string) => {
        const that = this;
        this.ruleForm[`${val}` as keyof typeof that.ruleForm] = value;
        // @ts-ignore
        publicConfigOpt.smallBigSet(that, val, 'small', value);
    }
}

export default EchartsOpt;