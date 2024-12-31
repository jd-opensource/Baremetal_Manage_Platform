<template>
    <div class="middle-dialog" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="list dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.bacthRestart')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="batch-table-content m0">
                <batch-table
                    v-bind="$attrs" 
                    batchTip="restart"
                    :multipleSelection="instanceInfo"
                >
                </batch-table>
            </div>
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
import {Ref, ref,watch} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import {batchRestartAPI} from 'api/request.ts'; // 批量重启接口
import batchTable from 'components/multipleOperate/batchTable.vue'; //实例选择表格
/**
 * 国际化
*/
const {t} = useI18n();

/**
 * loading加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);
    
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    instanceInfo: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    instanceInfo: []
});

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
 * 请求批量重启接口
*/
const requestBatchRestart = (): void => {
    let instanceIds = props.instanceInfo.map((item: any) => item.instanceId)
    batchRestartAPI({
        instanceIds: instanceIds,
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            closeDialog()
            emitClose('refresh');    
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
 const openDialog: () => Promise<void> = async () => {
    isLoading.value = true;
    await requestBatchRestart();
};
</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/batch-table';
@import 'assets/css/icon';
</style>