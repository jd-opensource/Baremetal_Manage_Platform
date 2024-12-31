<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'reset-password'"
        :is-loading="resetPassword.isLoading"
        :header-title="$t('resetPassword.header.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="resetPassword.cancelClick"
        @determine-click="resetPassword.determineClick"
    >
        <div class="reset-password-content">
            <p class="warnning-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 提示 -->
                <span>{{$t('resetPassword.header.tip')}}</span>
            </p>
            <p class="tip-title">
                {{$t('resetPassword.header.resetPassword1')}}
                <span>{{formRulesOperate.reactiveArr.tableData.length}}</span>
                {{$t('resetPassword.header.resetPassword2')}}
            </p>
            <!-- 表格数据 -->
            <device-reset-password-table :form-rules-operate="formRulesOperate"/>
            <device-reset-password-form :form-rules-operate="formRulesOperate"/>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import propsOperate from './utils/propsOperate';
import FormRulesOperate from './utils/formRules';
import ResetPassword from './utils/resetPassword';

// props 类
interface PropsType {
    diaLog: boolean;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

propsOperate((item) => formRulesOperate.setData(item))

// rulesCheck
const formRulesOperate = new FormRulesOperate();

const resetPassword = new ResetPassword(formRulesOperate, emitValue);
</script>
<style lang="scss">
@import './resetPassword.scss';
</style>
