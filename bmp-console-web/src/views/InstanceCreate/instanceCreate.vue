<template>
    <div class="page-position can-scroll">
        <div class="detail-header">
            <el-button
                class="back-btn"
                type= "text"
                :icon ="ArrowLeft"
                @click="jumpToList(ruleFormRef)"
            >
            </el-button>
            <span class="ml10">{{$t('instance.list.create') }}</span>
        </div>
        <div class="page-content t120 instance-create">
            <div class="form-content">
                <el-form
                    label-position="left"
                    :model="instanceForm"
                    ref="ruleFormRef"
                    :rules="rules"
                    label-width="110px"
                    @submit.native.prevent
                >
                    <!-- 机房 -->
                    <el-form-item 
                        :label="$t('instance.create.engine') "
                        class="lh36"
                    >
                        <el-radio-group v-model="instanceForm.az"
                            v-for="(item,key,index) in azs.tableData"
                            :key="index"
                            :label="key"
                            :value="item.idcId"
                        >
                            <el-radio-button :label="item.idcId" >
                                {{getLocationItem==='zh_CN' ? item.name : item.nameEn}}
                            </el-radio-button>
                        </el-radio-group>
                    </el-form-item>
                    <!-- 机型类型 -->
                    <el-form-item 
                        :label="$t('instance.create.model') "
                        class="lh36"
                    >
                        <div v-if="Object.keys(deviceTypes.tableData).length" style="display: contents">
                            <el-radio-group v-model="instanceForm.model"
                                v-for="(item,key,index) in deviceTypes.tableData"
                                :key="index"
                                :label="key"
                                @change="modelChange"
                            >
                                <el-radio-button :label="key" >{{seriesMap[key]}}</el-radio-button>
                            </el-radio-group>
                        </div>
                        <div v-else>{{'--'}}</div>
                    </el-form-item>              
                    <!-- 机型选择类型 -->
                    <div 
                        class="model-content" 
                        :class="Object.keys(deviceTypes.tableData[instanceForm.model]).length > 4 ? 'h265' : 'h240'"
                        v-if="Object.keys(deviceTypes.tableData).length"
                    >
                        <el-carousel 
                            v-if="deviceCarousel"                           
                            indicator-position="outside" 
                            ref="carouselRef"        
                            :autoplay="false" 
                            :class="modelUnit.list.length < 2 ? 'indicators-none' : modelNoneArrow"
                            @change="changeModel"
                        >
                            <el-carousel-item 
                                v-for="unit in modelUnit.list" 
                                :key="unit"
                            >
                                <el-radio-group v-model="instanceForm.modelType"
                                    v-for="(item,key,index) in unit"            
                                    :key="index"
                                    :label="key"
                                >
                                    <el-radio-button 
                                        :label="item.deviceTypeId" 
                                        :disabled="item.availableStock === 0 || item.availableStock < instanceForm.number"
                                    >
                                        <!-- 可售卖的机型 -->                                   
                                        <el-tooltip
                                            :disabled="lengthNum(item.name, 18)" 
                                            :content= item.name
                                        >
                                            <p  class="long-device-item">
                                                <span 
                                                    :class="[instanceForm.model, 'model-picture-grey']" 
                                                    class="t2"
                                                >
                                                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                                </span>
                                                {{item.name || $filter.emptyFilter()}}
                                            </p>
                                        </el-tooltip>
                                        <div class="raid-down">
                                            <div class="divider-line">
                                                <el-tooltip
                                                    :disabled="lengthNum($t('instance.detail.cpu') + ': ' + item.cpu, 28)" 
                                                    :content= item.cpu
                                                >
                                                    <p  class="long-item">
                                                        <span class="fw600">{{$t('instance.detail.cpu') + ': '}}</span>
                                                        {{item.cpu || $filter.emptyFilter()}}
                                                    </p>
                                                </el-tooltip>
                                                <el-tooltip
                                                    :disabled="lengthNum($t('instance.create.memoryShort') + ': ' + item.mem, 26)" 
                                                    :content= item.mem
                                                >
                                                    <p class="long-item">
                                                        <span class="fw600">{{$t('instance.create.memoryShort') + ': '}}</span>
                                                        {{item.mem || $filter.emptyFilter()}}
                                                    </p>
                                                </el-tooltip>
                                                <el-tooltip
                                                    :disabled="lengthNum($t('instance.create.systemDiskShort') + ': ' + item.system, 30)" 
                                                    :content= item.system
                                                >
                                                    <p class=" long-item">
                                                        <span class="fw600">{{$t('instance.create.systemDiskShort') + ': '}}
                                                        </span>{{item.system || $filter.emptyFilter()}}
                                                    </p>
                                                </el-tooltip>
                                                <!-- <el-tooltip
                                                    :disabled="lengthNum($t('instance.create.dataDiskShort') + ': ' + item.data, 28)" 
                                                    :content= item.data
                                                >
                                                    <p class=" long-item">
                                                        <span class="fw600">{{$t('instance.create.dataDiskShort') + ': '}}</span>
                                                        {{item.data || $filter.emptyFilter()}}
                                                    </p>
                                                </el-tooltip> -->
                                                <p class="long-item"><span class="fw600">{{$t('instance.detail.networkCard') + ': '}}
                                                </span>{{joinObjectKeyValues(item.nics, 'nicInfo') || $filter.emptyFilter()}}</p>
                                                <p class="long-item"><span class="fw600">{{$t('instance.detail.gpu') + ': '}}
                                                </span>{{item.gpu || $filter.emptyFilter()}}</p>    
                                                <span 
                                                    v-if="item.availableStock >= instanceForm.number" 
                                                    :class="instanceForm.modelType === item.deviceTypeId ? 'tick model-tick':''"
                                                >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                                </span>
                                                <div v-else>
                                                    <div v-if="getLocationItem==='zh_CN'" 
                                                        :class="instanceForm.modelType===item.deviceTypeId ? 'soldout-zh model-soldout' : 'soldout-zh-i model-soldout'"
                                                    >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                                    </div> 
                                                    <span 
                                                        v-else 
                                                        :class="instanceForm.modelType === item.deviceTypeId ? 'soldout-en model-soldout' : 'soldout-en-i model-soldout'"
                                                    >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                                    </span> 
                                                </div>                                              
                                            </div>      
                                        </div>
                                    </el-radio-button>    
                                </el-radio-group> 
                            </el-carousel-item>
                        </el-carousel>                                                   
                    </div> 
                    <!-- 镜像类型 -->
                    <el-form-item  
                        :label="$t('instance.create.image') "  
                        class="lh60 mb30"
                        v-if="Object.keys(osTypes.tableData).length"
                    >
                        <div class="image-content" :style="[{'max-height':imageHeight + 'px'}]">  
                            <el-radio-group 
                                v-model="instanceForm.imageType" 
                                v-for="(item,key,index) in osTypes.tableData"
                                :key="index"
                                :label="key"
                                @change="changeImageType"
                            >
                                <el-radio-button :label=key type="primary" plain>
                                    <div class="inline-position">    
                                        <div :class="[imageOtherIcon(key), 'icon-size']"></div>
                                    </div>
                                    <div>
                                        <div class="f14">{{key}}</div>
                                        <el-select 
                                            class="image-select-count"
                                            v-model="instanceForm.image[key]" 
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
                                    <!-- instanceForm.image[instanceForm.imageType] 作为最后所需的镜像id值（imageId）-->
                                    <span 
                                        class="tick image-tick" 
                                        v-if="(instanceForm.imageType === key.toString())"
                                    >
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                    </span>    
                                </el-radio-button>
                            </el-radio-group>
                        </div> 
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
                    <el-form-item  
                        :label="$t('instance.create.image') " 
                        class="mt-10"
                        v-else
                    >
                        <div >{{'--'}}</div>
                    </el-form-item>
                    <!-- 引导模式 -->
                    <el-form-item 
                        :label="$t('instance.create.guideMode') "
                        v-if="Object.keys(osTypes.tableData).length"
                        class="lh36"
                    >
                        <el-radio-group v-model="instanceForm.bootMode"
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
                        class="mt-10 mb10"  
                    >
                        <div >{{'--'}}</div>
                    </el-form-item>
                    <!-- 系统卷 -->
                    <el-form-item 
                        v-if="raids.tableData.length" 
                        :label="$t('instance.create.systemRoll') "
                    >
                          <el-select v-model="instanceForm.raid" placeholder="Select" style="width: 340px">
                            <el-option
                                v-for="item in raids.tableData"
                                :key="item.raidId"
                                :label="item.detail.split(')')[0] + ' | ' + item.raidName + ')'"
                                :value="item.raidId"
                            >
                            {{$filter.emptyFilter(item.detail.split(')')[0] + ' | ' + item.raidName + ')')}}
                            </el-option>
                        </el-select> 
                    </el-form-item>
                    <el-form-item  
                        v-else
                        :label="$t('instance.create.systemRoll') " 
                        class="mb20"
                    >
                        <div>--</div>
                    </el-form-item>
                    <!-- 系统卷分区 -->
                    <el-form-item  
                        v-if="systemDiskPartition.tableDataChange.length > 0"
                        :label="$t('instance.create.systemRollPartition') " 
                        :class="getLocationItem==='zh_CN' ? 'textarea-label' : 'system-partition-label'"
                        class="mb30"
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
                        class="mt-10"
                        :class="getLocationItem==='zh_CN' ? '' : 'empty-system-partition-label'"
                        :label="$t('instance.create.systemRollPartition') " 
                    >
                        <div>--</div>
                    </el-form-item>
                    <!-- 数据卷 -->
                    <el-form-item
                        v-if="datas.tableData.length" 
                        :label="$t('instance.create.dataRoll') " 
                    >   
                        <div class="table-content m0 h-max">
                            <el-table 
                                border
                                ref="datasTableRef"
                                :data="datas.tableData" 
                            >   
                                <!-- 卷名称 -->
                                <el-table-column 
                                    prop="volumeName" 
                                    :label="$t('instance.create.rollName')" 
                                    width="120"
                                />
                                <!-- RAID模式 -->
                                <el-table-column 
                                    v-if="raidIdShow"
                                    prop="raidModel" 
                                    :label="$t('instance.create.raidModel')" 
                                    min-width="200"
                                >
                                    <template v-slot="scope">
                                        <el-select 
                                            v-model="scope.row.raidModel" 
                                            @change="raidModelChange($event, scope.row)"
                                        >
                                            <el-option v-for="(item, index) in scope.row.raids"
                                                :key="item.raidId"
                                                :label="item.raidName"
                                                :value="index"
                                            />
                                        </el-select>
                                    </template>
                                </el-table-column>
                                <!-- 硬盘数量 -->
                                <el-table-column 
                                    prop="volumeAmount" 
                                    :label="$t('instance.create.harddiskNumber')" 
                                    min-width="100"
                                />
                                <!-- 单个硬盘容量 -->
                                <el-table-column 
                                    prop="harddiskVolume" 
                                    :label="$t('instance.create.harddiskVolume')" 
                                    min-width="200"
                                >
                                    <template v-slot="scope">
                                        <p>{{scope.row.volumeSize + scope.row.volumeUnit}}</p>
                                    </template>
                                </el-table-column>
                                <!-- 可用硬盘容量 -->
                                <el-table-column 
                                    prop="useVolume" 
                                    :label="$t('instance.create.useVolume')" 
                                    min-width="200"
                                >
                                    <template v-slot="scope">
                                        <p>{{scope.row.volumeSize + scope.row.volumeUnit}}</p>
                                    </template>
                                </el-table-column>
                                <!-- 硬盘类型 -->
                                <el-table-column 
                                    prop="diskType" 
                                    :label="$t('instance.create.harddiskType')" 
                                    min-width="100"
                                />
                                <template #empty>
                                    <div>{{$t('instance.create.noData')}}</div>
                                </template>
                            </el-table>
                        </div>
                    </el-form-item>
                    <div class="instance-message">
                        <!-- 实例名称 -->
                        <el-form-item 
                            prop="instanceName"
                            :label="$t('instance.create.instanceName') "
                        >
                            <el-input 
                                v-model="instanceForm.instanceName" 
                                maxlength="128"
                                :placeholder="$t('list.placeholder.instanceName')"
                            />                           
                        </el-form-item>                       
                        <!-- HostName -->
                        <el-form-item  
                            prop="hostName"
                            :label="$t('instance.create.hostName') "
                        >
                            <el-input 
                                v-model="instanceForm.hostName" 
                                maxlength="64"
                                :placeholder="$t('list.placeholder.hostName')" 
                            />
                        </el-form-item>
                        <!-- 用户名 -->
                        <el-form-item  :label="$t('instance.create.userName') ">
                            <span>{{instanceForm.imageType === 'Windows' ? 'administrator' : 'root'}}</span>
                        </el-form-item>
                        <!-- 设置密码 -->
                        <el-form-item  :label="$t('instance.create.setPassword') ">
                            <el-radio-group 
                                v-model="instanceForm.setPassword" 
                                class="password-size" @change="setPasswordChange(ruleFormRef)"
                            >
                                <el-radio label="1">{{$t('instance.create.customPassword')}}</el-radio>
                                <el-radio label="2">{{$t('instance.create.keyPasswordLogin')}}</el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <div v-if="instanceForm.setPassword==='1'">
                            <!-- 密码 -->
                            <el-form-item  
                                :class="getLocationItem==='zh_CN' ? 'mb42' : 'mb57'"
                                prop="password"
                                ref="password"
                                :label="$t('instance.create.password') " 
                            >
                                <el-input 
                                    v-model="instanceForm.password" 
                                    maxlength="30"
                                    minlength="8"
                                    type="password"
                                />
                                <div 
                                    class="password-tip el-form-item__error" 
                                    v-if="!newPasswordState">
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
                                    v-model="instanceForm.confirmPassword" 
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
                                            width="200"
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
                                        target="_blank">
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
                        <el-checkbox v-model="instanceForm.montiorAgent" :label="$t('alarm.monitor.downloadMonitorAgent')"/>
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
                    <!-- 数量 -->
                    <el-form-item  :label="$t('instance.create.quantity') " :class="instanceMountClass">
                        <el-input-number 
                            v-model="instanceForm.number" 
                            :min="1" 
                            :max="avaliableMaxCount"
                            @change="numberhandleChange" 
                        />   
                    </el-form-item>
                    <div class="create-confirm">
                        <el-button 
                            class="create-confirm-button cancel-button" 
                            @click="jumpToList(ruleFormRef)"
                        >
                            {{$t('instance.create.cancel')}}
                        </el-button>
                        <el-tooltip
                            :disabled="!createInstanceState" 
                            :content="$t('instance.tip.allSell')"
                            placement="top"
                        >
                            <el-button 
                                type="primary" 
                                :loading="isLoading"
                                class="create-confirm-button" 
                                :disabled="createInstanceState " 
                                @click="createInstance(ruleFormRef)"
                            >
                                {{$t('instance.create.create')}}
                            </el-button>
                        </el-tooltip>
                    </div>
                </el-form>
            </div>
        </div>        
    </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute} from 'vue-router';
