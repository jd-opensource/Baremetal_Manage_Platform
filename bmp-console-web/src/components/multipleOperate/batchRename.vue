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
                    <span class="my-title">{{$t('list.operate.bacthRename')}}</span>
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
                <div class="error-tip-content batch-tip-error">
                    <div class="error error-icon"></div>
                    {{$t('list.content.bacthRenameTip')}}
                </div>
                <batch-table
                    v-bind="$attrs" 
                    batchTip="rename"
                    :multipleSelection="instanceInfo"
                >
                </batch-table>
                <el-form
                    label-position="left"
                    :model="instanceForm"
                    ref="ruleFormRef"
                    :rules="rules"
                    label-width="110px"
                    style="margin-top: 20px;"
                    @submit.native.prevent
                >
                    <!-- 实例名称 -->
                    <el-form-item 
                        prop="newInstanceName"
                        label-width="125"
                        :label="$t('instance.list.newInstanceName') "
                    >
                        <el-input 
                            v-model="instanceForm.newInstanceName" 
                            maxlength="128"
                            :placeholder="$t('list.placeholder.instanceName')"
                        /> 
                        <p 
                            class="f12 c999 el-form-item__error" 
                            v-if="!instanceNameState"
                        >
                            {{$t('personCentre.content.instanceVerification')}}
                        </p>                          
                    </el-form-item> 
                </el-form>
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
import {Ref, ref, watch, reactive, nextTick} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    instanceNameReg, // 实例名称正则
} from 'utils/regular.ts';
import type {
    FormInstance,
    FormRules // element-plus ui类
} from 'element-plus';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import {batchRenameAPI} from 'api/request.ts'; // 批量重启接口
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
    allInstance: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    instanceInfo: [],
    allInstance: []
});

/**
 * 回传对话框关闭
 */
const emitClose = defineEmits(['close', 'refresh']);

/**
 * 实例类型
 */
 interface Instance {
    newInstanceName: string, // 实例名称
}
/**
 * 实例表单
 */
const instanceForm:Instance = reactive({
    newInstanceName: '', // 实例名称
})

/**
 * 用户信息类
*/
type UserInfoType = {
    required: boolean;
    message: string;
    trigger: string;
    validator?: unknown;
};

/**
 * 规则类
*/
interface RulesType {
    newInstanceName: UserInfoType[],//实例名称
};

/**
 * 实例名称状态
 */
 const instanceNameState: Ref<boolean> = ref<boolean>(false)

/**
 * 实例名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
 * 
*/
const instanceNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    instanceNameState.value = true
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyInstanceName')));
    }
    else if (props.allInstance.map((item:any) => item.instanceName).includes(value)) {
        callback(new Error(t('personCentre.content.instanceNameExist')));
    }
    else if(!instanceNameReg.test(value)) {
        callback(new Error(t('personCentre.content.instanceVerification')));
    }else {
        callback();
    }
};

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    newInstanceName: [ // 实例名称
        {
            required: false,
            trigger: 'blur',
            validator: instanceNameCheck
        }
    ]
});

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
    instanceForm.newInstanceName = '';
    instanceNameState.value = false; 
    emitClose('close', false) 
}

/**
 * 请求批量编辑实例名称接口
*/
const requestBatchRename = (): void => {
    let instanceIds = props.instanceInfo.map((item: any) => item.instanceId)
    batchRenameAPI({
        instanceIds: instanceIds,
        instanceName: instanceForm.newInstanceName
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
    })
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 点击确定
 */
 const openDialog: () => Promise<void> = async () => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        ruleFormRef.value.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                isLoading.value = true;
                requestBatchRename();
            }
        });
    });
};

document.onkeydown = (event: {keyCode: number;}) => {
    event.keyCode === 13 && openDialog();
};
</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/batch-table';
@import 'assets/css/icon';
.batch-tip-error {
    font-size: 12px;
    margin-bottom: 20px;
    text-align: center;
}

</style>