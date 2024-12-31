import {CurrencyType2, StatusParamsType, StatusType} from '@utils/publicType';

interface DataType {
    reactiveArr: {
        itemData: CurrencyType2;
    };
    securityVerificationDiaLog: Ref<boolean>;
};

interface EditIdcOperateType {
    editingComputerRoomDiaLog: {
        value: boolean;
    };
};

/**
 * 安全验证操作
 * @param {Object} editIdcOperate 当前编辑的这一项
*/
const securityVerificationOperate = (editIdcOperate: EditIdcOperateType) => {
    const data: DataType = {
        reactiveArr: reactive<{
            itemData: CurrencyType2;
        }>({
            itemData: {}
        }),
        securityVerificationDiaLog: ref<boolean>(false)
    };

    /**
     * 点击操作，控制安全验证弹窗显示隐藏
     * @return {boolean} securityVerificationDiaLog.value true-显示 false-隐藏
    */
    const operateClick = (item: string | number): boolean => {
        data.reactiveArr.itemData.value = item;
        return data.securityVerificationDiaLog.value = !data.securityVerificationDiaLog.value;
    };

    /**
     * 安全验证弹窗取消事件，弹窗关闭
     * @param {Array} type 控制弹窗显示隐藏 true 显示 false 隐藏
     * @return {boolean} securityVerificationDiaLog.value true-显示 false-隐藏
    */
    const securityVerificationCancel: StatusParamsType <(type: boolean) => boolean> = (type: boolean): boolean => {
        return data.securityVerificationDiaLog.value = type;
    };

    /**
     * 安全验证弹窗成功的回调-显示隐藏
     * @return {boolean} editingComputerRoomDiaLog.value true-显示 false-隐藏
    */
    const securityVerificationSure: StatusType<() => boolean> = (): boolean => {
        return editIdcOperate.editingComputerRoomDiaLog.value = !editIdcOperate.editingComputerRoomDiaLog.value;
    };

    return {
        ...data,
        operateClick,
        securityVerificationSure,
        securityVerificationCancel
    }
};

export default securityVerificationOperate;
