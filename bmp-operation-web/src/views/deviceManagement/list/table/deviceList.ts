/**
 * @file
 * @author
*/

import {AxiosError} from 'axios';
import {TimerOperate, paginationOperate, RouterOperate, CurrentInstance, language} from 'utils/publicClass.ts';
import {CurrencyType} from '@utils/publicType';
import store from 'store/index.ts';
import {devicesListAPI} from 'api/device/request.ts';
import {msgTip, decrypt, methodsTotal} from 'utils/index.ts';
import DeviceStaticData from 'staticData/device/index.ts'; // 中间态1 中间态2
import {TableType, OSSType, ReactiveArrType, FilterParamsType} from '../typeManagement';
import {removeLong} from 'request/RemoveLongReq/remove.ts';

let _self: any = null;
class DeviceList extends TimerOperate {
    routerOperate = new RouterOperate();
    searchTip: Ref<boolean> = ref<boolean>(false);
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(false);
    reactiveArr: {
        tableData: TableType[];
        selectArr: CurrencyType[];
        defTableData: TableType[];
        arr: any;
        currentData: any;
    } = reactive<{
        tableData: TableType[];
        selectArr: CurrencyType[];
        defTableData: TableType[];
        arr: TableType[];
        currentData: TableType[];
    }>({
        tableData: [], // 设备列表数据
        selectArr: [],
        defTableData: [],
        arr: [],
        currentData: []
    });
    monitorCode: Ref<number> = ref<number>(0);
    code: Ref<number> = ref<number>(0);
    tableRef: {
        [x: string]: unknown;
        value?: {
            toggleRowSelection(arg1: any, arg2: any): any;
            clearFilter(): unknown;
            clearSelection(): any;
        }
    } = {value: {clearFilter(): void{}, toggleRowSelection(): any{}, clearSelection(): any{}}};
    // hasRequest: boolean = false;
    filterEmptyInfo: { // store
        filterParams(
            arg0: {[x: string]: number | string},
            arg1: OSSType,
            arg2: ReactiveArrType,
            arg3: FilterParamsType
    ): Promise<void>;
        deleteEmtpyData(arg0: FilterParamsType): unknown;
    } = store.filterEmpty;
    isIntervalRequest: boolean = false;
    type: string = '';
    search: any = {};
    filter: any = {};
    hasCheckAll: Ref<boolean> = ref<boolean>(false);
    closeBtnDisabled: Ref<boolean> = ref<boolean>(true);
    openBtnDisabled: Ref<boolean> = ref<boolean>(true);
    restartBtnDisabled: Ref<boolean> = ref<boolean>(true);
    instanceNameBntDisabled: Ref<boolean> = ref<boolean>(true);
    resetPasswordBntDisabled: Ref<boolean> = ref<boolean>(true);
    recoveryInstanceBtnDisabled: Ref<boolean> = ref<boolean>(true);
    sessionId: Ref<string> = ref<string>('');
    instanceProxy = new CurrentInstance().proxy;

    constructor (search: any, filter: any) {
        super();
        this.search = search;
        this.filter = filter;
        const id: string | null = sessionStorage?.getItem('deviceTypeId');
        const encryptDecrypt: string[] = this.instanceProxy.$defInfo.encryptDecrypt(1);
        this.sessionId.value = id && decrypt(id, encryptDecrypt[0], encryptDecrypt[1]);
        _self = this;

        // 页面销毁时-触发，清空延时器
        onUnmounted(() => {
            clearTimeout((this.timer as number));
            removeLong(this.routerOperate.router.currentRoute.value.name);
            sessionStorage?.removeItem('deviceTypeId');
        });
        // 页面销毁前-触发，清空延时器
        onBeforeUnmount(() => {
            clearTimeout((this.timer as number));
            removeLong(this.routerOperate.router.currentRoute.value.name);
            sessionStorage?.removeItem('deviceTypeId');
        });
        this.getDataList();
        this.getHardStatusList();
    };

    getHardStatusList = async () => {
        try {
            await this.instanceProxy.$hardwareStatusApi.hardwareStatusAPI({pageNumber: 1});
        }
        catch(e) {
            const err = e as AxiosError;
            this.code.value = err.code as unknown as number;
        }
    }

