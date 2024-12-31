<template>
    <div v-if="operate.secondaryConfirmationDiaLog.value">
        <secondary-confirmation
            :name="operate.secondaryConfirmationName.value"
            :dia-log="operate.secondaryConfirmationDiaLog.value"
            @dia-log-close="operate.secondaryConfirmationCancel"
            @determine-click="operate.secondaryConfirmationSure"
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
