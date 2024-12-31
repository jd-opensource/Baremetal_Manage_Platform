<template>
    <div class="page-position" :class="personalCentre === 'sshKey' || personalCentre === 'apiKey' ? 'pr7': 'can-scroll'">        
        <div class="page-content item-position">
            <!-- 侧边栏 -->
            <el-tabs 
                tab-position= "left" 
                v-model="personalCentre" 
                @tab-click="tabChange(ruleFormRef,passwordRuleFormRef)"
            >
                <!-- 我的账户 -->
                <el-tab-pane name="account">
                    <template #label>
                        <span class="custom-tabs-label">
                            <span :class="personalCentre === 'account' ? 'user' : 'userGrey'">    
                                <span class="personal-title">{{$t('personCentre.account')}}</span>
                            </span>
                        </span>
                    </template>
                    <div class="pane-content">                          
                        <el-form
                            label-position="left"
                            :model="userForm.userData"
                            ref="ruleFormRef"
                            :rules="rules"
                            label-width="80px"
                            @submit.native.prevent
                        >
                            <div class="ml55 pt45">
                                <!-- 基本信息 -->
                                <h3 class="section-title">{{ $t('personCentre.form.basicInformation') }}</h3>
                                <!-- 用户名 -->
                                <el-form-item 
                                    prop="userName"
                                    class="input-change"
                                    :label="$t('personCentre.form.user') ">
                                    <el-input v-model="userForm.userData.userName" disabled 
                                />
                                </el-form-item>
                                <!-- 用户ID -->
                                <el-form-item 
                                    prop="userId"
                                    class="input-change"
                                    :label="$t('personCentre.form.userId') "
                                >
                                    <el-input v-model="userForm.userData.userId" disabled />
                                </el-form-item>
                                <!-- 邮箱 -->
                                <el-form-item 
                                    prop="email"
                                    :label="$t('personCentre.form.email') "
                                >
                                    <el-input v-model.trim="userForm.userData.email" />
                                </el-form-item>
                                <!-- 手机号 -->
                                <el-form-item 
                                    prop="phoneNumber"
                                    :class="getLocationItem === 'zh_CN' ? '' : 'lh18'"
                                    :label="$t('personCentre.form.phoneNumber') "
                                >
                                    <el-input v-model="userForm.userData.phoneNumber" class="phone-select">
                                        <template #prepend>
                                            <el-select 
                                                v-model="userForm.userData.phonePrefix" 
                                                @change="areaCodeChange(ruleFormRef)"
                                            >
                                                <el-option 
                                                    v-for="item in phoneData"
                                                    :key="item.value"
                                                    :label="item.label"
                                                    :value="item.value"
                                                />
                                            </el-select>
                                        </template>
                                    </el-input>
                                </el-form-item>
                            </div>  
                            <el-divider />
                            <div class="ml55">
                                <!-- 偏好设置 -->
                                <h3 class="section-title">{{ $t('personCentre.form.preferences') }}</h3>
                                <!-- 语言 -->
                                <el-form-item 
                                    :label="$t('personCentre.form.language')"
                                    class="language-width"
                                >
                                    <el-select 
                                        v-model="userForm.userData.language" 
                                        :placeholder="$t('personCentre.form.selectLanguage')" 
                                    >
                                        <el-option 
                                            v-for="item in languageList" :key="item.value"
                                            :label="item.label" 
                                            :value="item.value"
                                        />
                                    </el-select>
                                </el-form-item>
                                <!-- 时区 -->
                                <el-form-item :label="$t('personCentre.form.timeZone')">
                                    <el-select 
                                        v-model="userForm.userData.timezone" 
                                        :placeholder="$t('personCentre.form.selectTimeZone')"
                                    >
                                        <el-option 
                                            v-for="item in timeZoneList" 
                                            :key="item.value"
                                            :label="item.label" 
                                            :value="item.value"
                                        />
                                    </el-select>
                                </el-form-item>
                                <!-- 默认项目 -->
                                <el-form-item 
                                    :label="$t('personCentre.form.defaultProject') "
                                    :class="getLocationItem === 'zh_CN' ? '' : 'lh18'"
                                >
                                    <el-select 
                                        v-model="userForm.userData.defaultProjectId" 
                                        :placeholder="$t('personCentre.form.selectProject') "
                                    >
                                        <el-option 
                                            v-for="item in defaultProjectList" 
                                            :key="item.value"
                                            :label="item.label" 
                                            :value="item.value"
                                        />
                                    </el-select>
                                </el-form-item>
                            </div>
                            <el-divider />
                            <!-- 保存按钮 -->
                            <div class="ml55">
                                <el-button type="primary" 
                                    class="button-style"
                                    :loading="isUserLoading"
                                    @click="editUser(ruleFormRef)"
                                >
                                    {{isUserLoading?$t('personCentre.form.saving'):$t('personCentre.form.save') }}
                                </el-button>
                            </div>    
                        </el-form>
                    </div>
                </el-tab-pane>
                <!-- 安全设置 -->
                <el-tab-pane name="securitySettings">
                    <template #label>
                        <span class="custom-tabs-label">
                            <span 
                                :class="personalCentre==='securitySettings' ? 'security' : 'securityGrey'"
                            >
                                <span class="personal-title">{{$t('personCentre.securitySettings')}}</span>
                            </span>
                        </span>
                    </template>
                    <div class="pane-content">
                        <el-form
                            label-position="left"
                            ref="passwordRuleFormRef"
                            :model="securityForm"
                            :rules="passwordRules"
                            label-width="80px"
                        >
                            <div class="ml30 pt40">
                                <!-- 当前密码 -->
                                <el-form-item 
                                    prop="currentPassword"
                                    :class="getLocationItem === 'zh_CN' ? '' : 'lh18'"
                                    :label="$t('personCentre.form.currentPassword') "
                                >
                                    <el-input v-model="securityForm.currentPassword" type="password"/>
                                </el-form-item>
                                <!-- 新密码 -->
                                <el-form-item 
                                    prop="newPassword"
                                    :class="getLocationItem === 'zh_CN' ? '' : 'lh18'"
                                    :label="$t('personCentre.form.newPassword')"
                                >
                                    <el-input 
                                        v-model="securityForm.newPassword"
                                        maxlength="30"
                                        minlength="8"
                                        type="password" 
                                    />
                                    <p 
                                        class="password-tip lh22" 
                                        v-if="!newPasswordState"
                                    >
                                        {{$t('personCentre.content.passwordVerification')}}
                                    </p>
                                </el-form-item>
                                <!-- 确认密码 -->
                                <el-form-item 
                                    prop="confirmPassword"
                                    :class="getLocationItem === 'zh_CN' ? '' : 'lh18'"
                                    :label="$t('personCentre.form.confirmPassword') "
                                >
                                    <el-input v-model="securityForm.confirmPassword" type="password"/>
                                </el-form-item>
                            </div>
                        </el-form>
                        <el-divider />
                        <div class="ml55">
                            <!-- 更改按钮 -->
                            <el-button type="primary" 
                                class="button-style" 
                                :loading="isPasswordLoading"
                                @click="changePassword(passwordRuleFormRef)"
                            >
                                {{isPasswordLoading ? $t('personCentre.form.changing') : $t('personCentre.form.change')}}
                            </el-button>
                        </div>
                    </div>
                </el-tab-pane>
                <!-- 个人SSH密钥 -->
                <el-tab-pane name="sshKey">
                    <template #label>
                        <span class="custom-tabs-label">
                            <span 
                                :class="personalCentre === 'sshKey' ? 'sshKey' : 'sshKeyGrey'"
                            >
                                <span class="personal-title">{{$t('personCentre.sshKey')}}</span>
                            </span>
                        </span>
                    </template>
                    <div class="pane-content">
                        <div class="m21 pt40">
                            <!-- 创建密钥 -->
                            <el-tooltip
                                placement="bottom"
                                :disabled="sshTableData.tableData.length<=1"
                                :content="$t('personCentre.tip.createSshKey')"
                            >
                                <label>
                                    <el-button 
                                        type="primary" 
                                        class="add-button-style" 
                                        :disabled="sshTableData.tableData.length > 1"
                                        @click="sshKeyOperateCreate"
                                    >
                                        + &nbsp;&nbsp;{{$t('personCentre.form.createSshKey') }}
                                    </el-button>
                                </label>
                            </el-tooltip>
                            <div class="table-content">
                                <!-- 个人SSH密钥列表 -->
                                <el-table 
                                    border
                                    :data="sshTableData.tableData" 
                                    :max-height="sshTableData.tableData.length ? tableMaxHeight : 420" 
                                    v-loading="isSshLoading" 
                                    style="width: auto"
                                >
                                    <!-- 密钥名称 -->
                                    <el-table-column 
                                        prop="name" 
                                        :label="$t('personCentre.form.keyName')" 
                                        min-width="100"
                                    >
                                        <template v-slot="scope">
                                            <el-tooltip
                                                v-model="scope.row.showTooltip"
                                                :disabled="!scope.row.showTooltip"
                                                :content= scope.row.name
                                            >
                                                <p 
                                                    class="long-row"
                                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.name, 1.17, 'list')"
                                                >
                                                    {{scope.row.name || $filter.emptyFilter()}}
                                                </p>
                                            </el-tooltip>
                                        </template>
                                    </el-table-column>
                                    <!-- 创建时间 -->
                                    <el-table-column 
                                        prop="createdTime" 
                                        :label="$t('personCentre.form.createdTime')" 
                                        min-width="100"
                                    />
                                    <!-- 更新时间 -->
                                    <el-table-column 
                                        prop="updatedTime" 
                                        :label="$t('personCentre.form.updatedTime')" 
                                        min-width="100"
                                    />
                                    <!-- 操作 -->
                                    <el-table-column :label="$t('personCentre.form.operation')" >
                                        <template v-slot="scope">
                                            <!-- 编辑 -->
                                            <operate-unit
                                                :operateName="$t('personCentre.operate.edit')"
                                                :disabled="false"
                                                @handleClick="sshKeyOperateEdit(scope.row)"
                                                :instanceInfo="scope.row"
                                                operateTip=""
                                            >
                                            </operate-unit>
                                            <!-- 删除 -->
                                            <operate-unit
                                                :operateName="$t('personCentre.operate.delete')"
                                                :disabled="false"
                                                @handleClick="sshKeyOperateDelete(scope.row)"
                                                :instanceInfo="scope.row"
                                                operateTip=""
                                            >
                                            </operate-unit>
                                        </template>
                                    </el-table-column>
                                    <!-- 无数据展示 -->
                                    <template #empty>
                                        <div class="noData">                        
                                        </div>
                                        <p>
                                            <span class="c333">{{$t('personCentre.noData')}}</span>
                                            <el-button 
                                                type="text" 
                                                class="f12 pt7" 
                                                @click="sshKeyOperateCreate"
                                            >
                                                {{$t('personCentre.form.create') }}
                                            </el-button>
                                        </p>
                                    </template>
                                </el-table>
                            </div>
                        </div>
                        <div v-if="deleteSshKey">
                            <!-- 删除 -->
                            <ssh-key-delete
                                :openVisible="deleteSshKey"
                                :instanceInfo="instanceInfo"
                                @close="deleteSshKeyCancel"
                                @refresh="refreshSshData"
                            >
                            </ssh-key-delete>
                        </div>
                        <div v-if="editSshKey">
                            <!-- 编辑 -->
                            <ssh-key-edit
                                :openVisible="editSshKey"
                                :instanceInfo="instanceInfo"
                                :repeatData="sshTableData.repeatData"
                                :operateState="operateState"
                                @close="editSshKeyCancel"
                                @refresh="refreshSshData"
                            >
                            </ssh-key-edit>
                        </div>
                        <div class="footer-count">
                            <!-- 翻页 -->
                            <my-pagination
                                v-if="sshTableData.tableData.length > 0"
                                :hasUseDefault="false"
                                :page-sizes="[20, 50, 100]"
                                :total="sshTotalCount" 
                                :page-size="sshPageSize" 
                                :page-number="sshPageNumber" 
                                @change-page="sshChangePage" 
                                @page-sizes-change="sshPageSizesChange"       
                            >
                            </my-pagination>
                        </div>
                    </div>                    
                </el-tab-pane>
                <!-- 个人API密钥 -->
                <el-tab-pane name="apiKey">
                    <template #label>
                        <span class="custom-tabs-label">
                            <span :class="personalCentre === 'apiKey' ? 'apiKey' : 'apiKeyGrey'">
                                <span class="personal-title">{{$t('personCentre.apiKey')}}</span>
                            </span>
                        </span>
                    </template>
                    <div class="pane-content">
                        <div class="m21 pt40">
                            <!-- 创建密钥 -->
                            <el-tooltip
                                placement="bottom"
                                :disabled="apiTableData.tableData.length <= 1"
                                :content="$t('personCentre.tip.createApiKey')"
                            >
                                <label>
                                    <el-button 
                                        type="primary" 
                                        class="add-button-style"
                                        :disabled="apiTableData.tableData.length > 1"
                                        @click="apiKeyOperateCreate"
                                    >
                                        + &nbsp;&nbsp;{{$t('personCentre.form.createApiKey') }}
                                    </el-button>
                                </label>   
                            </el-tooltip>
                            <div class="table-content">
                                <!-- 个人API密钥列表 -->
                                <el-table 
                                    border
                                    :data="apiTableData.tableData" 
                                    :max-height="apiTableData.tableData.length ? tableMaxHeight : 420" 
                                    v-loading="isApiLoading"
                                    style="width: auto"
                                >   
                                    <!-- 密钥名称 -->
                                    <el-table-column 
                                        prop="name" 
                                        :label="$t('personCentre.form.keyName')" 
                                        min-width="85px"
                                    >
                                        <template v-slot="scope">
                                            <el-tooltip
                                                v-model="scope.row.showTooltip"
                                                :disabled="!scope.row.showTooltip" 
                                                :content= scope.row.name
                                            >
                                                <p 
                                                    class="long-row"
                                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.name, 1.17, 'list')"
                                                >
                                                    {{scope.row.name || $filter.emptyFilter()}}
                                                </p>
                                            </el-tooltip>
                                        </template>
                                    </el-table-column>
                                    <!-- Token -->
                                    <el-table-column 
                                        prop="token" 
                                        label="Token" 
                                        min-width="105px"
                                    >
                                        <template v-slot="scope">
                                            <el-tooltip
                                                v-model="scope.row.showTooltip"
                                                :disabled="!scope.row.showTooltip" 
                                                :content= scope.row.token
                                                placement="top-start"
                                            >
                                                <p 
                                                    class="long-row"
                                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.token, 1.17, 'list')"
                                                >
                                                {{scope.row.token}}</p>
                                            </el-tooltip>
                                        </template>
                                    </el-table-column>
                                    <!-- 权限 -->
                                    <el-table-column 
                                        prop="readOnly" 
                                        :label="$t('personCentre.form.power')" 
                                        width="110px"
                                    >
                                        <template v-slot="scope">
                                            <p>{{apiKeyTypeChange(scope.row.readOnly)}}</p>
                                        </template>
                                    </el-table-column>
                                    <!-- 创建时间 -->
                                    <el-table-column 
                                        prop="createdTime" 
                                        :label="$t('personCentre.form.createdTime')" 
                                        min-width="135px"
                                    />
                                    <!-- 操作 -->
                                    <el-table-column 
                                        prop="operation" 
                                        min-width="110px"
                                        :label="$t('personCentre.form.operation')" 
                                    >
                                        <template v-slot="scope">
                                            <!-- 复制 -->
                                            <operate-unit
                                                :operateName="$t('personCentre.operate.copy')"
                                                :disabled="false"
                                                @handleClick="clickOperateCopy(scope.row)"
                                                :instanceInfo="scope.row"
                                                operateTip=""
                                            >
                                            </operate-unit>
                                            <!-- 删除 -->
                                            <operate-unit
                                                :operateName="$t('personCentre.operate.delete')"
                                                :disabled="false"
                                                @handleClick="apiKeyOperateDelete(scope.row)"                                        
                                                :instanceInfo="scope.row"
                                                operateTip=""
                                            >
                                            </operate-unit>
                                        </template>
                                    </el-table-column>
                                    <template #empty>
                                        <!-- 无数就展示 -->
                                        <div class="noData">                        
                                        </div>
                                        <p>
                                            <span class="c333">{{$t('personCentre.noData')}}</span>
                                            <el-button 
                                                type="text" 
                                                class="f12 pt7" 
                                                @click="apiKeyOperateCreate"
                                            >
                                                {{$t('personCentre.form.create') }}
                                            </el-button>
                                        </p>
                                    </template>
                                </el-table>
                            </div>
                        </div>
                        <div v-if="deleteApiKey">
                            <!-- 删除 -->
                            <api-key-delete
                                :openVisible="deleteApiKey"
                                :instanceInfo="instanceInfo"
                                @close="deleteApiKeyCancel"
                                @refresh="refreshApiData"
                            >
                            </api-key-delete>
                        </div>
                        <!-- 创建 -->
                        <div v-if="createApiKey">
                            <api-key-create
                                :openVisible="createApiKey"
                                :repeatData="apiTableData.repeatData"
                                @close="createApiKeyCancel"
                                @refresh="refreshApiData"
                            >
                            </api-key-create>
                        </div>
                        <div class="footer-count">
                            <!-- 翻页 -->
                            <my-pagination
                                v-if="apiTableData.tableData.length > 0"
                                :hasUseDefault="false"
                                :page-sizes="[20, 50, 100]"
                                :total="apiTotalCount" 
                                :page-size="apiPageSize" 
                                :page-number="apiPageNumber" 
                                @change-page="apiChangePage"  
                                @page-sizes-change="apiPageSizesChange"     
                            >
                            </my-pagination>
                        </div>
                    </div>   
                </el-tab-pane>
                
            </el-tabs>
        </div>        
    </div>
