<template>
    <div v-if="deleteOperate.deleteUserDiaLog.value">
        <delete-user
            :dia-log="deleteOperate.deleteUserDiaLog.value"
            :list-data="deleteOperate.listData.value"
            @dia-log-close="deleteOperate.deleteCancel"
            @determine-click="deleteOperate.deleteSure"
        />
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    deleteOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>