<template>
    <el-select :popper-class="customClass">
        <slot></slot>
    </el-select>
</template>

<script setup lang="ts">

/**
 * props 类
*/
interface PropType {
    customClass?: string;
};

/**
 * withDefaults API 用来定义默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropType>(), {
    customClass: 'default-select'
});

</script>

<style lang="scss">
@import './select.scss';
</style>