</template>

<script setup lang="ts">
import { 
    reactive,
    ref,
    Ref,
    onMounted,
    onUnmounted,
    nextTick,
    watch,
    WritableComputedRef
} from 'vue';
import { useRoute, useRouter } from 'vue-router';
import useCurrentInstance from "hooks/useCurrentInstance.ts";
import useClipboard from 'vue-clipboard3'
import {encrypt} from 'utils/index.ts';
import operateUnit from 'components/OperateUnit/operateUnit.vue';
import { ElMessage } from 'element-plus';
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {
    emailReg, // 邮箱地址
    phoneReg, // 中国大陆手机号正则
    aOmenPhoneReg, // 中国澳门手机号正则
    taiWanPhoneReg, // 中国台湾手机号正则
    hongKongPhoneReg, // 中国香港手机号正则
    otherPhoneReg, //其他国家手机号正则
    passwordReg // 密码正则
} from 'utils/regular.ts';
import {
    hasShowTooltip // 是否显示提示
} from 'utils/index.ts';
import {useI18n} from 'vue-i18n'; // 国际化
import {cellPhoneType, languageType, apiKeyType} from 'utils/constants.ts'; // 类型数据
import apiKeyDelete from 'components/PersonalCentreOperate/apiKeyDelete.vue';
import apiKeyCreate from 'components/PersonalCentreOperate/apiKeyCreate.vue';
import sshKeyDelete from 'components/PersonalCentreOperate/sshKeyDelete.vue';
import sshKeyEdit from 'components/PersonalCentreOperate/sshKeyEdit.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import VueCookies from 'vue-cookies'; // cookie
import {
    setUserAPI,
    projectListAPI,
    timeZoneAPI,
    userAPI,
    changePasswordAPI,
    sshKeyAPI,
    apiKeyAPI
    
} from 'api/request.ts'; // 编辑用户，手机区号，语言，时区,项目,用户接口,更改密码接口,ssh,api密钥数据接口
import allProjectStore from '@/store/modules/headerDetail.ts';
/**
 * 路由带值
 */
