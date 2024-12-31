<template>
    <!-- 机型验证操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="['image-model-add currency', !customPagination.total.value ? 'set-height' : '']"
        :is-loading="addModel.isLoading"
        :header-title="$t('addModel.header.addModel')"
        :src="($defInfo.imagePath('imageManage') as unknown as string)"
        :disabled-status="imageDeviceTypes.btnDisabled.value <= 0"
        @dia-log-close="addModel.cancelClick"
        @determine-click="addModel.determineClick"
    >
        <!-- 添加机型数据 -->
        <div class="currency-count">
            <!-- 镜像名称 -->
            <image-model-name :image-name="imageName" :image-device-types="imageDeviceTypes"/>
            <div class="currency-count-table">
                <image-device-types-list :image-device-types="imageDeviceTypes"/>
                <!-- 分页组件 -->
                <image-model-pagination :operate="imageDeviceTypes" :custom-pagination="customPagination"/>
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import * as all from './all';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    imageId: string;
    architecture: string;
    imageName: string;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    imageId: '',
    architecture: '',
    imageName: ''
});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const customPagination = all.paginationOperate(() => {
    imageDeviceTypes.processingParameters();
});

const imageDeviceTypes = new all.ImageDeviceTypes(customPagination, props);

const addModel = all.addModelOperate(imageDeviceTypes, props, emitValue);

</script>
<style lang="scss">
@import './imageAddModel.scss';
</style>
