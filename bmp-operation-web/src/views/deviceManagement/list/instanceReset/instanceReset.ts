interface DataType {
    resetInstanceDiaLog: Ref<boolean>;
    instanceName: Ref<string>;
    instanceId: Ref<string>;
};

const resetInstance = (deviceList: any) => {
    const data: DataType = {
        resetInstanceDiaLog: ref<boolean>(false),
        instanceName: ref<string>(''),
        instanceId: ref<string>('')
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
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        ...data,
        resetInstanceClick,
        resetInstanceCancel,
        resetInstanceSure
    }
};

export default resetInstance;
