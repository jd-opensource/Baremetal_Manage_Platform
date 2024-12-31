import {msgTip} from 'utils/index.ts';
import {FormType} from '../typeManagement';
import {CurrencyType5} from '@utils/publicType';
import {language, CurrentInstance} from 'utils/publicClass.ts';
import IdcStaticData from 'staticData/idc/index.ts'; // 机房等级数据

/**
 * 编辑机房操作
*/
class EditIdcOperate {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    // 机房等级
    machineRoomGrade: CurrencyType5[] = IdcStaticData.machineRoomGradeData;
    formRuleOperate: any;
    emitValue: any;
    instanceProxy = new CurrentInstance().proxy;

    constructor (formRuleOperate: any, emitValue: any) {
        this.emitValue = emitValue;
        this.formRuleOperate = formRuleOperate;
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };
        onUnmounted(() => document.onkeyup = (() => {return}));
    };

    /**
     * 确定按钮弹窗
    */
    determineClick = async (): Promise<void> => {
        // 判断输入项是否符合条件
        await nextTick(() => {
            this.formRuleOperate.validate('customGrade', this.formRuleOperate.customGradeFlag)
            this.formRuleOperate.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.isLoading.value = true;
                    this.#requestVerificationInterface();
                }
            });
        });
    };

    /**
     * 请求编辑机房验证接口，成功后把事件回传，关闭弹窗
    */
    #requestVerificationInterface = async () => {
        try {
            const {level, customGrade}: FormType = this.formRuleOperate.ruleForm;
            const newLevel =  IdcStaticData.otherData.includes(level!) ? customGrade : level;
            const editApi = await this.instanceProxy.$idcApi.editIdcAPI({...this.formRuleOperate.ruleForm, level: newLevel})
            const status = (editApi?.data?.result?.success)?? false;
            status && this.#dealWithResponse();
        }
        finally {
            this.isLoading.value = false;
        }
    };

    #dealWithResponse = () => {
        msgTip.success(language.t('operate.success'));
        this.cancelClick();
        this.emitValue('determine-click');
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = () => {
        // 清空输入内容、异常提示状态
        this.formRuleOperate.ruleFormRef.value!.resetFields();
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
    };
};

export default EditIdcOperate;
