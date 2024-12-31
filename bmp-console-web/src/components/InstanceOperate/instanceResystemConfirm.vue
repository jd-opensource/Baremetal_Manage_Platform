<template>
    <div class="middle-dialog resystem-confirm" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="list dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.resystemSure')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="error-tip-content display-flex">
                <div class="error error-icon mt6"></div>
                <span class="notice-change">{{$t('list.content.notice')}} </span>
                <div class="ta-l lh24" v-html="$t('list.content.resystemTips')"></div>   
            </div>
            <div class="dialog-body-content mt45">{{$t('list.content.resystemContent', [instanceInfo?.instanceName])}}</div>
            <template #footer>
                <span>
                    <el-button      
                        class="cancel-button"
                        @click="closeDialog"
                    >
                        {{$t('list.cancel')}}
                    </el-button>
                    <el-button 
                        type="primary" 
                        :loading="isLoading" 
                        @click="openDialog"
                    >
                        {{$t('list.confirm')}}
                    </el-button>
                </span>
            </template>    
        </el-dialog>
    </div>
</template>
<script setup lang="ts">
import {Ref, ref, watch} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    instanceInfo: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    instanceInfo: {}
});

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 回传对话框关闭
 */
const emitClose = defineEmits(['close', 'refresh','confirm']);

/**
 * 监听蒙层开关
 * @param {boolean} newValue 蒙层显示隐藏 true false
*/
watch(() => props.openVisible, (newValue: Required<boolean>): Readonly<void> => {  
    emitClose('close', newValue);
});

/**
 * 点击取消
 */
const closeDialog:  () => void = () => {
    emitClose('close', false)
}

/**
 * 点击确定
 */
const openDialog: () => void = () => {
    emitClose('close', false)
    emitClose('confirm', props.instanceInfo)
    isLoading.value = true;
    
}
</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/icon';

.resystem-confirm {
    .el-dialog {
        .el-dialog__body {
            text-align: center !important;
        }
    }

    .notice-change {
        word-break: keep-all;
        margin: 4px 5px 0 0;
        color: #FF4D4F;
    }
}
</style>