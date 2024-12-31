<template>
    <!-- 分页组件 -->
    <ui-pagination
        v-if="operate.reactiveArr.tableData.length > 0 && !operate.isLoading.value"
        :has-use-default="false"
        :page-sizes="[20, 50, 100]"
        :total="customPagination.total.value" 
        :page-size="customPagination.pageSize.value"
        :page-number="customPagination.pageNumber.value"
        @change-page="customPagination.paginationChange"
        @page-sizes-change="customPagination.paginationChange"
    >
    </ui-pagination>
</template>

<script lang="ts" setup>

interface PropsType {
    operate: any;
    customPagination: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});


</script>