const route = useRoute();

// js中使用国际化
const {locale} : {locale: WritableComputedRef<string>} = useI18n();

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: string | null = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 改变的语言
*/
const changeLanguage: Ref<string> = (ref(getLocationItem) as any);

/**
 * 改变语言类型
 * @param {string} lang 语言类型
*/
const changeLang = (lang: string): void => {
    locale.value = lang;
    (VueCookies as any).set('X-Jdcloud-Language', locale.value);
    changeLanguage.value = (VueCookies as any).get('X-Jdcloud-Language');
    location.reload()
};

/**
 * 监控路由，选择不同项目获取数据
 */
watch(() => route, (): Readonly<void> => { 
    personalCentre.value = route.query.type
},{deep:true});

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 用户ID
*/
let userId: string = (VueCookies as any).get('user_id');

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 用户ID解码
 */
userId = window.atob(userId);

/**
 * 用户item类型
 */
type userType ={
    userName: string, // 用户名
    userId: string, // 用户ID
    email: string, // 邮箱
    phonePrefix: string, // 手机区号
    phoneNumber: number, // 手机号
    language: string, // 语言
    timezone: string, // 时区
    defaultProjectId: string, // 默认项目
    password: string, // 密码
};

/**
 * 用户类型
 */
interface user {
    userData: userType
}