import {
  ArrowLeft,
  ArrowDown,
    ArrowUp
} from '@element-plus/icons-vue';
import { 
    reactive,
    ref,
    Ref,
    onUnmounted,
    onMounted,
    computed,
    nextTick,
    watch,
    getCurrentInstance
} from 'vue';
import {
    passwordReg, // 密码正则
    instanceNameReg, // 实例名称正则
    hostNameReg // host名称正则
} from 'utils/regular.ts';
import {useI18n} from 'vue-i18n'; // 国际化
import type {
    FormInstance, // element-plus ui类
    FormRules, // element-plus ui类
    ElTable
} from 'element-plus';
import {
    joinObjectKeyValues,
    deepCopy, // 深拷贝
    hasShowTooltip, // 是否显示提示
} from 'utils/index.ts'; 
import { ElMessage } from 'element-plus';
import VueCookies from 'vue-cookies'; // cookie
import allProjectStore from '@/store/modules/headerDetail.ts';
import useProjectStore from '@/store/modules/projectId.ts';
import {
    imageType, // 镜像类型
} from 'utils/constants.ts'; 
import {
    azAPI,
    deviceTypeAPI,
    osTypeAPI,
    raidAPI,
    systemDiskPartitionAPI,
    createInstanceAPI,
    sshKeyAPI,   
    isDeviceStockEnoughAPI,
    instanceListAPI
} from 'api/request.ts'; // 机房, 机型, 镜像, 系统盘, 系统盘分区信息接口，创建实例接口

