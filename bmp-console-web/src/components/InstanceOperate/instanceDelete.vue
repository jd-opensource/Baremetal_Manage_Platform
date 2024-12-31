<template>
    <div class="small-dialog" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="list dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.delete')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="error-tip-content">
                <div class="error error-icon"></div>
                {{$t('list.content.deleteTip')}}
            </div>
            <div class="dialog-body-content">{{$t('list.content.delete', [instanceInfo?.instanceName])}}</div>
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
import {instanceDeleteAPI} from 'api/request.ts'; // 删除接口
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
const emitClose = defineEmits(['close', 'refresh']);

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
 * 请求关闭接口
*/
const requestDelete = (): void => {
    instanceDeleteAPI({
        instanceId: props.instanceInfo.instanceId
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            closeDialog()
            emitClose('refresh')
        }
    }).catch(() => {
        closeDialog()
        emitClose('refresh');
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 点击确定
 */
const openDialog: () => void = () => {
    isLoading.value = true;
    requestDelete()
}
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';
</style>