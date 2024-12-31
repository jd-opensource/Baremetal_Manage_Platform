import {FormType} from '../typeManagement';
import IdcStaticData from 'staticData/idc/index.ts';

class SetLevel {
    constructor (props: any, formRuleOperate: any, editIdcOperate: any) {
        this.init(props, formRuleOperate, editIdcOperate);
    };

    init = (props: any, formRuleOperate: any, editIdcOperate: any) => {
        const level: string | undefined = formRuleOperate.ruleForm!.level;
        if (IdcStaticData.levelData.includes(level!)) {
            (formRuleOperate.ruleForm.level as string | FormType) = props.itemData?.value?.level??props.itemData?.level;
            return;
        }
        formRuleOperate.ruleForm.customGrade = formRuleOperate.ruleForm.level;
        formRuleOperate.ruleForm.level = editIdcOperate.machineRoomGrade[4].label;
    }
};

export default SetLevel;
