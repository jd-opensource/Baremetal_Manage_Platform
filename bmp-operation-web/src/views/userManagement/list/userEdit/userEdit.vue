<template>
    <div v-if="editOperate.editingUserDiaLog.value">
        <editing-user
            :dia-log="editOperate.editingUserDiaLog.value"
            :prefix="editOperate.prefix.value"
            :list-data="editOperate.listData.value"
            @dia-log-close="editOperate.editUserCancel"
            @determine-click="editOperate.editUserSure"
        />
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    editOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>