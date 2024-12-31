import {RuleFormRefType} from '@utils/publicType';
import {revisePasswordAPI} from 'api/userCenter/request.ts';
import {language} from 'utils/publicClass';
import {encrypt, msgTip, customTip} from 'utils/index.ts';

class RevisePassword {
    isLoading: Ref<boolean> = ref<boolean>(false);
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单项
    ruleForm: any = reactive<any>({
        oldPassword: '',
        newPassword: '',
        confirmPassword: ''
    });
    // 规则
    rules: any = reactive<any>({
        oldPassword: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        newPassword: [ // 邮箱
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        confirmPassword: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ]
    });
    regExpCheck;
    formData = [
        {
            prop: 'oldPassword',
            style: 'marginBottom: 20px',
            label: language.t('userCenter.currectPassword'),
            class: 'set-password',
            flag: 'oldPasswordFlag',
            formItemOtherClass: ['set-password-empty', ''],
            itemOtherClass: ['', ''],
            maxlength: '30',
            type: 'password',
            model: 'oldPassword',
            placeholder: '',
            showMessage: true
        },
        {
            prop: 'newPassword',
            style: 'marginBottom: 20px',
            label: language.t('userCenter.newPassword'),
            class: 'set-password',
            flag: 'colorStatus',
            formItemOtherClass: [],
            itemOtherClass: ['error-color-tip', 'password-tip'],
            maxlength: '30',
            type: 'password',
            model: 'newPassword',
            placeholder: '',
            showMessage: false
        },
        {
            prop: 'confirmPassword',
            style: '',
            label: language.t('userCenter.confirmPassword'),
            class: 'set-confirm-password',
            flag: 'confirmPasswordFlag',
            formItemOtherClass: ['set-confirmpassword-empty', ''],
            itemOtherClass: ['', ''],
            maxlength: '30',
            type: 'password',
            model: 'confirmPassword',
            placeholder: '',
            showMessage: true
        }
    ]

    constructor (regExpCheck: any) {
        this.regExpCheck = regExpCheck;
        this.rules['oldPassword'][0].validator = regExpCheck.oldPasswordCheck;
        this.rules['newPassword'][0].validator = regExpCheck.newPasswordCheck;
        this.rules['confirmPassword'][0].validator = regExpCheck.confirmPasswordCheck;
    };

    formClass = (item: {prop: string; flag: string; formItemOtherClass: string[]}, type: string) => {
        return !Object.is(item.prop, type) ? this.regExpCheck[item.flag].value ? item.formItemOtherClass[0] : item.formItemOtherClass[1] : '';
    }

    customTipClass = (item: {flag: string; itemOtherClass: string[]}) => {
        return this.regExpCheck[item.flag].value ? item.itemOtherClass[0] : item.itemOtherClass[1];
    }

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    reviseClick = async () => {
        this.isLoading.value = true;
        await nextTick(() => {
            this.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.revisePassword();
                }
                else {
                    this.isLoading.value = false;
                }
            });
        });
    };

    revisePasswordIpt = (val: any) => {
        this.regExpCheck.newPassword.value = val;
    };

    revisePassword = () => {
        revisePasswordAPI(
            {
                oldPassword: encrypt(this.ruleForm.oldPassword),
                password: encrypt(this.ruleForm.newPassword)
            }
        )
        .then(({data} : {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                this.isLoading.value = false;
            }
        })
        .catch(({message}: {message: string}) => {
            customTip('error', message, 800, () => this.isLoading.value = false)
        });
    }
};

export default RevisePassword;
