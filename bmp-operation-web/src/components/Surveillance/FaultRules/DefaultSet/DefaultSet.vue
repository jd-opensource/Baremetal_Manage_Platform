<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'default-set'"
        :is-loading="restoreDefaultSet.isLoading"
        :header-title="$t('defaultSet.header.title')"
        :src="($defInfo.imagePath('surveillance') as unknown as string)"
        @dia-log-close="restoreDefaultSet.cancelClick"
        @determine-click="restoreDefaultSet.determineClick"
    >

        <div class="default-set-content">
            <p class="warnning-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 提示 -->
                <span>{{$t('defaultSet.header.tip1')}}</span>
            </p>
            <p class="tip-title">
                {{$t('defaultSet.header.tip2')}}
            </p>
            <div class="currency-count-table">
                <ui-config-provider :locale="uiLocale.locale">
                    <!-- 表格数据 -->
                    <ui-table
                        border
                        style="width: 100%"
                        :max-height="180"
                        :class="[tableClass(selectArr, false)]"
                        :data="selectArr"
                    >
                        <!-- 故障名称 -->
                        <ui-table-column
                            prop="alertName"
                            align="center"
                            min-width="120"
                            :label="$t('defaultSet.label.faultName')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.alertName}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.alertName)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.alertName)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 故障类型 -->
                        <ui-table-column
                            prop="alertType"
                            min-width="130"
                            align="center"
                            :label="$t('defaultSet.label.faultType')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.alertType}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.alertType)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.alertType)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- 故障等级 -->
                        <ui-table-column
                            prop="alertLevel"
                            align="center"
                            min-width="100"
                            :label="$t('defaultSet.label.faultLevel')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.alertLevel)}}
                            </template>
                        </ui-table-column>
                        <!-- 故障描述 -->
                        <ui-table-column
                            prop="alertDesc"
                            align="center"
                            min-width="120"
                            :label="$t('defaultSet.label.faultDesc')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.alertDesc}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.alertDesc)"
                                    >
                                        <span>
                                            {{$filter.emptyFilter(scope.row.alertDesc)}}
                                        </span>    
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                    </ui-table>
                </ui-config-provider>
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {faultResetAPI} from 'api/surveillance/faultRule/request.ts';
import {msgTip, hasShowTooltip, tableClass} from 'utils/index.ts';
import {uiLocale, language} from 'utils/publicClass.ts';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    selectArr: any;
    status?: string;
};


interface ResponseType {
    result: {
        success: boolean;
    }
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    title: '', // 标题
    diaLog: false, // 蒙层
    selectArr: []
});

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
        await this.#requestRestoreDefaultSet();
    };

    #requestRestoreDefaultSet = () => {
        faultResetAPI(
            this.#toLine({ruleIds: props.selectArr.map((item: {ruleId: string}) => item.ruleId).join(',')})
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
@import './defaultSet.scss';
</style>
