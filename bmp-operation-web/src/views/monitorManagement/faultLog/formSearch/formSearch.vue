<template>
    <div ref="searchRef">
        <ui-config-provider :locale="uiLocale.locale">
            <!-- 主体内容-表单操作 -->
            <ui-form
                class="fault-form"
                label-position="left"
                :model="ruleFormOperate.ruleForm"
                @submit.native.prevent
                @rule-form-ref="ruleFormOperate.getFormRef"
            >
                <!-- 机房名称 -->
                <ui-form-item
                    prop="idcName"
                    :label="$filter.withClon($t('faultLogList.search.label.idcName'))"
                >
                    <!-- 请选择机房 -->
                    <ui-select
                        class="custom-width custom-width1"
                        v-model="ruleFormOperate.ruleForm.idcId"
                        :loading="store.idcInfo.loading"
                        :placeholder="$t('faultLogList.search.placeholder.idcName')"
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
                <!-- sn -->
                <ui-form-item
                    prop="sn"
                    :label="$filter.withClon($t('faultLogList.search.label.sn'))"
                >
                    <ui-input
                        class="custom-width custom-width2"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm.sn"
                        :placeholder="$t('faultLogList.search.placeholder.sn')"
                        @change="ruleFormOperate.snChange"
                        @clear="ruleFormOperate.clearSn"
                    />
                </ui-form-item>
                <!-- 故障等级 -->
                <ui-form-item
                    :label="$filter.withClon($t('faultLogList.search.label.level'))"
                >
                    <!-- 请选择故障等级 -->
                    <ui-select
                        class="custom-width custom-width3"
                        clearable
                        v-model="ruleFormOperate.ruleForm.level"
                        :no-data-text="$t('importDevice.noData.tip')"
                        :loading="store.faultLevelInfo.loading"
                        :placeholder="$t('faultLogList.search.placeholder.level')"
                        @clear="ruleFormOperate.singleClear('level')"
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
                <!-- 故障提报时间 -->
                <ui-form-item
                    :label="$filter.withClon($t('faultLogList.search.label.time'))"
                >
                    <!-- 请选择故障提报时间 -->
                    <ui-date-picker
                        class="custom-width custom-width4"
                        type="daterange"
                        v-model="ruleFormOperate.ruleForm.faultReportingTime"
                        :disabled-date="ruleFormOperate.disabledDate"
                        :start-placeholder="$t('faultLogList.search.placeholder.startDate')"
                        :end-placeholder="$t('faultLogList.search.placeholder.endDate')"
                        @change="ruleFormOperate.dateChange"
                    />
                </ui-form-item>
                <!-- 故障配件 -->
                <ui-form-item
                    :label="$filter.withClon($t('faultLogList.search.label.accessory'))"
                >
                    <!-- 请选择故障配件 -->
                    <ui-select
                        class="custom-width custom-width5"
                        clearable
                        v-model="ruleFormOperate.ruleForm.spec"
                        :loading="store.faultSpecInfo.loading"
                        :placeholder="$t('faultLogList.search.placeholder.accessory')"
                        @clear="ruleFormOperate.singleClear('spec')"
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
                <!-- 状态 -->
                <ui-form-item
                    :label="$filter.withClon($t('faultLogList.search.label.status'))"
                >
                    <!-- 请选择状态 -->
                    <ui-select
                        class="custom-width custom-width3"
                        v-model="ruleFormOperate.ruleForm.dealStatus"
                        :placeholder="$t('faultLogList.search.placeholder.status')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in SurveillanceStaticeData.status"
                            :key="item.value"
                            :label="item.text"
                            :value="item.value"
                        />
                    </ui-select>
                </ui-form-item>
                <ui-form-item>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.searchClick"
                    >
                        {{$t('faultLogList.search.btn.search')}}
                    </ui-button>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.clearClick"
                    >
                        {{$t('faultLogList.search.btn.clear')}}
                    </ui-button>
                </ui-form-item>
        
            </ui-form>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale, CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import SurveillanceStaticeData from 'staticData/surveillance/index.ts';

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
const emitValue = defineEmits(['search-ref']);

const proxy = new CurrentInstance().proxy;

store.idcInfo.getIdc(proxy.$idcApi.idcListAPI);

store.faultLevelInfo.getFaultLevel();

store.faultSpecInfo.getFaultSpec();

const searchRef: Ref = ref(0);

onMounted(() => {
    emitValue('search-ref', searchRef)
})
</script>
