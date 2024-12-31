import {CurrentInstance} from 'utils/publicClass.ts';

const faultHandingConfirmOpt = (faultLogListOperate: {refresh(): void}) =>{
    const faultHandingConfirmDiaLog: Ref<boolean> = ref<boolean>(false);
    const mitt = new CurrentInstance();

    const faultHandingConfirmClick = async (item: any) => {
        if (item.status) return;
        faultHandingConfirmDiaLog.value = !faultHandingConfirmDiaLog.value;
        await nextTick(() => {
            mitt.instanceMitt?.proxy?.$Bus?.emit('fault-handing-confirm', item);
        });
    };

    const faultHandingConfirmCancel = (type: boolean) => {
        faultHandingConfirmDiaLog.value = type;
    };

    const faultHandingConfirmCancelSure = () => {
        faultLogListOperate.refresh();
    };

    return {
        faultHandingConfirmDiaLog,
        faultHandingConfirmClick,
        faultHandingConfirmCancel,
        faultHandingConfirmCancelSure
    }
};

export default faultHandingConfirmOpt;
