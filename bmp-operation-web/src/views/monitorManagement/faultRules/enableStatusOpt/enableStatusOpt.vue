<template>
    <div v-if="enableStatusOperate.enableStatusDiaLog.value">
        <enable-status
            :dia-log="enableStatusOperate.enableStatusDiaLog.value"
            :status="!enableStatusOperate.useItem.useStatus"
            :alert-name="enableStatusOperate.useItem.alertName"
            :role-id="enableStatusOperate.useItem.ruleId"
            @dia-log-close="enableStatusOperate.enableStatusCancel"
            @determine-click="enableStatusOperate.enableStatusSure"
        />
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    enableStatusOperate: any;
}

// defineEmits API 用来代替emit
const emitValue = defineEmits(['']);
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>
