<template>
    <div class="echarts-search-count">
        <div>
            <!-- 刷新 -->
            <div class="seach-opt-count">
                <!-- 刷新 -->
                <p
                    class="operate-refresh"
                    @click="echartsOperate.refresh"
                >
                    <span class="my-update-gray"></span>  
                </p>
                <!-- 小时、天筛选 -->
                <div class="radio-active">
                    <el-radio-group v-model="echartsOperate.hours.value">
                        <el-radio-button
                            v-for="(item, index) in InitStaticData.radioData"
                            :key="index"
                            :value="index"
                            :label="item.label"
                        />
                    </el-radio-group>
                </div>
                <!-- 日期筛选 -->
                <el-config-provider :locale="locale">
                    <el-date-picker
                        v-model="echartsOperate.value1.value"
                        type="daterange"
                        range-separator="-"
                        :start-placeholder="$t('monitorEcharts.datePicker.startDate')"
                        :end-placeholder="$t('monitorEcharts.datePicker.endDate')"
                        :default-time="[new Date(0, 0, 0, 0, 0, 0), new Date(0, 0, 0, 23, 59, 59)]"
                        :disabled-date="echartsOperate.disabledDate"
                    >
                    </el-date-picker>
                </el-config-provider>
                <div class="set-alarm-rules">{{$t('monitorEcharts.goRuleMangement')}} <span @click="jumpToRuleList">{{$t('monitorEcharts.goAlarmRule')}}</span></div>
            </div>
            <div class="search-radio-count">
                <el-config-provider :locale="locale">
                    <el-radio-group v-model="echartsOperate.readioSearch.value">
                        <el-radio :label="$t('monitorEcharts.radio.instance')">
                            {{$t('monitorEcharts.radio.instance')}}
                        </el-radio>
                        <el-radio :label="$t('monitorEcharts.radio.disk')">
                            {{$t('monitorEcharts.radio.disk')}}
                            <el-tooltip placement="bottom">
                                <template #content>
                                    <div>{{$t('monitorEcharts.diskTip')}}</div>
                                </template>
                                <img
                                    alt=""
                                    class="common-img"
                                    src="@/assets/img/tooltip.png"
                                />
                            </el-tooltip>
                            <el-select
                                v-if="Object.is(echartsOperate.readioSearch.value, $t('monitorEcharts.radio.disk'))"
                                v-model="echartsOperate.diskVal.value"
                                @change="echartsOperate.disk"
                            >
                                <el-option
                                    style="fontSize: 12px"
                                    v-for="item in echartsOperate.selectValForm.disk"
                                    :key="item"
                                    :label="item"
                                    :value="item"
                                />
                            </el-select>    
                        </el-radio>
                        <el-radio :label="$t('monitorEcharts.radio.diskPartition')">
                            {{$t('monitorEcharts.radio.diskPartition')}}
                            <el-tooltip placement="bottom">
                                <template #content>
                                    <div>{{$t('monitorEcharts.diskPartitionTip')}}</div>
                                </template>
                                <img
                                    alt=""
                                    class="common-img"
                                    src="@/assets/img/tooltip.png"
                                />
                            </el-tooltip>
                            <el-select
                                v-if="echartsOperate.readioSearch.value === $t('monitorEcharts.radio.diskPartition')"
                                v-model="echartsOperate.mountpointVal.value"
                                @change="echartsOperate.mountpoint"
                            >
                                <el-option
                                    style="fontSize: 12px"
                                    v-for="item in echartsOperate.selectValForm.mountpoint"
                                    :key="item"
                                    :label="item"
                                    :value="item"
                                />
                            </el-select>    
                        </el-radio>
                        <el-radio :label="$t('monitorEcharts.radio.netWorkInterface')">
                            {{$t('monitorEcharts.radio.netWorkInterface')}}
                            <el-select
                                v-if="echartsOperate.readioSearch.value === $t('monitorEcharts.radio.netWorkInterface')"
                                v-model="echartsOperate.nicVal.value"
                                @change="echartsOperate.nic"
                            >
                                <el-option
                                    style="fontSize: 12px"
                                    v-for="item in echartsOperate.selectValForm.nic"
                                    :key="item"
                                    :label="item"
                                    :value="item"
                                />
                            </el-select>
                        </el-radio>
                    </el-radio-group>
                </el-config-provider>
            </div>
        </div>
        <div class="echarts-count" >
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
import {
    onMounted,
    onUnmounted,
} from 'vue';
import EchartsOpt from './echartsData';
import BigEcharts from './echartsBig.vue';
import EchartsPage from './echartsPage.vue';
import EchartsBigOpt from './utils/echartsBigOpt';
import InitStaticData from './utils/initStaticData';
import { useRoute, useRouter } from 'vue-router';
import { 
    languageSwitch, // 语言切换
} from 'utils/index.ts';
import {newMyChart} from './options';
//import 'default-passive-events';
const route = useRoute();
const instanceId: string = route.query.instanceId as string;
/**
 * 语言
*/
const locale: any = languageSwitch();
const echartsOperate = new EchartsOpt(instanceId, (status: boolean) => {
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


const router = useRouter();
const jumpToRuleList = () => {
    router.push({
        path: `/instance/rules/list`,
            query: {
                projectId: route.query.projectId,
                projectName: route.query.projectName
            }
    });
}

</script>
<style lang="scss" scoped>
@import './css/deviceEcharts.scss';

</style>