    setCollected = (status: string) => {
        const type = new Map([
            ['1', language.t('deviceList.status.collected')],
            ['2', language.t('deviceList.status.noCollected')],
            ['3', language.t('deviceList.status.collecting')],
            ['4', language.t('deviceList.status.collectionFailed')]
        ]);
        return type.get(status);
    }

    jumpClick = <T extends {router: {push(arg0: string): void}}>(routerOperate: T, item: {sn: string; deviceId: string; instanceId: string}) => {
        const path = this.instanceProxy.$defInfo.routerPath('deviceDetail');
        const defaultPath = `${path}?sn=${item.sn}&deviceId=${item.deviceId}&instanceId=${item.instanceId}&type=basicInfo`;
        routerOperate.router.push(`${defaultPath}&monitorStatus=error`);
    }

    /**
     * 获取data数据
    */
    getDataList = () => {
        this.isLoading.value = true;
        this.processingParameter();
    };

    /**
     * 处理参数
    */
    processingParameter = () => {
        const params = {
            pageSize: paginationOperate.pageSize.value,
            pageNumber: paginationOperate.pageNumber.value,
            ...this.search?.reactiveArr.searchParams,
            ...this.filter?.reactiveArr.filterParams
        };
        this.filterEmptyInfo.deleteEmtpyData(params);
        this.getDeviceList(params);
    };

    getDeviceList = (params: {[x: string]: string}) => {
        clearTimeout((this.timer as number));
        if (sessionStorage?.getItem('deviceTypeId')) {
            params = {
                ...params,
                deviceTypeId: this.sessionId.value
            }
        }
        devicesListAPI(
            {
                ...params,
            }
        )
        .then(({data} : {data: {result: {totalCount: number; devices: TableType[]}}}) => {
            if (data?.result?.devices?.length) {
                const result: TableType[] = data.result.devices;
                const totalCount: number = data.result.totalCount;
                result.forEach((item: TableType, index: number) => {
                    DeviceStaticData.deviceTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                })
                this.reactiveArr.tableData = result;
                paginationOperate.total.value = totalCount;
                for (const key of this.reactiveArr.tableData) {
                    const hasLoopReq = [...DeviceStaticData.intermediate1, ...DeviceStaticData.intermediate2].some((item: string) => item === key.instanceStatus);
                    if (['putawaying'].some((item: string) => item === key.manageStatus) || hasLoopReq) {
                        this.isIntervalRequest = true;
                        this.timer = setTimeout(() => {
                            // if (this.hasRequest) return;
                            this.processingParameter();
                        }, 5000)
                        return;
                    }
                    this.isIntervalRequest = false;
                }
                return;
            }
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            methodsTotal.initScrollLeft(this.filter.tableRef.value);
        })
        .catch(({message} : {message: string;}) => {
            if (!this.isIntervalRequest) {
                this.reactiveArr.tableData = [];
                paginationOperate.total.value = 0;
                methodsTotal.initScrollLeft(this.filter.tableRef.value);
                message && msgTip.error(message);
                return;
            }
            this.timer = setTimeout(() => {
                // if (this.hasRequest) return;
                this.processingParameter();
            }, 5000);
        })
        .finally(() => {
            sessionStorage.removeItem('deviceTypeId')
            this.isLoading.value = false;
            this.search.hasClear.value = false;
            this.searchTip.value = Object.keys(params).length > 2;
        });
    };

    // 多选框回显
    getRowKeys = (row: {sn: string;}) => {
        if (row.sn) return row.sn;
    };

    rowClick = (row: any, _: any, event: any) => {
        this.filter.tableRef?.value?.toggleRowSelection(row, event);
        this.selectChange(this.reactiveArr.selectArr, row);
    };

    /**
     * 单选
    */
    selectChange = (rows: CurrencyType[], row: CurrencyType) => {
        const newRows: boolean | number = rows.filter((item) => !item.hasCheckbox).length &&  rows.indexOf(row) !== -1;
        if (newRows) {
            this.reactiveArr.arr.push(row);
        }
        else {
            for(let i = 0; i < this.reactiveArr.currentData.length; i++) {
                if(this.reactiveArr.currentData[i].sn === row.sn){
                    this.reactiveArr.currentData.splice(i, 1);
                }
            }
            for(let i = 0; i < this.reactiveArr.arr.length; i++){
                if(this.reactiveArr.arr[i].sn === row.sn){
                    this.reactiveArr.arr.splice(i, 1);
                }
            }
        }
    };

