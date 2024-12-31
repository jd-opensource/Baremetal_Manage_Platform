/**
 * @file
 * @author
*/

import {RuleFormRefType} from '@utils/publicType';
import {inspectuserPurview, CurrentInstance, language} from 'utils/publicClass.ts';
import {useRouter} from 'vue-router';
import {methodsTotal, msgTip} from 'utils/index.ts';

type RuleFormType = {
    emailPush: boolean;
    ipAddress: string;
    port: string;
    emailAddress: string;
    emailPassword: string;
    testEmail: string;
}

class MessageSetOpt {
    pageLoading: Ref<boolean> = ref<boolean>(true);
    isShowLoading: Ref<boolean> = ref<boolean>(false);
    isLoading: Ref<boolean> = ref<boolean>(false);
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单项
    ruleForm = reactive<RuleFormType>({
        emailPush: false,
        ipAddress: '',
        port: '',
        emailAddress: '',
        emailPassword: '',
        testEmail: ''
    });
    disabledStatus: Ref<boolean> = ref<boolean>(true);
    router = useRouter();
    // 规则
    rules = reactive({
        ipAddress: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        port: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        emailAddress: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        emailPassword: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        testEmail: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ]
    });
    validateInfo: Ref<string> = ref<string>('');
    regExpCheck: any;
    instanceProxy = new CurrentInstance().proxy;
    fn;

    constructor (regExpCheck: any, fn: Function) {
        this.fn = fn;
        const path: string = this.instanceProxy.$defInfo.routerPath('idcList');
        inspectuserPurview.inspectuserPurview()
        .then(() => {
            this.regExpCheck = regExpCheck;
            this.rules['ipAddress'][0].validator = regExpCheck.ipAddressCheck;
            this.rules['port'][0].validator = regExpCheck.portCheck;
            this.rules['emailAddress'][0].validator = regExpCheck.emailAddressCheck;
            this.rules['emailPassword'][0].validator = regExpCheck.emailPasswordCheck;
            this.rules['testEmail'][0].validator = regExpCheck.testEmailCheck;
            this.watchForm();
            this.getDescribeMail();
        })
        .catch(() => {
            this.router.push(path);
        });
    };

    getDescribeMail = async () => {
        this.pageLoading.value = true;
        try {
            const res = await this.instanceProxy.$messageApi.describeMailAPI({});
            const {result} = methodsTotal.lineConverting(res.data);
            const {isPush, serverPort, adminAddr, adminPass, serverAddr, isPass, to} = result;
            this.ruleForm.emailPush = this.#setPush(isPush);
            this.fn(this.ruleForm.emailPush);
            this.ruleForm.port = serverPort;
            this.ruleForm.emailAddress = adminAddr;
            this.ruleForm.emailPassword = adminPass;
            this.ruleForm.ipAddress = serverAddr;
            this.validateInfo.value = this.#setValidate(isPass) as string;
            this.ruleForm.testEmail = to;
        }
        finally {
            this.pageLoading.value = false;
        }
    }

    #setPush = (type: string): boolean => {
        const obj = {
            '0': false,
            '1': true
        };
        return (obj[`${type}` as keyof typeof obj] as boolean);
    }
    #setValidate = (type: string) => {
        const info = new Map([
            ['0', ''],
            ['1', language.t('messageSettings.footer.success')]
        ]);
        return info.get(type);
    }

    watchForm = () => {
        watch(() => this.ruleForm, (newValue) => {
            const val = Object.values(newValue).some(item => item === '');
            this.disabledStatus.value = newValue.emailPush ? val : true;
        }, {
            deep: true
        })
    }

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    reviseClick = async () => {
        if (this.disabledStatus.value || this.isShowLoading.value) return;
        await nextTick(() => {
            this.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.isLoading.value = true;
                    this.isShowLoading.value = true;
                    this.dialMail();
                }
            });
        });
    };

    dialMail = async () => {
        const {emailPush, ipAddress, port, emailAddress, emailPassword, testEmail} = this.ruleForm;
        const params = methodsTotal.toLine({
            serverAddr: ipAddress,
            serverPort: port,
            adminAddr: emailAddress,
            adminPass: emailPassword,
            isPush: emailPush ? '1' : '0',
            to: testEmail
        })
        try {
            const res = await this.instanceProxy.$messageApi.dialMailAPI(params);
            if (res?.data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                this.getDescribeMail();
            }
        }
        finally {
            this.isShowLoading.value = false;
            this.isLoading.value = false;
        }
    }
};

export default MessageSetOpt;
