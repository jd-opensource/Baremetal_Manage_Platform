interface DataType {
    lockDiaLog: Ref<boolean>;
    instanceName: Ref<string>;
    operate: Ref<string>;
    instanceId: Ref<string>;
};

const lockDeblocking = (deviceDetail: any) => {
    const data: DataType = {
        lockDiaLog: ref<boolean>(false),
        instanceName: ref<string>(''),
        operate: ref<string>(''),
        instanceId: ref<string>('')
    };

    const lock = (name: string, id: string, type: string) => {
        data.instanceName.value = name;
        data.operate.value = type;
        data.instanceId.value = id;
        data.lockDiaLog.value = !data.lockDiaLog.value;
    };

    const lockCancel = (type: boolean): boolean => {
        return data.lockDiaLog.value = type;
    };

    const lockSure = () => {
        deviceDetail.initData();
    };

    return {
        ...data,
        lock,
        lockCancel,
        lockSure
    };
};

export default lockDeblocking;
