<template>
    <div class="user-center-content">
        <header-info
            :type="'myProfile'"
        />
        <div
            class="my-info"
            v-loading="formRulesOperate.myInfoLoading.value"
        >
            <div class="basinfo-count">
                <p class="header-title m-l">
                    <span class="line-left"></span>
                    <span class="title">
                        {{$t('userCenter.header.title1')}}
                    </span>
                </p>
                <ui-form
                    class="user-center-ruleForm"
                    label-position="left"
                    :label-width="setDiffClass('95px', '70px')"
                    :model="formRulesOperate.ruleForm"
                    :rules="formRulesOperate.rules"
                    @rule-form-ref="formRulesOperate.getFormRef"
                >
                    <div class="m-l">
                        <!-- 角色 -->
                        <ui-form-item
                            prop="roleName"
                            required
                            :label="$t('userCenter.label.role')"
                        >
                            <ui-input
                                class="user-universal"
                                :disabled="true"
                                v-model.trim="formRulesOperate.ruleForm.roleName"
                            />
                            <!-- <span class="user-role">
                                {{$filter.emptyFilter(formRulesOperate.ruleForm.roleName)}}
                            </span> -->
                        </ui-form-item>
                        <!-- 用户名 -->
                        <ui-form-item
                            prop="userName"
                            required
                            :label="$t('userCenter.label.userName')"
                        >
                            <ui-input
                                class="user-universal"
                                :disabled="true"
                                v-model.trim="formRulesOperate.ruleForm.userName"
                            />
                            <!-- <span class="user-name">
                                {{$filter.emptyFilter(formRulesOperate.ruleForm.userName)}}
                            </span> -->
                        </ui-form-item>
                        <!-- setDiffClass('set-email', 'user-count'), -->

                        <!-- 邮箱 -->
                        <ui-form-item
                            prop="email"
                            style="marginTop: 6px;"
                            :label="$t('userCenter.label.email')"
                            :class="[
                                formRulesOperate.emailFlag.value ? 'set-empty' : ''
                            ]"
                        >
                            <ui-input
                                class="user-universal"
                                maxlength="128"
                                v-model.trim="formRulesOperate.ruleForm.email"
                                :placeholder="$t('userCenter.placeholder.email')"
                                @blur="formRulesOperate.validate('email', formRulesOperate.emailFlag)"
                            />
                        </ui-form-item>
                        <!-- :class="setDiffClass('set-select-count', 'input-select-count')" -->
                        <!-- 手机 -->
                        <ui-form-item
                            prop="phonePrefix"
                            required
                            style="marginBottom: 15px"
                            class="input-select-count"
                            :label="$t('userCenter.label.phoneNumber')"
                        >
                            <ui-select 
                                class="input-select"
                                v-model="formRulesOperate.ruleForm.phonePrefix"
                                @change="formRulesOperate.phonePrefixSelect"
                            >
                                <ui-option
                                    style="fontSize: 12px"
                                    v-for="item in UserStaticData.cellPhoneType"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.label"
                                />
                            </ui-select>
                        </ui-form-item>
                        <!-- 手机号输入框 -->
                        <ui-form-item
                            prop="phoneNumber"
                            style="marginBottom: 15px;"
                            :class="[
                                'input-phone-count',
                                formRulesOperate.phoneFlag.value ? 'set-empty' : ''
                            ]"
                        >
                            <ui-input
                                maxlength="32"
                                v-model.trim="formRulesOperate.ruleForm.phoneNumber"
                                :placeholder="$t('userCenter.placeholder.phoneNumber')"
                                @blur="formRulesOperate.validate('phoneNumber', formRulesOperate.phoneFlag)"
                            />
                        </ui-form-item>
                    </div>
                    <div class="preference-count">
                        <p class="header-title" style="marginBottom: 16px">
                            <span class="line-left"></span>
                            <span class="title">
                                {{$t('userCenter.header.title2')}}
                            </span>
                        </p>
                        <!-- 语言 -->
                        <ui-form-item
                            prop="language"
                            required
                            class="language"
                            :label="$t('userCenter.label.language')"
                        >
                            <ui-select 
                                class="input-select"
                                v-model="formRulesOperate.ruleForm.language"
                                :placeholder="$t('userCenter.placeholder.language')"
                            >
                                <ui-option
                                    style="fontSize: 12px"
                                    v-for="item in formRulesOperate.reactiveArr.languageData"
                                    :key="item"
                                    :label="item"
                                    :value="item"
                                />
                            </ui-select>
                        </ui-form-item>
                        <!-- 时区 -->
                        <ui-form-item
                            prop="timezonesVal"
                            required
                            class="time-one"
                            :label="$t('userCenter.label.time')"
                        >
                            <ui-select 
                                class="input-select"
                                v-model="formRulesOperate.ruleForm.timezonesVal"
                                :no-data-text="$t('importDevice.noData.tip')"
                                :placeholder="$t('userCenter.placeholder.time')"
                                @change="formRulesOperate.timezoneChange"
                            >
                                <ui-option
                                    style="fontSize: 12px"
                                    v-for="item in formRulesOperate.reactiveArr.timezonesData"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.label"
                                />
                            </ui-select>
                        </ui-form-item>
                        <div class="line2"></div>
                        <ui-form-item>
                            <ui-button
                                class="user-save"
                                type="primary"
                                :loading="formRulesOperate.isLoading.value"
                                @click="formRulesOperate.saveClick"
                            >
                                {{$t('userCenter.btn.save')}}
                            </ui-button>
                        </ui-form-item>
                    </div>
                </ui-form>
            </div>
            <div class="line"></div>
        </div>
    </div>
    
</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index';
import UserStaticData from 'staticData/user/index.ts'; // 手机号类型数据
import FormRulesOperate from './myProfile'
// interface PropsType {
//     formRulesOperate: any;
// };

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
// withDefaults(defineProps<PropsType>(), {});

const formRulesOperate = new FormRulesOperate();

</script>
<style lang="scss">
@import './myProfile.scss';
</style>