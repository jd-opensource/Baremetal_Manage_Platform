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
                    <div class="logout dialog-icon"></div>
                    <span class="my-title">{{$t('personCentre.logout')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="dialog-body-content">{{$t('personCentre.content.logOut')}}</div>
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
                        class="confirm-button"
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
import {ref, Ref, watch} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {logOutAPI} from 'api/request.ts'; // 关闭接口
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false

});

/**
 * 回传对话框关闭
 */
const emitClose = defineEmits(['close']);

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

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
 * 请求关闭接口
*/
const requestClose = (): void => {
    logOutAPI({}).then(({data}: {data: any}) => {
        if (data?.result?.success) {
            window.open('/Login/login', '_self');
        }
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 点击确定
 */
const openDialog = () => {
    isLoading.value = true;
    requestClose()
}
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';
</style>