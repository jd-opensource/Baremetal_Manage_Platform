<template>
    <div class="message-settings-content">
        <!-- set没权限要判断不能跳入这个页面 -->
        <header-info :type="'messageSettings'"/>
        <div class="message-settings" v-loading="messageSet.pageLoading.value">
            <ui-form
                class="user-center-ruleForm form"
                :label-width="setDiffClass('207px', '180px')"
                label-position="left"
                :model="messageSet.ruleForm"
                :rules="messageSet.rules"
                @rule-form-ref="messageSet.getFormRef"
            >
                <!-- 是否关联邮箱推送 -->
                <div class="email-push">
                    <label>{{$t('messageSettings.label.emailPush')}}</label>
                    <!-- 提示 -->
                    <ui-tooltip placement="right">
                        <template #content>
                            <div v-html="$t('messageSettings.tips.pushTip')">
                            </div>
                        </template>
                        <img
                            alt=""
                            class="tooltip-img"
                            :src="($defInfo.imagePath('toolTip') as unknown as string)"
                        />
                    </ui-tooltip>
                    <ui-switch
                        v-model="messageSet.ruleForm.emailPush"
                        :loading="pushEmailOpt.switchLoading.value"
                        :before-change="pushEmailOpt.switchBeforeChange"
                    />
                </div>
                <ui-form-item
                    prop="ipAddress"
                    style="marginBottom: 20px"
                    :label="$t('messageSettings.label.ipAddress')"
                    :class="[
                        regExpCheck.ipAddressFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <ui-input
                        maxlength="30"
                        v-model.trim="messageSet.ruleForm.ipAddress"
                        :disabled="!messageSet.ruleForm.emailPush"
                        :placeholder="$t('messageSettings.placeholder.ipAddress')"
                    />
                </ui-form-item>
                <ui-form-item
                    prop="port"
                    style="marginBottom: 20px"
                    :label="$t('messageSettings.label.port')"
                    :class="[
                        regExpCheck.portFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <ui-input
                        maxlength="30"
                        v-model.trim="messageSet.ruleForm.port"
                        :disabled="!messageSet.ruleForm.emailPush"
                        :placeholder="$t('messageSettings.placeholder.port')"
                    />
                </ui-form-item>
                <ui-form-item
                    prop="emailAddress"
                    style="marginBottom: 20px"
                    :label="$t('messageSettings.label.emailAddress')"
                    :class="[
                        regExpCheck.emailAddressFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <ui-input
                        maxlength="30"
                        v-model.trim="messageSet.ruleForm.emailAddress"
                        :disabled="!messageSet.ruleForm.emailPush"
                        :placeholder="$t('messageSettings.placeholder.emailAddress')"
                    />
                </ui-form-item>
                <ui-form-item
                    prop="emailPassword"
                    style="marginBottom: 20px"
                    :label="$t('messageSettings.label.emailPassword')"
                    :class="[
                        regExpCheck.emailPasswordFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <ui-input
                        maxlength="30"
                        type="password"
                        v-model.trim="messageSet.ruleForm.emailPassword"
                        :disabled="!messageSet.ruleForm.emailPush"
                        :placeholder="$t('messageSettings.placeholder.emailPassword')"
                    />
                </ui-form-item>
                <ui-form-item
                    prop="testEmail"
                    style="marginBottom: 20px"
                    :label="$t('messageSettings.label.testEmail')"
                    :class="[
                        setDiffClass('en-test-email', 'test-email'),
                        regExpCheck.testEmailFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <ui-input
                        maxlength="30"
                        v-model.trim="messageSet.ruleForm.testEmail"
                        :disabled="!messageSet.ruleForm.emailPush"
                        :placeholder="$t('messageSettings.placeholder.testEmail')"
                    />
                    <ui-tooltip placement="bottom">
                        <template #content>
                            <div v-html="$t('messageSettings.tips.checkTip')">
                            </div>
                        </template>
                        <img
                            alt=""
                            class="tooltip-img"
                            :src="($defInfo.imagePath('toolTip') as unknown as string)"
                        />
                    </ui-tooltip>
                </ui-form-item>
            </ui-form>
            <p class="footer-opt">
                <ui-button
                    class="user-save"
                    type="primary"
                    :disabled="messageSet.disabledStatus.value"
                    :loading="messageSet.isLoading.value"
                    @click="messageSet.reviseClick"
                >
                    {{$t('messageSettings.footer.save')}}
                </ui-button>
                
                <span
                    class="is-pass"
                    v-if="messageSet.validateInfo.value"
                >
                    {{messageSet.validateInfo.value}}
                </span>
            </p>
        </div>
        <push-email-info :push-email-opt="pushEmailOpt"></push-email-info>
    </div>
</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts';
import MessageSetOpt from './utils/messageSet';
import RegExpCheck from './utils/regExpReg';
import PushEmailOpt from './pushEmail/pushEmail';
import PushEmailInfo from './pushEmail/pushEmailInfo.vue'

const regExpCheck = new RegExpCheck();

const messageSet = new MessageSetOpt(regExpCheck, (status: boolean) => {
    pushEmailOpt.operate.value = status ? 'cancel' : 'push';
});

const pushEmailOpt = new PushEmailOpt(messageSet.ruleForm.emailPush, () => {
    messageSet.ruleFormRef.value?.resetFields();
    regExpCheck.resetFlag();
    messageSet.getDescribeMail()
});

</script>
<style lang="scss" scoped>
@import './index';
</style>