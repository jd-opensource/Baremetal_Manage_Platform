<template>
    <!-- 删除镜像操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'delete-image'"
        :is-loading="deleteImageOperate.isLoading"
        :header-title="$t('deleteImage.header.deleteImage')"
        :src="($defInfo.imagePath('imageManage') as unknown as string)"
        @dia-log-close="deleteImageOperate.cancelClick"
        @determine-click="deleteImageOperate.determineClick"
    >
        <!-- 删除提示 -->
        <div class="delete-image-count">
            <span class="delete-image-count-title">
                {{hasRelationImages ? $t('deleteImage.tip.delete') : $t('deleteImage.tip.text')}}：【{{imageInfo}}】？
            </span>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {imagesDeleteAPI} from 'api/image/request.ts'; // 删除镜像接口
import {deleteImageAPI} from 'api/model/request.ts';
import {SuccessType} from '@utils/publicType';

// props类
interface PropsType {
    hasRelationImages?: boolean;
    imageInfo: string
    deviceTypeId?: string
    imageId: string
    diaLog: boolean;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['error-click', 'dia-log-close', 'determine-click']);

class DeleteImageOperate {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestDeleteImage();
    };

    /**
     * 请求删除镜像接口，成功后把事件回传，关闭弹窗
    */
    requestDeleteImage = () => {
        const requestUrl = props.hasRelationImages ? imagesDeleteAPI : deleteImageAPI;
        const params = {
            deviceTypeId: props.deviceTypeId, // 设备ID
            imageId: props.imageId // 镜像ID
        };
        requestUrl(
            {
                ...params
            }
        ).then(({data}: {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            this.cancelClick();
            emitValue('determine-click');
        })
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
const deleteImageOperate = new DeleteImageOperate();

</script>
<style lang="scss">
@import './deleteImage.scss';
</style>
