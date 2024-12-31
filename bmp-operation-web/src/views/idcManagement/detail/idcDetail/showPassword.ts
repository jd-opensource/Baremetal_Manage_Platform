
import {PasswordType} from '../typeManagement';

/**
 * 是否显示密码
 * @param {Object} verificationOperate 验证操作
 * @param {Object} idcDetail 当前操作的这一项
*/
const hasShowPassword = (verificationOperate: PasswordType['validate'], idcDetail: PasswordType['idcDetail']) => {

    /**
     * 是否明文显示密码
     * @param {Object} item 当前点击的这一项
     * @param {string} type 当前点击的状态是 open or close
    */
    const hasPasswordClick = (item: {hasPassword: boolean; type: string; password: string;}, type: string): void => {
        verificationOperate.hasEditOperate.value = false;
        idcDetail.password.value = item.type;
        if (type === 'close') {
            item.hasPassword = !item.hasPassword;
            item.password = '********';
            idcDetail.password.value = '';
            return;
        }
        verificationOperate.securityVerificationDiaLog.value = !verificationOperate.securityVerificationDiaLog.value;
        verificationOperate.hasEditOperate.value = false;
    };

    return {
        hasPasswordClick
    }
};

export default hasShowPassword;
