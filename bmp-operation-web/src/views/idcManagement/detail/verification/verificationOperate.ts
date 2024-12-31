import {StatusType} from '@utils/publicType';

interface DataType {
    securityVerificationDiaLog: Ref<boolean>;
    hasEditOperate: Ref<boolean>;
};

interface EditIdcType {
    editIdcDiaLog: {
        value: boolean;
    }
};

/**
 * 安全验证操作
 * @param {Object} 编辑机房操作
*/
const verificationOperate = (editIdcOperate: EditIdcType) => {
    const data: DataType = {
        // 安全验证弹窗状态
        securityVerificationDiaLog: ref<boolean>(false),
        // 是否进行编辑操作-默认true
        hasEditOperate: ref<boolean>(true)
    };

    /**
     * 点击操作，控制安全验证弹窗显示隐藏
     * @return {boolean} securityVerificationDiaLog.value true-显示 false-隐藏
    */
    const operateClick: StatusType<() => boolean> = (): boolean => {
        data.hasEditOperate.value = true;
        return data.securityVerificationDiaLog.value = !data.securityVerificationDiaLog.value;
    };

    /**
     * 安全验证弹窗-显示隐藏
     * @return {boolean} editIdcDiaLog.value true-显示 false-隐藏
    */
    const securityVerificationSure: StatusType<() => boolean> = (): boolean => {
        return editIdcOperate.editIdcDiaLog.value = true;
    };

    /**
     * 安全验证弹窗取消事件，弹窗关闭
     * @param {Array} type 控制弹窗显示隐藏 true 显示 false 隐藏
     * @return {boolean} securityVerificationDiaLog.value true-显示 false-隐藏
    */
    const securityVerificationCancel = (type: boolean): boolean => {
        return data.securityVerificationDiaLog.value = type;
    };

    return {
        ...data,
        operateClick,
        securityVerificationSure,
        securityVerificationCancel
    }
};

export default verificationOperate;
