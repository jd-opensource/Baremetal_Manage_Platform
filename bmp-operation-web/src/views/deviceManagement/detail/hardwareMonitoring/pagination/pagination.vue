<template>
    <!-- 分页组件 -->
    <ui-pagination
        v-if="paginationOperate.total.value > 0 && activeOperate.activeName.value !== 'basicInfo' && !faultLogTable.isTableLoading.value"
        :has-use-default="false"
        :page-sizes="[20, 50, 100]"
        :total="paginationOperate.total.value"
        :page-size="paginationOperate.pageSize.value"
        :page-number="paginationOperate.pageNumber.value"
        @change-page="paginationOperate.paginationChange"
        @page-sizes-change="paginationOperate.paginationChange"
    />
</template>

<script lang="ts" setup>
import {paginationOperate} from 'utils/publicClass.ts';

interface PropsType {
    tableDetail: any;
    faultLogTable: any;
    activeOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>