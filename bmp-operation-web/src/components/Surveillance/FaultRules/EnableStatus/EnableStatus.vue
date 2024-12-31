<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'enable-status'"
        :is-loading="restoreDefaultSet.isLoading"
        :header-title="$t('faultOperation.header.title')"
        :src="($defInfo.imagePath('surveillance') as unknown as string)"
        @dia-log-close="restoreDefaultSet.cancelClick"
        @determine-click="restoreDefaultSet.determineClick"
    >

        <div class="enable-status-content">
            <p class="tip-title set-word" v-if="status">
                {{$t('faultOperation.header.tip1', [alertName])}}
            </p>
            <p class="tip-title set-word" v-else>
                {{$t('faultOperation.header.tip2', [alertName])}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {useI18n} from 'vue-i18n'; // 使用国际化
import {msgTip} from 'utils/index.ts';
import {faultUseAPI} from 'api/surveillance/faultRule/request.ts';

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
 * 使用国际化
*/
const {t} = useI18n();


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
        await this.#requestUseStatus();
    };

    #requestUseStatus = () => {
        faultUseAPI(
            this.#toLine({useStatus: props.status ? 1 : 0, ruleId: props.roleId})
        )
        .then(({data} : {data: ResponseType}) => {
            if (data?.result?.success) {
                msgTip.success(t('operate.success'));
            }
        }).finally(() => {
            this.isLoading.value = false;
            this.cancelClick();
            emitValue('determine-click');
        })
    };

    #toLine(params: any) {
        const conversionData = [
            [
                (params: string) => typeof params === 'string',
                () => this.#humpConversion(params)
            ],
            [
                (params: {[x: string]: string | number}) => params instanceof Object,
                () => {
                    const obj = [];
                    for (const index of Object.keys(params)) {
                        obj.push(
                            {
                                [this.#humpConversion(index)]: params[index]
                            }
                        )
                    }
                    return Object.assign({}, ...obj);
                }
            ]
        ];

        for (const key of conversionData) {
            if (key[0](params)) {
                return key[1](params);
            }
        }
    }

    #humpConversion = (params: string) => {
        return params.replace(/([A-Z])/g, '_$1').toLowerCase()
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
@import './enableStatus.scss';
</style>
