import {language} from 'utils/publicClass.ts';
import {CurrencyType3} from '@utils/publicType';

const rulesVerification = () => {

    /**
     * 机房名称的校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const computerRoomNameCheck = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck(value, language.t('editIdc.errTip.computerRoomNameEmpty'), language.t('editIdc.errTip.currency'), callback);
    };

    /**
     * 机房英文名称的校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const computerRoomCodeCheck = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck(value, language.t('editIdc.errTip.computerRoomCodeEmpty'), language.t('editIdc.errTip.currency'), callback);
    };

    /**
     * 自定义机房等级校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const computerRoomGradeCheck = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck(value, language.t('editIdc.errTip.customComputerRoomGrade'), language.t('editIdc.errTip.currency'), callback);
    };

    /**
     * 通用校验
     * @param {string} value 输入项/select选择项
     * @param {string} emptyTip 为空提示
     * @param {string} errTip 输入有误提示
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const currencyCheck = (value: string, emptyTip: string, errTip: string, callback: (arg0?: Error | string) => void) => {
        const check = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(emptyTip))
            ],
            [
                (value: string) => value.length > 64,
                () => callback(new Error(errTip))
            ],
            [
                (value: string) => value,
                () =>  callback()
            ]
        ];
        for (const key of check) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    /**
     * 带外登录账号校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const iloUserCheck = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck2(value, language.t('editIdc.errTip.currencyPassword'), callback);
    };

    /**
     * 带外登录密码校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const iloPasswordCheck = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck2(value, language.t('editIdc.errTip.currencyPassword'), callback);
    };

    /**
     * 交换机1登录账号校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const switchUser1Check = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck2(value, language.t('editIdc.errTip.currencyPassword'), callback);
    };

    /**
     * 交换机1登录密码校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const switchPassword1Check = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck2(value, language.t('editIdc.errTip.currencyPassword'), callback);
    };

    /**
     * 交换机2登录账号校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const switchUser2Check = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck2(value, language.t('editIdc.errTip.currencyPassword'), callback);
    };

    /**
     * 交换机2登录密码校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const switchPassword2Check = ( _: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        currencyCheck2(value, language.t('editIdc.errTip.currencyPassword'), callback);
    };

    /**
     * 通用校验
     * @param {string} value 输入项
     * @param {string} errTip 输入有误提示
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    const currencyCheck2 = (value: string, errTip: string, callback: (arg0?: Error | string) => void) => {
        if (value.length > 128) {
            callback(new Error(errTip));
        }
        else {
            // 正常情况
            callback();
        }
    };

    return {
        computerRoomNameCheck,
        computerRoomCodeCheck,
        computerRoomGradeCheck,
        iloUserCheck,
        iloPasswordCheck,
        switchUser1Check,
        switchPassword1Check,
        switchUser2Check,
        switchPassword2Check
    };
};

export default rulesVerification;
