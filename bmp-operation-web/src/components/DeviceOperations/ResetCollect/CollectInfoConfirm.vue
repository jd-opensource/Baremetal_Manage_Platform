<template>
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'currency collect-info-confirm'"
        :header-title="$t('collectInfoConfirm.header.title')"
        :src="($defInfo.imagePath('collect') as unknown as string)"
        :is-loading="collectInfoConfirmOpt.isLoading"
        @dia-log-close="collectInfoConfirmOpt.cancelClick"
        @determine-click="collectInfoConfirmOpt.determineClick"
    >
        <div class="currency-count">
            <!-- headerTitle提示 -->
            <p class="header-title">
                <span>{{$t('collectInfoConfirm.tips')}}</span>
            </p>
            <ui-checkbox
                size="small"
                class="collect-info-confirm-checkbox"
                v-if="!['2', ''].includes(hasClearRaid)"
                v-model="collectInfoConfirmOpt.checkboxStatus.value"
            >
                <span class="text-tip"> {{$t('collectInfoConfirm.checkbox')}} </span>
            </ui-checkbox>
            <div
                class="text-count"
                v-if="collectInfoConfirmOpt.checkboxStatus.value"
            >
                {{$t('collectInfoConfirm.tipsCount.tip0')}}
                <p v-html="$t('collectInfoConfirm.tipsCount.tip1')"></p>
            </div>
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
    hasClearRaid: string;
}

// defineEmits API 用来代替emit
const emitValue = defineEmits(['opt-click', "dia-log-close", "determine-click"]);

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {
    sn: '',
    hasClearRaid: ''
});

class CollectInfoConfirmOpt {
    isLoading: Ref<boolean> = ref<boolean>(false);
    checkboxStatus: Ref<boolean> = ref<boolean>(false);

    determineClick = () => {
        if (!this.checkboxStatus.value) {
            this.resetCollect();
            return;
        }
        emitValue("determine-click",this.checkboxStatus.value);
        this.cancelClick();
    };

    resetCollect = () => {
        this.isLoading.value = true;
        devicesCollectAPI({
            sn: props.sn,
            allowOverride: false
        })
        .then(({data}: {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            emitValue('opt-click')
            this.cancelClick();
        })
    }

    cancelClick = () => {
        emitValue("dia-log-close", false);
    };
}
const collectInfoConfirmOpt = new CollectInfoConfirmOpt();
</script>
<style lang="scss">
@import "./index.scss";
</style>
