import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {resetSystemAPI} from 'api/device/request.ts';
import {SuccessType} from '@utils/publicType';
import {FormRulesOptType, RulesCheckType} from '../type';

const resetSystemOpt = (formRulesOpt: FormRulesOptType, rulesCheck: RulesCheckType, emitValue: Function) => {
    const isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    const sureClick = async () => {
        await nextTick(() => {
            formRulesOpt.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    rulesCheck.hasKeyFlag.value = (formRulesOpt.ruleForm.setType === 'sshKey' && !formRulesOpt.ruleForm.sshKeyId);
                    if (rulesCheck.hasKeyFlag.value) return;
                    isLoading.value = true;
                    requestResetSystem();
                }
            });
        });
    };

    const requestResetSystem = () => {
        resetSystemAPI(
            {
                ...formRulesOpt.ruleForm
            }
        )
        .then(({data}: {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            isLoading.value = false;
            cancelClick();
            emitValue('determine-click');
        })
    };
  
    /**
     * 取消按钮点击操作
    */
    const cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

    return {
        isLoading,
        sureClick,
        cancelClick
    }
};

export default resetSystemOpt;
