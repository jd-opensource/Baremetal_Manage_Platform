<template>
    <div v-if="customOperate.customListDiaLog.value">
        <custom-list
            :page-key="customOperate.pageKey"
            :dia-log="customOperate.customListDiaLog.value"
            :check-list-arr="customOperate.reactiveArr.checkListArr"
            :has-custom-info="customOperate.reactiveArr.hasCustomInfo"
            @dia-log-close="customOperate.customCancelOperate"
            @determine-click="customOperate.getCustomList"
        />
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    customOperate: any;
}

// props参数
withDefaults(defineProps<PropsType>(), {});
</script>
