import {deepCopy} from 'utils/index.ts';
import {RuleFormRefType} from '@utils/publicType';
import {AddUserFormRulesType} from '../typeManagement';
import {language, locationItem, CurrentInstance} from 'utils/publicClass.ts';
import UserStaticData from 'staticData/user/index.ts';

interface RefType extends AddUserFormRulesType {
    ruleFormRef: Ref<RuleFormRefType | undefined>;
};

class AddUserOperate {
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(false);
    formRulesOperate: RefType;
    emitValue: (arg0: string, arg1?: boolean) => () => {};
    msgTip: {success(arg0: string): void};
    proxy = new CurrentInstance().proxy;

    constructor (formRulesOperate: RefType, emitValue: (arg0: string, arg1?: boolean) => () => {}, msgTip: {success(arg0: string): void}) {
        this.formRulesOperate = formRulesOperate;
        this.emitValue = emitValue;
        this.msgTip = msgTip;
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        // 判断输入项是否符合条件
        await nextTick(() => {
            this.formRulesOperate.validate('phoneNumber', this.formRulesOperate.phoneFlag);
            this.formRulesOperate.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.requestAddUser();
                }
                else {
                    this.isLoading.value = false;
                }
            });
        });
    };

    /**
     * 请求添加用户验证接口，成功后把事件回传，关闭弹窗
    */
    requestAddUser = async () => {
        let item = deepCopy(this.formRulesOperate.ruleForm as any);
        const type = UserStaticData.phoneType.find((item: any) => item.name === this.formRulesOperate.ruleForm.phonePrefix).type;
        item.phonePrefix = type;
        try {
            const addUserApi = await this.proxy.$userApi.addUserAPI({...item, language: locationItem.getLocationItem});
            if (addUserApi?.data?.result?.userId?.length) {
                this.cancelClick();
                this.msgTip.success(language.t('operate.success'));
                this.emitValue('determine-click');
            }
        }
        finally {
            this.isLoading.value = false;
            sessionStorage?.removeItem('roleId');
        }
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 清空输入内容、异常提示状态
        this.formRulesOperate.ruleFormRef.value!.resetFields();
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
    };
};

export default AddUserOperate;