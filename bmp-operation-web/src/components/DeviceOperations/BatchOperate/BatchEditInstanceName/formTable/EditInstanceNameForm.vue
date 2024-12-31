<template>
    <div class="instance-name-ruleform">
        <ui-form
            label-position="left"
            :model="formOpt.ruleForm"
            :rules="rulesCheck.rules"
            @submit.native.prevent
            @rule-form-ref="formOpt.getFormRef"
        >
            <!-- 实例名称 -->
            <ui-form-item
                prop="instanceName"
                :label="$t('batchOperate.instanceName.label.newInstanceName')"
                :show-message="false"
            >
                <ui-input
                    type="text"
                    v-model.trim="formOpt.ruleForm.instanceName"
                />
                <p :class="rulesCheck.errorFlag.value ? 'error-tip' : 'ipt-tip'">
                    {{rulesCheck.iptTipVal.value}}
                </p>
            </ui-form-item>
        </ui-form>
    </div>
</template>
<script lang="ts" setup>
import RulesCheck from './rulesCheck';

interface PropsType {
    formOpt: any;
}

const rulesCheck = new RulesCheck();

withDefaults(defineProps<PropsType>(), {});

</script>