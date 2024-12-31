<template>
    <div v-if="defaultSetOperate.defaultSetDiaLog.value">
        <default-set
            :dia-log="defaultSetOperate.defaultSetDiaLog.value"
            :select-arr="faultRulesListOpt.reactiveArr.selectArr"
            @dia-log-close="defaultSetOperate.defaultSetCancel"
            @determine-click="defaultSetOperate.defaultSetSure"
        />
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    defaultSetOperate: any;
    faultRulesListOpt: any;
}

// defineEmits API 用来代替emit
const emitValue = defineEmits(['']);
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>