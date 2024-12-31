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
                    <div class="apiKey dialog-icon"></div>
                    <span class="my-title">{{$t('personCentre.operate.createApiKey')}}</span>
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
                :model="apiKeyForm.apiData"
                ref="ruleFormRef"
                :rules="rules"
                :hide-required-asterisk="true"
                label-width="auto"
                @submit.native.prevent
            >
                    <el-form-item 
                        prop="name"
                        :label="$t('personCentre.form.keyName') "
                    >
                        <el-input 
                            v-model="apiKeyForm.apiData.name" 
                            maxlength="64"
                            :placeholder="$t('personCentre.placeholder.apiKey')"
                            @submit.native.prevent
                        >
                        </el-input>
                        <p class="f12 c999 w350 left-text" v-if="!keyNameState">{{$t('personCentre.content.stringVerification')}}</p>
                    </el-form-item>
                    <el-form-item
                        prop="power"
                        :label="$t('personCentre.form.power') "
                    >
                        <el-select 
                            v-model="apiKeyForm.apiData.power" 
                            :placeholder="$t('personCentre.placeholder.select')"
                        >
                            <el-option v-for="item in apiKeyData"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
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
import {Ref, ref, nextTick, reactive, watch} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    userNameReg // 项目名称正则
} from 'utils/regular.ts';
import {useI18n} from 'vue-i18n'; // 国际化
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {ElMessage} from 'element-plus';
import {apiKeyCreateType} from 'utils/constants.ts'; // api权限类型数据
import {apiKeyCreateAPI} from 'api/request.ts'; // api密钥创建接口

/**
 * api item类型
 */
 type ApiType ={
    name: string, //api密钥名
    power: string, //权限
};

/**
 * api类型
 */
interface api {
    apiData: ApiType
}

/**
 * api表单
 */
const apiKeyForm:api = reactive<api>({
    apiData: {
        name: '',
        power: '',
    }
})

/**
 * api数据类型类
*/
interface DataType {
    value: string;
    label: string;
};

/**
 * api类型数据
*/
const apiKeyData: DataType[] = reactive<DataType[]>(apiKeyCreateType);

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    repeatData: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    repeatData: []

});

/**
 * 密钥名称状态
 */
 const keyNameState: Ref<boolean> = ref<boolean>(false)

/**
 * 用户信息类
*/
type UserInfoType = {
    required: boolean;
    message?: string;
    trigger: string;
    validator?: unknown;
};

/**
 * 规则类
*/
interface RulesType {
    name: UserInfoType[];
    power: UserInfoType[]
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
    else if (props.repeatData.includes(value)) {
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
    power: [ // 权限
        {
            required: true,
            trigger: 'change',
            message: t('personCentre.toolTip.emptyKeyPower')
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
    apiKeyForm.apiData.name = '';
    formEl.resetFields();  
    emitClose('close', false)
}

/**
 * 请求创建接口
*/
const requestCreate = (): void => {
    apiKeyCreateAPI({
        name: apiKeyForm.apiData.name,
        readOnly: apiKeyForm.apiData.power,
        type: 'user'
    }).then((res: any) => {
        if(res) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            keyNameState.value = false;
            apiKeyForm.apiData.name = '';
            emitClose('close', false);
            emitClose('refresh')
        }
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
                requestCreate();
            }
        });
    });
};
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';

.left-text {
    text-align: left;
    line-height: 14px;
    margin-top: 5px;
}
</style>