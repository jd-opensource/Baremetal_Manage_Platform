import {CurrentInstance} from 'utils/publicClass.ts';

interface DataType {
    removeRecoveryDiaLog: Ref<boolean>;
    removeRecoveryTitle: Ref<string>;
};

interface DeviceListType {
    filter: any;
    getDataList: () => void;
};

const removeRecovery = (deviceList: DeviceListType) => {
    const data: DataType = {
        // 设备移除、实例回收组件状态
        removeRecoveryDiaLog: ref<boolean>(false),
        // 设备移除、实例回收标题名称
        removeRecoveryTitle: ref<string>('')
    };

    const mitt = new CurrentInstance();

    const removeRecoveryClick = async (item: {[x: string]: string}, type: string, status: boolean): Promise<void> => {
        if (!status) return;
        data.removeRecoveryDiaLog.value = !data.removeRecoveryDiaLog.value;
        data.removeRecoveryTitle.value = type;
        await nextTick(() => {
            mitt.instanceMitt?.proxy?.$Bus?.emit('device-table', {list: item, type: 'removeRecovery'});
        });
    };

    const removeRecoveryCancel = (type: boolean) => {
        data.removeRecoveryDiaLog.value = type;
    };

    const removeRecoverySure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        ...data,
        removeRecoveryClick,
        removeRecoveryCancel,
        removeRecoverySure
    }
};

export default removeRecovery;
