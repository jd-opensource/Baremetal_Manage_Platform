class DefaultSetOperate {
    defaultSetDiaLog: Ref<boolean> = ref<boolean>(false);
    faultRulesListOpt: {refresh(): void};

    constructor(faultRulesListOpt: {refresh(): void}) {
        this.faultRulesListOpt = faultRulesListOpt;
    }

    defaultSetClick = () => {
        this.defaultSetDiaLog.value = !this.defaultSetDiaLog.value;
    };

    defaultSetCancel = (type: boolean) => {
        this.defaultSetDiaLog.value = type;
    };

    defaultSetSure = () => {
        this.faultRulesListOpt.refresh();
    };
};

export default DefaultSetOperate;
