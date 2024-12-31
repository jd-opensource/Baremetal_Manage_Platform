import {CurrentInstance} from 'utils/publicClass.ts';

const resetPassword = (deviceList: any) => {
    const resetPasswordDiaLog: Ref<boolean> = ref<boolean>(false);
    const mitt = new CurrentInstance();

    const resetPasswordClick = async (item: any, _: string, status: boolean) => {
        if (!status) return;
        resetPasswordDiaLog.value = !resetPasswordDiaLog.value;
        await nextTick(() => {
            mitt.instanceMitt?.proxy?.$Bus?.emit('reset-password',item);
        });
    };

    const resetPasswordCancel = (type: boolean) => {
        resetPasswordDiaLog.value = type;
    };

    const resetPasswordSure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        resetPasswordDiaLog,
        resetPasswordClick,
        resetPasswordCancel,
        resetPasswordSure
    }
};

export default resetPassword;
