class EnableStatusOperate {
    enableStatusDiaLog: Ref<boolean> = ref<boolean>(false);
    useItem: {useStatus: boolean; ruleId: number; alertName: string} = {useStatus: false, ruleId: 0, alertName: ''};
    faultRulesListOpt: any;

    constructor(faultRulesListOpt: any) {
        this.faultRulesListOpt = faultRulesListOpt;
    };

    enableStatusClick = (item: any) => {
        item.useLoading = true;
        this.useItem.useStatus = item.useStatus;
        this.useItem.ruleId = item.ruleId;
        this.useItem.alertName = item.alertName
        nextTick(() => {
            this.enableStatusDiaLog.value = !this.enableStatusDiaLog.value;
        })
    };

    enableStatusCancel = (type: boolean) => {
        this.faultRulesListOpt.reactiveArr.tableData.map((item: {useLoading: boolean})=> item.useLoading = false);
        this.enableStatusDiaLog.value = type;
    };

    enableStatusSure = () => {
        this.faultRulesListOpt.refresh();
    };
};

export default EnableStatusOperate;
