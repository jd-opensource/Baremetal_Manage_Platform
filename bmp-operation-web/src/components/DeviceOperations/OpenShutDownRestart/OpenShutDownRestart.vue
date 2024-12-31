<template>
    <!-- 开关机重启操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'open-shut-down-restart'"
        :is-loading="openCloseRestart.isLoading"
        :header-title="openCloseRestart.titleTip()"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="openCloseRestart.cancelClick"
        @determine-click="openCloseRestart.determineClick"
    >
        <div class="open-shut-down-restart-count">
            <span class="open-shut-down-restart-count-title">
                {{openCloseRestart.tooltipInfo()}}
            </span>
            <span class="open-shut-down-restart-count-info">
                {{instanceInfo}}
            </span>
            <span>？</span>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {CurrencyType, SuccessType} from '@utils/publicType';
import {devicesStartAPI, devicesStopAPI, devicesRestartAPI} from 'api/device/request.ts';

/**
 * props 类
*/
interface PropsType {
    title: string;
    diaLog: boolean;
    instanceId: string;
    instanceInfo: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

class OpenCloseRestart {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    titleTip = () => {
        const titleTip: CurrencyType = {
            '开机': language.t('openShutDownRestart.titleTip.open'),
            '关机': language.t('openShutDownRestart.titleTip.close'),
            '重启': language.t('openShutDownRestart.titleTip.restart')
        };
        return titleTip[props.title];
    };

    /**
     * 开/关/重启操作返回对应的提示文案
     * @return {string} xxx 返回对应的提示信息
    */
    tooltipInfo = () => {
        const tooltipInfo: {[x: string]: string} = {
            '开机': language.t('openShutDownRestart.tooltipInfo.open'),
            '关机': language.t('openShutDownRestart.tooltipInfo.close'),
            '重启': language.t('openShutDownRestart.tooltipInfo.restart')
        };
        return tooltipInfo[props.title];
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestOpenDownRestart();
    };

    /**
     * 请求开关机重启接口，成功后把事件回传，关闭弹窗
    */
    requestOpenDownRestart = () => {
        const requestUrl = new Map([
            ['关机', devicesStopAPI],
            ['重启', devicesRestartAPI],
            ['开机', devicesStartAPI]
        ]);
        return this.#getDevicesOperate(requestUrl.get(props.title))!;
    };

    /**
     * 设备操作-开机、关机、重启
     * @param {Function} requestUrl url请求地址
    */
    #getDevicesOperate = (requestUrl: (arg0: {instanceId: string}) => Promise<{data: SuccessType}>) => {
        requestUrl(
            {
                instanceId: props.instanceId
            }
        )
        .then(({data} : {data: SuccessType}) => {
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

const openCloseRestart = new OpenCloseRestart();

</script>
<style lang="scss">
@import './openShutDownRestart.scss';
</style>
