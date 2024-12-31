/**
 * @file
 * @author
*/
import i18n from 'lib/i18n/index.ts'; // 国际化

import {
    Ref,
    ref,
    reactive,
    watch,
    nextTick
} from 'vue';
import echarts from './utils/publicQuote';
import RuleFormOpt from './utils/ruleForm';
import publicConfigOpt from './utils/publicConfig';
import InitStaticData from './utils/initStaticData';
import {intervalTime, getMonthHoursMinutes} from 'utils/index.ts';
import {publicOption, seriesDataObj, echartsOpt} from './options';
import {monitorDataAPI, deviceTagAPI} from 'api/request.ts';
import {NewDataType, NewTimeType, HourRangesType, ReactiveArrType, EchartsDataObjType} from './type';
const {global} = i18n;

class EchartsOpt extends RuleFormOpt {
    cycle: Ref<string> = ref<string>('1min');
    readioSearch: Ref<string> = ref<string>(global.t('monitorEcharts.radio.instance'));
    hours: Ref<string> = ref<string>(global.t('monitorEcharts.tabs.h1'));
    value1: Ref<string> = ref<string>('');
    selectValForm = reactive({
        disk: [],
        mountpoint: [],
        nic: []
    })
    diskVal: Ref<string> = ref<string>('');
    mountpointVal: Ref<string> = ref<string>('');
    nicVal: Ref<string> = ref<string>('');
    reactiveArr: ReactiveArrType = reactive<ReactiveArrType>({
        aggregationMethod: publicConfigOpt.aggregationMethod,
        isLoadingObj: {},
        echartsParamsData: InitStaticData.echartsParamsData,
        // @ts-ignore
        echartsData: InitStaticData.echartsData
    });
    instanceId
    fn;

    constructor( instanceId: string, fn: Function ) {
        super();
        const that = this;
        for (let index in this.ruleForm) {
            (this.reactiveArr.isLoadingObj[`${index}` as keyof typeof that.reactiveArr.isLoadingObj] as boolean) = true;
        }
        this.instanceId = instanceId;
        this.fn = fn;
        this.initEchartsParamsData();
        this.watchDetail();
        this.watchRadio();
        this.watchHours();
        this.watchValue1();
    }

