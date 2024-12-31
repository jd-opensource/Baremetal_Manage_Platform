/**
 * @file
 * @author
*/

class CollectConfirmOpt {
    resetCollectInfoDiaLog: Ref<boolean> = ref<boolean>(false);
    fn;
    deviceDetail: {initData(): void}

    constructor(deviceDetail: {initData(): void}, fn: Function) {
        this.fn = fn;
        this.deviceDetail = deviceDetail;
    }

    collectClick = () => {
        this.resetCollectInfoDiaLog.value = !this.resetCollectInfoDiaLog.value;
    }

    resetCollectInfoCancel = (type: boolean) => {
        this.resetCollectInfoDiaLog.value = type;
    }

    resetCollectInfoSure = (type: boolean) => {
        this.resetCollectInfoDiaLog.value = false;
        this.fn(type);
    }

    resetCollectSuccessClick = () => {
        this.deviceDetail.initData();
    }
}

export default CollectConfirmOpt;