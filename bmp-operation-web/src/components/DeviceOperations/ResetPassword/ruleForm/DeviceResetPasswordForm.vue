<template>
    <div class="rule-form">
        <ui-form
            label-position="left"
            :label-width="setDiffClass('130px', '100px')"
            :model="formRulesOperate.ruleForm"
            :rules="rulesCheck.rules"
            @submit.native.prevent
            @rule-form-ref="formRulesOperate.getFormRef"
        >
            <!-- 新密码 -->
            <ui-form-item
                prop="newPassword"
                class="custom-password"
                :label="$t('resetPassword.ipt.label.newPassword')"
                :show-message="false"
            >
                <ui-input
                    maxlength="30"
                    type="password"
                    class="custom-ipt"
                    placeholder=""
                    v-model.trim="formRulesOperate.ruleForm.newPassword"
                />
                <p :class="[rulesCheck.colorStatus.value ? 'error-color-tip' : 'password-tip']">
                    {{rulesCheck.errorTip.value}}
                </p>
            </ui-form-item>
            <!-- 确认密码 -->
            <ui-form-item
                prop="confirmPassword"
                :label="$t('resetPassword.ipt.label.confirmPassword')"
                :class="[
                    rulesCheck.confirmPasswordFlag.value ? 'set-empty' : ''
                ]"
            >
                <ui-input
                    maxlength="30"
                    class="custom-ipt"
                    type="password"
                    v-model.trim="formRulesOperate.ruleForm.confirmPassword"
                    :placeholder="''"
                />
            </ui-form-item>
        </ui-form>
    </div>
</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts';
import RulesCheck from '../utils/rulesCheck';
// props 类
interface PropsType {
    formRulesOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

const rulesCheck = new RulesCheck();

</script>