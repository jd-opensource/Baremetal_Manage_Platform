<template>
    <div class="middle-dialog resystem-shape" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="list dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.resystemConfi')}}</span>
                    <el-button     
                        circle class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="error-tip-content display-flex">
                <div class="error error-icon mt6"></div>
                <span class="notice-change">{{$t('list.content.notice')}} </span>
                <div class="lh24" v-html="$t('list.content.resystemTips')"></div>                
            </div>
            <div class="resystem-form">
                <el-form
                    label-position="left"
                    :model="instanceResystemForm"
                    ref="ruleFormRef"
                    :rules="rules"
                    label-width="110px"
                    @submit.native.prevent
                >
                    <!-- 实例信息 -->
                    <el-form-item 
                        :label="$t('instance.detail.instanceMessage') "
                        class="mb10 mt10"
                    >
                    </el-form-item>
                    <div class="table-content m0">
                        <el-table 
                            border
                            ref="tableRef"
                            :data="[props.instanceInfo]" 
                            style="width: 100%"
                            class="mb30"
                            max-height="155"
                        >
                            <!-- 实例名称 -->
                            <el-table-column 
                                prop="instanceName" 
                                :label="$t('instance.detail.instanceName')" 
                            >
                                <template v-slot="scope">
                                    <el-tooltip    
                                        v-model="scope.row.showTooltip"
                                        :disabled="!scope.row.showTooltip"                           
                                        :content= scope.row.instanceName
                                        placement="top-start"
                                    >
                                        <p 
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row, scope.row.instanceName, 1.17, 'list')"
                                        >
                                            {{scope.row.instanceName}}
                                        </p>
                                    </el-tooltip>
                                </template>
                            </el-table-column>
                            <!-- 实例ID -->
                            <el-table-column 
                                prop="instanceId" 
                                :label="$t('instance.detail.instanceId')" 
                            >
                                <template v-slot="scope">
                                    <el-tooltip    
                                        v-model="scope.row.showTooltip"
                                        :disabled="!scope.row.showTooltip"                           
                                        :content= scope.row.instanceId
                                        placement="top-start"
                                    >
                                        <p 
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row, scope.row.instanceId, 1.17, 'list')"
                                        >
                                            {{scope.row.instanceId}}
                                        </p>
                                    </el-tooltip>
                                </template>
                            </el-table-column>
                            <!-- 操作系统 -->
                            <el-table-column 
                                prop="imageName" 
                                :label="$t('instance.list.osSystem')" 
                            >
                                <template v-slot="scope">
                                    <el-tooltip
                                        v-model="scope.row.showTooltip"
                                        :disabled="!scope.row.showTooltip"
                                        :content= scope.row.imageName
                                    >
                                        <div
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row, scope.row.imageName, 1.17, 'list')"
                                        >
                                            <span>{{scope.row.imageName || $filter.emptyFilter()}}</span>
                                        </div>
                                    </el-tooltip>
                                </template>
                            </el-table-column>
                            <!-- 创建时间 -->
                            <el-table-column 
                                prop="createdTime" 
                                min-width="100"
                                :label="$t('instance.list.createdTime')" 
                            >
                                <template v-slot="scope">
                                    {{scope.row.createdTime || $filter.emptyFilter()}}
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                    <!-- 镜像类型 -->
                    <div class="image-gap">
                        <el-form-item  
                            :label="$t('instance.create.image') "  
                        >
                            <div 
                                v-if="Object.keys(osTypes.tableData).length"
                                class="image-content" 
                                :style="[{'max-height':imageHeight + 'px'}]"
                            >  
                                <el-radio-group 
                                    v-model="instanceResystemForm.imageType" 
                                    v-for="(item,key,index) in osTypes.tableData"
                                    :key="index"
                                    :label="key"
                                    @change="changeImageType"
                                >
                                    <el-radio-button :label=key type="primary" plain>
                                        <div  class="inline-position">    
                                            <div :class="[imageOtherIcon(key), 'icon-size']"></div>
                                        </div>
                                        <div>
                                            <div class="f14">{{key}}</div>
                                            <el-select 
                                                class="image-select-count"
                                                v-model="instanceResystemForm.image[key]" 
                                                size="small" 
                                                :placeholder="$t('instance.create.select')"
                                            >
                                                <el-option 
                                                    v-for="item in osTypes.tableData[key]" 
                                                    :key="item.imageId"
                                                    :label="item.osName"
                                                    :value="item.imageId"
                                                >
                                                </el-option>
                                            </el-select>   
                                        </div>
                                        <!-- instanceResystemForm.image[instanceResystemForm.imageType] 作为最后所需的镜像id值（imageId）-->
                                        <span 
                                            class="tick image-tick" 
                                            v-if="(instanceResystemForm.imageType === key.toString())"
                                        >
                                            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                        </span>    
                                    </el-radio-button>
                                </el-radio-group>
                            </div> 
                            <div v-else>{{'--'}}</div>
                            <el-button 
                                v-if="changeArrow && Object.keys(osTypes.tableData).length > 3"
                                class="arrow-text"
                                :icon="ArrowDown" 
                                circle
                                @click="imageDownText"
                            >
                            </el-button>
                            <el-button 
                                v-else-if="!changeArrow  && Object.keys(osTypes.tableData).length > 3"
                                class="arrow-text"
                                :icon="ArrowUp" 
                                circle
                                @click="imageUpText"
                            >
                            </el-button>   
                        </el-form-item> 
                    </div> 
                    <!-- 系统盘 -->
                    <el-form-item  
                        :label="$t('instance.create.systemDisk') " 
                    >
                        <div 
                            v-if="raids.tableData.length" 
                            class="system-raid-content"
                        >
                        <el-select v-model="instanceResystemForm.raid" placeholder="Select" style="width: 340px">
                                <el-option
                                    v-for="item in raids.tableData"
                                    :key="item.systemVolumeRaidId"
                                    :label="item.systemInfo.split(',')[0] +' | ' + $filter.emptyFilter(item.systemVolumeRaidName) + ')'"
                                    :value="item.systemVolumeRaidId"
                                >
                                    {{item.systemInfo.split(',')[0] +' | ' + $filter.emptyFilter(item.systemVolumeRaidName) + ')' || $filter.emptyFilter()}}
                                </el-option>
                            </el-select>  
                        </div> 
                        <div v-else>{{'--'}}</div>   
                    </el-form-item>
                    <!-- 系统盘分区 -->
                    <el-form-item  
                        v-if="systemDiskPartition.tableDataChange.length > 0"
                        :label="$t('instance.create.systemDiskPartition') " 
                    >
                        <div v-for="(item) in systemDiskPartition.tableDataChange" class="chamfer-position">
                            <div class="chamfer">
                                <div class="chamfer-inner">
                                    <div class="chamfer-up fw600">{{item.point + ' :' || $filter.emptyFilter()}}</div>
                                    <div 
                                        class="chamfer-down" 
                                        :class="getLocationItem==='en_US' && item.size === 999999999 ? 'lh12' : ''"
                                        >
                                        {{item.sizeChange || $filter.emptyFilter()}}
                                    </div>
                                    <div class="chamfer-down">{{item.format || $filter.emptyFilter()}}</div>
                                </div>   
                            </div>                  
                        </div>                                 
                    </el-form-item>
                    <el-form-item  
                        v-else
                        :label="$t('instance.create.systemDiskPartition') " 
                    >
                        <div>{{'--'}}</div>
                    </el-form-item>
                    <!-- 引导模式 -->
                    <el-form-item 
                        :label="$t('instance.create.guideMode') "
                        v-if="Object.keys(osTypes.tableData).length"
                    >
                        <el-radio-group 
                            v-model="instanceResystemForm.bootMode"
                            v-for="(item,key,index) in guideModes.tableData"
                            :key="index"
                            :label="key"
                            :value="item"
                        >
                            <el-radio-button :label="item" >{{item}}</el-radio-button>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item  
                        v-else
                        :label="$t('instance.create.guideMode') " 
                    >
                        <div >{{'--'}}</div>
                    </el-form-item>
                    <div class="instance-message">                     
                        <!-- HostName -->
                        <el-form-item  
                            prop="hostName"
                            :label="$t('instance.create.hostName') "
                        >
                            <el-input 
                                v-model="instanceResystemForm.hostName" 
                                maxlength="64"
                                :placeholder="$t('list.placeholder.hostName')" 
                            />
                        </el-form-item>
                        <!-- 用户名 -->
                        <el-form-item  
                            :label="$t('instance.create.userName')"
                            class="mt-10"
                        >
                            <span>{{instanceResystemForm.imageType === 'Windows' ? 'administrator' : 'root'}}</span>
                        </el-form-item>
                        <!-- 设置密码 -->
                        <el-form-item  
                            :label="$t('instance.create.setPassword')"
                            class="mt-10"
                        >
                            <el-radio-group 
                                v-model="instanceResystemForm.setPassword" 
                                class="password-size" @change="setPasswordChange()"
                            >
                                <el-radio label="1">{{$t('instance.create.customPassword')}}</el-radio>
                                <el-radio label="2">{{$t('instance.create.keyPasswordLogin')}}</el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <div v-if="instanceResystemForm.setPassword==='1'" class="password-change">
                            <!-- 密码 -->
                            <el-form-item  
                                prop="password"
                                ref="password"
                                :label="$t('instance.create.password') " 
                            >
                                <el-input 
                                    v-model="instanceResystemForm.password" 
                                    maxlength="30"
                                    minlength="8"
                                    type="password"
                                />
                                <div 
                                    class="password-tip el-form-item__error" 
                                    v-if="!newPasswordState"
                                >
                                    {{$t('personCentre.content.passwordVerification')}}
                                </div>
                            </el-form-item>
                            <!-- 确认密码 -->
                            <el-form-item  
                                prop="confirmPassword"
                                ref="confirmPassword"
                                :class="getLocationItem==='zh_CN' ? '' : 'password-label'"
                                :label="$t('instance.create.confirmPassword') "
                            >
                                <el-input 
                                    v-model="instanceResystemForm.confirmPassword" 
                                    type="password"
                                />
                            </el-form-item>
                        </div>
                        <div v-else>
                            <!-- 密钥密码 -->
                            <el-form-item  
                                class="mb42"
                                prop="keyPassword"
                                :label="$t('instance.create.keyPassword') " 
                            >
                                <div class="table-content m0">
                                    <el-table 
                                        border
                                        ref="sshTableRef"
                                        :data="sshTableData.tableData" 
                                        style="width: 631px"
                                        @row-click="rowClick"
                                        @selection-change="handleSelectionChange"
                                    >
                                        <el-table-column type="selection" width="55" />
                                        <el-table-column 
                                            prop="name" 
                                            :label="$t('personCentre.form.keyName')" 
                                            min-width="200"
                                        >
                                            <template v-slot="scope">
                                                <el-tooltip
                                                    v-model="scope.row.showTooltip"
                                                    :disabled="!scope.row.showTooltip"
                                                    :content= scope.row.name
                                                    placement="top-start"
                                                >
                                                    <p 
                                                        class="long-name-title"
                                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.name, 1.17, 'list')"
                                                    >
                                                        {{scope.row.name}}
                                                    </p>
                                                </el-tooltip>
                                            </template>
                                        </el-table-column>
                                        <el-table-column 
                                            prop="createdTime" 
                                            :label="$t('personCentre.form.createdTime')" 
                                            min-width="200"
                                        />
                                        <template #empty>
                                            <div>{{$t('instance.create.noData')}}</div>
                                        </template>
                                    </el-table>
                                </div>
                                <div 
                                    class="create-key-tip el-form-item__error" 
                                    :class="chooseEmptyKeyPasword ? 'mt25' : ''"
                                >
                                    <span>
                                        {{$t('instance.create.noMatchKey')}}
                                    </span>
                                    <a 
                                        class="mouse-point"
                                        @click="jumpToKey"
                                        target="_blank"
                                    >
                                        {{$t('instance.create.createdKey')}}
                                    </a> 
                                    <img
                                        alt=""
                                        @click="requestSshKey"
                                        class="refresh-img"
                                        src="@/assets/img/refresh.png"
                                    /> 
                                </div>
                            </el-form-item>
                        </div>
                    </div>
                    <!-- 安装监控Agent -->
                    <el-form-item  
                        prop="montiorAgent"
                        ref="montiorAgent"
                        class="download-monitorAgent-conetnt"
                        :label="$t('instance.detail.performanceMonitoring') "
                    >
                        <el-checkbox v-model="instanceResystemForm.montiorAgent" :label="$t('alarm.monitor.downloadMonitorAgent')"/>
                        <el-tooltip
                        :content= "$t('alarm.monitor.monitorTip')"
                        placement="right-start">
                        <img
                            class="tooltip-img"
                            src="@/assets/img/tooltip.png"
                            alt="暂无"
                        />
                    </el-tooltip>
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