/**
 * 用户表单
 */
const userForm: user = reactive<user>({
    userData: {
        userName: '', // 用户名
        userId:'', // 用户ID
        email: '', // 邮箱
        phonePrefix: '', // 手机区号
        phoneNumber: 0, // 手机号
        language: '', // 语言
        timezone: '', // 时区
        defaultProjectId: '', // 默认项目
        password:'', // 密码
    }
})

// 计算表格的最大高度
const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 290)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
  // 触发响应式更新
  tableMaxHeight.value = window.innerHeight - 290
};

onMounted(() => {
  window.addEventListener('resize', updateSize);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateSize);
});

/**
 * 用户表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 密码表单ref
*/
let passwordRuleFormRef: Ref<any> = ref<FormInstance>();

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
 * 邮箱正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const emailCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyEmail')));
    }
    else if (!emailReg.test(value)) {
        callback(new Error(t('personCentre.toolTip.errorEmail')));
    }
    else {
        callback();
    }
};

/**
 * 判断手机号类型
*/
const useWhichPhoneReg: Ref<string> = ref<string>('086');  

/**
 * 手机号类型类
*/
interface PhoneDataType {
    value: string;
    label: string;
};

/**
 * 手机号类型数据
*/
const phoneData: PhoneDataType[] = reactive<PhoneDataType[]>(cellPhoneType);

