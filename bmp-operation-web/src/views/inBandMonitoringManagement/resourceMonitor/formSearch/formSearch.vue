<template>
    <div ref="searchRef">
        <ui-config-provider :locale="uiLocale.locale">
            <!-- 主体内容-表单操作 -->
            <ui-form
                class="in-band-monitoring"
                label-position="left"
                label-width="90px"
                :model="ruleFormOperate.ruleForm"
                :rules="rulesCheckOpt.rules"
                @submit.native.prevent
                @rule-form-ref="ruleFormOperate.getFormRef"
            >
                <!-- 用户名 -->
                <ui-form-item
                    prop="userName"
                    class="set-empty"
                    :label="$filter.withClon($t('resourceMonitor.search.userName'))"
                >
                    <ui-input
                        class="custom-width custom-width7"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm.userName"
                        :placeholder="$t('resourceMonitor.placeholder.userName')"
                        @change="ruleFormOperate.userNameChange"
                        @clear="ruleFormOperate.clearUserName"
                    />
                </ui-form-item>
                <!-- 资源类型 -->
                <ui-form-item
                    prop="resourceType"
                    class="set-empty"
                    :label="$filter.withClon($t('resourceMonitor.search.resourceType'))"
                >
                    <!-- 请选择资源类型 -->
                    <ui-select
                        class="custom-width custom-width7"
                        v-model="ruleFormOperate.ruleForm.resourceType"
                        :placeholder="$t('resourceMonitor.placeholder.resourceType')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="(item, index) in InitStaticData.resourceData"
                            :key="index"
                            :label="item"
                            :value="item"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 机房 -->
                <ui-form-item
                    prop="idcId"
                    class="set-empty"
                    :label="$filter.withClon($t('resourceMonitor.search.idcName'))"
                >
                    <!-- 请选择机房 -->
                    <ui-select
                        class="custom-width custom-width7"
                        v-model="ruleFormOperate.ruleForm.idcId"
                        :placeholder="$t('resourceMonitor.placeholder.idcName')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in store.idcInfo.idcData"
                            :key="item.name"
                            :label="item.newIdcName"
                            :value="item.idcId"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 实例ID -->
                <ui-form-item
                    prop="instanceId"
                    class="set-empty"
                    :label="$filter.withClon($t('resourceMonitor.search.instanceId'))"
                >
                    <ui-input
                        class="custom-width custom-width7"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm.instanceId"
                        :placeholder="$t('resourceMonitor.placeholder.instanceId')"
                        @change="ruleFormOperate.instanceIdChange"
                        @clear="ruleFormOperate.clearInstanceId"
                    />
                </ui-form-item>
                <!-- 维度 -->
                <ui-form-item
                    prop="dimension"
                    class="set-empty"
                    :label="$filter.withClon($t('resourceMonitor.search.dimensions'))"
                >
                    <!-- 请选择维度 -->
                    <ui-select
                        class="custom-width custom-width7"
                        v-model="ruleFormOperate.ruleForm.dimension"
                        :placeholder="$t('resourceMonitor.placeholder.dimensions')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="(item, index) in InitStaticData.dimensionsData"
                            :key="index"
                            :label="item"
                            :value="item"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 监控指标 -->
                <ui-form-item
                    prop="monitoringIndicators"
                    :class="[
                        'multiple-choice',
                        'set-empty',
                        setDiffClass('', 'multiple-choice-short')
                    ]"
                    :label="$filter.withClon($t('resourceMonitor.search.monitoringMetrics'))"
                >
                    <!-- 请选择监控指标 -->
                    <ui-select
                        v-model="ruleFormOperate.ruleForm.monitoringIndicators"
                        multiple
                        collapse-tags
                        popper-class="select-seleted-icon"
                        :class="[
                            'custom-width',
                            setDiffClass('custom-width6', 'custom-width7')
                        ]"
                        :placeholder="$t('resourceMonitor.placeholder.monitoringMetrics')"
                    >
                        <div class="select-checkbox-all">
                            <ui-checkbox
                                v-model="ruleFormOperate.checkAll.value"
                                :indeterminate="ruleFormOperate.indeterminate.value"
                                @change="ruleFormOperate.handleCheckAll"
                            >
                                {{$t('size.all')}}
                            </ui-checkbox>
                        </div>
                        <ui-option
                            style="fontSize: 12px"
                            v-for="(item, index) in ruleFormOperate.monitoringIndicators"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </ui-select>
                </ui-form-item>
                <ui-form-item
                    prop="hours"
                    class="hours-select set-empty"
                    :label="''"
                >
                    <div class="radio-active">
                        <ui-radio-group v-model="echartsOperate.hours.value">
                            <ui-radio-button
                                v-for="(item, index) in InitStaticData.radioData"
                                :key="index"
                                :value="index"
                                :label="item.label"
                            />
                        </ui-radio-group>
                    </div>
                </ui-form-item>
                <ui-form-item
                    prop="value1"
                    class="daterange-select set-empty"
                    label=""
                >
                    <ui-config-provider :locale="uiLocale.locale">
                        <ui-date-picker
                            v-model="echartsOperate.value1.value"
                            type="daterange"
                            range-separator="-"
                            :start-placeholder="$t('monitorEcharts.datePicker.startDate')"
                            :end-placeholder="$t('monitorEcharts.datePicker.endDate')"
                            :default-time="[new Date(0, 0, 0, 0, 0, 0), new Date(0, 0, 0, 23, 59, 59)]"
                            :disabled-date="echartsOperate.disabledDate"
                        >
                        </ui-date-picker>
                    </ui-config-provider>
                </ui-form-item>
                <ui-form-item class="opt-btn set-empty" prop="btn" label="">
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.searchClick"
                    >
                        {{$t('resourceMonitor.search.btn.search')}}
                    </ui-button>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.clearClick"
                    >
                        {{$t('resourceMonitor.search.btn.clear')}}
                    </ui-button>
                </ui-form-item>
            </ui-form>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
// import EchartsOpt from '../echartsData';
// import EchartsBigOpt from '../utils/echartsBigOpt';
import {uiLocale, CurrentInstance} from 'utils/publicClass.ts';
import {setDiffClass} from 'utils/index.ts';
import InitStaticData from '../utils/initStaticData';
import store from 'store/index.ts';
import RulesCheckOpt from './rules';

// props类型
type PropsType = {
    ruleFormOperate: any;
    echartsOperate: any;
};

/**
 * props参数
 * @param {boolean} diaLog 弹窗状态
*/
withDefaults(defineProps<PropsType>(), {});

// emit 事件
const emitValue = defineEmits(['search-ref']);

const proxy = new CurrentInstance().proxy;

store.idcInfo.getIdc(proxy.$idcApi.idcListAPI);

const searchRef = ref(0);

onMounted(() => {
    emitValue('search-ref', searchRef)
})

const rulesCheckOpt = new RulesCheckOpt();
</script>