import {Ref, ref, watch, nextTick, reactive} from 'vue';
import {
    Close,
    ArrowDown,
    ArrowUp
} from '@element-plus/icons-vue';
import type {
    FormInstance, // element-plus ui类
    FormRules, // element-plus ui类
    ElTable
} from 'element-plus';
import {
    passwordReg, // 密码正则
    hostNameReg // host名称正则
} from 'utils/regular.ts';
import {
    deepCopy, // 深拷贝
    hasShowTooltip, // 是否显示提示
} from 'utils/index.ts';
import VueCookies from 'vue-cookies'; // cookie
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import {instanceDeleteAPI, 
        osTypeAPI, 
        raidAPI, 
        systemDiskPartitionAPI,
        sshKeyAPI,
        resystemInstanceAPI,
        instanceDetailAPI
} from 'api/request.ts'; // 删除接口
import {
    imageType, // 镜像类型
} from 'utils/constants.ts'; 
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
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

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
 * 实例类型
 */
 interface Instance {
    imageType: string, //  镜像类型
    bootMode: string, // 引导模式
    image: object, // 镜像
    raid: string, // 系统盘
    instanceName: string, // 实例名称
    hostName: string, // HostName
    setPassword: string, // 设置密码
    password: string, //密码
    confirmPassword: string, // 确认密码
    keyPassword: string, // 密钥
    montiorAgent: boolean, // 安装监控Agent
}
/**
 * 实例表单
 */
