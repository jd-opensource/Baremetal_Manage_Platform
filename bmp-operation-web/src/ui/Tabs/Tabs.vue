<template>
    <el-tabs
        :class="setDiffClass('tabs-en-width', 'tabs-zh-width')"
        :active-name="activeName">
        <el-tab-pane
            v-for="(item, index) in data"
            :key="index"
            :label="item.label"
            :name="item.name"
        >
        </el-tab-pane>
    </el-tabs>
</template>

<script setup lang="ts">
import {setDiffClass} from 'utils/index.ts';

/**
 * 数据类
*/
type DataType = {label: string; name: string;};

/**
 * props 类
*/
interface PropsType {
    activeName: string;
    data: DataType[];
};

/**
 * withDefaults API 用来定义默认值，defineProps API 用来代替props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    activeName: 'essentialInformation', // v-model
    data: () => []
});

/**
 * defineEmits API 来替代 emits
*/
const emitValue= defineEmits(['on-change']);

/**
 * 监听props.model
 * @return {string} newValue,每次更新触发
*/
watch(() => props.activeName, (newValue: string) => {
    emitValue('on-change', newValue);
}, {deep: true});

</script>
<style lang="scss" scoped>
@import './tabs';
</style>
