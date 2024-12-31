interface DataType {
    instanceId: Ref<string>;
    resetInstanceDiaLog: Ref<boolean>;
    instanceName: Ref<string>;
};

const resetInstanceOperate = (deviceDetail: any) => {
    const data: DataType = {
        instanceId: ref<string>(''),
        resetInstanceDiaLog: ref<boolean>(false),
        instanceName: ref<string>('')
    };

    const resetInstanceClick = (id: string, name: string, status: boolean) => {
        if (!status) return;
        data.instanceId.value = id;
        data.instanceName.value = name;
        data.resetInstanceDiaLog.value = !data.resetInstanceDiaLog.value;
    };

    const resetInstanceCancel = (type: boolean) => {
        data.resetInstanceDiaLog.value = type;
    };

    const resetInstanceSure = () => {
        deviceDetail.initData();
    };

    return {
        ...data,
        resetInstanceClick,
        resetInstanceCancel,
        resetInstanceSure
    };
};

export default resetInstanceOperate;
