import {TableType} from '../typeManagement';
import {CurrencyType} from '@utils/publicType';
import {CurrentInstance, paginationOperate, RouterOperate, language} from 'utils/publicClass';
import {methodsTotal, deepCopy, msgTip} from 'utils/index.ts';
// OSSType, ReactiveArrType, FilterParamsType
import store from 'store/index.ts';
import MessageStaticData from 'staticData/message/index.ts';
import {AxiosError} from 'axios';
import {removeLong} from 'request/RemoveLongReq/remove.ts';
// let d: any;
let _self: any = null;
class MyMessageListOpt {
    listTimer: Ref<any> = ref<any>(null);
    ossStore = store.ossDataInfo();
    routerOperate = new RouterOperate();
    searchTip: Ref<boolean> = ref<boolean>(false);
    deleteBtnDisabled: Ref<boolean> = ref<boolean>(true);
    readBtnDisabled: Ref<boolean> = ref<boolean>(true);
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(false);
    reactiveArr: {
        // filterParams: {messageSubType: string};
        tableData: {[x: string]: string}[];
        selectArr: CurrencyType[];
        defTableData: TableType[];
        arr: any;
        currentData: any;
        copyDataVal: {label: string; value: string; checkboxStatus: boolean}[];
        customFilterList: {label: string; value: string; checkboxStatus: boolean;}[];
    } = reactive<{
        // filterParams: {messageSubType: string};
        tableData: {[x: string]: string}[];
        selectArr: CurrencyType[];
        defTableData: TableType[];
        arr: TableType[];
        currentData: TableType[];
        copyDataVal: {label: string; value: string; checkboxStatus: boolean}[];
        customFilterList: {label: string; value: string; checkboxStatus: boolean}[];
    }>({
        tableData: [],
        selectArr: [],
        defTableData: [],
        arr: [],
        currentData: [],
        copyDataVal: [],
        customFilterList: []
    });
    isIntervalRequest: boolean = false;
    hasCheckAll: Ref<boolean> = ref<boolean>(false);
    total: Ref<number> = ref<number>(0);
    noReadTotal: Ref<number> = ref<number>(0);
    proxy = new CurrentInstance().proxy;
    filter: any = {};
    search: any;
    instanceMitt: any = new CurrentInstance();
    setMessageSubTypeClass: Ref<boolean> = ref<boolean>(false);
    readonly mittData = [
        {
            type: 'custom-filter-click',
            fn: 'getOpt'
        },
        {
            type: 'custom-filter-reset',
            fn: 'customResetClick'
        },
        {
            type: 'custom-filter-dia-log',
            fn: 'getStatus'
        }
    ]

    constructor(filter: any, search: any) {
        // super();
        this.filter = filter;
        this.search = search;
        this.init();
        // this.handleVisibility();
        this.watchStoreCustomList();
        // 标为已读和删除需要调用这个接口
        this.getMessageCountNoReadTotal();

        onBeforeUnmount(() => this.unMountOpt())
        onUnmounted(() => this.unMountOpt());
        for (const key of this.mittData) {
            // @ts-ignore
            this.onOffMitt(key.type, this[key.fn], true);
        }
        _self = this;
    }

    unMountOpt = () => {
        clearTimeout(this.listTimer.value as number);
        removeLong(this.routerOperate.router.currentRoute.value.name)
        for (const key of this.mittData) {
            // @ts-ignore
            this.onOffMitt(key.type, this[key.fn]);
        }
    }

    onOffMitt = (type: string, fn: Function, hasOn: boolean = false) => {
        if (!hasOn) {
            this.instanceMitt?.proxy?.$Bus?.off(type, fn);
            return;
        }
        this.instanceMitt?.proxy?.$Bus?.on(type, fn);
    }

    handleVisibility = () => {
        // document.addEventListener('visibilitychange', () => {
        //     const status = {
        //         'hidden': () => clearTimeout(this.listTimer as number),
        //         'visible': () => this.getMessageList()
        //     }
        //     status[document.visibilityState]();
        // })
    }

    getOpt = (item: {label: string; value: string; checkboxStatus: boolean;}[]) => {
        this.customFilterClick(item);
    }

    watchStoreCustomList = () => {
        watch(() => this.ossStore.customFilterList, (newValue) => {
            this.reactiveArr.customFilterList = newValue;
            this.reactiveArr.copyDataVal = deepCopy(this.reactiveArr.customFilterList);
        })
    }

    rowClick = (row: {[x: string]: string}, _: any, event: Event) => {
        this.filter.tableRef?.value?.toggleRowSelection(row, event);
    };

