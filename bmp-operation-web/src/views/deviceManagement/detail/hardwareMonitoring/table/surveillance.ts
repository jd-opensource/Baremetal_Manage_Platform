import {deviceStatusAPI} from 'api/device/request.ts';
import {CurrencyType} from '@utils/publicType';
import {language, CurrentInstance} from 'utils/publicClass.ts';
import {methodsTotal} from 'utils/index.ts';
import router from 'router/index.ts';
import AllStaticData from 'staticData/staticData.ts';

class SurveillanceOpt {
    isLoading: Ref<boolean> = ref<boolean>(true);
    titlePrefix: string = 'deviceDetail.hardwareMonitoring.status.title';
    normal: string = language.t('deviceDetail.hardwareMonitoring.status.normal');
    error: string = language.t('deviceDetail.hardwareMonitoring.status.error');
    total: Ref<number> = ref<number>(0);
    reactiveArr: any = reactive({
        tableData: [
            {
                label: language.t(`${this.titlePrefix}.cpu`),
                status: '--'
            },
            {
                label: language.t(`${this.titlePrefix}.storage`),
                status: '--'
            },
            {
                label: language.t(`${this.titlePrefix}.hardDisk`),
                status: '--'
            },
            {
                label: language.t(`${this.titlePrefix}.networkCard`),
                status: '--'
            },
            {
                label: language.t(`${this.titlePrefix}.powerSupply`),
                status: '--'
            },
            {
                label: language.t(`${this.titlePrefix}.other`),
                status: '--'
            }
        ]
    });
    query: {sn: string;};
    instanceProxy = new CurrentInstance().proxy;
    faultLogTable: {refresh(): void};
    mitt = new CurrentInstance();
    deviceDetail;

    constructor (props: {query: {sn: string}, deviceDetail: {initData(): void}}, faultLogTable: {refresh(): void;}) {
        this.query = props.query;
        this.deviceDetail = props.deviceDetail;
        this.faultLogTable = faultLogTable;
        this.deviceDetail.initData();
        this.getDeviceStatus();
        this.watchLoading();
    };

    watchLoading = () => {
        watchEffect(() => {
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('device-detail-hardwareMonitoring-loading', this.isLoading)
        })
    }

    getDeviceStatus = () => {
        this.isLoading.value = true;
        deviceStatusAPI({
            sn: this.query.sn
        }).then(({data}: any) => {
            if (data?.result?.detail?.length) {
                const result = methodsTotal.lineConverting(data.result.detail[0]);
                const newData = this.reactiveArr.tableData.map((item: {label: string; status: string;}) => {
                    return {
                        ...item,
                        status: this.#setLabel(item, result)
                    }
                })
                this.reactiveArr.tableData = newData;
                this.total.value = data.result.total_count;
                return;
            }
            return Promise.reject('');
        })
        .catch(({message}: {message: string}) => {
            const path: string = this.instanceProxy.$defInfo.routerPath('deviceList');
            if (AllStaticData.noFoundData.indexOf(message) > -1) {
                router.push(path);
                return;
            }
            this.total.value = 0;
        })
        .finally(() => {
            this.isLoading.value = false;
            this.faultLogTable.refresh()
        });
    };

    #setLabel = (item: {label: string; status: string;}, keys: CurrencyType) => {
        const key: Map<string, string> = new Map([
            [language.t(`${this.titlePrefix}.cpu`), this.#valueConvert(keys['cpuStatus'])],
            [language.t(`${this.titlePrefix}.storage`), this.#valueConvert(keys['memStatus'])],
            [language.t(`${this.titlePrefix}.hardDisk`), this.#valueConvert(keys['diskStatus'])],
            [language.t(`${this.titlePrefix}.networkCard`), this.#valueConvert(keys['nicStatus'])],
            [language.t(`${this.titlePrefix}.powerSupply`), this.#valueConvert(keys['powerStatus'])],
            [language.t(`${this.titlePrefix}.other`), this.#valueConvert(keys['otherStatus'])]
        ]);
        return key.get(item.label);
    };

    #valueConvert = (item: string) => {
        return !item ? this.normal : this.error;
    };

    setBackgourdColor = (status: string) => {
        const key: Map<string, string> = new Map([
            [this.normal, 'normal-background'],
            [this.error, 'error-background']
        ]);
        return key.get(status);
    };

    setTextColor = (status: string) => {
        const key: Map<string, string> = new Map([
            [this.normal, 'normal-title'],
            [this.error, 'error-title']
        ]);
        return key.get(status);
    };
};

export default SurveillanceOpt;