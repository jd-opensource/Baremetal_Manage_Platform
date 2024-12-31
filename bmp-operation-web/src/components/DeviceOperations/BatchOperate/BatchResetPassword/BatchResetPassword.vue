<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'batch-reset-password'"
        :is-loading="batchResetPasswordOpt.isLoading"
        :header-title="$t('batchOperate.resetPassword.header.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="batchResetPasswordOpt.cancelClick"
        @determine-click="batchResetPasswordOpt.determineClick"
    >
        <div class="batch-reset-password-content">
            <p class="warnning-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 提示 -->
                <div class="tip-count">
                    <p class="reset-system-tip" v-html="$t('batchOperate.resetPassword.header.tip')"></p>
                </div>
            </p>
            <p class="tip-title">
                {{$t('batchOperate.resetPassword.header.resetPassword1')}}
                <span>{{checkArr.length}}</span>
                {{$t('batchOperate.resetPassword.header.resetPassword2')}}
            </p>
            <!-- 表格数据 -->
            <div class="currency-count-table">
                <ui-config-provider :locale="uiLocale.locale">
                    <ui-table
                        border
                        style="width: 100%"
                        :max-height="180"
                        :class="[tableClass(checkArr, false)]"
                        :data="checkArr"
                    >
                        <!-- 实例名称 -->
                        <ui-table-column
                            prop="instanceName"
                            align="center"
                            min-width="120"
                            :label="$t('batchOperate.resetPassword.table.label.instanceName')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.instanceName}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.instanceName)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.instanceName)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 实例ID -->
                        <ui-table-column
                            prop="instanceId"
                            min-width="120"
                            align="center"
                            :label="$t('batchOperate.resetPassword.table.label.instanceId')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.instanceId}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.instanceId)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.instanceId)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 带外IP -->
                        <ui-table-column
                            prop="iloIp"
                            align="center"
                            min-width="100"
                            :label="$t('batchOperate.resetPassword.table.label.outOfBandIP')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.iloIp)}}
                            </template>
                        </ui-table-column>
                        <!-- 内网IPv4 -->
                        <ui-table-column
                            prop="privateIpv4"
                            align="center"
                            min-width="120"
                            :label="$t('batchOperate.resetPassword.table.label.intranceIPv4')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.privateIpv4)}}
                            </template>
                        </ui-table-column>
                    </ui-table>
                </ui-config-provider>
            </div>
            <div class="rule-form">
                <ui-form
                    label-position="left"
                    :label-width="setDiffClass('130px', '100px')"
                    :model="formRulesOperate.ruleForm"
                    :rules="formRulesOperate.rules"
                    @submit.native.prevent
                    @rule-form-ref="formRulesOperate.getFormRef"
                >
                    <!-- 新密码 -->
                    <ui-form-item
                        prop="newPassword"
                        class="custom-password"
                        :label="$t('batchOperate.resetPassword.ipt.label.newPassword')"
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
                        :label="$t('batchOperate.resetPassword.ipt.label.confirmPassword')"
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
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {hasShowTooltip, tableClass, setDiffClass} from 'utils/index.ts';
import {uiLocale} from 'utils/publicClass.ts';
import RulesCheck from './utils/rulesCheck';
import FormRulesOperate from './utils/formRules';
import BatchResetPasswordOpt from './utils/batchResetPasswordOpt'

// props 类
interface PropsType {
    diaLog: boolean;
    checkArr: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const rulesCheck = new RulesCheck();

const formRulesOperate = new FormRulesOperate(rulesCheck);

const batchResetPasswordOpt = new BatchResetPasswordOpt(formRulesOperate, props.checkArr, emitValue);
</script>
<style lang="scss">
@import './batchResetPassword.scss';
</style>