<template>
    <el-table
        ref="tableRef"
        :header-cell-style="headerCellStyle"
    >
        <slot></slot>
        <template #empty>
            <slot name="empty"></slot>
        </template>
    </el-table>
</template>

<script setup lang="ts">

/**
 * Header 类
*/
type HeaderCellStyleType = {
    backgroundColor: string;
    color: string;
    fontWeight: string;
    fontSize: string;
    textAlign: string;
};

/**
 * props 类
*/
interface PropType {
    headerCellStyle?: HeaderCellStyleType | any;
};

// withDefaults API 用来定义默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropType>(), {
    headerCellStyle: () => {
        return {
            // backgroundColor: '#e8f4fd',
            color: '#333',
            fontWeight: '600',
            fontSize: '12px',
            textAlign: 'center'
        }
    }
});

// 表格ref
const tableRef: Ref<any> = ref<any>();

// defineEmits API 来替代 emits
const emitRefVal = defineEmits(['get-table-ref']);

// 生命周期，回传数据
onMounted(async () => {
    await emitRefVal('get-table-ref', tableRef);
});

</script>

<style lang="scss">
@import './table.scss';
</style>
