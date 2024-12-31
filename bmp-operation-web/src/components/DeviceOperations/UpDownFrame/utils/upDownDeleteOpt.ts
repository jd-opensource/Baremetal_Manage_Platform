import {msgTip} from 'utils/index.ts';
import {SuccessType} from '@utils/publicType';
import {PropsType, ParamsType, TableDataType} from './type';
import {language, CurrentInstance} from 'utils/publicClass.ts';
import {devicesUpAPI, devicesDownAPI, devicesDeleteAPI, associateDisksAPI, associateAPI} from 'api/device/request.ts';
import {AxiosError} from 'axios';

let s: HTMLElement | null;

class UpDownDel {
    nextLoading: Ref<boolean> = ref<boolean>(false);
    errorFlag: Ref<boolean> = ref<boolean>(false);
    instanceMitt = new CurrentInstance().instanceMitt;
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    // 表格数据
    upDownFrameTable: TableDataType[] = reactive<TableDataType[]>([]);
    props: PropsType = {title: '', deviceId: ''};
    emitValue: any;
    active: Ref<number> = ref<number>(1);
    setUpTitle: Ref<string> = ref<string>(language.t('upDownFrame.btn.next'));
    setTitle: Ref<string> = ref<string>(language.t('upDownFrame.btn.cancel'));
    upFormOpt;

    constructor(props: PropsType, emitValue: any, upFormOpt: any) {
        this.props = props;
        this.emitValue = emitValue;
        this.upFormOpt = upFormOpt;
        // 监听mitt-device-table，获取表格数据
        this.instanceMitt?.proxy?.$Bus?.on('device-table', this.getTableFn);
        watch((this.active), (newValue) => {
            s = document.querySelector('.up-down-frame .default-step .el-step__icon.is-text')
            if (newValue > 1) {
                nextTick(() => s!.innerHTML = '2');
                return;
            }
            s!.innerHTML = '';
        })
    };

    preClick = () => {
        if (this.setTitle.value === language.t('upDownFrame.btn.cancel')) {
            this.cancelClick();
        }
        else {
            this.active.value = 1;
            this.setUpTitle.value = language.t('upDownFrame.btn.next');
            this.setTitle.value = language.t('upDownFrame.btn.sure');
        }
    }

    nextClick = () => {
        if (this.active.value === 1) {
            this.errorFlag.value = true;
            this.upFormOpt.formRef.value.validate((valid: boolean) => {
                if (valid) {
                    this.associateDisks();

                }
            })
        }
        else {
            // 调用
           this.determineClick()
        }
    }

    associateDisks = async () => {
        const status = this.upFormOpt.reactiveArr.tableData.some((item: {disks: string[]}) => !item.disks.length);
        if (status) {
            for (const key of this.upFormOpt.reactiveArr.tableData) {
                key.disksDataFlag = !key.disks.length ? true : false;
            }
            return;
        }
        this.nextLoading.value = true;
        const valData = this.upFormOpt.reactiveArr.tableData.map((item: {volumeId: string; disks: string[]}) => {
            return {
                volumeId: item.volumeId,
                diskId: typeof item.disks === 'string' ? [item.disks] : item.disks
            }
        });
        try {
            const res = await associateDisksAPI({
                deviceTypeId: this.upFormOpt.deviceTypeId.value,
                deviceId: this.props.deviceId,
                volumes: valData
            });
            if (res?.data?.result?.success) {
                this.associate();
                return;
            }
            this.nextLoading.value = false;
        }
        catch (e) {
            this.nextLoading.value = false;
            const err = e as AxiosError;
            err.message && msgTip.error(err.message);
        }
        // const newReq = this.upFormOpt.reactiveArr.tableData.map((item: {volumeId: string; disks: string[]}) => {
        //     return associateDisksAPI({
        //         deviceTypeId: this.upFormOpt.deviceTypeId.value,
        //         deviceId: this.props.deviceId,
        //         volumes: [
        //             {
        //                 volumeId: item.volumeId,
        //                 diskId: item.disks
        //             }
        //         ]
        //     })
        // });
        // Promise.all([newReq])
        // .then((res) => {
        //     if (res?.length) {
        //         this.associate();
        //     }
        // })
        // .catch(() => {
        //     this.nextLoading.value = false;
        // })
    }

    associate = () => {
        associateAPI({
            deviceTypeId: this.upFormOpt.deviceTypeId.value,
            deviceId: this.props.deviceId
        })
        .then(({data}: any) => {
            if (data?.result?.success) {
                this.active.value = 2;
                this.setUpTitle.value = language.t('upDownFrame.btn.sure');
                this.setTitle.value = language.t('upDownFrame.btn.pre');
                this.errorFlag.value = false;
            }
        })
        .finally(() => {
            this.nextLoading.value = false;
        })
    }

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestUpDownFrame();
    };

    /**
     * 请求上架下架删除接口，成功后把事件回传，关闭弹窗
    */
    requestUpDownFrame = () => {
        const requestUrl = new Map([
            ['up', devicesUpAPI],
            ['down', devicesDownAPI],
            ['delete', devicesDeleteAPI]
        ]);
        return this.#getDevicesOperate(requestUrl.get(this.props.title), this.props.title)!;
    };

    /**
     * 设备操作-上架、下架、删除
     * @param {Function} requestUrl url请求地址
     * @param {string} type 设备操作状态-上架、下架、删除
    */
    #getDevicesOperate = (requestUrl: (arg0: ParamsType) => Promise<{data: SuccessType}>, type: string) => {
        const params: ParamsType = this.#setParams(type);
        requestUrl(
            {
                ...params
            }
        )
        .then(({data}: {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            this.cancelClick();
            this.emitValue('determine-click', type);
        })
    };

    #setParams = (type: string) => {
        const param = {deviceId: this.props.deviceId};
        const newParams = new Map([
            ['delete', param]
        ]);
        return newParams.get(type)?? {...param, sns: this.upDownFrameTable.map((item) => item.sn).join(',')};
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
        // 取消device-table mitt事件，第一个参数-事件名，第二个参数，需要取消的函数
        this.instanceMitt.proxy.$Bus.off('device-table', this.getTableFn);
    };

    /**
     * 获取表格数据函数
     * @param {Object} list 表格数据
    */
    getTableFn: unknown = ({list}: {list: TableDataType}) => {
        this.upDownFrameTable.push(list);
    };
};

export default UpDownDel;
