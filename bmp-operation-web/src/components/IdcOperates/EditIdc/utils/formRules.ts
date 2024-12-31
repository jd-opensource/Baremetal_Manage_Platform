import {language} from 'utils/publicClass.ts';
import {FormType, RulesType, RuleFormRefType} from '../typeManagement';

interface DataType {
    value: FormType;
    idcNameFlag: Ref<boolean>;
    nameEnFlag: Ref<boolean>;
    customGradeFlag: Ref<boolean>;
    iloUserFlag: Ref<boolean>;
    iloPasswordFlag: Ref<boolean>;
    switchUser1Flag: Ref<boolean>;
    switchPassword1Flag: Ref<boolean>;
    switchUser2Flag: Ref<boolean>;
    switchPassword2Flag: Ref<boolean>;
    ruleFormRef: Ref<RuleFormRefType | undefined>;
    ruleForm: FormType;
    rules: RulesType;
};

const formRule = (props: any, rulesCheck: any) => {
    const data: DataType = {
        // props参数对象
        value: props?.itemData?.value??props?.itemData,
        idcNameFlag: ref<boolean>(false),
        nameEnFlag: ref<boolean>(false),
        customGradeFlag: ref<boolean>(false),
        iloUserFlag: ref<boolean>(false),
        iloPasswordFlag: ref<boolean>(false),
        switchUser1Flag: ref<boolean>(false),
        switchPassword1Flag: ref<boolean>(false),
        switchUser2Flag: ref<boolean>(false),
        switchPassword2Flag: ref<boolean>(false),
        // 表单ref
        ruleFormRef: ref<RuleFormRefType>(),
        // 表单-输入项ts类型
        ruleForm: reactive<DataType['ruleForm']>(
            {
                ...props?.itemData?.value??props?.itemData,
                customGrade: ''
            }
        ),
        // 表单-规则
        rules: reactive<RulesType>({
            name: [ // 机房名称
                {
                    required: true,
                    trigger: 'blur',
                    validator: rulesCheck.computerRoomNameCheck
                }
            ],
            nameEn: [ // 机房英文名称
                {
                    required: true,
                    trigger: 'blur',
                    validator: rulesCheck.computerRoomCodeCheck
                }
            ],
            level: [ // 机房等级
                {
                    required: true,
                    trigger: 'change'
                }
            ],
            customGrade: [ // 自定义机房等级
                {
                    required: true,
                    trigger: 'blur',
                    validator: rulesCheck.computerRoomGradeCheck
                }
            ],
            address: [ // 机房地址
                {
                    required: false,
                    message: language.t('editIdc.errTip.computerRoomAddressEmpty'),
                    trigger: 'blur'
                }
            ],
            iloUser: [ // 带外登录账号
                {
                    required: false,
                    trigger: 'blur',
                    validator: rulesCheck.iloUserCheck
                }
            ],
            iloPassword: [ // 带外登录密码
                {
                    required: false,
                    trigger: 'blur',
                    validator: rulesCheck.iloPasswordCheck
                }
            ],
            switchUser1: [ // 交换机1登录账号
                {
                    required: false,
                    trigger: 'blur',
                    validator: rulesCheck.switchUser1Check
                }
            ],
            switchPassword1: [ // 交换机1登录密码
                {
                    required: false,
                    trigger: 'blur',
                    validator: rulesCheck.switchPassword1Check
                }
            ],
            switchUser2: [ // 交换机2登录账号
                {
                    required: false,
                    trigger: 'blur',
                    validator: rulesCheck.switchUser2Check
                }
            ],
            switchPassword2: [ // 交换机2登录密码
                {
                    required: false,
                    trigger: 'blur',
                    validator: rulesCheck.switchPassword2Check
                }
            ]
        })
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    const getFormRef = (formEl: {value: RuleFormRefType}): void => {
        data.ruleFormRef.value = formEl.value;
    };

    const validate = (type1: string, type2: {value: boolean}) => {
        data.ruleFormRef.value!.validateField([type1] as never, (valid: string) => !valid ? type2.value = false : type2.value = true);
    };

    return {
        ...data,
        getFormRef,
        validate
    };
};

export default formRule;
