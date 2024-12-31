<template>
    <!-- 开机、关机、重启 -->
    <div v-if="operate.openShutDownRestartDiaLog.value">
        <open-shut-down-restart
            :dia-log="operate.openShutDownRestartDiaLog.value"
            :title="operate.title.value"
            :instance-id="operate.instanceId.value"
            :instance-info="operate.instanceInfo.value"
            @dia-log-close="operate.openShutDownRestartCancel"
            @determine-click="operate.openShutDownRestartSure"
        />
    </div>
</template>
<script setup lang="ts">
interface PropsType {
    operate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