/**
 * 手机号正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const phoneCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    useWhichPhoneReg.value = userForm.userData.phonePrefix
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyPhone')))
    }
    else if (!phoneReg.test(value) && useWhichPhoneReg.value === '086') {
        callback(new Error(t('personCentre.toolTip.errorPhone')));
    }
    else if (!hongKongPhoneReg.test(value) && useWhichPhoneReg.value === '852') {
        callback(new Error(t('personCentre.toolTip.errorPhoneHongkong')));
    }
    else if (!aOmenPhoneReg.test(value) && useWhichPhoneReg.value === '853') {
        callback(new Error(t('personCentre.toolTip.errorPhoneMacao')));
    }
    else if (!taiWanPhoneReg.test(value) && useWhichPhoneReg.value === '886') {
        callback(new Error(t('personCentre.toolTip.errorPhoneTaiwan')));
    }
    else if (!otherPhoneReg.test(value) && (useWhichPhoneReg.value !== '086' && '852' && '853' && '886')) {
        callback(new Error(t('personCentre.toolTip.errorPhone')));
    }
    else {
        callback();
    }
};

/**
 * 当前密码正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const currentPasswordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyCurrentPassword')));
    }
    else {
        callback();
    }
};

/**
 * 新密码正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const newPasswordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    newPasswordState.value = true
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyNewPassword')));
    }
    else if (!passwordReg.test(value)) {
        callback(new Error(t('personCentre.content.passwordVerification')));
    }
    else {
        callback();
    }
};

/**
 * 确认密码正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const confirmPasswordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyconfirmPassword')));
    }
    else if (value !== securityForm.newPassword) {
        callback(new Error(t('personCentre.toolTip.passNotSame')));
    }
    else {
        callback();
    }
};

/**
 * 账户规则类
*/
interface RulesType {
    userName: UserInfoType[];
    email: UserInfoType[];
    phonePrefix: UserInfoType[];
    phoneNumber: UserInfoType[];
    language: UserInfoType[];
    timezone: UserInfoType[];
    defaultProject: UserInfoType[];
};

/**
 * 账户规则
*/
const rules: RulesType = reactive<FormRules>({
    email: [ // 邮箱
        {
            required: false,
            trigger: 'blur',
            validator: emailCheck
        }
    ],
    phonePrefix: [ // 手机号类型
        {
            required: false,
            trigger: 'change'
        }
    ],
    phoneNumber: [ // 手机号
        {
            required: false,
            trigger: 'blur',
            validator: phoneCheck
        }
    ],
});

/**
 * 安全设置类型
 */
