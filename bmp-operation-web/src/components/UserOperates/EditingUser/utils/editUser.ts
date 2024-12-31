import {CurrencyType} from '@utils/publicType';
import {deepCopy, msgTip, methodsTotal} from 'utils/index.ts';
import {language, locationItem, CurrentInstance} from 'utils/publicClass.ts';

const editUserOperate = (formRules: any, emitValue: any) => {
    // loading态
    const isLoading: Ref<boolean> = ref<boolean>(false);
    const proxy = new CurrentInstance().proxy;

    onMounted(() => {
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && determineClick();
        };
        onUnmounted(() => {
            document.onkeyup = () => {return;}
        });
    });

    /**
     * 确定按钮弹窗
    */
    const determineClick = async (): Promise<void> => {
        // 判断输入项是否符合条件
        await nextTick(() => {
            formRules.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    isLoading.value = true;
                    requestEditUser();
                }
                else {
                    if (!formRules.ruleForm.phoneNumber) {
                        formRules.phoneFlag.value = true;
                    }
                }
            });
        });
    };

    /**
     * 请求编辑用户验证接口，成功后把事件回传，关闭弹窗
    */
    const requestEditUser = async () => {
        const copyRuleForm = deepCopy(formRules.ruleForm as never);
        copyRuleForm.phonePrefix = formRules.ruleForm.type;
        try {
            const editApi = await proxy.$userApi.editUserAPI({...copyRuleForm, language: locationItem.getLocationItem});
            const status: boolean = (editApi?.data?.result?.success)?? false;
            status && dealWithResponse(copyRuleForm);
        }
        finally {
            dealWithFinally();
        }
    };

    const dealWithResponse = (copyRuleForm: CurrencyType) => {
        methodsTotal.sendMsg('update-user-info', {...copyRuleForm});
        msgTip.success(language.t('operate.success'));
    };

    const dealWithFinally = () => {
        isLoading.value = false;
        emitValue('determine-click');
        cancelClick();
    };

    /**
     * 取消按钮点击操作
    */
    const cancelClick = (): void => {
        // 清空输入内容、异常提示状态
        formRules.ruleFormRef.value!.resetFields();
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

    return {
        isLoading,
        determineClick,
        cancelClick
    }
};

export default editUserOperate;
