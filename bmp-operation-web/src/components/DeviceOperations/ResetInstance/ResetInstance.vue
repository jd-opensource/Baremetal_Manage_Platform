<template>
    <!-- 重置实例操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'reset-instance'"
        :is-loading="resetInstance.isLoading"
        :header-title="$t('resetInstance.header.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="resetInstance.cancelClick"
        @determine-click="resetInstance.determineClick"
    >
        <div class="tip-count">
            <img
                class="tip-img"
                alt=""
                :src="($defInfo.imagePath('tip') as unknown as string)"
            />
            <span
                v-html="$t('resetInstance.tooltip.title')"
            >
            </span>
        </div>
        
        <div class="reset-instance-count">
            <span class="reset-instance-count-info">
                <!-- 确定将xxx重置实例状态？ -->
                {{$t('resetInstance.tooltip.count.tip1')}}
                <span>
                    【{{instanceName}}】
                </span>
                {{$t('resetInstance.tooltip.count.tip2')}}
            </span>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {SuccessType} from '@utils/publicType';
import {resetInstanceAPI} from 'api/device/request.ts';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    instanceName: string;
    instanceId: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

class ResetInstance {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.resetDevicesOperate();
    };

    /**
     * 重置实例操作
     * @param {Function} requestUrl url请求地址
    */
    resetDevicesOperate = () => {
        resetInstanceAPI(
            {
                instanceId: props.instanceId
            }
        )
        .then(({data}: {data: SuccessType}) => {
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
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
const resetInstance = new ResetInstance();

</script>
<style lang="scss">
@import './resetInstance.scss';
</style>
