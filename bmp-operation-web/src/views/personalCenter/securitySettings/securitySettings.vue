<template>
    <div class="user-center-content">
        <header-info
            :type="'securitySettings'"
        />
        <div class="security-setting">
            <ui-form
                class="user-center-ruleForm form"
                :label-width="setDiffClass('130px', '70px')"
                label-position="left"
                :model="revisePassword.ruleForm"
                :rules="revisePassword.rules"
                @rule-form-ref="revisePassword.getFormRef"
            >
                <!-- 当前密码 -->
                <ui-form-item
                    v-for="(item, index) in revisePassword.formData"
                    :key="index"
                    :prop="item.prop"
                    :style="item.style"
                    :label="item.label"
                    :show-message="item.showMessage"
                    :class="[
                        item.class,
                        revisePassword.formClass(item, 'newPassword')
                    ]"
                >
                    <ui-input
                        v-model.trim="revisePassword.ruleForm[item.model]"
                        :maxlength="item.maxlength"
                        :type="item.type"
                        :placeholder="item.placeholder"
                    />
                    <p
                        v-if="Object.is(item.prop, 'newPassword')"
                        :class="[revisePassword.customTipClass(item)]"
                    >
                        {{regExpCheck.errorTip.value}}
                    </p>
                </ui-form-item>
            </ui-form>
            <ui-button
                class="user-revise"
                type="primary"
                :loading="revisePassword.isLoading.value"
                @click="revisePassword.reviseClick"
            >
                {{$t('userCenter.btn.saveChange')}}
            </ui-button>
        </div>
    </div>
 
</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts';
import RevisePassword from './revisePassword';
import RegExpCheck from './regExpReg';

const regExpCheck = new RegExpCheck();

const revisePassword = new RevisePassword(regExpCheck);

// interface PropsType {
//     regExpCheck: any;
//     revisePassword: any;
// };

// // withDefaults API 用来代替默认值，defineProps API 用来替代 props
// withDefaults(defineProps<PropsType>(), {});

</script>
<style lang="scss" scoped>
@import './index';
</style>