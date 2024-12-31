<template>
    <!-- 镜像验证操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="['model-image-add currency', !customPagination.total.value ? 'set-height' : '']"
        :header-title="$t('addImage.header.addImage')"
        :src="($defInfo.imagePath('imageManage') as unknown as string)"
        :disabled-status="imageList.btnDisabled.value <= 0"
        :is-loading="addImageOperate.isLoading"
        @dia-log-close="addImageOperate.cancelClick"
        @determine-click="addImageOperate.determineClick"
    >
        <!-- 添加镜像数据 -->
        <div class="currency-count">
            <!-- 机型名称 -->
            <model-image-add-name
                :model-name="modelName"
                :image-list="imageList"
            />
            <model-image-add-reset
                v-if="imageList.searchTip.value"
                :reset-operate="reset"
                :custom-pagination="customPagination"
            />
            <!-- :class="imageList.reactiveArr.tableData.length ? 'currency-count-table' : 'no-table'" -->
            <div class="currency-count-table">
                <model-image-list-table
                    :image-list="imageList"
                    :filter-operate="filter"
                />
                <!-- 分页组件 -->
                <model-image-add-pagination
                    :operate="imageList"
                    :custom-pagination="customPagination"
                />
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import * as all from './all';

// props 类
interface PropsType {
    diaLog: boolean;
    deviceTypeId: string;
    modelName: string;
    architecture: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    deviceTypeId: '',
    modelName: '',
    architecture: ''
});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const customPagination = all.paginationOperate(() => {
    imageList.init();
});

const imageList = new all.ImageList(customPagination, props);

const addImageOperate = all.addImageOperate(imageList, props, emitValue);

const filter = new all.FilterOperate(customPagination, imageList);

const reset =  all.resetOperate(imageList, filter);

</script>
<style lang="scss">
@import './addImage.scss';
</style>
