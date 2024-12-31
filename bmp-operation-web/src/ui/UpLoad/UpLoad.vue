<template>
    <el-upload ref="upload">
        <template v-slot="trigger">
            <slot name="trigger" :trigger="trigger"></slot>
        </template>
        <template #tip>
            <slot name="tip"></slot>
        </template>
    </el-upload>
</template>

<script setup lang="ts">

/**
 * upload实例
*/
const upload = ref();

/**
 * defineEmits API 来替代 emits
*/
const emitRefVal = defineEmits(['uploadEvent']);

/**
 * 生命周期触发，数据回传
*/
onMounted(() => {
    emitRefVal('uploadEvent', upload);
});
</script>