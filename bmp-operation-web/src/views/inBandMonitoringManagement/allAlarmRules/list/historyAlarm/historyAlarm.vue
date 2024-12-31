<template>
    <div v-if="alaramHistoryOpt.alaramHistoryDiaLog.value">
        <alarm-history
            :dia-log="alaramHistoryOpt.alaramHistoryDiaLog.value"
            :rule-id="alaramHistoryOpt.ruleId.value"
            @dia-log-close="alaramHistoryOpt.alarmHistoryCancel"
        />
    </div>
</template>

<script setup lang="ts">
interface PropsType {
    alaramHistoryOpt: {
        alaramHistoryDiaLog: {
            value: boolean;
        };
        ruleId: {
            value: string;
        };
        alarmHistoryCancel: () => void;
    };
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