/**
 * 路由
 */
const route = useRoute();

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 设置类
*/
type SetType<T> = T extends {} ? any : T;

/**
 * 使用mitt传值
*/
const instanceMitt: Exclude<Required<SetType<{}>> | null, never> = getCurrentInstance();

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 实例类型
 */
interface Instance {
    az: string, // 机房
    model: string, // 机型选择
    modelType: string, // 机型
    imageType: string, //  镜像类型
    bootMode: string, // 引导模式
    image: object, // 镜像
    raid: string, // 系统盘
    systemChoose: Object, // 系统卷选择
    instanceName: string, // 实例名称
    hostName: string, // HostName
    setPassword: string, // 设置密码
    password: string, //密码
    confirmPassword: string, // 确认密码
    montiorAgent: boolean, // 安装监控Agent
    keyPassword: string, // 密钥
    number: number, // 数量
}

/**
 * 实例表单
 */
const instanceForm: Instance = reactive({
    az: '', // 机房
    model: '', // 机型选择
    modelType: '', // 机型
    imageType: '', // 镜像类型
    image: {}, // 镜像
    bootMode: '', // 引导模式
    raid: '', // 系统盘
    systemChoose: {},
    dataRaid: '', //数据卷
    instanceName: '', // 实例名称
    hostName: '', // HostName
    setPassword: '1', // 设置密码
    password: '', // 密码
    confirmPassword: '', // 确认密码
    montiorAgent: true,
    keyPassword: '', // 密钥
    number: 1, // 数量
})

