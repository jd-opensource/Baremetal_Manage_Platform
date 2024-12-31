/**
 * @file
 * @author
*/

// uiLocale, locationItem, 
import {CurrentInstance, language} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import {RuleFormRefType} from '@utils/publicType';
// import {ElForm} from 'element-plus';
import {msgTip, deepCopy} from 'utils/index.ts';
import {describeVolumesAPI, describeAssociateDisksAPI} from 'api/device/request.ts';
import {modelDetailAPI} from 'api/model/request.ts';
import {AxiosError} from 'axios';

class UpFormOpt {
    isTableLoading: Ref<boolean> = ref<boolean>(false);
    instanceProxy = new CurrentInstance().proxy;
    ruleForm = reactive({
        modelName: ''
    })
    arrayCard: Ref<string> = ref<string>('');
    formRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType | undefined>()
    reactiveArr: {
        modelData: {text: string; filterParams: string; value: number;}[];
        tableData: any
        disksData: {name: string; size: string; sizeUnit: string; diskId: string; select: boolean;}[][];
        uniquenessData: string[][];
    } = reactive<{
        modelData: {text: string; filterParams: string; value: number;}[];
        tableData: any
        disksData: {name: string; size: string; sizeUnit: string; diskId: string; select: boolean;}[][];
        uniquenessData: string[][];
    }>({
        modelData: store.ossDataInfo().deviceTypeId,
        tableData: [],
        disksData: [],
        uniquenessData: []
    });
    rules = reactive({
        modelName: [
            {
                required: true,
                trigger: 'change',
                message: '请选择机型'
            }
        ],
    });
    deviceId: Ref<string> = ref<string>('');
    deviceTypeId: Ref<string> = ref<string>('');

    constructor(props: {deviceId: string}) {
        this.deviceId.value = props.deviceId;
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
            msgTip.error(err.message);
        }
    }

    getDescribeVolumes = (val: string) => {
        describeVolumesAPI({
            deviceTypeId: val
        })
        .then(({data}: any) => {
            if (data?.result?.length) {
                this.reactiveArr.tableData = data.result.map((item: {disksDataFlag: boolean}) => {
                    return {
                        ...item,
                        disksDataFlag: false
                    }
                });
                const ids = this.reactiveArr.tableData.map((item: {volumeId: string}) => item.volumeId);
                this.describeAssociateDisks(ids, val);
                return;
            }
            this.reactiveArr.tableData = [];
        })
        .finally(() => {
            this.isTableLoading.value = false;
        })
    }

    setRaid = (arr: any) => {
        return arr?.map((item: {name: string}) => item.name).join(',')
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
            this.reactiveArr.uniquenessData = [];
            if (res?.length) {
                const data = res.map((t: any) => t.data).map((g: any) => g.result);
                this.reactiveArr.disksData = data;
                for (let i = 1; i <= this.reactiveArr.disksData.length; i++) {
                    this.reactiveArr.uniquenessData.push([]);
                }
            }
        })
        .catch(() => {
            this.reactiveArr.disksData = [];
        })
    }

    setDiskLabel = (item: {name: string; size: string; sizeUnit: string;}) => {
        return `${item.name}: ${item.size}${item.sizeUnit}`;
    }

    setRaidCan = (item: string) => {
        return item === 'RAID0-stripping' ? language.t('modelForm.raid') : item;
    }

    jumpPage = () => {
        const path: string = this.instanceProxy!.$defInfo.routerPath('modelList') as unknown as string;
        window.open(path, '_blank')
    }

    getFormRef = (forEl: {value: RuleFormRefType}) => {
        this.formRef.value = forEl.value;
    }

    setDiskInterfaceType = (item: string) => {
        return item === 'notLimited' ? language.t('modelForm.unrestricted') : item;
    }

    // disksDataChange = (item: {disks: string[]; disksDataFlag: boolean}) => {
    //     item.disksDataFlag = item.disks.length ? false : true;
    // }
  
    disksDataChange = (item: {newDisks: string; raidCan: string; disks: string[]; disksDataFlag: boolean; select: boolean}, index: number) => {
        item.disksDataFlag = item.disks.length ? false : true;
        this.reactiveArr.uniquenessData[index] = [];
        if (['RAID0-stripping', 'NO RAID'].includes(item.raidCan)) {
            item.newDisks = deepCopy([item.disks]);
            this.reactiveArr.uniquenessData[index] = [...this.reactiveArr.uniquenessData[index], ...item.newDisks];
        }
        else {
            this.reactiveArr.uniquenessData[index] = [...this.reactiveArr.uniquenessData[index], ...item.disks];
        }
        for (let i = this.reactiveArr.disksData.length; i >= 0; i--) {
            if (this.reactiveArr.disksData[i] !== void 0) {
                this.reactiveArr.disksData[i].forEach((t) => {
                    i --;
                    if (this.reactiveArr?.uniquenessData[i]?.includes(t.diskId)) {
                        t.select = true;
                    }
                    else {
                        t.select = false;
                    }
                    i ++;
                })
            }
        }

        for (let i = 0; i < this.reactiveArr.disksData.length; i++) {
            if (this.reactiveArr.disksData[i] !== void 0) {
                this.reactiveArr.disksData[i].forEach((t) => {
                    i ++;
                    if (this.reactiveArr?.uniquenessData[i]?.includes(t.diskId)) {
                        t.select = true;
                    }
                    i --;
                })
            }
        }
    }
}

export default UpFormOpt;