    cellClick = (row: CurrencyType, b: {property: string | undefined}, event: any) => {
        // prop必传，否则为undefined
        const type = ['content', undefined];
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

    // 多选框回显
    getRowKeys = (row: {messageId: string;}) => {
        if (row.messageId) return row.messageId;
    };

    setMessageCount = (item: any) => {
        // Object.is(language.t('myMessage.table.inbondAlert'), item.messageType) ? this.#instanceId(item) + this.#setDays(item) :
        const val =  Object.is(language.t('myMessage.table.inbondAlert'), item.messageType) ? this.#instanceId(item) : this.#setSn(item);
        return item.messageType + ' - ' + this.#setNotice(item) + val + this.#setDays(item);
    }

    handleSelectionChange = async (val: {[x: string]: string}[]) => {
        _self.reactiveArr.selectArr = val;
        _self.hasCheckAll.value = _self.reactiveArr.selectArr.length === _self.reactiveArr.tableData.length;
        _self.deleteBtnDisabled.value = !(val?.length);
        _self.readBtnDisabled.value = !(val?.length);
    }

    readMessageClick = async () => {
        try {
            // @ts-ignore
            const res = await this.proxy.$messageApi.doReadAPI({messageIds: this.reactiveArr.selectArr.map(item => item.messageId)});
            if (res?.data?.result?.success) {
                this.restoreDefaultStatus();
            }
        }
        catch (e) {
            const err = e as AxiosError;
            err.message && msgTip.error(err.message);
        }
    }

    restoreDefaultStatus = () => {
        this.reactiveArr.tableData.forEach(item => {
            this.filter.tableRef.value?.toggleRowSelection(item, false);
        });
        this.init();
        this.getMessageCountNoReadTotal();
    }

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

    #setInstanceName = (name: undefined) => {
        const type = new Map([
            [void 0, ''],
        ]);
        return type.get(name)?? name + '/';
    }

