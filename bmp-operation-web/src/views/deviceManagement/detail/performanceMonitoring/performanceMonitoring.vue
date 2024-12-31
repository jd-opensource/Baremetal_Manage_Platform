<template>
    <div class="echarts-search-count">
        <div>
            <!-- 刷新 -->
            <div class="seach-opt-count">
                <refresh-custom-export
                    :has-btn-operate="false"
                    :has-export="false"
                    :custom-status="false"
                    @refresh="echartsOperate.refresh"
                />
                <!-- 小时、天筛选 -->
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
                <!-- 日期筛选 -->
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
            </div>
            <div class="search-radio-count radio-style1">
                <ui-config-provider :locale="uiLocale.locale">
                    <ui-radio-group v-model="echartsOperate.readioSearch.value">
                        <!-- 实例 -->
                        <ui-radio :label="$t('monitorEcharts.radio.instance')">
                            {{$t('monitorEcharts.radio.instance')}}
                        </ui-radio>
                        <!-- 磁盘 -->
                        <ui-radio :label="$t('monitorEcharts.radio.disk')">
                            {{$t('monitorEcharts.radio.disk')}}
                            <ui-tooltip placement="bottom">
                                <template #content>
                                    <div>{{$t('monitorEcharts.diskTip')}}</div>
                                </template>
                                <img
                                    alt=""
                                    class="common-img"
                                    :src="($defInfo.imagePath('toolTip') as unknown as string)"
                                />
                            </ui-tooltip>
                            <div v-if="Object.is(echartsOperate.readioSearch.value, $t('monitorEcharts.radio.disk'))">
                                <ui-select
                                    v-model="echartsOperate.diskVal.value"
                                    @change="echartsOperate.disk"
                                >
                                    <ui-option
                                        style="fontSize: 12px"
                                        v-for="item in echartsOperate.selectValForm.disk"
                                        :key="item"
                                        :label="item"
                                        :value="item"
                                    />
                                </ui-select>
                            </div>
                        </ui-radio>
                        <!-- 磁盘用量 -->
                        <ui-radio :label="$t('monitorEcharts.radio.diskPartition')">
                            {{$t('monitorEcharts.radio.diskPartition')}}
                            <ui-tooltip placement="bottom">
                                <template #content>
                                    <div>{{$t('monitorEcharts.diskPartitionTip')}}</div>
                                </template>
                                <img
                                    alt=""
                                    class="common-img"
                                    :src="($defInfo.imagePath('toolTip') as unknown as string)"
                                />
                            </ui-tooltip>
                            <div v-if="Object.is(echartsOperate.readioSearch.value, $t('monitorEcharts.radio.diskPartition'))">
                                <ui-select
                                    v-model="echartsOperate.mountpointVal.value"
                                    @change="echartsOperate.mountpoint"
                                >
                                    <ui-option
                                        style="fontSize: 12px"
                                        v-for="item in echartsOperate.selectValForm.mountpoint"
                                        :key="item"
                                        :label="item"
                                        :value="item"
                                    />
                                </ui-select>
                            </div>
                        </ui-radio>
                        <!-- 网口 -->
                        <ui-radio :label="$t('monitorEcharts.radio.netWorkInterface')">
                            {{$t('monitorEcharts.radio.netWorkInterface')}}
                            <div v-if="Object.is(echartsOperate.readioSearch.value, $t('monitorEcharts.radio.netWorkInterface'))">
                                <ui-select
                                    v-model="echartsOperate.nicVal.value"
                                    @change="echartsOperate.nic"
                                >
                                    <ui-option
                                        style="fontSize: 12px"
                                        v-for="item in echartsOperate.selectValForm.nic"
                                        :key="item"
                                        :label="item"
                                        :value="item"
                                    />
                                </ui-select>
                            </div>
                        </ui-radio>
                    </ui-radio-group>
                </ui-config-provider>
            </div>
        </div>
        <div class="echarts-count">
            <div
                v-for="(item, index) in echartsOperate.reactiveArr.echartsParamsData"
                :key="index"
            >
                <div v-loading="echartsOperate.isLoading(item.type)">
                    <echarts-page
                        :title="item.title"
                        :show-type="item.showType"
                        :readio-search="echartsOperate.readioSearch.value"
                        :cycle="echartsOperate.cycle.value"
                        :model-data="item.model"
                        :type="item.type"
                        :id="item.id"
                        @select-change="echartsOperate.selectChange"
                        @big-click="echartsBigOpt.bigClickOpt(item.bigClickVal.type, item.bigClickVal.title)"
                    />
                </div>
            </div>
        </div>
        <!-- 放大操作 -->
        <div v-if="echartsBigOpt.diaLog.value">
            <big-echarts
                :echarts-operate="echartsOperate"
                :echarts-big-opt="echartsBigOpt"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import EchartsOpt from './echartsData';
import BigEcharts from './echartsBig.vue';
import EchartsPage from './echartsPage.vue';
import EchartsBigOpt from './utils/echartsBigOpt';
import InitStaticData from './utils/initStaticData';
import {uiLocale} from 'utils/publicClass.ts';
import {newMyChart} from './options';
// import 'default-passive-events';

interface PropsType {
    deviceDetail: {
        reactiveArr: {
            detail: {
                instanceId: string;
                deviceId: string;
            };
        };
        initData(): void;
    };
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

props.deviceDetail.initData();

const echartsOperate = new EchartsOpt(props, (status: boolean) => {
    echartsBigOpt.bigSizeLoading.value = status;
});

const echartsBigOpt = new EchartsBigOpt(echartsOperate);

class ResizeOpt {

    constructor() {
        onMounted(() => window.addEventListener('resize', this.resize));
        onUnmounted(() => window.removeEventListener('resize', this.resize));
    }

    resize = () => {
        newMyChart?.resize();
    }
}

new ResizeOpt();

</script>
<style lang="scss" scoped>
@import './css/deviceEcharts.scss';
</style>