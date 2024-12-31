<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'fault-push'"
        :is-loading="restoreDefaultSet.isLoading"
        :header-title="$t('faultPush.header.title')"
        :src="($defInfo.imagePath('surveillance') as unknown as string)"
        @dia-log-close="restoreDefaultSet.cancelClick"
        @determine-click="restoreDefaultSet.determineClick"
    >

        <div class="fault-push-content">
            <p class="tip-title set-word" v-if="status">
                {{$t('faultPush.header.tip1', [alertName])}}
            </p>
            <p class="tip-title set-word" v-else>
                {{$t('faultPush.header.tip2', [alertName])}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {faultPushAPI} from 'api/surveillance/faultRule/request.ts';
import {msgTip, methodsTotal} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    status: boolean;
    alertName: string;
    roleId: number;
};


interface ResponseType {
    result: {
        success: boolean;
    }
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

/**
 * defineEmits API 用来代替emit
*/
const emitValue = defineEmits(['dia-log-close', 'determine-click']);


class RestoreDefaultSet {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.#requestRaultPush();
    };

    #requestRaultPush = () => {
        faultPushAPI(
            methodsTotal.toLine({pushStatus: props.status ? 1 : 0, ruleId: props.roleId})
        )
        .then(({data} : {data: ResponseType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        }).finally(() => {
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

const restoreDefaultSet = new RestoreDefaultSet();
</script>
<style lang="scss">
@import './faultPush.scss';
</style>
