<template>
    <!-- 添加用户验证操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'add-user'"
        :is-loading="addUserOperate.isLoading"
        :header-title="$t('addUser.header.addUser')"
        :src="($defInfo.imagePath('user') as unknown as string)"
        @dia-log-close="addUserOperate.cancelClick"
        @determine-click="addUserOperate.determineClick"
    >
        <!-- 主体内容-表单操作 -->
        <ui-form
            class="add-user-ruleForm"
            label-width="80px"
            label-position="left"
            :model="formRules.ruleForm"
            :rules="formRules.rules"
            @rule-form-ref="formRules.getFormRef"
        >
            <!-- 角色 -->
            <ui-form-item
                prop="role"
                class="add-user-ruleForm-role"
                :label="$t('addUser.label.role')"
            >
                <ui-radio-group v-model="formRules.ruleForm.roleId">
                    <!-- 平台拥有者 -->
                    <!-- role-admin-uuid-01 -->
                    <ui-radio
                        v-if="!hasAdmin"
                        :label="$defInfo.purview('admin')"
                    >
                        {{$t('addUser.role.owner')}}
                    </ui-radio>
                    <!-- 用户 -->
                    <!-- role-user-uuid-01 -->
                    <ui-radio :label="$defInfo.purview('user')">
                        {{$t('addUser.role.user')}}
                    </ui-radio>
                    <!-- 运营人员 -->
                    <!-- role-operator-uuid-01 -->
                    <ui-radio
                        style="marginRight: 0"
                        :label="$defInfo.purview('operator')"
                    >
                    {{$t('addUser.role.operator')}}
                    </ui-radio>
                </ui-radio-group>
                <!-- 提示 -->
                <ui-tooltip placement="right">
                    <template #content>
                        <div v-html="$t('addUser.tooltip.count')"></div>
                    </template>
                    <img
                        alt=""
                        class="tooltip-img"
                        :src="($defInfo.imagePath('toolTip') as unknown as string)"
                    />
                </ui-tooltip>
            </ui-form-item>
            <!-- 用户名 -->
            <ui-form-item
                prop="userName"
                class="user-name-c"
                :label="$t('addUser.label.userName')"
                :show-message="false"
            >
                <ui-input
                    type="text"
                    v-model.trim="formRules.ruleForm.userName"
                    :disabled="regExpCheck.isShowLoading.value"
                    :placeholder="$t('addUser.placeholder.userName')"
                />
                <ui-icon
                    class="is-loading"
                    v-if="regExpCheck.isShowLoading.value"
                >
                    <Loading />
                </ui-icon>
                <span
                    :class="[regExpCheck.userNameTip.value ? 'custom-tip' : 'default-tip']"
                >
                    {{regExpCheck.userNameTip.value || $t('addUser.default.userName')}}
                </span>
            </ui-form-item>
            <!-- 密码 -->
            <ui-form-item
                prop="password"
                :class="[
                    regExpCheck.heightFlag.value ? 'height-user-password' : 'user-password',
                    regExpCheck.passwordFlag.value ? 'set-empty' : ''
                ]"
                :label="$t('addUser.label.password')"
            >
                <!-- 密码是否明文显示 -->
                <div :class="setDiffClass('set-add-user-ipt', 'add-user-ipt')">
                    <!-- open-eye -->
                    <img
                        class="add-user-ipt-open-eye"
                        v-if="formRules.passwordType.value !== 'password'"
                        :src="($defInfo.imagePath('openEye') as unknown as string)"
                        @click="formRules.passwordStatusChange('password')"
                    />
                    <!-- close-eye -->
                    <img
                        class="add-user-ipt-close-eye"
                        v-else
                        :src="($defInfo.imagePath('closeEye') as unknown as string)"
                        @click="formRules.passwordStatusChange('text')"
                    />
                    <!-- 密码输入框 -->
                    <ui-input
                        v-model.trim="formRules.ruleForm.password"
                        :type="formRules.passwordType.value"
                        :placeholder="$t('addUser.placeholder.password')"
                    >
                    </ui-input>
                    <!-- clear图片 -->
                    <img
                        class="add-user-ipt-clear"
                        alt=""
                        v-if="formRules.ruleForm.password.length"
                        :src="($defInfo.imagePath('loginClear') as unknown as string)"
                        @click="formRules.clearInput('password')"
                    />
                </div>
            </ui-form-item>
            <!-- 描述 -->
            <ui-form-item
                prop="description"
                class="add-user-ruleForm-desc"
                :label="$t('addUser.label.desc')"
            >
                <ui-input
                    type="textarea"
                    maxlength="256"
                    show-word-limit
                    v-model.trim="formRules.ruleForm.description"
                    :placeholder="$t('addUser.placeholder.desc')"
                >
                </ui-input>
            </ui-form-item>
            <!-- 手机 -->
            <ui-form-item
                prop="phonePrefix"
                required
                :class="setDiffClass('set-select-count', 'input-select-count')"
                :label="$t('addUser.label.cellPhone')"
            >
                <ui-select
                    class="input-select"
                    v-model="formRules.ruleForm.phonePrefix"
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
                style="marginBottom: 20px;"
                :class="[
                    setDiffClass('set-ipt-phone', 'input-phone-count'),
                    formRules.phoneFlag.value ? 'set-empty' : ''
                ]"
            >
                <ui-input
                    v-model.trim="formRules.ruleForm.phoneNumber"
                    :placeholder="$t('addUser.placeholder.cellPhone')"
                    @blur="formRules.validate('phoneNumber', formRules.phoneFlag)"
                />
            </ui-form-item>
            <!-- 邮箱 -->
            <ui-form-item
                prop="email"
                style="marginTop: 6px;"
                :class="[formRules.emailFlag ? 'set-empty-email' : 'user-count']"
                :label="$t('addUser.label.email')"
            >
                <ui-input
                    class="user-mailbox"
                    v-model.trim="formRules.ruleForm.email"
                    :placeholder="$t('addUser.placeholder.email')"
                    @blur="formRules.validate('email', formRules.emailFlag)"
                />
            </ui-form-item>
        </ui-form>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {setDiffClass, msgTip} from 'utils/index.ts'; // 设置不同class message提示
import AddUserOperate from './utils/addUser';
import RegExpCheck from './utils/regExpCheck';
import FormRulesOperate from './utils/formRules';
import UserStaticData from 'staticData/user/index.ts'; // 手机号类型数据

// props 类
interface PropsType {diaLog: boolean; hasAdmin?: boolean;};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: Required<PropsType> = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    hasAdmin: true
});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const regExpCheck = new RegExpCheck();
const formRules = new FormRulesOperate(props.hasAdmin, regExpCheck);
const addUserOperate = new AddUserOperate(formRules, emitValue, msgTip);

</script>
<style lang="scss">
@import './addUser.scss';
</style>
