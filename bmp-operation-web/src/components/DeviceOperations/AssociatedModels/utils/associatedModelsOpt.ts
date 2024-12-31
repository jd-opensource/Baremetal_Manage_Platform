/**
 * @file
 * @author
*/

// RuleFormRefType, 
// import {SuccessType} from '@utils/publicType';
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {AxiosError} from 'axios';
import {associateDisksAPI} from 'api/device/request.ts';
// associateAPI
class AssociatedModelsOpt {
    errorFlag: Ref<boolean> = ref<boolean>(false);
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    upFormOpt;
    emitValue;
    deviceId: string;

    constructor(upFormOpt: any, emitValue: any, deviceId: string) {
        this.upFormOpt = upFormOpt;
        this.emitValue = emitValue;
        this.deviceId = deviceId;
    }

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.errorFlag.value = true;
        await this.upFormOpt.formRef.value!.validate((valid: boolean) => {
            if (valid) {
                this.associateDisks();
            }
        })
    };

    associateDisks = async () => {
        const res = this.upFormOpt.reactiveArr.tableData.map((item: any) => {
            // if (['RAID0-stripping', 'NO RAID'].includes(item.raidCan)) {
            //     item.newDisks = deepCopy([item.disks]);
            //     // this.reactiveArr.uniquenessData[index] = [...this.reactiveArr.uniquenessData[index], ...item.newDisks];
            // }
            return {
                ...item,
                disks: this.upFormOpt.reactiveArr.selectObj[item.volumeId]
            }
        })
        this.upFormOpt.reactiveArr.tableData = res;
        const status = res.some((item: {disks: string[]}) => !item.disks.length);
        if (status) {
            for (const key of this.upFormOpt.reactiveArr.tableData) {
                key.disksDataFlag = !key.disks.length ? true : false;
            }
            return;
        }
        this.isLoading.value = true;
        const valData = this.upFormOpt.reactiveArr.tableData.map((item: {volumeId: string; disks: string[]}) => {
            return {
                volumeId: item.volumeId,
                diskId: typeof item.disks === 'string' ? [item.disks] : item.disks
            }
        });
        try {
            const res = await associateDisksAPI({
                deviceTypeId: this.upFormOpt.deviceTypeId.value,
                deviceId: this.deviceId,
                volumes: valData
            });
            if (res?.data?.result?.success) {
                msgTip.success(language.t('operate.success'))
                this.emitValue('determine-click')
                // this.associate();
                return;
            }
            this.isLoading.value = false;
        }
        catch (e) {
            const err = e as AxiosError;
            err.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
            this.cancelClick();
        }
    //     const newReq = this.upFormOpt.reactiveArr.tableData.map((item: {volumeId: string; disks: string[]}) => {
    //         return associateDisksAPI({
                // deviceTypeId: this.upFormOpt.deviceTypeId.value,
                // deviceId: this.deviceId,
    //             volumes: [
    //                 {
    //                     volumeId: item.volumeId,
    //                     diskId: item.disks
    //                 }
    //             ]
    //         })
    //     })
    //     Promise.all(newReq).then((res: any) => {
    //         if (res?.length) {
    //             this.associate();
    //         }
    //     })
        // .catch(() => {
        //     this.isLoading.value = false;
        // })
    }

    // associate = () => {
    //     associateAPI({
    //         deviceTypeId: this.upFormOpt.deviceTypeId.value,
    //         deviceId: this.deviceId
    //     })
    //     .then(({data}: {data: SuccessType}) => {
    //         if (data?.result?.success) {
    //             msgTip.success(language.t('operate.success'))
    //             this.emitValue('determine-click')
    //         }
    //     })
    //     .finally(() => {
    //         this.isLoading.value = false;
    //         this.cancelClick();
    //     })
    // }

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
    };
};

export default AssociatedModelsOpt;
