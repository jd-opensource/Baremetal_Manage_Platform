import {CurrentInstance} from 'utils/publicClass.ts';
import {CurrencyType} from '@utils/publicType';
// import {DeviceDetail} from '../../methods';

interface DataType {
    removeRecoveryTitle: Ref<string>;
    removeRecoveryDiaLog: Ref<boolean>;
};

const removeRecoveryOperate = (deviceDetail: any) => {
    const data: DataType = {
        // 设备移除、实例回收标题名称
        removeRecoveryTitle: ref<string>(''),
        // 设备移除、实例回收组件状态
        removeRecoveryDiaLog: ref<boolean>(false)
    };

    const mitt = new CurrentInstance();

    const removeRecoveryClick = async (item: CurrencyType, type: string, status: boolean): Promise<void> => {
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
        deviceDetail.initData();
    };

    return {
        ...data,
        removeRecoveryClick,
        removeRecoveryCancel,
        removeRecoverySure
    };
};

export default removeRecoveryOperate;
