import {RuleFormRefType} from '@utils/publicType';
import {language, UserCheck} from 'utils/publicClass.ts'; // cookie、国际化、用户验证
import {RuleFormType, RuleInfoType} from '../typeManagement';
import UserStaticData from 'staticData/user/index.ts'; // 手机号类型数据

interface DataType {
    phoneFlag: Ref<boolean>;
    emailFlag: Ref<boolean>;
    ruleFormRef: Ref<RuleFormRefType | undefined>;
    ruleForm: RuleFormType;
    rules: RuleInfoType;
};

const formRulesOperate = (props: any) => {
    const data: DataType = {
        phoneFlag: ref<boolean>(false),
        emailFlag: ref<boolean>(false),
        // 表单ref
        ruleFormRef: ref<RuleFormRefType>(),
        // 表单项
        ruleForm: reactive<RuleFormType>({
            ...props.listData,
            type: ''
        }),
        // 规则
        rules: reactive<RuleInfoType>({
            description: [ // 描述
                {
                    required: false,
                    message: language.t('editUser.emptyTip.desc'),
                    trigger: 'blur'
                }
            ],
            phoneNumber: [ // 手机号
                {
                    required: true,
                    trigger: 'blur',
                    validator: {}
                }
            ],
            email: [ // 邮箱
                {
                    required: true,
                    trigger: 'blur',
                    validator: new UserCheck().emailChcek
                }
            ]
        })
    };
    
    onMounted(() => {
        nextTick(() => {
            data.rules.phoneNumber[0].validator = new UserCheck(data.ruleForm.phonePrefix).phoneChcek;
        })
        if (props.prefix === '1') {
            data.ruleForm.phonePrefix = language.t('editUser.phoneData.USA');
            data.ruleForm.type = '1';
        }
        else {
            for(const key of UserStaticData.phoneType) {
                if (key.type === props.prefix) {
                    data.ruleForm.phonePrefix = key.name;
                    const value = UserStaticData.cellPhoneType.find((item: {label: string}) => item.label === data.ruleForm.phonePrefix);
                    data.ruleForm.type = value.text;
                }
            }
        }
    });

    const selectChange = (newValue: string) => {
        data.ruleForm.phonePrefix = newValue;
        data.ruleForm.phoneNumber = '';
        const value = UserStaticData.cellPhoneType.find((item: {label: string}) => item.label === newValue);
        data.ruleForm.type = value.text;
        data.rules.phoneNumber[0].validator = new UserCheck(data.ruleForm.phonePrefix).phoneChcek;
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    const getFormRef = (formEl: {value: RuleFormRefType}): void => {
        data.ruleFormRef.value = formEl.value;
    };

    const validate = (type1: never | string, type2: {value: boolean}) => {
        data.ruleFormRef.value!.validateField(([type1] as never), (valid: string) => !valid ? type2.value = false : type2.value = true);
    };

    return {
        ...data,
        getFormRef,
        validate,
        selectChange
    };
};

export default formRulesOperate;
