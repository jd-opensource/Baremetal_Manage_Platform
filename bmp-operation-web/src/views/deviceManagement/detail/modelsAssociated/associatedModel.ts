/**
 * @file
 * @author
*/

class AssociatedModelOpt {
    associatedModelDiaLog: Ref<boolean> = ref<boolean>(false);
    deviceDetail;

    constructor(deviceDetail: {initData(): void}) {
        this.deviceDetail = deviceDetail;
    }

    associatedModelClick = () => {
        this.associatedModelDiaLog.value  = !this.associatedModelDiaLog.value;
    }

    associatedModelCancel = (type: boolean) => {
        this.associatedModelDiaLog.value = type;
    }

    associatedModelSure = (diskTableOpt: {init(): void}) => {
        this.deviceDetail.initData();
        diskTableOpt.init();
    }
}

export default AssociatedModelOpt;