    watchValue1 = () => {
        watch(() => this.value1, (newValue) => {
            this.value1Val(newValue.value);
            if (newValue.value?.length) {
                this.hours.value = '';
                const {diskVal, mountpointVal, nicVal} = this;
                let obj: Record<string, any> = {};
                obj['diskVal'] = diskVal.value;
                obj['mountpointVal'] = mountpointVal.value;
                obj['nicVal'] = nicVal.value;
                let str: string = '';
                let newValS: string = '';
                if (Object.is(this.readioSearch.value, global.t('monitorEcharts.radio.instance'))) {
                    str = '';
                }
                else {
                    for (const index in obj) {
                        if (obj[index] !== '') {
                            str = index;
                            newValS = obj[index];
                        }
                    }
                }
                this.strVal(str, newValS);
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
            outer: for (const key of objData) {
                // @ts-ignore
                if (key[0]()) {
                    key[1]();
                    break outer;
                }
            }
            return;
        }
        this.hours.value = global.t('monitorEcharts.tabs.h1');
        this.cycle.value = '1min';
    }

    isLoading = (type: string) => {
        const that = this;
        return this.reactiveArr.isLoadingObj[`${type}` as keyof typeof that.reactiveArr.isLoadingObj];
    }

    strVal = (str: string, newValS: string) => {
        if (!str?.length) {
            this.initData();
            return;
        }
        const newData = this.reactiveArr.echartsData.filter((item) => item.status?.includes(str));
        for (const key of newData) {
            this.getEcharts(key, 'small', {device: newValS});
        }
    }

    watchHours = () => {
        watch(() => this.hours, (newValue: {value: string}) => {
            if (newValue.value.length) {
                if (this.value1.value?.length) {
                    this.value1.value = ''
                }
            }
            const val: {value: string} = InitStaticData.radioData.filter((item: {label: string}) => item.label === newValue.value)[0];
            if (!Object.is(val, void 0)) {
                this.cycle.value = val.value;
                this.refresh();
            }
        }, {
            deep: true
        })
    }

    initEchartsParamsData = () => {
        const that = this;
        for (const key of this.reactiveArr.echartsParamsData) {
            key.model = this.ruleForm[`${key.type}` as keyof typeof that.ruleForm];
            key.bigClickVal.model = this.ruleForm[`${key.type}` as keyof typeof that.ruleForm];
        }
    }

    watchDetail = () => {
        this.initData();
        // watch(() => this.props.detail.instanceId, () => {
        //     this.initData();
        // }, {
        //     deep: true
        // })
    }

    watchRadio = () => {
        watch(() => this.readioSearch, (newValue: {value: string}) => {
            const objData = [
                [
                    (value: string) => Object.is(value, global.t('monitorEcharts.radio.instance')),
                    () => {
                        this.diskVal.value = '';
                        this.nicVal.value = '';
                        this.mountpointVal.value = '';
                        this.initData();
                    }
                ],
                [
                    (value: string) => Object.is(value, global.t('monitorEcharts.radio.disk')),
                    (value: string) => {
                        this.nicVal.value = '';
                        this.mountpointVal.value = '';
                        this.initTagName(this.getTagName(value)!, 'diskVal')
                    }
                ],
                [
                    (value: string) => Object.is(value, global.t('monitorEcharts.radio.diskPartition')),
                    (value: string) => {
                        this.diskVal.value = '';
                        this.nicVal.value = '';
                        this.initTagName(this.getTagName(value)!, 'mountpointVal')
                    }
                ],
                [
                    (value: string) => Object.is(value, global.t('monitorEcharts.radio.netWorkInterface')),
                    (value: string) => {
                        this.diskVal.value = '';
                        this.mountpointVal.value = '';
                        this.initTagName(this.getTagName(value)!, 'nicVal')
                    }
                ]
            ]
            outer: for (const key of objData) {
                // @ts-ignore
                if (key[0](newValue.value)) {
                    key[1](newValue.value)
                    break outer;
                }
            }
        }, {deep: true})
    }

    initData = () => {
        if (Array.isArray(this.reactiveArr.echartsData)) {
            for (const key of this.reactiveArr.echartsData) {
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
                () => Object.is(value, global.t('monitorEcharts.radio.instance')),
                () => this.initData()
            ],
            [
                () => Object.is(value, global.t('monitorEcharts.radio.disk')),
                () => this.disk(this.diskVal.value)
            ],
            [
                () => Object.is(value, global.t('monitorEcharts.radio.diskPartition')),
                () => this.mountpoint(this.mountpointVal.value)
            ],
            [
                () => Object.is(value, global.t('monitorEcharts.radio.netWorkInterface')),
                () => this.nic(this.nicVal.value)
            ]
        ];
        for (const key of objData) {
            // @ts-ignore
            if (key[0]()) {
                return key[1]();
            }
        }
    }

    getTagName = (value: string) => {
        const tagName = new Map([
            [global.t('monitorEcharts.radio.disk'), 'disk'],
            [global.t('monitorEcharts.radio.diskPartition'), 'mountpoint'],
            [global.t('monitorEcharts.radio.netWorkInterface'), 'nic']
        ]);
        return tagName.get(value);
    }

    disabledDate = (time: any) => {
        return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 30;
    }

    disk = (val: string) => {
        this.selectData('diskVal', val);
    }

    mountpoint = (val: string) => {
        this.selectData('mountpointVal', val);
    }

    nic = (val: string) => {
        this.selectData('nicVal', val);
    }

    selectData = (type: string, val: string) => {
        const newData = this.reactiveArr.echartsData.filter((item) => item.status?.includes(type));
        for (const key of newData) {
            this.getEcharts(key, 'small', {device: val});
        }
    }

    initTagName = async <T>(tagName: string, value: T) => {
        const that = this;
        const params = {
            instanceId: this.instanceId,
            tagName
        };
        try {
            const res = await deviceTagAPI(params);
            if (res?.data?.result?.length) {
                this.selectValForm[`${tagName}` as keyof typeof that.selectValForm] = res.data.result;
                // @ts-ignore
                this[value].value = res.data.result[0];
                // @ts-ignore
                const val = this[value].value;
                this.mountpointVal.value?.length && this.mountpoint(val);
                this.diskVal.value?.length && this.disk(val);
                this.nicVal.value?.length && this.nic(val);
            }
        }
        catch {
            throw new Error('error');
        }
    }

    getEcharts = async <T extends EchartsDataObjType>(key: T, size = 'small', selectParams: any = {}) => {
        Object.is(size, 'small') ? this.setLoading(key.typeVal, true) : this.fn(true);
        const param = {
            metricName: this.getMetricName(key),
            instanceId: this.instanceId,
            ...selectParams,
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
                    const initTimeData = data.data.map((item: {timestamp: number}) => (item.timestamp));
                    key.data = data.data.map((item: {value: string}) => parseFloat(Number(item.value).toFixed(2)));
                    key.timeData = data.data.map((item: {timestamp: number}) => getMonthHoursMinutes(item.timestamp));
                    this.setData(key, size, initTimeData);
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
                        return item.map((t) => parseFloat(Number(t).toFixed(2)))
                    });
                    const initTimes: NewTimeType[] = Array.from(res.data.result, (item: {data: {map: NewTimeType['map']}}) => {
                        return item.data.map((ite: {timestamp: number}) => (ite.timestamp));
                    });
                    key.data = newArrData;
                    key.timeData = [newTimes[0]];
                    this.setData(key, size, initTimes[0]);
                }
                return;
            }
            throw new Error('');
        }
        catch {
            key.data = [];
            key.timeData = [];
            this.setData(key, size, []);
        }
        finally {
            Object.is('bigSize', size) ? this.fn(false) : this.setLoading(key.typeVal, false);
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

    getMetricName = <T extends EchartsDataObjType>(key: T) => {
        if (this.mountpointVal.value?.length) {
            return key.paramsMetricName;
        }
        return typeof key.metricName === 'string' ? key.metricName : key.metricName.join(',')
    }

    setDateParams = () => {
        return {
            // @ts-ignore
            startTime: this.value1.value[0].getTime() / 1000, endTime: this.value1.value[1].getTime() / 1000
        };
    }

    setTabsParams = () => {
        const obj = new Map([
            [global.t('monitorEcharts.tabs.h1'), '1'],
            [global.t('monitorEcharts.tabs.h6'), '6'],
            [global.t('monitorEcharts.tabs.h12'), '12'],
            [global.t('monitorEcharts.tabs.d1'), '24'],
            [global.t('monitorEcharts.tabs.d3'), '72'],
            [global.t('monitorEcharts.tabs.d7'), '168'],
            [global.t('monitorEcharts.tabs.d14'), '336']
        ])
        return obj.get(this.hours.value);
    }

    setData = (key: any, size: string, data: number[] | NewTimeType) => {
        let id: string = '';
        nextTick(() => {
            if (document.getElementById(key.id)) {
                id = Object.is(size, 'small') ? key.id : 'bigSize';
                this.reqEcharts(id, key.data, key.type, key.timeData, key.unit, key.hoverUnit, data)
            }
        })
    }

    reqEcharts = (id: string, ...args: any) => {
        const newData = !Array.isArray(args[1]) ? [
            seriesDataObj(args[1], args[0], publicConfigOpt.singleColor)
        ] : [
            ...this.setEcharts(args[1], args[0])
        ];
        const myChart: echarts.ECharts | undefined = echarts?.getInstanceByDom(document.getElementById(id) as HTMLElement);
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