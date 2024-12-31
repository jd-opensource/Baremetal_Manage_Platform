<template>
    <div
        class="echart-bar"
        v-show="showType.includes(readioSearch)"
    >
        <div class="header">
            <p class="title">{{title}}</p>
            <div class="count">
                <p>{{$filter.withClonZh($t('monitorEcharts.echartsCount.period'))}}{{cycle}}</p>
                <p>{{$filter.withClonZh($t('monitorEcharts.echartsCount.aggregationMethod'))}}
                    <el-select
                        v-model="modelData"
                        @change="emitValue('select-change', type, modelData)"
                    >
                        <el-option
                            style="fontSize: 12px"
                            v-for="item in publicConfigOpt.aggregationMethod"
                            :key="item"
                            :label="item"
                            :value="item"
                        />
                    </el-select>
                </p>
            </div>
            <p
                class="big common-img"
                v-if="hasEnlarge"
                @click="emitValue('big-click')"
            >
            </p>
        </div>
        <div
            class="echart-context"
            :id="id"
        >
        </div>
    </div>
</template>
<script lang="ts" setup>
import publicConfigOpt from './utils/publicConfig';

interface PropsType {
    hasEnlarge?: boolean;
    title: string;
    showType: string[];
    readioSearch: string;
    cycle: string;
    modelData: string;
    id: string;
    type: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    hasEnlarge: true
});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['select-change', 'big-click']);

</script>
<style lang="scss" scoped>
@import './css/echartsPage.scss';
</style>