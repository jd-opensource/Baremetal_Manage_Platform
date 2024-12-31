<template>
    <!-- 移除机型操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'remove-model'"
        :is-loading="removeOperate.isLoading"
        :header-title="$t('removeOperate.header.remove')"
        :src="($defInfo.imagePath('imageManage') as unknown as string)"
        @dia-log-close="removeOperate.cancelClick"
        @determine-click="removeOperate.determineClick"
    >
        <!-- 主体内容 -->
        <div class="remove-model-count">
            <p class="remove-model-count-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 移除提示 -->
                <span>
                    {{$t('removeOperate.content.deleteTip')}}
                </span>
            </p>
            <!-- 移除确认 -->
            <p class="remove-model-count-title">
                {{$t('removeOperate.content.tip1', [imageName, modelName, '(' + deviceType + ')'])}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts';
import {deleteImageModelAPI} from 'api/image/request.ts';
import {language} from 'utils/publicClass.ts';

// props 类
interface PropsType {
    imageName: string;
    modelName: string;
    imageId: string;
    deviceType: string;
    deviceTypeId: string;
    diaLog: boolean;
};

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

class RemoveOperate {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick = () => {
        this.isLoading.value = true;
        this.requestDeleteModel();
    };

    /**
     * 请求移除接口，成功后把事件回传，关闭弹窗
    */
    requestDeleteModel = () => {
        deleteImageModelAPI(
            {
                imageId: props.imageId,
                deviceTypeId: props.deviceTypeId
            }
        )
        .then(({data} : {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                emitValue('determine-click');
                this.cancelClick();
            }
        })
        .finally(() => {
            this.isLoading.value = false;
        })
        .catch(() => {
            emitValue('determine-click');
            this.cancelClick();
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = () => {
        // 关闭蒙层
        emitValue('dia-log-close', false);
    };
};

const removeOperate: RemoveOperate = new RemoveOperate();

</script>
<style lang="scss">
@import './removeModel.scss';
</style>