interface security {
    currentPassword: string,
    newPassword: string,
    confirmPassword: string,
}

/**
 * 安全设置表单
 */
const securityForm:security = reactive<security>({
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
})

/**
 * 密码规则类
*/
interface passwordRulesType {
    currentPassword: UserInfoType[];
    newPassword: UserInfoType[];
    confirmPassword: UserInfoType[];
};

/**
 * 账户规则
*/
const passwordRules: passwordRulesType = reactive<FormRules>({
    currentPassword: [ // 当前密码
        {
            required: false,
            trigger: 'blur',
            validator: currentPasswordCheck
        }
    ],
    newPassword: [ // 新密码
        {
            required: false,
            trigger: 'blur',
            validator: newPasswordCheck
        }
    ],
    confirmPassword: [ // 确认密码
        {
            required: false,
            trigger: 'blur',
            validator: confirmPasswordCheck
        }
    ],
});

/**
 * 语言菜单类型类
*/
interface LanguageType {
    value: number;
    label: string;
};

/**
 * 语言菜单
 */
let languageList: LanguageType[] = reactive<LanguageType[]>(languageType);

/**
 * 时区菜单
 */
let timeZoneList: ProjectType[] = reactive<ProjectType[]>([]);

/**
 * 项目菜单类型类
*/
interface ProjectType {
    value:string;
    label:string;
};

/**
 * 默认项目菜单
 */
let defaultProjectList: ProjectType[] = reactive<ProjectType[]>([]);

/**
 * 新密码状态
 */
const newPasswordState: Ref<boolean> = ref<boolean>(false)

/**
 * 加载态-用于登录请求时，账户按钮的加载态
*/
const isUserLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 加载态-用于登录请求时，密码按钮的加载态
*/
const isPasswordLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 个人中心跳转标志
 */
let personalCentre: Ref<any> = ref('');

/**
 * 个人ssh密钥item类型
 */
 type SshType = {
    createdTime: string,
    name: string,
    operation: string,
}

/**
 * 个人ssh密钥类型
 */
 interface ssh {
    tableData: SshType[],
    repeatData: [],
}

/**
 * 个人ssh密钥数据
 */
const sshTableData: ssh = reactive<ssh>({    
    tableData: [],
    repeatData: [],
})

/**
 * 判断是ssh密钥操作编辑还是创建
 */
 const operateState: Ref<string> = ref<string>('')

/**
 * ssh当前页面页数条数
*/
const sshPageSize: Ref<number> = ref<number>(20);

/**
 * ssh当前页面页数
*/
const sshPageNumber: Ref<number> = ref<number>(1);

/**
 * ssh总条数
*/
const sshTotalCount: Ref<number> = ref<number>(0);

/**
 * 个人api密钥item类型
 */
 type ApiType = {
    createdTime: string,
    name: string,
    readOnly: number,
    token: string,
    operation: string,
}

/**
 * 个人api密钥类型
 */
 interface api {
    tableData: ApiType[],
    repeatData: [],
}

/**
 * 个人api密钥数据
 */
const apiTableData: api = reactive<api>({
    tableData: [],
    repeatData: [],
})

/**
 * api当前页面页数条数
*/
const apiPageSize: Ref<number> = ref<number>(20);

/**
 * api当前页面页数
*/
const apiPageNumber: Ref<number> = ref<number>(1);

/**
 * api总条数
*/
const apiTotalCount: Ref<number> = ref<number>(0);

/**
 * 个人中心点击跳转到指定页面
 */
onMounted(() => {
    if(route.query.type) {
        personalCentre.value = route.query.type
    } else{
        personalCentre.value = 'account'
    }
});

/**
 * 手机区号改变
 * 
*/
const areaCodeChange =(formEl: FormInstance | undefined)=> {
    formEl.clearValidate(['phone']);  
    useWhichPhoneReg.value = userForm.userData.phonePrefix
}

/**
 * tab页改变
 * @param formEl 
 */
const router = useRouter();
const tabChange = (formEl: FormInstance | undefined, passwordFormEl: FormInstance | undefined)=>{
    router.push(`/user?type=${personalCentre.value}`);
    formEl.resetFields();
    passwordFormEl.resetFields();
    if(personalCentre.value === 'account') {
        requestUser()
        route.query.type = 'account'
    }else if(personalCentre.value === 'securitySettings') {
        route.query.type = 'securitySettings'
    }else if(personalCentre.value === 'sshKey') {
        requestSshKey()
        requestUser()
        route.query.type = 'sshKey'
    }else if(personalCentre.value === 'apiKey') {
        requestApiKey()
        requestUser()
        route.query.type = 'apiKey'
    }
    newPasswordState.value = false;
}

watch(() => personalCentre.value, (): Readonly<void> => { 
    if(personalCentre.value === 'account') {
       // requestUser()
    }
   
});

/**
 * 当前this
 */
const { globalProperties } = useCurrentInstance();

/**
 * 当前行值
 */
 const instanceInfo: Ref<any> = ref();

/**
 * SSH密钥编辑操作打开标志
 */
const editSshKey: Ref<boolean> = ref<boolean>(false)