/**
 * 机房item类型
 */
 type IdcType = {
    idcId: string;
    name: string;
    nameEn: string;
    shortname: string;
};

/**
 * 机房类型
 */
 interface idc {
    tableData: IdcType[]
}

/**
 * 机房列表数据
 */
const azs: idc = reactive<idc>({
    tableData: []
})

/**
 * 机型item类型
 */
type DeviceType =  {
    compute?: any,
    gpu?: any,
    storage?: any,
    other?: any 
}

/**
 * 机房类型
 */
interface device {
    tableData: DeviceType[]
}

/**
 * 机型数据
 */
const deviceTypes: device = reactive<device>({
    tableData: []
})

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
 * 机型系列对应的中文名称
 */
const seriesMap = {
    computer: t('instance.create.series.computeType'),
    storage: t('instance.create.series.storageType'),
    gpu: t('instance.create.series.gpuType'),
    other: t('instance.create.series.otherType'),
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
 * 系统盘item类型
 */
type RaidType =  {
    diskType: string,
    raidId: string,
    raidName: string,
    systemPartitionCount: 0,
    volumeDetail: string,
    volumeType: string,
    volumeId: string,
    raidCan: string
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
 * 数据卷item类型
 */
 type DataType =  {
    volumeName: string,
    volumeAmount: 0,
    deviceTypeId: string,
    diskType: string,
    raidModelChoose: {}
}

/**
 * 数据卷类型
 */
interface data {
    tableData: DataType[]
}

/**
 * 系统盘数据
 */
const datas: data = reactive<data>({
    tableData: []
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

const instanceMountClass = computed(() => {   
    return 'change-input-number';
});

/**
 * 创建按钮状态
 */
const createInstanceState: Ref<boolean> = ref<boolean>(false)

/**
 * 规则类
*/
interface RulesType {
    instanceName: UserInfoType[],//实例名称
    hostName: UserInfoType[],//HostName
    password: UserInfoType[],//密码
    confirmPassword: UserInfoType[],//确认密码
    keyPassword: UserInfoType[],//密钥
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 实例名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
 * 
*/
const instanceNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyInstanceName')));
    }
    else if (instanceData.repeatInstanceData.includes(value)) {
        callback(new Error(t('personCentre.content.instanceNameExist')));
    }
    else if(!instanceNameReg.test(value)) {
        callback(new Error(t('personCentre.content.instanceVerification')));
    }else {
        callback();
    }
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
    else if (value !== instanceForm.password) {
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
    instanceName: [ // 实例名称
        {
            required: false,
            trigger: 'blur',
            validator: instanceNameCheck
        }
    ],
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
 * 页面销毁时-触发，清空延时器
*/
onUnmounted(() => {
    clearTimeout(timer);
});

onMounted(() => {
    changeModel(0);
});

// 切换键头样式
const modelNoneArrow: Ref<string> = ref<string>('')

/**
 * 机型轮播图到第一页和最后一页时，切换箭头取消
 * @param event 
 */
const changeModel = (event: any) => {
    if(event === 0) {
        modelNoneArrow.value = 'indicators-none-left'
    }
    else if(event === modelUnit.list.length-1){
        modelNoneArrow.value = 'indicators-none-right'
    }else {
        modelNoneArrow.value = ''
    }
}

// 延时器
let timer = ref<any>(null);

// 判断展示机型轮播图
const deviceCarousel: Ref<boolean> = ref<boolean>(false)

/**
 * 机型选择改变时底下型号置为1
 */
const modelChange  = () => {
    deviceCarousel.value = false
    if (deviceTypes.tableData[instanceForm.model]) {
        dividedUnit(deviceTypes.tableData[instanceForm.model])
        // 延时取数据
        timer = setTimeout(() => {
            deviceCarousel.value = true
            // 将走马灯指示器（原type为submit）按钮type设为button，防止点击提交页面
            let carouselButton:any = document.getElementsByClassName('el-carousel__button');
            for(var i = 0; i < carouselButton.length; i++) {
                carouselButton[i].type = 'button';    
            }
            changeModel(0); 
        }, 200);
    }
};

/**
 * 镜像类型变化，设为第一项
 * @param event 
 */
const changeImageType = (event: any) => {
    if(!instanceForm.image[event]) {
        instanceForm.image[event] = osTypes.tableData[event][0].imageId; // 初始化镜像类型为第一项
    }    
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
    instanceForm.keyPassword = arr
}

/**
 * 数量变化
 * @param value 
 */
const numberhandleChange = () => {
    requestIsDeviceStockEnough()
};

/**
 * 设置密码切换
 * @param formEl
 */
const setPasswordChange=(formEl: FormInstance | undefined)=> {    
    instanceForm.password = '';
    instanceForm.confirmPassword = '';
    instanceForm.keyPassword = '';
    newPasswordState.value = false;
    chooseEmptyKeyPasword.value = false;
    formEl.clearValidate(['password', 'confirmPassword']);  
}

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

// 机型选择一个轮播图中的单元
let modelUnit = reactive({
    list: [] as any[]
});

// 将机型选择分组映入轮播图中的单元
const dividedUnit = (arr: any) => {
    modelUnit.list = [];
    var unit = [];
    if(arr.length === 1) {
        unit.push(arr)
    }else {
        for(var i = 0, j = arr.length; i < j; i += 4) {
            unit.push(arr.slice(i, i + 4))
        }
    }
    modelUnit.list = unit
};

/**
 * 跳转到列表页
 */
const router = useRouter();
const projectStore = useProjectStore();
const jumpToList = (formEl: FormInstance | undefined) => {
    formEl.resetFields();
    projectStore.setProject(route.query.projectId, route.query.projectName);
    router.push({
        path: `/instance/list`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName
        }
    });
};

/**
 * 跳转到ssh密钥创建
 */
const jumpToKey = () => {
    window.open('/user?type=sshKey', '_blank')
    // router.push({
    //     path: `/user`,
    //     query: {
    //         type: 'sshKey',   
    //     }
    // });
};

/**
 * 请求机房信息接口
*/
const requestAzData = (): void => {
    azAPI({        
    }).then(({data} : {data: any}) => {
        if(data?.result?.idcs){
            azs.tableData = data.result.idcs; 
            instanceForm.az = azs.tableData[0].idcId; // 初始化机房选择为第一项
        }                  
    }).catch(({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
        ElMessage.error(message);
        azs.tableData=[];
        instanceForm.az = ''
    })
    .finally(() => {
        store.requestObject();
        //store.requestUser();
        (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName);
    });
};
requestAzData()

/**
 * 监听机房信息，每次改变，重新获取机型信息
 */
watch(() => instanceForm.az, (): void => {
    if(azs?.tableData?.length) {
        requestDeviceTypeData()
        // 将走马灯指示器（原type为submit）按钮type设为button，防止点击提交页面
        timer = setTimeout(() => {
            let carouselButton:any = document.getElementsByClassName('el-carousel__button');
            for(var i = 0; i < carouselButton.length; i++) {
                carouselButton[i].type = 'button'
            }
            changeModel(0);
        }, 200);    
    }
});

// 机型选择未售罄的默认选项
const deviceTypesListNumber: Ref<number> = ref<number>(0);
// 机型未售罄的默认选项
const deviceTypeNumber: Ref<number> = ref<number>(0);
// 机型轮播图ref   
let carouselRef: Ref<any> = ref<FormInstance>();
/**
 * 寻找可用机型
 * @param data 
 */
const findDeviceTypes =(data:any) :void => {   
    for(let i = 0; i < Object.keys(data).length; i++) {
        for(let j = 0; j < data[Object.keys(data)[i]].length; j++) {
            if(data[Object.keys(data)[i]][j].availableStock > 0) {
                deviceTypesListNumber.value = i;
                deviceTypeNumber.value = j; 
                return; 
            } else {
                deviceTypesListNumber.value = 0;
                deviceTypeNumber.value = 0;
            }
        }
    }
};

/**
 * 查找并返回第一个对象中特定键的值大于0的项及其对应数组和数组的键值。
 * 如果找不到，则返回第一个属性的第一项及其数组和数组的键值。
 *
 * @param obj - 输入的对象，其每个属性值为数组，数组中的每个元素都是对象。
 *              这些对象中有一个特定键的值可能大于0。
 * @param key - 要检查的特定键。这个键的值将被用来确定哪个对象满足条件。
 * @returns 返回一个包含三个元素的数组：
 *           第一个元素是找到的对象或第一个属性的第一项，
 *           第二个元素是对应的数组，
 *           第三个元素是数组的键值。
 *           如果没有找到任何特定键值大于0的对象，将返回第一个数组的第一项及其数组和数组的键值。
 */
const findFirstPositiveByKey = <T extends Record<string, any>>(
    obj: Record<string, T[]>, // 使用 Record 类型来约束 obj 参数，确保它是一个对象，其属性值为数组。
    key: keyof T // keyof T 确保 key 参数必须是 T 类型对象中存在的键。
): [T | undefined, T[], string] => {
    // 遍历对象的每个属性（键值对）
    for (const prop in obj) {
        if (obj.hasOwnProperty(prop)) { // 检查属性是否直接属于对象，而不是通过原型链继承的。
            const array = obj[prop]; // 获取当前属性的值（一个数组）
            // 在当前数组中查找第一个特定键值大于0的对象
            const foundItem = array.find(item => item[key] > 0);
            // 如果找到了符合条件的对象，返回这个对象、对应的数组和数组的键值
            if (foundItem) {
                return [foundItem, array, prop];
            }
        }
    }
    // 如果没有找到任何符合条件的对象，获取并返回第一个数组的第一项、这个数组和数组的键值
    const firstKey = Object.keys(obj)[0]; // 获取对象的第一个键
    const firstArray = obj[firstKey] || []; // 获取第一个键对应的数组，如果不存在则默认为空数组
    return [firstArray[0], firstArray, firstKey]; // 返回第一个数组的第一项、这个数组和数组的键值
};

/**
 * 获取所有实例名称
 */
const getinstanceNameData = (): void => {
    instanceListAPI({
        isAll:'1',
        projectId: route.query.projectId ? route.query.projectId : '',
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            instanceData.repeatInstanceData = data?.result?.instances.map((item:any) => item.instanceName);
        }
 
    }).catch(() => {
        instanceData.repeatInstanceData = [];
    })
};
getinstanceNameData()

/**
 * 判断是否有raid模式
 */
 const raidIdShow: Ref<boolean> = ref<boolean>(false);

/**
 * 请求机型信息接口
*/
const requestDeviceTypeData = (): void => {
    deviceTypes.tableData = [];
    instanceForm.model = '';
    instanceForm.modelType = '';
    deviceTypeAPI({
        idc_id: instanceForm.az,
    }).then(({data} : {data: any}) => {
        if(Object.keys(data?.result)?.length){
            deviceTypes.tableData = data.result;
            const [foundItem, correspondingArray, arrayKey] = findFirstPositiveByKey(data.result, 'availableStock');
            // 初始化机型选择为第一项
            instanceForm.model = arrayKey;
            // 将机型选择分组映入轮播图中的单元
            dividedUnit(correspondingArray);
            // 展示机型轮播图
            deviceCarousel.value = true;
            // 初始化机型为第一项
            // @ts-ignore
            instanceForm.modelType = foundItem?.deviceTypeId;    
            return
        } 
        deviceTypes.tableData = [];
        instanceForm.model = '';
        instanceForm.modelType = '';
    }).finally(() => {
    });
};

interface DataItem {
  [key: string]: any; // 定义一个接口，允许任意字符串键映射到任意类型的值
}

/**
 * 在一个对象的所有数组属性中查找第一个具有指定属性值的对象。
 * 
 * @param data - 一个键值对的对象，其中每个值都是一个对象数组。
 * @param searchKey - 要在数组的对象中查找的属性名称。
 * @param searchValue - 要匹配的属性值。
 * @returns 如果找到，则返回第一个匹配的对象；否则返回undefined。
 */
const findObjectInArrayGroups = <T extends DataItem>(
  data: Record<string , Array<T>>,
  searchKey: keyof T,
  searchValue: any
): T | undefined => {
  // 使用Object.values遍历对象的每个数组，然后使用find查找符合条件的对象
  // 使用flatMap将找到的对象（如果有的话）放入一个新数组，然后过滤掉所有的undefined值
  // 最后，取第一个找到的对象（如果有的话）
  const foundObject = Object.values(data).flatMap(group => group.find(item => item[searchKey] === searchValue)).find(result => result !== undefined);
  return foundObject
};

/**
 * 监听机型信息，每次改变，重新获取镜像信息
 */
 watch(() => instanceForm.modelType, (): void => {   
    if (instanceForm.modelType) {
        // @ts-ignore
        const foundObject = findObjectInArrayGroups(deviceTypes.tableData, 'deviceTypeId', instanceForm.modelType)
        // @ts-ignore
        if(foundObject?.isNeedRaid === 'need_raid') {
            raidIdShow.value = true
        }else {
            raidIdShow.value = false
        }
        requestOsTypeData();
        requestIsDeviceStockEnough();
    }else {
        osTypes.tableData = [];
        instanceForm.imageType = '';
        instanceForm.image[instanceForm.imageType] = '';
        raids.tableData = [];
        datas.tableData = []
        instanceForm.raid = '';
        systemDiskPartition.tableData = [];
        systemDiskPartition.tableDataChange = [];
    }
});

/**
 * 镜像类型div高度
 */
const imageHeight: Ref<number> = ref<number>(80);

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
    imageHeight.value = 80
    changeArrow.value = !changeArrow.value

}

/**
 * 请求镜像信息接口
*/
const requestOsTypeData = (): void => {
    osTypeAPI({
        idcId: instanceForm.az,
        deviceTypeId: instanceForm.modelType
    }).then(({data} : {data: any}) => {
        if(Object.keys(data?.result)?.length) {
            osTypes.tableData = data?.result;
            instanceForm.imageType = Object.keys(osTypes.tableData)[0]; // 初始化镜像为第一项
            let osTypeFirst = osTypes.tableData[Object.keys(osTypes.tableData)[0]]; // 选出镜像第一项
            instanceForm.image[instanceForm.imageType] = osTypeFirst[0].imageId; // 初始化镜像类型为第一项
            guideModes.tableData = osTypeFirst[0].bootMode.split(',') // 初始化引导模式
            instanceForm.bootMode = guideModes.tableData[0] //选出引导模式第一项
            return
        }
        osTypes.tableData = [];
        instanceForm.imageType = '';
        instanceForm.image[instanceForm.imageType] = ''
    }).finally(() => {
        requestRaidData();
        requestDatasData();
        if(deviceTypeNumber.value) {
            carouselRef.value.setActiveItem(Math.trunc(deviceTypeNumber.value/4)) //机型跳转到已选项对应位置
        }       
    });
}; 

/**
 * 寻找镜像的引导模式
 */
const findImageForBootMode = (): void => {
    let item: any = []
    for(item of Object.values(osTypes.tableData)) {
        item.find((ele: any) => {
            if(ele.imageId === instanceForm.image[instanceForm.imageType]) {
                guideModes.tableData = ele.bootMode.split(',');
                instanceForm.bootMode = guideModes.tableData[0]
            }
        })
    }
}

/**
 * 监听镜像信息，每次改变，重新获取系统盘信息
 */
watch(() => instanceForm.image[instanceForm.imageType], (): void => {
    if(instanceForm.modelType && instanceForm.image[instanceForm.imageType]) {
        requestSystemDiskPartitionData();
        findImageForBootMode()
        
    }else {
        systemDiskPartition.tableData = [];
        systemDiskPartition.tableDataChange = [];
    }
});

/**
 * 请求系统卷信息接口
*/
const requestRaidData = (): void => {
    raids.tableData = [];
    raidAPI({
        deviceTypeId: instanceForm.modelType,
        volumeType: 'system'
    }).then(({data} : {data: any}) => {
        if(Object.keys(data?.result)?.length) {
            raids.tableData = data.result[0].raids;
            instanceForm.raid = raids.tableData[0].raidId; // 初始化系统盘为第一项
            instanceForm.systemChoose = {
                'volumeId': data.result[0].volumeId,
                'raidId': raids.tableData[0].raidId,
                'raidCan': raids.tableData[0].raidCan,
                'volumeType': data.result[0].volumeType
            }
            return
        }
        raids.tableData = [];
        instanceForm.raid =''
    }).finally(() => {
        if(instanceForm.modelType && instanceForm.image[instanceForm.imageType]) {
            requestSystemDiskPartitionData()
        } else {
            systemDiskPartition.tableData = [];
            systemDiskPartition.tableDataChange = [];
        }
    });
};

/**
 * 系统卷改变
 */
const systemRollChange = (val: string):void => {
    (instanceForm.systemChoose as any).raidId = val

};

/**
 * 请求数据卷信息接口
*/
const requestDatasData = (): void => {
    datas.tableData = [];
    raidAPI({
        deviceTypeId: instanceForm.modelType,
        volumeType: 'data'
    }).then(({data} : {data: any}) => {
        if(Object.keys(data?.result)?.length) {
            datas.tableData = data.result.map((item:any) => {
                return {
                    ...item,
                    raidModel: 0,
                    raidModelChoose: {
                        'volumeId': item.volumeId,
                        'raidId': item.raids[0].raidId,
                        'raidCan': item.raids[0].raidCan,
                        'volumeType': item.volumeType
                    },
                    raidModels: item.raids.map((obj: any) => {
                        return {
                            'volumeId': item.volumeId,
                            'raidId': obj.raidId,
                            'raidCan': obj.raidCan,
                            'volumeType': item.volumeType
                        }
                        
                    })
                }
            });
            return
        }
        datas.tableData = [];
    }).finally(() => {
       
    });
};

/**
 * raid改变
 */
const raidModelChange = (key: number, row:any ): void => {
    datas.tableData.map((item: any) => {
        if(item.volumeId === row.volumeId) {
            row.raidModelChoose = row.raidModels[key]
        }

    })
}

/**
 * 请求系统盘分区信息接口
*/
const requestSystemDiskPartitionData = (): void => {
    systemDiskPartitionAPI({
        imageId: instanceForm.image[instanceForm.imageType],
        deviceTypeId: instanceForm.modelType
    }).then(({data} : {data: any}) => {
        if(data?.result) {
            systemDiskPartition.tableData = deepCopy(data.result);
            systemDiskPartition.tableDataChange = data.result;
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

// 库存最大数
const avaliableMaxCount: Ref<number> = ref<number>(1)
/**
 * 查看库存是否足够
*/
const requestIsDeviceStockEnough = (): void => {
    isDeviceStockEnoughAPI({
        idcId: instanceForm.az,
        deviceTypeId: instanceForm.modelType,
        count: instanceForm.number
    }).then(({data} : {data: any}) => {
        if (data?.result?.success) {
            createInstanceState.value = false
            avaliableMaxCount.value = data.result.avaliableCount
        }else {
            createInstanceState.value = true
            avaliableMaxCount.value = 1
        }
    }).finally(() => {
    });
};
/**
 * 创建实例接口
*/
const requestCreateInstance = (formEl: FormInstance | undefined): void => {
    let volumeRaids: any = [];
    for(let i = 0; i < datas.tableData.length; i++) {
        volumeRaids.push(datas.tableData[i].raidModelChoose)
    }
    const volumeRaidAmount = volumeRaids.concat([instanceForm.systemChoose])
    let params = {
        count: instanceForm.number,
        deviceTypeId: instanceForm.modelType,
        hostname: instanceForm.hostName,
        idcId: instanceForm.az,
        imageId: instanceForm.image[instanceForm.imageType],
        bootMode: instanceForm.bootMode,
        instanceName: instanceForm.instanceName,
        password: instanceForm.password,
        sshKeyId: instanceForm.keyPassword,
        systemVolumeRaidId: instanceForm.raid,
        VolumeRaids: volumeRaidAmount,
        projectId: route.query.projectId,
        systemPartition: systemDiskPartition.tableData,
        installBmpAgent: instanceForm.montiorAgent,
        username: instanceForm.imageType === 'Windows' ? 'administrator' : 'root'
    }
    createInstanceAPI({
        ...params
    }).then(({data} : {data: any}) => {
        if(data?.result?.instanceIds) {
            ElMessage({
                message: t('personCentre.content.createSuccess'),
                type:'success'
            });
            jumpToList(formEl)
        }
    }).finally(() => {
        isLoading.value = false;
    });
};

document.onkeydown = (event: {keyCode: number;}) => {
    if(createInstanceState.value) return
    event.keyCode === 13 && createInstance(ruleFormRef.value);
};

/**
 * 点击立即创建
 */
const createInstance = async (formEl: FormInstance | undefined): Promise<void> => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        formEl.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                isLoading.value = true;
                requestCreateInstance(formEl);
            }
        });
    });
};
</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/icon';
@import './instanceCreate.scss';
</style>