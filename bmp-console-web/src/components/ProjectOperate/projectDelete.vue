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
                    <div class="folder dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.deleteProject')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="dialog-body-content">{{$t('list.content.deleteProject', [instanceInfo?.projectName])}}</div>
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
import {projectDeleteAPI} from 'api/request.ts'; // 项目删除接口
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
    instanceInfo: {}
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
 * 请求删除接口
*/
const requestDelete = (): void => {
    projectDeleteAPI({
        projectId: props.instanceInfo.projectId,
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            emitClose('refresh');
            closeDialog()
        }
    }).catch(() => {
        emitClose('refresh');
        closeDialog()
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 点击确定
 */
 const openDialog: () => Promise<void> = async () => {
    isLoading.value = true;
    await requestDelete();
};
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';
</style>