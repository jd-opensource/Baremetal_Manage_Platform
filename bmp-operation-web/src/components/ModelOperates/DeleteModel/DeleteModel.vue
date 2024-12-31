<template>
    <!-- 删除机型操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'delete-model'"
        :is-loading="deleteModelOperate.isLoading"
        :header-title="$t('deleteModel.header.deleteModel')"
        :src="($defInfo.imagePath('model') as unknown as string)"
        @dia-log-close="deleteModelOperate.cancelClick"
        @determine-click="deleteModelOperate.determineClick"
    >
        <!-- 主体内容 -->
        <div class="delete-model-count">
            <p class="delete-model-count-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 删除提示 -->
                <span>{{$t('deleteModel.content.deleteTip')}}</span>
            </p>
            <!-- 删除确认 -->
            <p class="delete-model-count-title">
                <span>{{$t('deleteModel.content.deleteSure', [modelInfo])}}</span>
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {deleteModelAPI} from 'api/model/request.ts'; // 删除机型接口
import {language} from 'utils/publicClass.ts';
import {msgTip, methodsTotal} from 'utils/index.ts';

/**
 * props 类
*/
interface PropsType {
    modelInfo: string;
    deviceTypeId: string;
    diaLog: boolean;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    modelInfo: '', // 机型信息
    deviceTypeId: '', // deviceTypeId
    diaLog: false, // 蒙层
});

/**
 * defineEmits API 来替代 emits
*/
const emitValue = defineEmits(['error-click', 'dia-log-close', 'determine-click']);

class DeleteModelOperate {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestDeleteModel();
    };

    /**
     * 请求删除机型接口，成功后把事件回传，关闭弹窗
    */
    requestDeleteModel = () => {
        deleteModelAPI(
            {
                id: props.deviceTypeId
            }
        )
        .then(({data}: any) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                methodsTotal.sendMsg('model-delete-success', 'success');
                emitValue('determine-click');
                this.cancelClick();
            }
        }).finally(() => {
            this.isLoading.value = false;
        }).catch(() => {
            this.cancelClick();
            emitValue('error-click');
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = () => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

};
const deleteModelOperate = new DeleteModelOperate();

</script>
<style lang="scss">
@import './deleteModel.scss';
</style>
