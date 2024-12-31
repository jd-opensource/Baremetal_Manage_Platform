import {deviceBatchResetPasswordAPI} from 'api/device/request.ts';
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';

class BatchResetPasswordOpt {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    formRulesOperate: any;
    emitValue: any;
    checkArr: any;

    constructor(formRulesOperate: any, checkArr: any, emitValue: any) {
        this.checkArr = checkArr;
        this.emitValue = emitValue;
        this.formRulesOperate = formRulesOperate;
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };
        onUnmounted(() => document.onkeyup = (() => {return}));
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        // 判断输入项是否符合条件
        await nextTick(() => {
            this.formRulesOperate.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.isLoading.value = true;
                    this.#requestBatchResetPssword();
                }
            });
        });
    };

    /**
     * 请求接口，成功后把事件回传，关闭弹窗
    */
    #requestBatchResetPssword = (): string | void => {
        const {ruleForm} = this.formRulesOperate;
        deviceBatchResetPasswordAPI(
            {
                instanceIds: this.checkArr.map((item: {instanceId: string}) => item.instanceId),
                password: ruleForm.newPassword
            }
        )
        .then(({data}: {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.emitValue('determine-click');
            this.cancelClick();
            this.isLoading.value = false;
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
    };
};

export default BatchResetPasswordOpt;
