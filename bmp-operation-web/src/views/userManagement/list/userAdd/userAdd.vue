<template>
    <div v-if="addUserOperate.addUserDialog.value">
        <add-user
            :dia-log="addUserOperate.addUserDialog.value"
            @dia-log-close="addUserOperate.addUserCancel"
            @determine-click="addUserOperate.addUserSure"
        />
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    addUserOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>