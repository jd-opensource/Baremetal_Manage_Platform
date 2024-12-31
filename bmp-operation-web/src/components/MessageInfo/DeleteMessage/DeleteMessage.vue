<template>
    <!-- 删除消息操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'delete-message'"
        :is-loading="deleteMessageOperate.isLoading"
        :header-title="$t('messageDelete.header.messageDelete')"
        :src="($defInfo.imagePath('messageLight') as unknown as string)"
        @dia-log-close="deleteMessageOperate.cancelClick"
        @determine-click="deleteMessageOperate.determineClick"
    >
        <!-- 主体内容 -->
        <div class="delete-message-count">
            <!-- 删除确认 -->
            <p class="delete-message-count-title">
                {{$t('messageDelete.tips')}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {language, CurrentInstance} from 'utils/publicClass.ts';
import {msgTip} from 'utils/index.ts';

/**
 * props 类
*/
interface PropsType {
    data: {messageId: string}[];
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

class DeleteMessageOperate {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    proxy = new CurrentInstance().proxy;

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
    requestDeleteModel = async () => {
        try {
            const res = await this.proxy.$messageApi.deleteAPI({messageIds: props.data.map((item: {messageId: string}) => item.messageId)});
            if (res?.data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        }
        finally {
            this.isLoading.value = false;
            emitValue('determine-click');
            this.cancelClick();
        }
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = () => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

};
const deleteMessageOperate = new DeleteMessageOperate();

</script>
<style lang="scss">
@import './deleteMessage.scss';
</style>
