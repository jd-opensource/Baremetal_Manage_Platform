<template>
    <!-- 分页组件 -->
    <my-message-list-pagination
        :other-class="'flex-end-right'"
        :pagination-opt="myMessageList"
    >
        <div class="check-all checkbox-text-ellipsis">
            <ui-checkbox
                slot="check-box"
                v-model="myMessageList.hasCheckAll.value"
                @change="myMessageList.checkAll"
            />
        </div>
        <div class="batch-operate">
            <!-- 删除 -->
            <ui-button
                type="primary"
                :disabled="myMessageList.deleteBtnDisabled.value"
                @click="emitValue('delete-message')"
            >
                <ui-tooltip
                    placement="bottom"
                    :content="$t('myMessage.footer.tips')"
                    :disabled="!myMessageList.deleteBtnDisabled.value"
                >
                    {{$t('myMessage.footer.delete')}}
                </ui-tooltip>
            </ui-button>
            <!-- 标为已读 -->
            <ui-button
                type="primary"
                :disabled="myMessageList.readBtnDisabled.value"
                @click="emitValue('read-message')"
            >
                
                <ui-tooltip
                    placement="bottom"
                    :content="$t('myMessage.footer.tips')"
                    :disabled="!myMessageList.readBtnDisabled.value"
                >
                    {{$t('myMessage.footer.read')}}
                </ui-tooltip>
            </ui-button>
        </div>
    </my-message-list-pagination>
</template>

<script lang="ts" setup>
// import BatchOperate from './batchOpt';
// import {paginationOperate} from 'utils/publicClass.ts';
import usePagination from 'hooks/pagination/usePagination.ts';
import myMessageListPagination from 'hooks/pagination/paginationSelect.vue';

interface PropsType {
    myMessageList: any;
    filter: any;
};


// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

usePagination(props.myMessageList.init, props.myMessageList.filter.tableRef?.value);

const emitValue = defineEmits(['delete-message', 'read-message']);

</script>
<style lang="scss" scoped>
@import 'assets/css/batchOperatePagination.scss';
</style>