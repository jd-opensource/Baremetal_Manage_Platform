<template>
    <div class="small-dialog ta-l" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="desc dialog-icon"></div>
                    <span class="my-title">{{$t('instance.detail.description')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="dialog-body-content">
                <p style="margin-bottom: 10PX;">{{$t('instance.detail.editDescription') + $filter.withClon(' ')}}</p>  
                <div>
                    <el-input
                        v-model="description"
                        maxlength="256"
                        :rows="5"
                        type="textarea"
                        show-word-limit
                        :placeholder="$t('list.placeholder.description')"
                    />
                </div>               
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
import {Ref, ref, watch, onMounted} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import {editInstanceAPI} from 'api/request.ts'; // 编辑描述接口
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
 * 编辑项目时，项目名称赋值
 */
onMounted(() => {
    description.value = Object.keys(props.instanceInfo).length ? props.instanceInfo.description : ''
})
/**
 * 描述信息
 */
const description: Ref<string> = ref<string>('')
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
    emitClose('close', false);
}

/**
 * 请求编辑描述接口
*/
const requestClose = (): void => {
    editInstanceAPI({
        instanceId: props.instanceInfo.instanceId,
        idcId: props.instanceInfo.idcId,
        description: description.value
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
        emitClose('refresh');
        closeDialog()
    }).finally(() => {
        isLoading.value = false;
    });
};

document.onkeydown = (event: {keyCode: number;}) => {
    event.keyCode === 13 && openDialog();
};

/**
 * 点击确定
 */
const openDialog= () => {
    isLoading.value = true;
    requestClose()
}
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';


.ta-l {
    .el-dialog {
        .el-dialog__body {
            text-align: left !important;
            padding-bottom: 0;
            .dialog-body-content {
                margin-top: 30px !important;
            }

        }
    }
}

</style>