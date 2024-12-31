<template>
    <div
        ref="infoHeader"
        class="info-header"
    >
        <span class="info-line"></span>
        <!-- 操作日志 -->
        <span class="info-title">
            {{$t('deviceDetail.operatLog.header.title')}}
        </span>
        <span style="margin-left: 20px">
            {{$t('deviceDetail.operatLog.header.tips')}}
        </span>
    </div>
</template>
<script lang="ts" setup>

const infoHeader = ref();
// defineEmits API 用来代替emit
const emitValue = defineEmits(['info-header-ref']);
onMounted(() => {
    emitValue('info-header-ref', infoHeader)
})
</script>