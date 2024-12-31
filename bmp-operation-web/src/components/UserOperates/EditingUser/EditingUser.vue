<template>
    <!-- 编辑用户验证操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'editing-user'"
        :is-loading="editUser.isLoading"
        :header-title="$t('editUser.header.editUser')"
        :src="($defInfo.imagePath('user') as unknown as string)"
        @dia-log-close="editUser.cancelClick"
        @determine-click="editUser.determineClick"
    >
        <!-- 主体内容-表单操作 -->
        <ui-form
            class="editing-user-ruleForm"
            label-width="80px"
            label-position="left"
            :model="formRules.ruleForm"
            :rules="formRules.rules"
            @rule-form-ref="formRules.getFormRef"
        >
            <!-- 角色 -->
            <ui-form-item
                style="marginBottom: 5px;"
                prop="roleName"
                required
                :label="$t('editUser.label.role')"
            >
                <span class="user-roles">
                    {{$filter.emptyFilter(formRules.ruleForm.roleName)}}
                </span>
            </ui-form-item>
            <!-- 用户名 -->
            <ui-form-item
                style="marginBottom: 5px"
                prop="userName"
                required
                :label="$t('editUser.label.userName')"
            >
                <span class="user-names">
                    {{$filter.emptyFilter(formRules.ruleForm.userName)}}
                </span>
            </ui-form-item>
            <!-- 用户id -->
            <ui-form-item
                style="marginBottom: 7px"
                prop="userId"
                required
                :label="$t('editUser.label.userId')"
            >
                <span class="user-names">
                    {{$filter.emptyFilter(formRules.ruleForm.userId)}}
                </span>
            </ui-form-item>
            <!-- 描述 -->
            <ui-form-item
                prop="description"
                class="editing-user-ruleForm-desc"
                :label="$t('editUser.label.desc')"
            >
                <ui-input
                    type="textarea"
                    maxlength="256"
                    show-word-limit
                    v-model.trim="formRules.ruleForm.description"
                    :placeholder="$t('editUser.placeholder.desc')"
                >
                </ui-input>
            </ui-form-item>
            <!-- 手机 -->
            <ui-form-item
                prop="phonePrefix"
                required
                style="marginBottom: 15px"
                :class="setDiffClass('set-select-count', 'input-select-count')"
                :label="$t('editUser.label.cellPhone')"
            >
                <ui-select 
                    class="input-select"
                    v-model="formRules.ruleForm.phonePrefix"
                    @change="formRules.selectChange"
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
                    setDiffClass('set-ipt-phone', 'input-phone-count'),
                    formRules.phoneFlag.value ? 'set-empty' : ''
                ]"
            >
                <ui-input
                    maxlength="32"
                    v-model.trim="formRules.ruleForm.phoneNumber"
                    :placeholder="$t('editUser.placeholder.cellPhone')"
                    @blur="formRules.validate('phoneNumber', formRules.phoneFlag)"
                />
            </ui-form-item>
            <!-- 邮箱 -->
            <ui-form-item
                prop="email"
                style="marginTop: 6px;"
                :label="$t('editUser.label.email')"
                :class="[
                    setDiffClass('set-email', 'user-counts'),
                    formRules.emailFlag.value ? 'set-empty' : ''
                ]"
            >
                <ui-input
                    class="user-mailbox"
                    maxlength="128"
                    v-model.trim="formRules.ruleForm.email"
                    :placeholder="$t('editUser.placeholder.email')"
                    @blur="formRules.validate('email', formRules.emailFlag)"
                />
            </ui-form-item>
        </ui-form>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import UserStaticData from 'staticData/user/index.ts';
import {RuleFormType} from './typeManagement';
import {setDiffClass} from 'utils/index.ts';
import formRulesOperate from './utils/formRules';
import editUserOperate from './utils/editUser';

/**
 * props类
*/
interface PropsType {
    diaLog: boolean;
    prefix: string;
    listData: RuleFormType;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false, // 蒙层状态
    prefix: '086' // 手机号类型
});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const formRules = formRulesOperate(props);

const editUser = editUserOperate(formRules, emitValue);

</script>
<style lang="scss">
@import './editingUser.scss';
</style>
