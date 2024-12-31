<template>
    <div v-if="editOperate.editingComputerRoomDiaLog.value">
        <edit-idc
            :dia-log="editOperate.editingComputerRoomDiaLog.value"
            :item-data="verificationOperate.reactiveArr.itemData"
            @dia-log-close="editOperate.editIdcCancel"
            @determine-click="editOperate.editIdcSure"
        />
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    editOperate: any;
    verificationOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
