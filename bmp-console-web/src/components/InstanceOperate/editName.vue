<template>
    <div class="small-dialog no-center" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="desc dialog-icon"></div>
                    <span class="my-title">{{$t('instance.detail.instanceName')}}</span>
                    <el-button                      
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog(ruleFormRef)"
                    >
                    </el-button>
                </div>
            </template>
            <div class="dialog-body-content">
                <el-form
                    label-position="left"
                    :model="instanceForm"
                    ref="ruleFormRef"
                    :rules="rules"
                    label-width="auto"
                    @submit.native.prevent
                >
                    <el-form-item 
                        prop="instanceName"
                        :label="$t('instance.detail.instanceName') "
                    >
                        <el-input 
                            v-model="instanceForm.instanceName"
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
                        @click="closeDialog(ruleFormRef)"
                    >
                        {{$t('list.cancel')}}
                    </el-button>
                    <el-button 
                        type="primary" 
                        :loading="isLoading"
                        @click="openDialog(ruleFormRef)"
                    >
                        {{$t('list.confirm')}}
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>
<script setup lang="ts">
import {
    Ref, 
    ref, 
    reactive, 
    watch, 
    nextTick, 
    onMounted
} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    instanceNameReg // 实例名称正则
} from 'utils/regular.ts';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {instanceListAPI, editInstanceAPI} from 'api/request.ts'; // 实例创建接口
import VueCookies from 'vue-cookies'; // cookie

/**
 * 国际化
*/
const {t} = useI18n();
const instanceForm = reactive({
    instanceName: '',
})
/**
 * 用户ID
*/
let userId: string = (VueCookies as any).get('user_id');

/**
 * 用户ID解码
 */
userId = window.atob(userId);

/**
 * loading加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 实例名称类型
 */
 interface instancename {
    repeatInstanceData: any
}
/**
 * 实例名称数据
 */
 const instanceData: instancename = reactive<instancename>({
    repeatInstanceData: []
})
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    instanceInfo?: any;
};
/**
 * prop组件传值
 */
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    instanceInfo: {}
});

/**
 * 实例名称状态
 */
 const instanceNameState: Ref<boolean> = ref<boolean>(false)

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
    instanceName: UserInfoType[];
};
/**
 * 实例名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const instanceNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        instanceNameState.value = true
        callback(new Error(t('personCentre.toolTip.emptyInstanceName')));
    }
    else if (instanceData.repeatInstanceData.includes(value)) {
        instanceNameState.value = true
        callback(new Error(t('personCentre.content.instanceNameExist')));
    }
    else if(!instanceNameReg.test(value)) {
        instanceNameState.value = true
        callback(new Error(t('personCentre.content.instanceVerification')));
    }else {
        instanceNameState.value = true
        callback();
    }
};
/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();
/**
 * 实例规则
*/
const rules: RulesType = reactive<FormRules>({
    instanceName: [ // 
        {
            required: false,
            trigger: 'blur',
            validator: instanceNameCheck
        }
    ],
});

/**
 * 编辑实例时，实例名称赋值
 */
onMounted(() => {
    instanceForm.instanceName = Object.keys(props.instanceInfo).length ? props.instanceInfo.instanceName : '';
})
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
const closeDialog = (formEl: FormInstance | undefined)=> {
    formEl.resetFields();
    instanceNameState.value = false; 
    isLoading.value = false;
    emitClose('close', false)
}

/**
 * 请求编辑接口
*/
const requestEdit = (): void => {
    isLoading.value = true;   
    editInstanceAPI({
        instanceId: props.instanceInfo.instanceId,
        idcId: props.instanceInfo.idcId,
        name: instanceForm.instanceName,
    }).then(({data}: any) => {
        if (data?.result?.success) {
            instanceNameState.value = false;
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            emitClose('refresh');
            emitClose('close', false)
        }
    }).catch(() => {
        emitClose('refresh');
        emitClose('close', false)
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 获取所有实例名称
 */
 const getinstanceNameData = (): void => {
    instanceListAPI({
        isAll:'1',
        projectId: props.instanceInfo.projectId,
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            instanceData.repeatInstanceData = data?.result?.instances.map((item: any) => item.instanceName);
        }
 
    }).catch(() => {
        instanceData.repeatInstanceData = [];
    })
};
getinstanceNameData()

document.onkeydown = (event: {keyCode: number;}) => {
    event.keyCode === 13 && openDialog(ruleFormRef.value);
};

/**
 * 点击确定
 */
const openDialog = async (formEl: FormInstance | undefined): Promise<void> => {
    
    // 判断输入项是否符合条件
    await nextTick(() => {
        formEl.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                requestEdit();   
            }
        });
    });
};
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';
</style>