const instanceResystemForm:Instance = reactive({
    imageType: '', // 镜像类型
    image: {}, // 镜像
    bootMode: '', // 引导模式
    raid: '', // 系统盘
    instanceName: '', // 实例名称
    hostName: '', // HostName
    setPassword: '1', // 设置密码
    password: '', // 密码
    confirmPassword: '', // 确认密码
    keyPassword: '', // 密钥
    montiorAgent: true,
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
    hostName: UserInfoType[],//HostName
    password: UserInfoType[],//密码
    confirmPassword: UserInfoType[],//确认密码
    keyPassword: UserInfoType[],//密钥
};

/**
 * host名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const hostNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (!hostNameReg.test(value)) {
        callback(new Error(t('personCentre.content.hostnameVerification')));
    }
    else {
        callback();
    }
};
/**
 * 密码正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const passwordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    newPasswordState.value = true
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyCurrentPassword')));
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
    else if (value !== instanceResystemForm.password) {
        callback(new Error(t('personCentre.toolTip.passNotSame')));
    }
    else {
        callback();
    }
};
/**
 * 判断密钥为空时表格底下样式
 */
const chooseEmptyKeyPasword: Ref<boolean> = ref<boolean>(false);
/**
 * 密钥正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const keyPasswordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        chooseEmptyKeyPasword.value = true;
        callback(new Error(t('personCentre.toolTip.emptyKeyPassword')));
    }
    else {
        chooseEmptyKeyPasword.value = false;
        callback();
    }
};

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    hostName: [ // host名称
        {
            required: false,
            trigger: 'blur',
            validator: hostNameCheck
        }
    ],
    password: [ // 密码
        {
            required: false,
            trigger: 'blur',
            validator: passwordCheck
        }
    ],
    confirmPassword: [ // 确认密码
        {
            required: false,
            trigger: 'blur',
            validator: confirmPasswordCheck
        }
    ],
    keyPassword: [ // 密钥
        {
            required: false,
            trigger: 'change',
            validator: keyPasswordCheck
        }
    ],
});


/**
 * 镜像类型div高度
 */
 const imageHeight: Ref<number> = ref<number>(160);
/**
 * 获取镜像类型div
 */
 var imageDiv: any = document.getElementsByClassName('image-content');
/**
 * 判断镜像类型下拉上提方向
 */
 const changeArrow: Ref<boolean> = ref<boolean>(true);
/**
 * 镜像类型下拉操作
 */
 const imageDownText : () => void = () => {
    imageHeight.value = imageDiv[0].scrollHeight // 下拉高度为元素内容的高度，包括溢出的不可见内容
    changeArrow.value = !changeArrow.value
}
/**
 * 镜像类型上提操作
 */
const imageUpText : () => void = () => {
    imageHeight.value = 160
    changeArrow.value = !changeArrow.value

}

/**
 * 镜像类型有特殊时，使用Other图片
 * @param key 
 */
 const imageOtherIcon = (key: any) => {
    if ([...imageType].some((item: string) => item === key)) {
        return key;
    } else {
        return 'Other';
    }
}

/**
 * 镜像item类型
 */
 type OsType =  {
    CentOS?: any,
    Ubuntu?: any,
    Windows?: any,
}
/**
 * 镜像类型
 */
 interface os {
    tableData: OsType[]
}
/**
 * 镜像类型数据
 */
const osTypes: os = reactive<os>({
    tableData: []
})

/**
 * 引导模式类型
 */
 interface guide {
    tableData: string
}
/**
 * 引导模式列表数据
 */
const guideModes: guide = reactive<guide>({
    tableData: ''
})

/**
 * 镜像类型变化，设为第一项
 * @param event 
 */
 const changeImageType = (event: any) => {
    if(!instanceResystemForm.image[event]) {
        instanceResystemForm.image[event] = osTypes.tableData[event][0].imageId; // 初始化镜像类型为第一项
    }    
}
/**
 * 请求镜像信息接口
*/
const requestOsTypeData = (): void => {
    osTypeAPI({
        idcId: props.instanceInfo.idcId,
        deviceTypeId: props.instanceInfo.deviceTypeId
    }).then(({data} : {data: any}) => {
        if(Object.keys(data?.result)?.length) {
            osTypes.tableData = data?.result;
            instanceResystemForm.imageType = Object.keys(osTypes.tableData)[0]; // 初始化镜像为第一项
            let osTypeFirst = osTypes.tableData[Object.keys(osTypes.tableData)[0]]; // 选出镜像第一项
            instanceResystemForm.image[instanceResystemForm.imageType] = osTypeFirst[0].imageId; // 初始化镜像类型为第一项
            guideModes.tableData = osTypeFirst[0].bootMode.split(',') // 初始化引导模式
            instanceResystemForm.bootMode = guideModes.tableData[0] //选出引导模式第一项
            return
        }
        osTypes.tableData = [];
        instanceResystemForm.imageType = '';
        instanceResystemForm.image[instanceResystemForm.imageType] = ''
    }).finally(() => {
        //requestRaidData();
        if(instanceResystemForm.image[instanceResystemForm.imageType]) {
            requestSystemDiskPartitionData()
        } else {
            systemDiskPartition.tableData = [];
            systemDiskPartition.tableDataChange = [];
        }
        //raids.tableData = [props.instanceInfo];
        requestInstanceDetailData()
        instanceResystemForm.raid = props.instanceInfo.systemVolumeRaidId;
        //carouselRef.value.setActiveItem(Math.trunc(deviceTypeNumber.value/4))
    });
};
requestOsTypeData();

/**
 * 请求实例详情数据接口
*/
const requestInstanceDetailData = (): void => {
    instanceDetailAPI({
        instanceId: props.instanceInfo.instanceId
    }).then(({data} : {data: any}) => {
        if (data?.result?.instance) {
            raids.tableData = [data.result.instance];
            return;
        }
        return Promise.reject();
    }).catch(({message} : {message: string;}) => {      
    }).finally(() => {
   
    });
};

/**
 * 寻找镜像的引导模式
 */
 const findImageForBootMode = (): void => {
    let item: any = []
    for(item of Object.values(osTypes.tableData)) {
        item.find((ele: any) => {
            if(ele.imageId === instanceResystemForm.image[instanceResystemForm.imageType]) {
                guideModes.tableData = ele.bootMode.split(',');
                instanceResystemForm.bootMode = guideModes.tableData[0]
            }
        })
    }
}
/**
 * 监听镜像信息，每次改变，重新获取系统盘信息
 */
 watch(() => instanceResystemForm.image[instanceResystemForm.imageType], (): void => {
    if(instanceResystemForm.image[instanceResystemForm.imageType]) {
        requestSystemDiskPartitionData();
        findImageForBootMode()
    }else {
        systemDiskPartition.tableData = [];
        systemDiskPartition.tableDataChange = [];
    }
});

/**
 * 计算字符串长度,判断是否弹出tooltip
 * @param str 
 */
const lengthNum = (str: string, num: number): boolean => {
    var length = 0;
    if(str) {
        Array.from(str).map((char: any)=>{
            if(char.charCodeAt(0) > 255) {// 字符编码大于255，说明是双字节字符
                length += 2;
            }else {
                length ++;
            }
        });
    }
    let contentWidth = length;
    // 判断是否开启tooltip功能
    if (contentWidth > num) {
        return false;
    } else {
        return true;
    }
};
/**
 * 系统盘item类型
 */
 type RaidType =  {
    diskType: string,
    raidId: string,
    raidName: string,
    systemPartitionCount: 0,
    volumeDetail: string,
    volumeType: string
}
/**
 * 系统盘类型
 */
interface raid {
    tableData: RaidType[]
}
/**
 * 系统盘数据
 */
const raids: raid = reactive<raid>({
    tableData: []
})


/**
 * 请求系统盘信息接口
*/
const requestRaidData = (): void => {
    raidAPI({
        deviceTypeId: props.instanceInfo.deviceTypeId,
        volumeType: 'system'
    }).then(({data} : {data: any}) => {
        if(Object.keys(data?.result)?.length) {
            raids.tableData = data.result;
            instanceResystemForm.raid = raids.tableData[0].raidId; // 初始化系统盘为第一项
            return
        }
        raids.tableData = [];
        instanceResystemForm.raid =''
    }).finally(() => {
        if(instanceResystemForm.image[instanceResystemForm.imageType]) {
            requestSystemDiskPartitionData()
        } else {
            systemDiskPartition.tableData = [];
            systemDiskPartition.tableDataChange = [];
        }
    });
};

/**
 * 系统盘分区item类型
 */
 type PartitionType =  {
    format: string,
    point: string,
    size: number
}

/**
 * 系统盘分区item类型
 */
type PartitionChangeType =  {
    format: string,
    point: string,
    size: number,
    sizeChange: number
}
/**
 * 系统盘分区类型
 */
interface partition {
    tableData: PartitionType[],
    tableDataChange: PartitionChangeType[]
}
/**
 * 系统盘分区数据
 */
const systemDiskPartition: partition = reactive<partition>({
    tableData: [],
    tableDataChange: []
})

/**
 * 请求系统盘分区信息接口
*/
const requestSystemDiskPartitionData = (): void => {
    systemDiskPartitionAPI({
        imageId: instanceResystemForm.image[instanceResystemForm.imageType],
        deviceTypeId: props.instanceInfo.deviceTypeId
    }).then(({data} : {data: any}) => {
        if(data?.result) {
            systemDiskPartition.tableData = deepCopy(data.result);
            systemDiskPartition.tableDataChange = data?.result;
            systemDiskPartition.tableDataChange.forEach((item: any) => {
                if(item.size === 999999999) {
                    item.sizeChange = t('instance.create.remain')
                } else {
                    item.sizeChange = Math.floor( item.size / 1024) + ' GiB'
                }
                
            })           
        }
    }).finally(() => {
    });
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 设置密码切换
 */
 const setPasswordChange=()=> {    
    instanceResystemForm.password = '';
    instanceResystemForm.confirmPassword = '';
    instanceResystemForm.keyPassword = '';
    newPasswordState.value = false;
    chooseEmptyKeyPasword.value = false;
    ruleFormRef.value.clearValidate(['password', 'confirmPassword']);  
}

/**
 * 表格ref
*/
const sshTableRef = ref<InstanceType<typeof ElTable>>()

/**
 * 表格点击某一行触发选中事件
 * @param {Object} row 当前选中的这一项
 */
 const rowClick = (row: any) => {
    row.rowChange = !row.rowChange;
    sshTableRef.value!.toggleRowSelection(row, row.rowChange);
};

/**
 * ssh密钥参数结合
 */
const handleSelectionChange = (val: any) => {
    let arr: string = ''
    arr = val.map((item: any) => {
         return item.keypairId
    }).join(',')
    instanceResystemForm.keyPassword = arr
}

/**
 * 密码状态
 */
 const newPasswordState: Ref<boolean> = ref<boolean>(false)

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
    tableData: SshType[]
}
/**
 * 个人ssh密钥数据
 */
 const sshTableData: ssh = reactive<ssh>({
    tableData: []
})

/**
 * 请求密钥接口
*/
const requestSshKey = (): void => {
    sshKeyAPI({
        isAll:'1'
    }).then(({data} : {data: any}) => {
        if (data?.result?.keypairs?.length) {
            sshTableData.tableData = data.result.keypairs.map((item:any) => {
                return {
                    ...item,
                    showTooltip: false
                };
            });
        return;
    }
    sshTableData.tableData  = [];
    }).catch(() => {
        sshTableData.tableData  = [];
    })
};
requestSshKey()

/**
 * 跳转到ssh密钥创建
 */
const jumpToKey = () => {
    window.open('/user?type=sshKey', '_blank')
};

/**
 * 点击取消
 */
 const closeDialog:  () => void = () => {
    emitClose('close', false)
}

/**
 * 重装系统接口
*/
const requestResystemInstance = (): void => {
    let params = {
        hostname: instanceResystemForm.hostName,
        instanceId: props.instanceInfo.instanceId,
        imageId: instanceResystemForm.image[instanceResystemForm.imageType],
        bootMode: instanceResystemForm.bootMode,
        instanceName: props.instanceInfo.instanceName,
        password: instanceResystemForm.password,
        sshKeyId: instanceResystemForm.keyPassword,
        systemVolumeRaidId: instanceResystemForm.raid,
        systemPartition: systemDiskPartition.tableData,
        username: instanceResystemForm.imageType === 'Windows' ? 'administrator' : 'root',
        installBmpAgent: instanceResystemForm.montiorAgent,
    }
    resystemInstanceAPI({
            ...params
    }).then(({data} : {data: any}) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type:'success'
            });
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

document.onkeydown = (event: {keyCode: number;}) => {
    event.keyCode === 13 && openDialog();
};

/**
 * 点击确定
 */
 const openDialog = async (): Promise<void> => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        ruleFormRef.value.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                requestResystemInstance();
            }
        });
    });
};
</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/icon';
@import './instanceResystem.scss';
@import 'assets/css/table';
@import 'assets/css/icon';
</style>