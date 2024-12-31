import {language} from 'utils/publicClass.ts';
import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType, RulesType} from '../typeManagement';

/**
 * 表单验证
*/
class RuleFormCheck {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 密码type切换
    passwordType: Ref<string> = ref<string>('password');
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        loginUserName: '', // 用户登录账号
        loginPassword: '' // 用户登录密码
    });
    // 规则类型
    rules: RulesType = reactive<RulesType>({
        loginUserName: [ // 用户登录账号
            {
                required: true,
                message: language.t('login.formSubmit.toolTip'),
                trigger: 'click'
            }
        ],
        loginPassword: [ // 用户登录密码
            {
                required: true,
                message: language.t('login.formSubmit.toolTip'),
                trigger: 'click'
            }
        ]
    });

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    /**
     * 清空输入项
     * @param {string} value 输入框输入的值
    */
    clearInput = (value: string) => {
        (this.ruleForm[`${value}` as keyof typeof this.ruleForm] as string) = '';
    };

    /**
     * 密码是否明文显示，状态切换
     * @param {string} type 密码明文显示的type值
     * @return {string} passwordType.value 密码明文显示的type值
    */
    passwordStatusChange = (type: string): string => {
        return this.passwordType.value = type;
    };
};

export default RuleFormCheck;
