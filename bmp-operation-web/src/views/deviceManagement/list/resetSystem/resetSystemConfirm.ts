import {CurrentInstance} from 'utils/publicClass.ts';

const resetSystemConfirmOperate = (resetSystemOpt: any) => {
    const secondaryConfirmationDiaLog: Ref<boolean> = ref<boolean>(false);
    const secondaryConfirmationName: Ref<string> = ref<string>('');
    const mitt = new CurrentInstance();
    let items: any = {}
    const resetSystemConfirmClick = async(item: {instanceStatus: string; instanceName: string;}) => {
        if (!['Reinstall failed', 'stopped'].includes(item.instanceStatus)) return;
        secondaryConfirmationName.value = item.instanceName;
        secondaryConfirmationDiaLog.value = !secondaryConfirmationDiaLog.value;
        items = item;
    };

    const secondaryConfirmationCancel = (type: boolean) => {
        secondaryConfirmationDiaLog.value = type;
    };

    const secondaryConfirmationSure = async () => {
        resetSystemOpt.resetSystemDiaLog.value= true;
            await nextTick(() => {
            mitt.instanceMitt?.proxy?.$Bus?.emit('reset-system', {...items});
        })
        // fn();
    };

    return {
        resetSystemConfirmClick,
        secondaryConfirmationName,
        secondaryConfirmationDiaLog,
        secondaryConfirmationCancel,
        secondaryConfirmationSure
    };
};

export default resetSystemConfirmOperate;
