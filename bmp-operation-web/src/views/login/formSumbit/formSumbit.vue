<template>
    <main class="operate-login">
        <!-- 登录标题 -->
        <p class="operate-login-title">
            {{$t('login.formSubmit.title')}}
        </p>
        <!-- 表单输入提交 -->
        <ui-form
            class="operate-login-ruleForm"
            :model="ruleFormCheck.ruleForm"
            :rules="ruleFormCheck.rules"
            @rule-form-ref="ruleFormCheck.getFormRef"
        >
            <!-- 登录用户名 -->
            <ui-form-item
                prop="loginUserName"
                :show-message="false"
            >
                <div class="operate-content-ipt">
                    <!-- 用户icon -->
                    <img
                        class="operate-content-ipt-user-img"
                        :src="($defInfo.imagePath('userLogin') as unknown as string)"
                    />
                    <!-- 用户名输入框 -->
                    <ui-input
                        v-model.trim="ruleFormCheck.ruleForm.loginUserName"
                        maxlength="64"
                        :placeholder="$t('login.formSubmit.placeholder.userName')"
                    >
                    </ui-input>
                    <!-- 清空icon -->
                    <img
                        class="operate-content-ipt-clear"
                        alt=""
                        v-if="ruleFormCheck.ruleForm.loginUserName.length"
                        :src="($defInfo.imagePath('loginClear') as unknown as string)"
                        @click="ruleFormCheck.clearInput('loginUserName')"
                    />
                </div>
            </ui-form-item>
            <!-- 登录密码 -->
            <ui-form-item prop="loginPassword">
                <div class="operate-content-ipt">
                    <!-- open eye icon -->
                    <img
                        class="operate-content-ipt-open-eye"
                        v-if="ruleFormCheck.passwordType.value !== 'password'"
                        :src="($defInfo.imagePath('openEye') as unknown as string)"
                        @click="ruleFormCheck.passwordStatusChange('password')"
                    />
                    <!-- close eye icon -->
                    <img
                        class="operate-content-ipt-close-eye"
                        v-else
                        :src="($defInfo.imagePath('closeEye') as unknown as string)"
                        @click="ruleFormCheck.passwordStatusChange('text')"
                    />
                    <!-- 密码输入框 -->
                    <ui-input
                        v-model.trim="ruleFormCheck.ruleForm.loginPassword"
                        maxlength="64"
                        :type="ruleFormCheck.passwordType.value"
                        :placeholder="$t('login.formSubmit.placeholder.passWord')"
                    >
                    </ui-input>
                    <!-- 清空icon -->
                    <img
                        class="operate-content-ipt-clear"
                        alt=""
                        v-if="ruleFormCheck.ruleForm.loginPassword.length"
                        :src="($defInfo.imagePath('loginClear') as unknown as string)"
                        @click="ruleFormCheck.clearInput('loginPassword')"
                    />
                </div>
            </ui-form-item>
            <!-- 异常提示 -->
            <div
                class="error-tooltip"
                v-if="loginOperate.errMessage.value"
            >
                {{loginOperate.errMessage.value}}
            </div>
        </ui-form>
        <!-- 登录按钮 -->
        <ui-button
            class="operate-login-button"
            type="primary"
            :loading="loginOperate.isLoading.value"
            @click="loginOperate.loginOperate"
        >
            {{$t('login.formSubmit.login')}}
        </ui-button>
    </main>
</template>
<script setup lang="ts">
import RuleFormCheck from './ruleFormCheck';
import LoginOperate from './loginOperate';
import {RouterOperate} from 'utils/publicClass.ts';

const router = new RouterOperate();

const ruleFormCheck = new RuleFormCheck();
const loginOperate = new LoginOperate(ruleFormCheck, router.router);

onUnmounted(() => {
    document.onkeyup = (() => {return});
    loginOperate.loadingKeyDownStatus();
})
</script>
