<template>
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'currency clear-info-confirm'"
        :header-title="$t('clearRaidConfirm.header.title')"
        :src="($defInfo.imagePath('collect') as unknown as string)"
        :is-loading="clearRaidConfirmOpt.isLoading"
        @dia-log-close="clearRaidConfirmOpt.cancelClick"
        @determine-click="clearRaidConfirmOpt.determineClick"
    >
        <div class="currency-count">
          
            <div class="currency-count-title">
                <ui-icon class="icon">
                    <CircleCloseFilled />
                </ui-icon>
                <div class="text-count">
                    {{$t('clearRaidConfirm.tipsCount.tip0')}}
                    <p v-html="$t('clearRaidConfirm.tipsCount.tip1')"></p>
                </div>
            </div>
              <!-- headerTitle提示 -->
              <p class="header-title">
                <span>
                    {{$t('clearRaidConfirm.tipsCount.tip2', [sn])}}
                </span>
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {SuccessType} from '@utils/publicType';
import {msgTip} from 'utils/index.ts'; // message提示
import {language} from 'utils/publicClass.ts'; // 国际化
import {devicesCollectAPI} from 'api/device/request.ts';
// props 类
interface PropsType {
    diaLog: boolean;
    sn: string;
    checkboxStatus: boolean;
};

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {
    sn: ''
});


class ClearRaidConfirmOpt {
    isLoading: Ref<boolean> = ref<boolean>(false);

    determineClick = () => {
        this.isLoading.value = true;
        devicesCollectAPI({
            sn: props.sn,
            allowOverride: props.checkboxStatus
        })
        .then(({data}: {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            emitValue('determine-click')
            this.cancelClick();
        })
    }

    cancelClick = () => {
        emitValue('dia-log-close', false);
    }
}
const clearRaidConfirmOpt = new ClearRaidConfirmOpt();
</script>
<style lang="scss">
@import './clearRaidConfirm.scss';
</style>
