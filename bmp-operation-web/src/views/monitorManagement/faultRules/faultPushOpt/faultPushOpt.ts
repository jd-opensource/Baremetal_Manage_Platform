class FaultPushOperate {
    faultPushDiaLog: Ref<boolean> = ref<boolean>(false);
    pushItem: {pushStatus: boolean; ruleId: number; alertName: string} = {pushStatus: false, ruleId: 0, alertName: ''};
    faultRulesListOpt: {
        reactiveArr: {
            tableData: any;
        };
        refresh(): void;
    };

    constructor (faultRulesListOpt: {reactiveArr: {tableData: any}; refresh(): void;}) {
        this.faultRulesListOpt = faultRulesListOpt;
    };

    faultPushClick = (item: any) => {
        item.pushLoading = true;
        this.pushItem.pushStatus = item.pushStatus;
        this.pushItem.ruleId = item.ruleId;
        this.pushItem.alertName = item.alertName
        
        return new Promise((resolve) => {
            nextTick(() => {
                this.faultPushDiaLog.value = !this.faultPushDiaLog.value;
            })
            return resolve(true);
        })
    };

    faultPushCancel = (type: boolean) => {
        this.faultRulesListOpt.reactiveArr.tableData.map((item: {pushLoading: boolean})=> item.pushLoading = false);
        this.faultPushDiaLog.value = type;
    };

    faultPushSure = () => {
        this.faultRulesListOpt.refresh();
    };
};

export default FaultPushOperate;
