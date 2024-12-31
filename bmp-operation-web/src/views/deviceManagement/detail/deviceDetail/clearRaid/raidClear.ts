/**
 * @file
 * @author
*/

class RaidClearOpt {
    clearInfoDiaLog: Ref<boolean> = ref<boolean>(false);
    checkboxStatus: Ref<boolean> = ref<boolean>(false);
    deviceDetail: {initData(): void}
    collectConfirmOpt: {
        resetCollectInfoDiaLog: {
            value: boolean;
        }
    };

    constructor(collectConfirmOpt: {resetCollectInfoDiaLog: {value: boolean}}, deviceDetail: {initData(): void}) {
        this.collectConfirmOpt = collectConfirmOpt;
        this.deviceDetail = deviceDetail;
    }

    clearInfoClick = (type: boolean) => {
        this.checkboxStatus.value = type;
        this.clearInfoDiaLog.value = true;
    }

    clearRaidCancel = (type: boolean) => {
        this.clearInfoDiaLog.value = type;
        // this.collectConfirmOpt.resetCollectInfoDiaLog.value = true;
    }

    clearRaidSure = () => {
        this.clearInfoDiaLog.value = false;
        this.deviceDetail.initData();
    }
}

export default RaidClearOpt;