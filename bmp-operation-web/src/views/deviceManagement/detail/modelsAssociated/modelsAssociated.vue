<template>
    <div v-if="operate.associatedModelDiaLog.value">
        <associated-models
            :device-id="deviceId"
            :dia-log="operate.associatedModelDiaLog.value"
            @dia-log-close="operate.associatedModelCancel"
            @determine-click="operate.associatedModelSure(diskTableOpt)"
        />
    </div>
</template>
<script lang="ts" setup>

// defineEmits API 来替代 emits
interface PropsType {
    operate: any;
    deviceId: string;
    diskTableOpt: {init(): void;};
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
});

// const routerOperate = new RouterOperate();
</script>