/**
 * SSH密钥删除操作打开标志
 */
const deleteSshKey: Ref<boolean> = ref<boolean>(false)

/**
 * api密钥创建操作打开标志
 */
const createApiKey: Ref<boolean> = ref<boolean>(false)

/**
 * api密钥删除操作打开标志
 */
const deleteApiKey: Ref<boolean> = ref<boolean>(false)

/**
 * sshKey创建操作
 * @param value 
 */
const sshKeyOperateCreate: () => void = () => {
    instanceInfo.value = {};
    operateState.value ='create';
    editSshKey.value = !editSshKey.value;
} 

/**
 * sshKey编辑操作
 * @param value 
 */
const sshKeyOperateEdit: (value: object) => void = (value: object) => {
    instanceInfo.value = value;
    operateState.value = 'edit';
    editSshKey.value = !editSshKey.value;
}

/**
 * sshKey编辑弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} editSshKey.value 弹窗关闭
 */
 const editSshKeyCancel = (type: boolean): boolean => {
    return editSshKey.value = type;
}; 

/**
 * sshKey删除操作
 * @param value 
 */
const sshKeyOperateDelete: (value: object) => void = (value: object) => {
    instanceInfo.value = value;   
    deleteSshKey.value = true;
}

/**
 * sshKey删除弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} deleteSshKey.value 弹窗关闭
 */
 const deleteSshKeyCancel = (type: boolean): boolean => {
    return deleteSshKey.value = type;
}; 

/**
 * ssh每页展示条数切换
 * @param {number} size 每页展示条数
*/
const sshPageSizesChange = (size: number) => {
    sshPageNumber.value = 1;
    sshPageSize.value = size;
    requestSshKey();
};

/**
 * ssh页数切换
 * @param {number} num 页数切换
*/
const sshChangePage = (num: number): void => {
    sshPageNumber.value = num;
    requestSshKey();
};

/**
 * api每页展示条数切换
 * @param {number} size 每页展示条数
*/
const apiPageSizesChange = (size: number) => {
    apiPageNumber.value = 1;
    apiPageSize.value = size;
    requestApiKey();
};

/**
 * api页数切换
 * @param {number} num 页数切换
*/
const apiChangePage = (num: number): void => {
    apiPageNumber.value = num;
    requestApiKey();
};

/**
 *  apiKey复制操作
 * @param value 
 */
const {toClipboard} = useClipboard();
const clickOperateCopy: (value: any) => void = (value: any) => {
    toClipboard(value.token);  
    ElMessage({
        message: globalProperties.$t('personCentre.content.copySuccess'),
        type: 'success'
    })
}

/**
 * apiKey创建操作
 * @param value 
 */
const apiKeyOperateCreate: () => void = () => {
    createApiKey.value = !createApiKey.value;
}

/**
 * apikey创建弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} createApiKey.value 弹窗关闭
 */
 const createApiKeyCancel = (type: boolean): boolean => {
    return createApiKey.value = type;
}; 

/**
 * apiKey删除操作
 * @param value 
 */
const apiKeyOperateDelete: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    deleteApiKey.value = !deleteApiKey.value;
}

/**
 * apikey删除弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} deleteApiKey.value 弹窗关闭
 */
 const deleteApiKeyCancel = (type: boolean): boolean => {
    return deleteApiKey.value = type;
}; 

/**
 * apikey权限展示转换
 * @param value 
 */
const apiKeyTypeChange: (value: number) => void = (value: number) => {
    let power = ''
    if(!value) {
        value = 0
    }
    power = apiKeyType.find((item: any) => item.value === value).label
    return power
}

/**
 * 请求项目列表数据接口
*/
const requestProjectListData = (): void => {
    projectListAPI({
        isAll:'1'
    }).then(({data} : {data: any}) => {
        if (data?.result?.projects?.length) {
            data.result.projects.map((item: any) => {
                let obj = {
                    value: item.projectId,
                    label: item.projectName,
                }
                defaultProjectList.push(obj)
            });
        return;
    }
    defaultProjectList  = [];
    }).catch(({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
        defaultProjectList  = [];
    })
    .finally(() => {
    });
};
requestProjectListData()

/**
 * 请求用户接口
*/
const requestUser = (): void => {
    userAPI({
    }).then(({data} : {data: any}) => {
        if(data?.result){
            userForm.userData = data.result; 
        }                  
    }).catch (({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
    })
    .finally(() => {
        store.requestObject();
        store.requestUser();
    });
};
requestUser()

/**
 * 请求时区接口
*/
const requestTimezone = (): void => {
    timeZoneAPI({
    }).then(({data} : {data: any}) => {
        if(data?.result?.timezones){
            Object.keys(data.result.timezones).forEach((item: string) => {
                let obj = {
                    value: item,
                    label: data.result.timezones[item],
                }
                timeZoneList.push(obj)
            })
        }                  
    })
    .catch(()=>{
        timeZoneList = [] 
    })
    .finally(() => {
    });
};
requestTimezone()

/**
 * 请求保存账户接口
*/
const requestSetUser = (): void => {
    isUserLoading.value = true;
    setUserAPI({
        userName: userForm.userData.userName,
        email: userForm.userData.email,
        language: userForm.userData.language,
        phoneNumber: userForm.userData.phoneNumber,
        phonePrefix: userForm.userData.phonePrefix,
        timezone: userForm.userData.timezone,
        defaultProjectId: userForm.userData.defaultProjectId,
    }).then((res: any) => {
        if (res) {
            //requestUser();
            ElMessage({
                message: globalProperties.$t('personCentre.content.operateSuccess'),
                type: 'success'
            })
        }

    }).finally(() => {
        isUserLoading.value = false;
        changeLang(userForm.userData.language)
    });
};

/**
 * 保存账户信息
 */
const editUser = async (formEl: FormInstance | undefined): Promise<void> => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        formEl.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                requestSetUser();               
            }
        });
    });
};

