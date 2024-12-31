<template>
    <div
        class="big-size-echarts"
        v-if="echartsBigOpt.diaLog.value"
    >
        <div class="size-continer">
            <div class="header">
                <p class="title">{{echartsBigOpt.sizeBigTitle.value}}</p>
                <div class="count">
                    <p>
                        {{$filter.withClonZh($t('monitorEcharts.echartsCount.period'))}}{{echartsOperate.cycle.value}}
                    </p>
                    <p>{{$filter.withClonZh($t('monitorEcharts.echartsCount.aggregationMethod'))}}
                        <!-- @vue-ignore -->
                        <ui-select
                            v-model="echartsOperate.ruleForm[echartsBigOpt.sizeRef.value]"
                            @change="echartsBigOpt.bigSizeChange(echartsBigOpt.sizeRef.value, echartsOperate.ruleForm, echartsBigOpt.sizeRef.value)"
                        >
                            <ui-option
                                style="fontSize: 12px"
                                v-for="item in Object.freeze(publicConfigOpt.aggregationMethod)"
                                :key="item"
                                :label="item"
                                :value="item"
                            />
                        </ui-select>
                    </p>
                </div>
                <!-- 标题关闭icon -->
                <ui-button
                    class="close-button"
                    circle
                    :icon="Close"
                    @click="echartsBigOpt.closeClick(echartsBigOpt.sizeRef.value)"
                >
                </ui-button>
            </div>
            <div
                id="bigSize"
                v-loading="echartsBigOpt.bigSizeLoading.value"
            >
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {Close} from '@element-plus/icons-vue';
import EchartsBigOpt from './utils/echartsBigOpt';
import EchartsOpt from './echartsData';
import publicConfigOpt from './utils/publicConfig';

interface PropsType {
    echartsOperate: EchartsOpt;
    echartsBigOpt: EchartsBigOpt;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
<style lang="scss" scoped>
@import './css/echartsBig.scss';
</style>