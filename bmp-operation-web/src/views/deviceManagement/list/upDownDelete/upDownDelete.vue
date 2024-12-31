<template>
    <div v-if="upDownDelete.upDownDeleteDiaLog.value">
        <up-down-frame
            :dia-log="upDownDelete.upDownDeleteDiaLog.value"
            :device-id="upDownDelete.deviceId.value"
            :title="upDownDelete.upDownDeleteTitle.value"
            @dia-log-close="upDownDelete.upDownDeleteCancel"
            @determine-click="upDownDelete.upDownDeleteSure"
        />
    </div>
</template>

<script setup lang="ts">
interface PropsType {
    upDownDelete: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
