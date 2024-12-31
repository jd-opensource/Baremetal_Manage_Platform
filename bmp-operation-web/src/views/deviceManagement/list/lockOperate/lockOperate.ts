interface DataType {
    lockDiaLog: Ref<boolean>;
    instanceName: Ref<string>;
    operate: Ref<string>;
    instanceId: Ref<string>;
};

const lockDeblocking = (deviceList: any) => {
    const data: DataType = {
        lockDiaLog: ref<boolean>(false),
        instanceName: ref<string>(''),
        operate: ref<string>(''),
        instanceId: ref<string>('')
    };

    const lock = (name: string, id: string, type: string) => {
        data.lockDiaLog.value = !data.lockDiaLog.value;
        data.instanceName.value = name;
        data.instanceId.value = id;
        data.operate.value = type;
    };

    const lockCancel = (type: boolean): boolean => {
        return data.lockDiaLog.value = type;
    };

    const lockSure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        ...data,
        lock,
        lockCancel,
        lockSure
    }
};

export default lockDeblocking;