    #setDays = (item: {detail: string; messageType: string}) => {
        if (item?.detail === null) return '';
        if ([language.t('myMessage.table.faultMessages', language.t('myMessage.table.operationMessages'))].includes(item.messageType)) return '';
        const obj = new Map([
            ['90', language.t('myMessage.day90')],
            ['30', language.t('myMessage.day30')],
            ['15', language.t('myMessage.day15')],
            ['0', language.t('myMessage.expired')]
        ]);
        return obj.get(item.detail)?? '';
    }

    #setNotice = (item: {messageType: string, faultType: string, messageSubType: string}) => {
        const type = new Map([
            [language.t('myMessage.table.systemMessages'), `【${item.messageSubType}】`],
            [language.t('myMessage.table.faultMessages'), `【${item.faultType}】${language.t('myMessage.table.remind')}`],
            [language.t('myMessage.table.operationMessages'), `【${item.messageSubType}】 ${language.t('myMessage.table.failureNotification')}`],
            [language.t('myMessage.table.inbondAlert'), `【${item.messageSubType}】 ${language.t('myMessage.table.alarmTips')}`]
        ]);
        return type.get(item.messageType)
    }

    refresh = () => {
        if (!this?.reactiveArr?.tableData?.length) {
            this.reset();
            return;
        }
        this.init();
    }

    setClass = (messageSubType: string) => {
        if (messageSubType || this.setMessageSubTypeClass.value) {
            return 'def-cls lignt-filter';
        }
        return 'def-cls default-filter';
    }

    reset = () => {
        let hasReq: boolean = false;
        const {hasClear, reactiveArr} = this.search;
        const {messageType, messageSubType} = this.filter.reactiveArr.filterParams;
        const status = ['0', '1'].includes(reactiveArr.radioGroup.hasRead) 
        hasReq = (!status  && (reactiveArr.searchParams.detail || messageType || messageSubType));
        clearTimeout(this.listTimer.value)
        this.filter.tableRef.value?.clearFilter();
        this.filter.reactiveArr.filterParams = {};
        this.setMessageSubTypeClass.value = false;
        hasClear.value = true;
        reactiveArr.detail = '';
        reactiveArr.searchParams = {};
        reactiveArr.radioGroup.hasRead = '';
        reactiveArr.radioGroup.hasReadName = language.t('myMessage.radioGroup.all')
        const valData = this.reactiveArr.copyDataVal.map(item => {
            return {
                ...item,
                checkboxStatus: false
            }
        })
        this.reactiveArr.customFilterList = valData;
        this.reactiveArr.copyDataVal = valData;
        this.instanceMitt?.proxy?.$Bus?.emit('custom-filter-input-clear');
        this.getMessageCountNoReadTotal();
        if (hasReq) {
            if (paginationOperate.pageNumber.value > 1) {
                paginationOperate.pageNumber.value = 1;
                return;
            }
            this.init();
        }
    }

    #instanceId = (item: {instanceName: string | undefined; instanceId: string | undefined}) => {
        if ([item.instanceName, item.instanceId].includes(void 0)) return;
        return '（' + item.instanceName + '/' + item.instanceId + '）';
    }
 
    #setSn = (item: {sn: undefined, instanceName: undefined}) => {
        const type = new Map([
            [void 0, '']
        ])
        return type.get(item.sn)?? '（' + this.#setInstanceName(item.instanceName) + item.sn + '）';
    }

    customFilterClick = (val: {label: string; value: string}[]) => {
        const data = val.map(item => item.label).join(',');
        this.filter.reactiveArr.filterParams.messageSubType = data;
        // paginationOperate.pageNumber.value = 1;
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.init();
    }

    customResetClick = () => {
        this.reactiveArr.customFilterList = this.reactiveArr.copyDataVal.map((item: {label: string; value: string; checkboxStatus: boolean}) => {
            return {
                ...item,
                checkboxStatus: false
            }
        })
        this.setMessageSubTypeClass.value = false;
        this.reactiveArr.copyDataVal = this.reactiveArr.customFilterList;
        this.filter.reactiveArr.filterParams.messageSubType = '';
        this.init();
    }

    jump = (item: {messageId: string; detail: string; messageType: string; licenseEndTime: number;}) => {
        const path = this.proxy!.$defInfo.routerPath('myMessageDetail');
        clearTimeout(this.listTimer.value as number);
        const obj = {
            type: this.getQueryType(item.messageType),
            id: item.messageId,
            newDate: item.detail,
            endTime: item.licenseEndTime
        }
        const val = encodeURIComponent(JSON.stringify(obj))
        localStorage.setItem('messageInfo', val);
        this.routerOperate.router.push(path as unknown as string)
    }

    getStatus = (val: boolean) => {
        this.setMessageSubTypeClass.value = val;
    }

    setDate = (detail: string) => {
        if (detail === null) return;
        return detail;
    }

    getQueryType = (messageType: string) => {
        const type = new Map([
            [language.t('myMessage.table.faultMessages'), 'faultMessage'],
            [language.t('myMessage.table.systemMessages'), 'unexpiredExpired'],
            [language.t('myMessage.table.operationMessages'), 'optMessage'],
            [language.t('myMessage.table.inbondAlert'), 'inbondAlert']
        ]);
        return type.get(messageType)
    }

    init = async () => {
        this.isLoading.value = true;
        this.getMessageList();
    }

    getMessageList = async () => {
        clearTimeout((this.listTimer.value as number));
        const path = this.proxy!.$defInfo.routerPath('myMessage');
        if (this.routerOperate.router.currentRoute.value.path !== path as unknown as string) {
            return;
        }
        const param = {
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value,
            ...this.filter.reactiveArr.filterParams,
            ...this.search.reactiveArr.searchParams,
            hasRead: this.search.reactiveArr.radioGroup.hasRead
        };
        this.filter.filterEmptyInfo.deleteEmtpyData(param);
        const status = Object.keys(param).length > 2;
        try {
            // @ts-ignore
            const res = await this.proxy.$messageApi.messageListAPI({...param});
            if (![void 0].includes(res?.data?.result?.messages)) {
                const {messages, totalCount} = methodsTotal.lineConverting(res.data.result);
                messages.forEach((item: TableType, index: number) => {
                    MessageStaticData.messageTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                })
                this.reactiveArr.tableData = messages;
                paginationOperate.total.value = totalCount;
                if (status) {
                    clearTimeout(this.listTimer.value);
                }
                else {
                    this.listTimer.value = setTimeout(() => {
                        if (status) return;
                        this.isIntervalRequest = true;
                        this.getMessageList();
                        this.getMessageCountNoReadTotal();
                    }, 5000)
                }
                return;
            }
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
        }
        catch(e) {
            if (!this.isIntervalRequest) {
                const err = e as AxiosError;
                if (Object.is(err?.code, 402)) return;
                err?.message && msgTip.error(err.message);
                return;
            }
            // if (status) {
            //     clearTimeout(this.listTimer.value);
            // }
            // else {
            this.listTimer.value = setTimeout(() => {
                if (status) {
                    clearTimeout(this.listTimer.value);
                    return;
                };
                this.getMessageList();
                this.getMessageCountNoReadTotal();
            }, 5000)
            // }
        }
        finally {
            this.isLoading.value = false;
            this.search.hasClear.value = false;
            this.searchTip.value = status;
            this.instanceMitt?.proxy?.$Bus?.emit('search-tip-value',  this.searchTip.value ? 15 : 0);
            this.instanceMitt?.proxy?.$Bus?.emit('message-list', this.reactiveArr.tableData);
        }
    }

    getMessageCountNoReadTotal = async () => {
        try {
            // @ts-ignore
            const res = await this.proxy.$messageApi.messagesStatisticAPI({});
            const {totalCount, unreadCount} = res.data.result;
            this.total.value = totalCount;
            this.noReadTotal.value = unreadCount;
        }
        catch {
            this.total.value = 0;
            this.noReadTotal.value = 0;
        }
    }
}

export default MyMessageListOpt;