/**
 * 请求更改密码接口
*/
const requestChangePassword = (): void => {
    isPasswordLoading.value = true;
    changePasswordAPI({
        oldPassword: encrypt(securityForm.currentPassword, 'ABCDEFGHIJKLMNOP', '0123456789ABCDEF'),
        password: encrypt(securityForm.confirmPassword, 'ABCDEFGHIJKLMNOP', '0123456789ABCDEF')
    }).then((res: any) => {
        if (res) {
            newPasswordState.value = false;
            ElMessage({
                message: globalProperties.$t('personCentre.content.fixSuccess'),
                type: 'success'
            })
            securityForm.currentPassword = '';
            securityForm.newPassword = '';
            securityForm.confirmPassword = '';
        }

    }).finally(() => {        
        isPasswordLoading.value = false;
    });
};

/**
 * 更改密码
 */
const changePassword = async (formEl: FormInstance | undefined): Promise<void> => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        formEl.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                requestChangePassword();
            }
        });
    });
};

/**
 * ssh loadig态
*/
const isSshLoading: Ref<boolean> = ref<boolean>(true);

/**
 * ssh密钥数据接口
*/
const requestSshKey = (): void => {
    isSshLoading.value = true;
    sshKeyAPI({
        pageNumber: sshPageNumber.value,
        pageSize: sshPageSize.value

    }).then(({data} : {data: any}) => {
        if (data?.result?.keypairs?.length) {
            sshTableData.tableData = data.result.keypairs.map((item: any) => {
                return {
                    ...item,
                    showTooltip: false
                }
            });
            sshTotalCount.value = data.result.totalCount;
        return;
    }
    sshTableData.tableData  = [];
    sshTotalCount.value = 0;
    })
    .catch(() => {
        sshTableData.tableData  = [];
        sshTotalCount.value = 0;
    })
    .finally(() => {
        isSshLoading.value = false;
    });
};

/**
 * ssh所有密钥数据接口
*/
const requestAllSshKey = (): void => {
    sshKeyAPI({
        isAll:1
    }).then(({data} : {data: any}) => {
        if (data?.result?.keypairs?.length) {
            sshTableData.repeatData = data.result.keypairs.map((item: any) => item.name);
        } else {
            sshTableData.repeatData = [];
        }
    }).catch(() => {
        sshTableData.repeatData = [];
    })
};
requestSshKey();
requestAllSshKey();

/**
 * 刷新ssh密钥
 */
const refreshSshData = () => {
    requestSshKey();
    requestAllSshKey();
}

/**
 * api loadig态
*/
const isApiLoading: Ref<boolean> = ref<boolean>(true);

/**
 * api密钥数据接口
*/
const requestApiKey = (): void => {
    isApiLoading.value = true;
    apiKeyAPI({
        pageNumber:apiPageNumber.value,
        pageSize: apiPageSize.value

    }).then(({data} : {data: any}) => {
        if (data?.result?.apikeys?.length) {
            apiTableData.tableData = data.result.apikeys.map((item: any) => {
                return {
                    ...item,
                    showTooltip: false
                }
            });
            apiTotalCount.value = data.result.totalCount;
        return;
    }
    apiTableData.tableData  = [];
    apiTotalCount.value = 0;
    }).catch(() => {
        apiTableData.tableData  = [];
        apiTotalCount.value = 0;
    })
    .finally(() => {
        isApiLoading.value = false;
    });
};

/**
 * api所有密钥数据接口
*/
const requestAllApiKey = (): void => {
    isSshLoading.value = true;
    apiKeyAPI({
        isAll:1
    }).then(({data} : {data: any}) => {
        if (data?.result?.apikeys?.length) {
            apiTableData.repeatData = data.result.apikeys.map((item: any) => item.name);
        } else {
            apiTableData.repeatData = [];
        }
    }).catch(() => {
        apiTableData.repeatData = [];
    })
};
requestApiKey();
requestAllApiKey();

/**
 * 刷新api密钥
 */
 const refreshApiData = () => {
    requestApiKey();
    requestAllApiKey();
}


</script>

<style lang="scss">
@import 'assets/css/table';
@import 'assets/css/page';
@import 'assets/css/addButton';
@import 'assets/css/icon';
@import './personalCentre.scss';
@import 'assets/css/tagColor';
</style>