<template>
    <div :class="operateState==='create' ? 'middle-dialog' : 'small-dialog'" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="sshKey dialog-icon"></div>
                    <span class="my-title">{{Object.keys(props.instanceInfo).length ? $t('personCentre.operate.editSshKey') : $t('personCentre.operate.createSshKey')}}</span>
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
                    :model="sshKeyForm.sshData"
                    ref="ruleFormRef"
                    :rules="rules"
                    label-width="auto"
                    @submit.native.prevent
                >
                    <el-form-item 
                        prop="name"
                        :label="$t('personCentre.form.keyPairName') "
                    >
                        <el-input 
                            v-model="sshKeyForm.sshData.name" 
                            type="text"
                            maxlength="64"
                            :placeholder="$t('personCentre.placeholder.keyPair')"
                        />
                        <p class="f12 c999 w350 left-text" v-if="!keyNameState">{{$t('personCentre.content.stringVerification')}}</p>
                    </el-form-item>
                    <el-form-item 
                        v-if="operateState==='create'"
                        prop="key"
                        :label="$t('personCentre.form.publicKey') "
                    >
                        <el-input 
                            type="textarea" 
                            :rows="6" 
                            :placeholder="$t('personCentre.content.publicKeyExample')" 
                            v-model="sshKeyForm.sshData.key" 
                        />
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
import {Ref, ref, reactive, watch, nextTick, onMounted} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    userNameReg // 项目名称正则
} from 'utils/regular.ts';
import {deepCopy} from 'utils/index.ts';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {sshKeyEditAPI, sshKeyCreateAPI} from 'api/request.ts'; // ssh密钥编辑,创建接口

/**
 * ssh item类型
 */
 type SshType ={
    name: string, // ssh密钥名
    key: string, // 公钥
};

/**
 * ssh类型
 */
interface ssh {
    sshData: SshType
}

/**
 * ssh表单
 */
const sshKeyForm: ssh = reactive<ssh>({
    sshData: {
        name: '',
        key: '',
    }
})

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 在编辑时，名称赋值
 */
onMounted(() => {
    sshKeyForm.sshData.name = props.operateState === 'edit' ? props.instanceInfo.name : '';
    const repeatDeepCopyData = deepCopy(props.repeatData);
    repeatDeepCopyData.map((item:any, index:number)=>{
        if(item === props.instanceInfo.name){
            repeatDeepCopyData.splice(index, 1)
        }
    })
    repeatDataNotSelf.value = repeatDeepCopyData;
})

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    operateState: string;
    repeatData: any;
    instanceInfo?: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    operateState: '',
    repeatData: [],
    instanceInfo: {}
});

// 名称重复除自己以外的数组
const repeatDataNotSelf = reactive<any>([])

/**
 * 密钥名称状态
 */
const keyNameState: Ref<boolean> = ref<boolean>(false)

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
    name: UserInfoType[];
    key: UserInfoType[];
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();
/**
 * 密钥名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const keyNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        keyNameState.value = true
        callback(new Error(t('personCentre.toolTip.emptyKeyName')));
    }
    else if (repeatDataNotSelf.value.includes(value)) {
        keyNameState.value = true
        callback(new Error(t('personCentre.content.keyNameExist')));
    }
    else if (!userNameReg.test(value)) {
        keyNameState.value = true
        callback(new Error(t('personCentre.content.stringVerification')));
    }
    else {
        keyNameState.value = true
        callback();
    }
};

/**
 * 公钥正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const publicKeyCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyPublicKey')));
    }
    else {
        callback();
    }
};

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    name: [ // 密钥名称
        {
            required: false,
            trigger: 'blur',
            validator: keyNameCheck
        }
    ],
    key: [ // 公钥
        {
            required: false,
            trigger: 'blur',
            validator: publicKeyCheck
        }
    ],
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
const closeDialog=(formEl: FormInstance | undefined)=> {    
    keyNameState.value = false;
    sshKeyForm.sshData.name = '';
    sshKeyForm.sshData.key = '';
    formEl.resetFields();  
    emitClose('close', false)
}

/**
 * 请求创建接口
*/
const requestCreate = (): void => {
    sshKeyCreateAPI({
        name: sshKeyForm.sshData.name,
        key: sshKeyForm.sshData.key
    }).then((res: any) => {
        if(res) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            sshKeyForm.sshData.name = '';
            sshKeyForm.sshData.key = '';
            emitClose('close', false);
            emitClose('refresh')
        }
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 请求编辑接口
*/
const requestEdit = (): void => {
    sshKeyEditAPI({
        name: sshKeyForm.sshData.name,
        keypairId:props.instanceInfo.keypairId
    }).then(({data}: any) => {
        if(data?.result) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            sshKeyForm.sshData.name = '';
            sshKeyForm.sshData.key = '';
            emitClose('close', false);
            emitClose('refresh')
        }
    }).catch(() => {
        sshKeyForm.sshData.name = '';
        sshKeyForm.sshData.key = '';
        emitClose('close', false);
        emitClose('refresh') 
    }).finally(() => {
        isLoading.value = false;
    });
};

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
                isLoading.value = true;
                if(props.operateState === 'create') {
                    requestCreate();
                } else if(props.operateState === 'edit') {
                    requestEdit();
                }
            }
        });
    });
};
</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/icon';
.left-text {
    text-align: left;
    line-height: 14px;
}
</style>