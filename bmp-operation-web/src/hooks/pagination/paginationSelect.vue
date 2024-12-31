<template>
    <!-- 分页组件 -->
    <ui-pagination
        v-if="paginationOpt.reactiveArr.tableData.length > 0 && !paginationOpt.isLoading.value"
        :has-use-default="false"
        :total="paginationOperate.total.value" 
        :page-size="paginationOperate.pageSize.value"
        :page-sizes="pageSizes"
        :page-number="paginationOperate.pageNumber.value"
        @change-page="paginationOperate.paginationChange"
        @page-sizes-change="paginationOperate.paginationChange"
    >
        <slot></slot>
    </ui-pagination>
</template>

<script lang="ts" setup>
import {paginationOperate} from 'utils/publicClass.ts';

interface PropsType {
    paginationOpt: any;
    pageSizes?: number[]
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    paginationOpt: {},
    // @ts-ignore
    pageSizes: [20, 50, 100]
});

</script>
