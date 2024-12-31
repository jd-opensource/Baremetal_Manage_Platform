/**
 * @file
 * @author
*/

// uiLocale, locationItem, 
import {CurrentInstance, language} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import {RuleFormRefType} from '@utils/publicType';
// import {ElForm} from 'element-plus';
import {describeVolumesAPI, describeAssociateDisksAPI} from 'api/device/request.ts';
import {modelDetailAPI} from 'api/model/request.ts';
import {AxiosError} from 'axios';
import {msgTip} from 'utils/index.ts';
import {CurrencyType} from '@utils/publicType';

class UpFormOpt {
    instanceProxy = new CurrentInstance().proxy;
    ruleForm = reactive({
        model: ''
    });
    arrayCard: Ref<string> = ref<string>('');
    isTableLoading: Ref<boolean> = ref<boolean>(false);
    formRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType | undefined>()
    reactiveArr: {
        selectObj: any;
        modelData: {text: string; filterParams: string; value: number}[];
        tableData: CurrencyType[];
        disksData: {name: string; size: string; sizeUnit: string; diskId: string; select: boolean;}[][];
    } = reactive<{
        selectObj: any;
        modelData: {text: string; filterParams: string; value: number}[];
        tableData: CurrencyType[];
        disksData: {name: string; size: string; sizeUnit: string; diskId: string; select: boolean;}[][];
    }>({
        selectObj: {},
        modelData: [],
        tableData: [],
        disksData: [],
    });
    rules = reactive({
        model: [
            {
                required: true,
                trigger: 'change',
                message: language.t('associatedModel.content.placeholder')
            }
        ]
    });
    deviceTypeId: Ref<string> = ref<string>('');
    deviceId: Ref<string> = ref<string>('');

    constructor(props: {deviceId: string}) {
        this.deviceId.value = props.deviceId;
        store.ossDataInfo().getModelListRes().then((data: {text: string; filterParams: string; value: number}[]) => {
            this.reactiveArr.modelData = data;
        })
    }

    modelChange = (val: string) => {
        this.isTableLoading.value = true;
        this.deviceTypeId.value = val;
        this.getModelDetail(val);
    };

    getModelDetail = async (val: string) => {
        try {
            const res = await modelDetailAPI({
                id: val
            });
            if (res?.data?.result && Object.keys(res.data.result)?.length) {
                this.arrayCard.value = res.data.result.isNeedRaid;
                this.getDescribeVolumes(val);
            }
        }
        catch(e) {
            const err = e as AxiosError;
            err?.message && msgTip.error(err.message);
        }
    }

    setRaid = (arr: any) => {
        return arr?.map((item: {name: string}) => item.name).join(',')
    }

    getDescribeVolumes = (val: string) => {
        this.reactiveArr.selectObj = {};
        describeVolumesAPI({
            deviceTypeId: val
        })
        .then(<T extends CurrencyType>({data}: {data: {result: T[]}}) => {
            if (data?.result?.length) {
                data.result.forEach((t) => {
                    this.reactiveArr.selectObj[t.volumeId] = [];
                })
                const newData = data.result.map((item) => {
                    return {
                        ...item,
                        disksDataFlag: false
                    }
                });
                this.reactiveArr.tableData = newData;
                const ids = data.result.map(item => item.volumeId);
                this.describeAssociateDisks(ids, val);
                return;
            }
            return Promise.reject();
        })
        .catch(() => {
            this.reactiveArr.tableData = [];
        })
        .finally(() => {
            this.isTableLoading.value = false;
        })
    }

    describeAssociateDisks = (ids: string[], val: string) => {
        const newReq = ids.map((item: string) => {
            return describeAssociateDisksAPI({
                deviceId: this.deviceId.value,
                volumeId: item,
                deviceTypeId: val
            })
        });
        Promise.all(newReq).then((res: any[]) => {
            if (res?.length) {
                const data = res.map((t: any) => t.data).map((g: any) => g.result);
                this.reactiveArr.disksData = data;
            }
        })
        .catch(() => {
            this.reactiveArr.disksData = [];
        })
    }

    setDiskLabel = (item: {name: string; size: string; sizeUnit: string;}) => {
        return `${item.name}: ${item.size}${item.sizeUnit}`;
    }

    jumpPage = () => {
        const path: string = this.instanceProxy!.$defInfo.routerPath('modelList') as unknown as string;
        window.open(path, '_blank')
    }

    getFormRef = (forEl: {value: RuleFormRefType}) => {
        this.formRef.value = forEl.value;
    }

    setRaidCan = (item: string) => {
        return item === 'RAID0-stripping' ? language.t('modelForm.raid') : item;
    }

    getAllSelectDisks = (item: string) => {
        return Object.keys(this.reactiveArr.selectObj).reduce((pre: string[], cur: string) => {
            const c = item == cur ? [] : this.reactiveArr.selectObj[cur];
            return pre.concat(c);
        }, [])
    }

    // 下拉勾选
    disksDataChange = (item: {volumeId: string; disksDataFlag: boolean; select: boolean}) => {
        item.disksDataFlag = this.reactiveArr.selectObj[item.volumeId]?.length ? false : true;
    }

    setDiskInterfaceType = (item: string) => {
        return item === 'notLimited' ? language.t('modelForm.unrestricted') : item;
    }

};

export default UpFormOpt;
