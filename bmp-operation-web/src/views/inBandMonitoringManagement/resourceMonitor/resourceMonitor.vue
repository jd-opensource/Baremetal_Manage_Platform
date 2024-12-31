<template>
    <div class="operate-management resource-monitor-content">
        <!-- 头部标题内容 -->
        <header-info
            :type="'resourceMonitor'"
        />
        <!-- 主体内容 -->
        <main class="operate-management-content resource-monitor-count">
            <form-search
                :rule-form-operate="ruleFormOperate"
                :echarts-operate="echartsOperate"
            />
            <div class="page-count" v-loading="echartsOperate.pageLoading.value">
                <div
                    class="echarts-count"
                    v-if="echartsOperate.showStatus() && ruleFormOperate.ruleForm.keyData.length"
                >
                    <div
                        v-for="(item, index) in echartsOperate.reactiveArr.echartsParamsData"
                        :key="index"
                    >
                        <div v-if="echartsOperate.echartShowStatus(item.type)">
                            <div v-loading="echartsOperate.isLoading(item.type)">
                                <echarts-page
                                    :title="item.title"
                                    :cycle="echartsOperate.cycle.value"
                                    :model-data="item.model"
                                    :has-enlarge="false"
                                    :type="item.type"
                                    :id="item.id"
                                    @select-change="echartsOperate.selectChange"
                                />
                            </div>
                        </div>
                    </div>
                </div>
                <div v-if="!echartsOperate.pageLoading.value && !ruleFormOperate.ruleForm.keyData?.length">
                    <div class="no-empty-img"></div>
                    <p v-if="Object.is(echartsOperate.errorCode.value, 0)" class="no-ower-text" v-html="echartsOperate.str"></p>
                </div>
            </div>
        </main>
</div>
</template>
<script lang="ts" setup>
import RuleFormOperate from './formSearch/rulesForm';
import FormSearch from './formSearch/formSearch.vue';
import EchartsOpt from './echartsData';
import EchartsPage from './echartsPage.vue';
import {myChartD} from './options';

const ruleFormOperate = new RuleFormOperate((res: string, objRuleForm: any) => {
    if (Object.is(res, 'search')) echartsOperate.searchInfo(objRuleForm);
    else echartsOperate.clearSearchInfo(ruleFormOperate, objRuleForm);
});

const echartsOperate = new EchartsOpt(ruleFormOperate);

class ResizeOpt {
    constructor() {
        onMounted(() => window.addEventListener('resize', this.resize));
        onUnmounted(() => window.removeEventListener('resize', this.resize));
    }
    
    resize = () => {
        for (const key of myChartD) {
            key?.resize()
        }
    }
}

new ResizeOpt();

</script>
<style lang="scss" scoped>
@import './css/index.scss';

</style>