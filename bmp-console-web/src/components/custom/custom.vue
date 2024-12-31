<template>
    <div class="small-dialog smaller-dialog no-center" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="settings dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.customList')}}</span>
                    <el-button      
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog"
                    >
                    </el-button>
                </div>
            </template>
            <div class="custom-list-count">
                <p class="custom-list-count-title">{{$t('list.content.customListMessage')}}</p>
                <div class="custom-list-count-checkbox">
                    <el-checkbox
                        v-for="(item, index) in checkListArr"
                        v-model="item.selected"
                        :key="index"
                        :label="item.name"
                        :disabled="item.disabled"
                        :title="item.name"
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
import {Ref, ref, watch} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {deepCopy} from 'utils/index.ts';
import {setCustomListAPI, setAlarmCustomListAPI} from 'api/request.ts'; // 自定义列表数据接口

/**
 * 需要显示的列表信息类
*/
type CheckListArrType = {
    name?: string;
    selected: boolean;
    disabled: boolean;
    label?: string;
};

/**
 * props类
*/
interface PropsType {
    pageKey?: string;
    openVisible: boolean;
    checkListArr: CheckListArrType[];
    hasCustomInfo: any;
    sign?:string;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    pageKey: 'projectList',
    openVisible: false, // 蒙层显示
    checkListArr: () => [], // 需要显示的列表
    hasCustomInfo: () => {},
    sign: '' //判断是否是报警日志列表
});

/**
 * loading加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 回传对话框关闭,事件回传
 */
const emitClose = defineEmits(['close', 'determine']);

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
 * 请求自定义列表接口
 */
const setCustomList = (): void => {
    // 深拷贝数据-自定义列表信息
    const pageValue = deepCopy(props.hasCustomInfo);
    const page_value = deepCopy(props.hasCustomInfo);
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
    }
    for (const key of props.checkListArr) {
        for (const index in page_value) {
            // 点击的当前项等于对应数据的那一项
            if (key.label === index) {
                // 状态赋值
                page_value[index].required = key.disabled;
                page_value[index].selected = key.selected;
            }
        }
    }
    if(props.sign === 'alarm') {
        setAlarmCustomListAPI(
            {
                page_key: props.pageKey,
                page_value
            }
        ).then(({data} : {data: any}) => {
            if (data?.result?.success) {
                closeDialog();
                emitClose('determine');
                return;
            }
        }).finally(() => {
            isLoading.value = false;
        });
    } else {
        setCustomListAPI(    
            {
                pageKey: props.pageKey,
                pageValue
            }
        ).then(({data} : {data: any}) => {
            if (data?.result?.success) {
                closeDialog();
                emitClose('determine');
                return;
            }
        }).finally(() => {
            isLoading.value = false;
        });
    }
    
}
/**
 * 点击确定
 */
const openDialog = (): void => {
    isLoading.value = true;
    setCustomList();
}
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';
.smaller-dialog{
    .el-dialog {
        --el-dialog-width: 446px !important;
        height: auto;
        .el-dialog__body {
            height: 142px !important;
            padding: 0 20px 20px 20px !important;
        }
    }
}
.custom-list-count {
    &-title {
        font-size: 12px;
        color: #333;
        font-weight: 500;
        text-align: left;
        margin: 19.5px 0 16px;
    }
    &-checkbox {
        display: flex;
        flex-wrap: wrap;
        width: 100%;
        .el-checkbox {
            width: 25%;
            height: 17px;
            margin-right: 0px;
            margin-bottom: 16px;
            display: flex !important;
            align-items: flex-start !important;
        }
        .is-checked,
        .is-focus {
            .el-checkbox__inner {
                border: 1px solid #108EF9;
            }
        }   
        .el-checkbox__inner {
            width: 10px;
            height: 10px;
            margin-top: 1px;
            border-radius: 2px;
            border: 1px solid #333;

            &:hover,
            &:focus,
            &:active {
                border-color: #108EF9  !important;
            }

            &::after {
                height: 5px;
                left: 2px;
                top: 0px;
            }
        }

        .is-disabled {

            .el-checkbox__label {
                font-size: 12px;
                color: #999 !important;
                font-weight: 400;
            }
            .el-checkbox__inner {
                background: #999 !important;
                border: 1px solid #999 !important;
                border-radius: 2px;

                &:hover,
                &:focus,
                &:active {
                    border-color: #999 !important;
                }

                &::after {
                    border-color: #fff !important;
                }
            }
        }

        .el-checkbox__label {
            overflow: hidden;
            text-overflow: ellipsis;
            font-size: 12px;
            color: #333 !important;
            font-weight: 400;
            max-width: 80px;
        }
    }
}
</style>