    cellClick = (row: CurrencyType, b: {property: string | undefined}, event: any) => {
        const type = ['sn', undefined];
        if (!type.includes(b.property)) return;
        this.toggleSelection([row], event);
    };

    toggleSelection = (rows: CurrencyType[], event: any) => {
        if (rows) {
            rows.forEach((row: any) => {
                this.filter.tableRef.value?.toggleRowSelection(row, event);
            });
            return;
        }
        this.filter.tableRef.value?.clearSelection();
    };

    handleSelectionChange (val: CurrencyType[]) {
        _self.reactiveArr.selectArr = val;
        let newDefTableData: CurrencyType[] = [];
        for (const key of val) {
            newDefTableData.push(
                {
                    instanceName: key.instanceName,
                    instanceId: key.instanceId,
                    iloIp: key.iloIp,
                    privateIpv4: key.privateIpv4
                }
            )
        }
        _self.reactiveArr.defTableData = newDefTableData;
        _self.hasCheckAll.value = _self.reactiveArr.selectArr.length === _self.reactiveArr.tableData.length;
        const runningStatus = val.every((item: CurrencyType) => ['running'].includes(item.instanceStatus));
        const stoppedStatus = val.every((item: CurrencyType) => ['stopped'].includes(item.instanceStatus));
        const recoveryInstanceStatus = val.every((item: CurrencyType) => ['Creation failed', 'stopped'].includes(item.instanceStatus) && item.locked === 'unlocked');
        const instanceNameStatus = val.every((item: CurrencyType) => item.instanceId?.length > 0);
        const arrData = [
            [
                (val: CurrencyType[]) => val?.length,
                () => {
                    _self.closeBtnDisabled.value = !runningStatus;
                    _self.openBtnDisabled.value = !stoppedStatus;
                    _self.restartBtnDisabled.value = !runningStatus;
                    _self.recoveryInstanceBtnDisabled.value = !recoveryInstanceStatus;
                    _self.instanceNameBntDisabled.value = !instanceNameStatus;
                    _self.resetPasswordBntDisabled.value = !stoppedStatus;
                }
            ],
            [
                (val: CurrencyType[]) => !val?.length,
                () => {
                    _self.closeBtnDisabled.value = true;
                    _self.openBtnDisabled.value = true;
                    _self.restartBtnDisabled.value = true;
                    _self.recoveryInstanceBtnDisabled.value = true;
                    _self.instanceNameBntDisabled.value = true;
                    _self.resetPasswordBntDisabled.value = true;
                }
            ]
        ];
        for (const key of arrData) {
            if (key[0](val)) {
                key[1](val);
                break;
            }
        }
    };

    checkAll = (event: boolean) => {
        const newData =  this.reactiveArr.tableData;
        this.reactiveArr.tableData.map(item => {
            // toggleRowSelection接收两个参数，第一个是被勾选的数据，第二个是选中状态（注：需要注册 ref 来引用）
            if (event) {
                for(let i = 0; i < newData.length; i++) {
                    // 未添加的才插入，否则数据会冗余
                    if (!this.reactiveArr.arr.some((s: {instanceId: string;}) => s.instanceId === newData[i].instanceId)) {
                        this.reactiveArr.arr.push(newData[i]);
                    }
                }
            }
            else {
                for (let i = 0; i< this.reactiveArr.tableData.length; i++) {
                    for (let j = 0; j< this.reactiveArr.currentData.length; j++){
                        // 移除当前页currentData的数据
                        if (this.reactiveArr.currentData[j].instanceId === this.reactiveArr.tableData[i].instanceId) {
                            this.reactiveArr.currentData.splice(j, 1);
                        }
                    }
                    for(let k = 0; k < this.reactiveArr.arr.length; k++){
                        // 移除当前页currentData的数据
                        if (this.reactiveArr.arr[k].instanceId === this.reactiveArr.tableData[i].instanceId) {
                            this.reactiveArr.arr.splice(k, 1);
                        }
                    }
                }
            }
            this.filter.tableRef?.value?.toggleRowSelection(item, event);
            // this.tableRef.value?.toggleRowSelection(item, event);
        });
    };
};

export default DeviceList;
