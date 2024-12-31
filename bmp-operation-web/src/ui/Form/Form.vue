<template>
    <el-form ref="ruleFormRef">
        <slot></slot>
    </el-form>
</template>
  
<script setup lang="ts">

/**
 * 表单ref
*/
const ruleFormRef = ref();

/**
 * defineEmits API 来替代 emits
*/
const emitRefVal = defineEmits(['ruleFormRef']);

/**
 * 生命周期，回传数据
*/
onMounted(async () => {
    await emitRefVal('ruleFormRef', ruleFormRef);
});


</script>
<style lang="scss">
@import './form.scss';
</style>
