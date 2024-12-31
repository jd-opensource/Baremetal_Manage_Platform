<template>
    <!-- 主体内容-表单操作 -->
    <div ref="searchRef">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-form
                class="fault-form"
                label-position="left"
                :model="ruleFormOperate.ruleForm"
                @submit.native.prevent
                @rule-form-ref="ruleFormOperate.getFormRef"
            >
                <!-- 故障名称 -->
                <ui-form-item
                    prop="faultName"
                    :label="$filter.withClon($t('faultRulesList.search.label.faultName'))"
                >
                    <ui-input
                        class="custom-width custom-width2"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm.alertName"
                        :placeholder="$t('faultRulesList.search.placeholder.faultName')"
                        @change="ruleFormOperate.faultNameChange"
                        @clear="ruleFormOperate.singleClear('alertName')"
                    />
                </ui-form-item>
                <!-- 故障配件 -->
                <ui-form-item
                    :label="$filter.withClon($t('faultRulesList.search.label.faultAccessories'))"
                >
                    <!-- 请选择故障配件 -->
                    <ui-select
                        class="custom-width custom-width5"
                        clearable
                        v-model="ruleFormOperate.ruleForm.alertSpec"
                        :loading="store.faultSpecInfo.loading"
                        :placeholder="$t('faultRulesList.search.placeholder.faultAccessories')"
                        @clear="ruleFormOperate.singleClear('alertSpec')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in store.faultSpecInfo.faultSpecData"
                            :key="item"
                            :label="item"
                            :value="item"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 故障等级 -->
                <ui-form-item
                    :label="$filter.withClon($t('faultRulesList.search.label.faultLevel'))"
                >
                    <!-- 请选择故障等级 -->
                    <ui-select
                        class="custom-width custom-width3"
                        clearable
                        v-model="ruleFormOperate.ruleForm.alertLevel"
                        :loading="store.faultLevelInfo.loading"
                        :no-data-text="$t('importDevice.noData.tip')"
                        :placeholder="$t('faultRulesList.search.placeholder.faultLevel')"
                        @clear="ruleFormOperate.singleClear('alertLevel')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in store.faultLevelInfo.faultLevelData"
                            :key="item"
                            :label="item"
                            :value="item"
                        />
                    </ui-select>
                </ui-form-item>
                <ui-form-item>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.searchClick"
                    >
                        {{$t('faultRulesList.search.btn.search')}}
                    </ui-button>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.clearClick"
                    >
                        {{$t('faultRulesList.search.btn.clear')}}
                    </ui-button>
                </ui-form-item>
            </ui-form>
        </ui-config-provider>
        <div class="restore-default-set">
            <ui-tooltip
                placement="bottom"
                :disabled="!ruleFormOperate.btnDisabled.value"
                :content="$t('operate.tip.default')"
            >
                <ui-button
                    type="primary"
                    :disabled="ruleFormOperate.btnDisabled.value"
                    @click="emitValue('def-set-opt')"
                >
                    {{$t('faultRulesList.search.btn.restoreDefaultSet')}}
                </ui-button>
            </ui-tooltip>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';
import store from 'store/index.ts';

// props类型
type PropsType = {
    ruleFormOperate: any;
};

/**
 * props参数
 * @param {boolean} diaLog 弹窗状态
*/
withDefaults(defineProps<PropsType>(), {});

// emit 事件
const emitValue = defineEmits(['search-ref', 'def-set-opt']);

store.faultLevelInfo.getFaultLevel();

store.faultSpecInfo.getFaultSpec();

const searchRef: Ref = ref(0);

onMounted(() => {
    emitValue('search-ref', searchRef)
})
</script>