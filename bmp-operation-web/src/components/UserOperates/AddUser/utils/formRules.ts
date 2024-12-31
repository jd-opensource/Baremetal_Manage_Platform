import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType, RuleInfoType} from '../typeManagement';
import {decrypt} from 'utils/index.ts';
import {language, UserCheck, CurrentInstance} from 'utils/publicClass.ts';

class FormRulesOperate {
    emailFlag: Ref<boolean> = ref<boolean>(false);
    phoneFlag: Ref<boolean> = ref<boolean>(false);
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    ruleForm: RuleFormType = reactive<RuleFormType>({
        userName: '', // 用户名
        password: '', // 密码
        roleId: '', // 角色选择
        description: '', // 描述
        phonePrefix: language.t('addUser.phoneData.China'), // 手机号类型
        phoneNumber: '', // 手机号
        email: '' // 邮箱
    });
    // 规则
    rules: RuleInfoType = reactive<RuleInfoType>({
        userName: [ // 用户名
            {
                required: true,
                trigger: 'blur',
                // validator: regExpCheck.userNameCheck
            }
        ],
        password: [ // 密码
            {
                required: true,
                trigger: 'blur',
                // validator: regExpCheck.passwordCheck
            }
        ],
        phoneNumber: [ // 手机号
            {
                required: true,
                trigger: 'blur',
                validator: new UserCheck(this.ruleForm.phonePrefix, 'addUser').phoneChcek
            }
        ],
        email: [ // 邮箱
            {
                required: true,
                trigger: 'blur',
                validator: new UserCheck('', 'addUser').emailChcek
            }
        ]
    });
    // 密码是否明文显示，密码类型
    passwordType: Ref<string> = ref<string>('password');
    instanceProxy = new CurrentInstance().proxy;

    constructor (status: boolean, regExpCheck: {userNameCheck: unknown; passwordCheck: unknown;}) {
        const userPurview: string = this.instanceProxy.$defInfo.purview('user');
        const adminPurview: string = this.instanceProxy.$defInfo.purview('admin');
        const roleId = sessionStorage?.getItem('roleId')??'';
        const encryptDecrypt: string[] = this.instanceProxy.$defInfo.encryptDecrypt(1);
        this.ruleForm.roleId = roleId ? decrypt(roleId, encryptDecrypt[0], encryptDecrypt[1]) : status ? userPurview : adminPurview; // 角色选择
        this.rules.userName[0].validator = regExpCheck.userNameCheck;
        this.rules.password[0].validator = regExpCheck.passwordCheck;
        watch(() =>  this.ruleForm.phonePrefix, (newValue: string) => {
            this.ruleForm.phonePrefix = newValue;
            this.ruleForm.phoneNumber = '';
            this.rules.phoneNumber[0].validator = new UserCheck(this.ruleForm.phonePrefix).phoneChcek;
        }, {deep: true});
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    /**
     * 清空输入项
     * @param {string} inputValue 输入项内容
    */
    clearInput = (inputValue: string): void => {
        (this.ruleForm[`${inputValue}` as keyof typeof this.ruleForm] as string) = '';
    };

    /**
     * 清空输入项
     * @param {string} status 密码类型
     * @return {string} passwordType.value 密码类型
    */
    passwordStatusChange = (status: string): string => {
        return this.passwordType.value = status;
    };

    validate = (type1: string, type2: {value: boolean}) => {
        this.ruleFormRef.value!.validateField(([type1] as never), (valid: string) => !valid ? type2.value = false : type2.value = true);
    };
};

export default FormRulesOperate;
