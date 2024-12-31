<template>
    <!-- 分页组件 -->
    <ui-pagination
        v-if="faultLog.reactiveArr.tableData.length > 0 && !faultLog.isLoading.value"
        :has-use-default="false"
        :total="paginationOperate.total.value" 
        :page-size="paginationOperate.pageSize.value"
        :page-sizes="[20, 50, 100]"
        :page-number="paginationOperate.pageNumber.value"
        @change-page="paginationOperate.paginationChange"
        @page-sizes-change="paginationOperate.paginationChange"
    />
</template>

<script lang="ts" setup>
import {paginationOperate} from 'utils/publicClass.ts';

interface PropsType {
    faultLog: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
