<template>
    <user-list-pagination
        :pagination-opt="userList"
    />
</template>

<script setup lang="ts">
import PaginationWatch from './pagination';
import UserListPagination from 'hooks/pagination/paginationSelect.vue';

interface PropsType {
    userList: any;
    search: any;
    filter: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});


new PaginationWatch(props);

</script>
