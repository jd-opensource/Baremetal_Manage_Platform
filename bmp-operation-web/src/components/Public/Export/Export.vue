<template>
    <!-- 导出操作 -->
    <div
        class="export-data"
        v-if="diaLog"
    >
        <!-- 导出弹窗 -->
        <custom-dia-log
            :dia-log="diaLog"
            :header-title="$t('export.title')"
            :is-loading="exportData.isLoading"
            :src="($defInfo.imagePath('export') as unknown as string)"
            @dia-log-close="exportData.cancelClick"
            @determine-click="exportData.exportDataClick"
        >
            <!-- 主体内容 -->
            <div class="export-content radio-style1">
                <div class="export-flex">
                    <span class="export-scope">{{$t('export.content.scope')}}</span>
                    <ui-radio-group v-model="radio1">
                        <ui-radio label="1" size="large">
                            {{$t('export.content.all')}}
                        </ui-radio>
                        <!-- <ui-radio label="0" size="large" v-if="selectArr?.length">
                            {{$t('export.content.Selected')}}
                        </ui-radio> -->
                        <ui-radio
                            label="2"
                            size="large"
                            v-if="hasSearch"
                        >
                            {{$t('export.content.search')}}
                        </ui-radio>
                    </ui-radio-group>
                </div>
                <!-- <ui-progress
                    v-if="exportData.hasProgress.value"
                    :percentage="store.progressInfo.progressNum"
                    :color="[
                        {
                            color: '#108ee9',
                            percentage: store.progressInfo.progressNum
                        }
                    ]"
                /> -->
            </div>
        </custom-dia-log>
    </div>
</template>
  
<script lang="ts" setup>
import store from 'store/index.ts';
import {RouterOperate} from 'utils/publicClass.ts';
import {removeLong} from 'request/RemoveLongReq/remove.ts';

/**
 * props类
*/
interface PropsType {
    diaLog: boolean;
    hasSearch?: boolean;
    event?: {
        exportLoading: {
            value: boolean;
        };
        exportData(arg0: never): any;
       
    };
    search?: {
        paramsProcessing(arg0: never): void;
    };
    // selectArr?: any;
};

const radio1: Ref<string> = ref<string>('1');

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
 * @param {boolean} diaLog 弹窗状态
 * @param {boolean} hasEditOperate 是否进行编辑操作 默认true
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    hasSearch: false,
    // selectArr: () => []
});

// defineEmits API 来替代emit，进行事件回传
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// 路由
const routerOperate = new RouterOperate();

store.progressInfo.progressNum = 0;

/**
 * 导出
*/
class ExportDataOperate {
    isLoading: Ref<boolean> = ref<boolean>(false);
    hasProgress: Ref<boolean> = ref<boolean>(false);
    
    /**
     * 请求导出接口，成功后把事件回传，关闭弹窗
    */
    exportDataClick = (): void => {
        emitValue('determine-click');
        this.hasProgress.value = true;
        this.isLoading.value = true;
        if (radio1.value === '1') {
            // @ts-ignore
            props.event!.exportData(this);
            return;
        }
        // @ts-ignore
        props.search!.paramsProcessing(this);
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // store.progressInfo.progressNum = 0;
        this.isLoading.value = false;
        // this.hasProgress.value = false;
        emitValue('dia-log-close', false);
        // 移除请求
        removeLong(routerOperate.router.currentRoute.value.name);
    };
};
// 实例化
const exportData = new ExportDataOperate();

</script>
<style lang="scss">
@import './export.scss';

</style>
