<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'reset-system'"
        :header-title="$t('resetSystem.header.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        :is-loading="resetSystemOpt.isLoading"
        @dia-log-close="resetSystemOpt.cancelClick"
        @determine-click="resetSystemOpt.sureClick"
    >

        <div class="reset-system-content">
            <!-- 须知 -->
            <div class="warnning-tip">
                <!-- 提示icon -->
                <div style="display: flex; align-items: center; margin-top: 5px;">
                    <img
                        alt=""
                        :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                    />
                    <p style="width: 40px;font-size:12px;color: #ff4D4f; text-align: left;">{{$t('resetSystem.header.tip0')}}</p>
                </div>
                <!-- 提示 -->
                <div class="tip-count">
                    <p class="reset-system-tip" v-html="$t('resetSystem.header.tip')"></p>
                </div>
            </div>
            <!-- 实例信息 -->
            <reset-system-instance-info :form-rules-operate="formRulesOpt"/>
            <reset-system-form-submit
                :form-rules-operate="formRulesOpt"
                :rules-check="rulesCheck"
            />
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import * as allType from './type';
import RulesCheck from './utils/rulesCheck';
import propsOperate from './utils/propsOperate';
import FormRulesOperate from './FormSubmit/submitIpt';
import systemReset from './utils/resetSystem';

// props 类
interface PropsType {
    diaLog: boolean;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    diaLog: false, // 蒙层
});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const rulesCheck = new RulesCheck();

propsOperate((item: allType.ItemType, query: Pick<allType.QueryType, 'deviceTypeId' | 'idcId'>) => formRulesOpt.setData(item, query))

const formRulesOpt = new FormRulesOperate(rulesCheck);

const resetSystemOpt = systemReset(formRulesOpt, rulesCheck, emitValue);
</script>
<style lang="scss">
@import './resetSystem.scss';
</style>
