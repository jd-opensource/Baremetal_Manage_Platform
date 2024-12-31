<template>
    <!-- 自定义列表操作 -->
    <div
        class="custom-list"
        v-if="diaLog"
    >
        <custom-dia-log
            :dia-log="diaLog"
            :is-loading="customListOperate.isLoading"
            :header-title="$t('customList.header.name')"
            :src="($defInfo.imagePath('setUp') as unknown as string)"
            @dia-log-close="customListOperate.cancelClick"
            @determine-click="customListOperate.determineClick"
        >
            <!-- 主体内容 -->
            <div class="custom-list-count">
                <p class="custom-list-count-title">
                    {{$filter.withClon($t('customList.text.tip'))}}
                </p>
                <!-- 可以选择需要展示的列表信息 -->
                <div class="unify-checkbox">
                    <ui-tooltip
                        placement="bottom"
                        v-for="(item, index) in checkListArr"
                        v-model="item.showTooltip"
                        :disabled="!item.showTooltip"
                    >
                        <template #content>
                            {{item.name}}
                        </template>
                        <div
                            class="checkbox-text-ellipsis"
                            @mouseenter="hasShowTooltip($event, item, item.name, .725)"
                        >
                            <ui-checkbox
                                v-model="item.selected"
                                :key="index"
                                :label="item.name"
                                :disabled="item.disabled"
                            >
                                {{item.name}}
                            </ui-checkbox>
                        </div>
                    </ui-tooltip>
                </div>
            </div>
        </custom-dia-log>
    </div>
</template>
  
<script lang="ts" setup>
import {language} from 'utils/publicClass.ts';
import {ParamsType, CheckListArrType, HasCustomInfoType} from './type';
import {deepCopy, hasShowTooltip, msgTip, methodsTotal} from 'utils/index.ts';
import {setCustomListAPI, surveillanceSetCustomListAPI} from 'api/public/request.ts';


/**
 * props类
*/
type PropsType = {
    pageKey?: string;
    diaLog: boolean;
    checkListArr: CheckListArrType[];
    hasCustomInfo: HasCustomInfoType;
};


// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    pageKey: 'idcList',
    diaLog: false, // 蒙层显示
    checkListArr: () => [], // 需要显示的列表
    hasCustomInfo: () => []
});

// defineEmits API 来替代 emits，子传父，事件回传
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

/**
 * 自定义列表操作
*/
class CustomListOperate {
    // loading加载态，默认false，用于发送请求之前的loading
    isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 确定按钮弹窗
    */
    determineClick = (): void => {
        this.isLoading.value = true;
        this.#setParams();
    };

    /**
     * 请求自定义列表接口，成功后把事件回传，关闭弹窗
    */
    #setParams = (): void => {
        // 深拷贝数据-自定义列表信息
        const pageValue: HasCustomInfoType = deepCopy(props.hasCustomInfo);
        // 遍历自定义列表数据
        for (const key of props.checkListArr) {
            for (const index in pageValue) {
                // 点击的当前项等于对应数据的那一项
                if (key.label === index) {
                    // 状态赋值
                    pageValue[index].required = key.disabled;
                    pageValue[index].selected = key.selected;
                }
            }
        };
        let params: ParamsType;
        nextTick(() => {
            this.#setApi(params, pageValue);
        })
    };

    #setApi = (params: ParamsType, pageValue: HasCustomInfoType) => {
        let requestAPI: Function;
        const requestParamsAPI = [
            [
                (key: string) => key.indexOf('_') > -1,
                () => {
                    params = methodsTotal.toLine({pageKey: props.pageKey, pageValue: pageValue});
                    requestAPI = surveillanceSetCustomListAPI;
                }
            ],
            [
                (key: string) => !(key.includes('_')),
                () => {
                    params = {pageKey: props.pageKey, pageValue};
                    requestAPI = setCustomListAPI;
                }
            ]
        ];
        for (const key of requestParamsAPI) {
            if (key[0](props.pageKey!)) {
                key[1](props.pageKey!);
                break;
            }
        }
        nextTick(() => {
            this.#requestSetCustom(params, requestAPI);
        })
    };

    #requestSetCustom = (params: ParamsType, requestAPI: Function) => {
        nextTick(() => {
            requestAPI(
                {
                    ...params
                }
            ).then(({data} : {data: {result: {success: boolean;}}}) => {
                if (data?.result?.success) {
                    msgTip.success(language.t('operate.success'));
                    emitValue('dia-log-close', false);
                    emitValue('determine-click');
                    return;
                }
            }).finally(() => {
                this.isLoading.value = false;
            });
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 关闭蒙层
        emitValue('dia-log-close');
        this.isLoading.value = false;
    };
};
// 实例化-自定义列表操作
const customListOperate: CustomListOperate = new CustomListOperate();

</script>
<style lang="scss">
@import './customList.scss';
</style>
