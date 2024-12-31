<template>
    <!-- 锁定、解锁操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'currency lock-instance'"
        :is-loading="lockOperate.isLoading"
        :header-title="lockOperate.tipTitle()"
        :src="($defInfo.imagePath('lock') as unknown as string)"
        @dia-log-close="lockOperate.cancelClick"
        @determine-click="lockOperate.determineClick"
    >
        <!-- 主体内容 -->
        <div class="currency-count">
            <p
                class="lock-count"
                v-if="!operate.localeCompare('lock')"
            >
                <img
                    class="lock-count-icon"
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 锁定实例后“删除”操作不可用，如需使用需提前解锁 -->
                <span class="lock-count-tip">
                    {{$t('lockInstance.tip')}}
                </span>
            </p>
            <!-- 锁定、解锁提示 -->
            <p class="currency-count-title">
                {{lockOperate.tipContent()}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script setup lang="ts">
import {msgTip} from 'utils/index.ts'; // message提示
import {language} from 'utils/publicClass.ts'; // 国际化
import {SuccessType} from '@utils/publicType';
import {lockAPI, unlockAPI} from 'api/device/request.ts'; // 锁定、解锁接口

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    instanceName: string;
    operate: string;
    instanceId: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

class LockUnLockOperate {
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick = (): void => {
        this.isLoading.value = true;
        this.requestLockUnLock(this.#setUrl());
    };

    #setUrl = () => {
        const urlInfo = new Map([
            ['lock', lockAPI]
        ]);
        return urlInfo.get(props.operate)?? unlockAPI;
    };

    tipTitle = () => {
        const info = new Map([
            ['lock', language.t('lockInstance.header.lock')]
        ]);
        return info.get(props.operate)?? language.t('lockInstance.header.unlock')
    };

    tipContent = () => {
        const info = new Map([
            ['lock', language.t('lockInstance.lockTip', [props.instanceName])]
        ]);
        return info.get(props.operate)?? language.t('lockInstance.unlockTip', [props.instanceName])
    };

    /**
     * 请求锁定、解锁接口，成功后把事件回传，关闭弹窗
    */
    requestLockUnLock = (url: (arg0: {instanceId: string;}) => Promise<{data: SuccessType;}>): void => {
        url(
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
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
const lockOperate = new LockUnLockOperate();

</script>
<style lang="scss">
@import './lock.scss';
</style>
