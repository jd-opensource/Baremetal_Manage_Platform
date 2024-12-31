<template>
    <!-- 编辑机房验证弹窗 -->
    <custom-dia-log
        :custom-class="'edit-idc'"
        :dia-log="diaLog"
        :header-title="$t('editIdc.header.name')"
        :is-loading="editIdcOperate.isLoading"
        :src="($defInfo.imagePath('idcManage') as unknown as string)"
        @dia-log-close="editIdcOperate.cancelClick"
        @determine-click="editIdcOperate.determineClick"
    >
        <!-- 主体内容-表单操作 -->
        <edit-idc-form
            :form-rule-operate="formRuleOperate"
            :edit-idc-operate="editIdcOperate"
        />
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {FormType} from './typeManagement';
import formRule from './utils/formRules';
import IdcDetail from './utils/idcDetail';
import SetLevel from './utils/setLevel';
import EditIdcOperate from './utils/editIdc';
import rulesVerification from './utils/ruleVerification';


/**
 * props类型
*/
interface PropsType {
    diaLog: boolean;
    itemData: {[x: string]: FormType;};
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    itemData: () => {return {}}
});

// defineEmits API 来替代emit，进行事件回传
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// rules-规则-实例化
const rulesCheck = rulesVerification();

const formRuleOperate = formRule(props, rulesCheck);

new IdcDetail(props, formRuleOperate);

const editIdcOperate: EditIdcOperate = new EditIdcOperate(formRuleOperate, emitValue);

new SetLevel(props, formRuleOperate, editIdcOperate);

</script>
<style lang="scss">
@import './editIdc.scss'; 
</style>
