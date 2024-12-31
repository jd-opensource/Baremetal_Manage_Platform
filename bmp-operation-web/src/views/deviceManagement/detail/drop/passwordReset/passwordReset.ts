import {CurrentInstance} from 'utils/publicClass.ts';

const resetPassword = (deviceDetail: any) => {
    const resetPasswordDiaLog: Ref<boolean> = ref<boolean>(false);
    const mitt = new CurrentInstance();

    const resetPasswordClick = async (item: any) => {
        if (!['stopped', 'Resetpasswd failed']) return;
        resetPasswordDiaLog.value = !resetPasswordDiaLog.value;
        await nextTick(() => {
            mitt.instanceMitt?.proxy?.$Bus?.emit('reset-password',item);
        });
    };

    const resetPasswordCancel = (type: boolean) => {
        resetPasswordDiaLog.value = type;
    };

    const resetPasswordSure = () => {
        deviceDetail.initData();
    };

    return {
        resetPasswordDiaLog,
        resetPasswordClick,
        resetPasswordCancel,
        resetPasswordSure
    }
};

export default resetPassword;
