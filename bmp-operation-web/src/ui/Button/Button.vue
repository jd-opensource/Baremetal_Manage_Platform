<template>
    <el-button v-if="slotDefault">
        <slot><span></span></slot>
    </el-button>
    <el-button v-else>
    </el-button>
</template>
  
<script setup lang="ts">

/**
 * 判断是否使用默认插槽
*/
const slotDefault: boolean = !!useSlots().default;

</script>
<style lang="scss">
@import './button.scss';
</style>
