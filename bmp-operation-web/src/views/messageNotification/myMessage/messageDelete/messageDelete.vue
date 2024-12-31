<template>
    <div v-if="operate.deleteMsgDiaLog.value">
        <delete-message
            :dia-log="operate.deleteMsgDiaLog.value"
            :data="data"
            @dia-log-close="operate.deleteMsgCancel"
            @determine-click="operate.deleteMsgSure"
        />
    </div>
</template>
<script lang="ts" setup>
interface PropsType {
    operate: any;
    data: {[x: string]: string}[];
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>