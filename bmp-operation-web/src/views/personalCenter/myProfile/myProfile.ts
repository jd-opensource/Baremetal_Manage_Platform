import {RuleFormRefType} from '@utils/publicType';
import {language, UserCheck, locationItem} from 'utils/publicClass.ts';
import {userInfoAPI, timeOneAPI, setUserInfoAPI} from 'api/userCenter/request.ts';
import {msgTip, customTip} from 'utils/index.ts';
import UserStaticData from 'staticData/user/index.ts'; // 手机号类型数据

class FormRulesOperate {
    isLoading: Ref<boolean> = ref<boolean>(false);
    myInfoLoading: Ref<boolean> = ref<boolean>(true);
    phoneFlag: Ref<boolean> = ref<boolean>(false);
    emailFlag: Ref<boolean> = ref<boolean>(false);
    reactiveArr: any = reactive({
        languageData: [
            language.t('userCenter.language.zh'),
            language.t('userCenter.language.en')
        ],
        timezonesData: []
    });
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单项
    ruleForm: any = reactive<any>({
        phonePrefix: '',
        type: '',
        language: '',
        timezonesVal: '',
        roleName: '--',
        email: '',
        userName: '--',
        phoneNumber: ''
    });
    val: Ref<string> = ref<string>('');
    timeVal: string = '';
    // 规则
    rules: any = reactive<any>({
        phoneNumber: [ // 手机号
            {
                required: true,
                trigger: 'blur',
                // validator: new UserCheck(this.ruleForm.phonePrefix).phoneChcek
            }
        ],
        email: [ // 邮箱
            {
                required: true,
                trigger: 'blur',
                validator: new UserCheck().emailChcek
            }
        ]
    });

    constructor () {
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.saveClick();
        };
        setTimeout(() => {
            this.getUserName();
        }, 200)
        // if (tabsOperate.activeName.value === 'myProfile') this.getUserName();
    };

    getUserName = () => {
        userInfoAPI({})
        .then(({data}: any) => {
            if (data?.result && Object.keys(data.result)?.length) {
                const {roleName, userName, email, phoneNumber, phonePrefix} = data.result;
                const languageInfo = data.result.language;
                this.ruleForm.roleName = roleName;
                this.ruleForm.email = email;
                this.ruleForm.userName = userName;
                this.ruleForm.phoneNumber = phoneNumber;
                this.ruleForm.language = languageInfo === 'en_US' ? language.t('userCenter.language.en') : language.t('userCenter.language.zh');
                this.val.value = data.result.timezone;
                localStorage.setItem('email', window.btoa(this.ruleForm.email));
                // locationItem.cookie.set('X-Jdcloud-Language', languageInfo);
                for(const key of UserStaticData.phoneType) {
                    if (key.type === phonePrefix) {
                        const value = UserStaticData.cellPhoneType.find((item: {label: string}) => item.label === key.name);
                        this.ruleForm.phonePrefix = key.name;
                        this.ruleForm.type = value.text;
                    }
                }
                this.rules.phoneNumber[0].validator = new UserCheck(this.ruleForm.phonePrefix).phoneChcek;
                this.getTimeOne();
            }
        })
        .catch(({message} : {message: string}) => {
            msgTip.error(message);
        })
        .finally(() => {
            this.myInfoLoading.value = false;
        });
    };

    getTimeOne = () => {
        timeOneAPI({})
        .then(({data} : any) => {
            if (data?.result && Object.keys(data.result?.timezones).length) {
                this.ruleForm.timezonesVal = data.result.timezones[this.val.value];
                for (const index in data.result.timezones) {
                    this.reactiveArr.timezonesData.push(
                        {
                            value: index,
                            label: data.result.timezones[index]
                        }
                    )
                }
                this.timeVal = this.val.value;
            }
        })
    };

    phonePrefixSelect = (val: string) => {
        const value = UserStaticData.cellPhoneType.find((item: {label: string}) => item.label === val);
        this.ruleForm.type = value.text;
        this.ruleForm.phonePrefix = val;
        this.ruleForm.phoneNumber = '';
        this.rules.phoneNumber[0].validator = new UserCheck(this.ruleForm.phonePrefix).phoneChcek;
    };

    timezoneChange = (val: string) => {
        this.timeVal = this.reactiveArr.timezonesData.find((item: {label: string}) => item.label === val).value;
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    validate = (type1: never | string, type2: {value: boolean}) => {
        this.ruleFormRef.value!.validateField(([type1] as never), (valid: string) => !valid ? type2.value = false : type2.value = true);
    };

    saveClick = async () => {
        this.isLoading.value = true;
        await nextTick(() => {
            this.validate('phoneNumber', this.phoneFlag);
            this.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.setUserInfo();
                }
                else {
                    this.isLoading.value = false;
                }
            });
        });
    };

    setUserInfo = () => {
        const languages = this.ruleForm.language === language.t('userCenter.language.en') ? 'en_US' : 'zh_CN';
        setUserInfoAPI(
            {
                language: languages,
                phoneNumber: this.ruleForm.phoneNumber,
                phonePrefix: this.ruleForm.type,
                timezone: this.timeVal,
                email: this.ruleForm.email,
                userName: this.ruleForm.userName,
            }
        )
        .then(({data} : {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                msgTip({
                    type: 'success',
                    message: language.t('operate.success'),
                    duration: 1000,
                    onClose: () => {
                        localStorage.setItem('email', window.btoa(this.ruleForm.email));
                        locationItem.cookie.set('X-Jdcloud-Language', languages);
                        location.reload();
                        this.isLoading.value = false;
                    }
                });
            }
        })
        .catch(({message}: {message: string}) => {
            customTip('error', message, 800, () => this.isLoading.value = false);
        });
    };
};

export default FormRulesOperate;
