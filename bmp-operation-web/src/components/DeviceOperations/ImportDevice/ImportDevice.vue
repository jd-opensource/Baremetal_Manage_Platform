<template>
    <!-- 导入设备操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'import-device'"
        :src="($defInfo.imagePath('device') as unknown as string)"
        :is-loading="importDeviceOperate.isLoading"
        :header-title="$t('importDevice.header.title')"
        @dia-log-close="importDeviceOperate.cancelClick"
        @determine-click="importDeviceOperate.determineClick"
    >
        <!-- 主体内容-表单操作 -->
        <ui-form
            class="import-device-ruleForm"
            label-position="left"
            :label-width="setDiffClass('120px', '80px')"
            :model="formRulesOperate.ruleForm"
            :rules="formRulesOperate.rules"
            @submit.native.prevent
            @rule-form-ref="formRulesOperate.getFormRef"
        >
            <!-- 机房名称 -->
            <ui-form-item
                prop="idcName"
                :class="[
                    setDiffClass('set-computer-room', 'import-device-ruleForm-computer-room'),
                    formRulesOperate.ruleForm.idcName ? '' : heightStatus.hasClick.value ? 'set-empty' : ''
                ]"
                :label="$t('importDevice.label.computerRoomName')"
            >
                <!-- 请选择机房 -->
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.idcName"
                    :no-data-text="$t('importDevice.noData.tip')"
                    :placeholder="$t('importDevice.placeholder.computerRoomName')"
                    :loading="idcList.idcListLoading.value"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in idcList.reactiveArr.idcData"
                        :key="item.idcId"
                        :label="item.newIdcName"
                        :value="item.idcId"
                    />
                </ui-select>
            </ui-form-item>
            <!-- 机型名称 -->
            <!-- <ui-form-item
                prop="modelName"
                :class="[
                    setDiffClass('set-model-name', 'import-device-ruleForm-model-name'),
                    formRulesOperate.ruleForm.modelName ? '' : heightStatus.hasClick.value ? 'set-empty' : ''
                ]"
                :label="$t('importDevice.label.modelName')"
            >
                请选择机型
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.modelName"
                    :no-data-text="$t('importDevice.noData.tip')"
                    :placeholder="$t('importDevice.placeholder.modelName')"
                    :loading="modelList.modelNameLoading.value"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in modelList.reactiveArr.modelData"
                        :key="item.deviceTypeId"
                        :label="item.name"
                        :value="item.deviceTypeId"
                    />
                    
                </ui-select>
                <p
                    class="operate-refresh"
                    @click="modelList.resetData"
                >
                    <span :class="[modelList.hasRotate.value ? 'refresh-rotate' : 'refresh-img']"></span>
                </p>
            </ui-form-item> -->
            <!-- 跳转机型列表页面 -->
            <!-- <p :style="locationItem.getLocationItem === 'zh_CN' ? {marginLeft: '110px'} : {marginLeft: '150px'}" :class="heightStatus.hasClick.value ? 'error-add-new-model' : 'add-new-model'">
                <span>
                    {{$t('importDevice.operate.tip1')}}
                </span>
                <span @click="modelList.jumpPage">
                    {{$t('importDevice.operate.addNewModel')}}
                </span>
            </p> -->
            <!-- 设备信息 -->
            <ui-form-item
                prop="deviceInfo"
                :class="setDiffClass('set-device-info', 'import-device-ruleForm-device-info')"
                :show-message="false"
                :label="$t('importDevice.label.deviceInfo')"
            >
                <ui-upload
                    class="upload-demo"
                    action="/operation-web/devices/upload"
                    name="deviceFile"
                    :limit="1"
                    :auto-upload="false"
                    :on-exceed="upLoadOperate.handleExceed"
                    :on-change="upLoadOperate.handleChange"
                    :file-list="formRulesOperate.ruleForm.deviceInfo"
                    :on-error="upLoadOperate.errorChange"
                    :on-progress="upLoadOperate.progressChange"
                    :on-remove="upLoadOperate.removeChange"
                    :on-success="upLoadOperate.uploadSuccess"
                    @uploadEvent="upLoadOperate.uploadEvent"
                >
                    <template #trigger>
                        <!-- 选取文件 -->
                        <ui-button
                            class="import-device-footer-confirm-btn"
                            type="primary"
                        >
                            {{$t('importDevice.btn.selectionFile')}}
                        </ui-button>
                    </template>
                    <template
                        v-if="upLoadOperate.hasError.value"
                        #tip
                    >
                        <!-- 异常提示 -->
                        <div
                            name="tip"
                            class="error-tip"
                        >
                            {{$t('importDevice.errTip.deviceInfo')}}
                        </div>
                    </template>
                </ui-upload>
                <!-- 模板下载 -->
                <span
                    class="template-down-load"
                    @click="downLoadOperate.templateDownLoad"
                >
                    {{$t('importDevice.operate.templateDownLoad')}}
                </span>

                <!-- 上传文件提示信息 -->
                <div class="upload-tip">
                    <p v-html="$t('importDevice.tooltip.first')"></p>
                    <p v-html="$t('importDevice.tooltip.two')"></p>
                    <p v-html="$t('importDevice.tooltip.third')"></p>
                </div>
            </ui-form-item>
        </ui-form>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {locationItem} from 'utils/publicClass.ts';
import {setDiffClass} from 'utils/index.ts';
import downLoad from './utils/down';
import IdcList from './utils/idcList';
import HeightStatus from './utils/utils';
import formRules from './utils/formRules';
// import ModelList from './utils/modelList';
import UpLoadOperate from './utils/upLoad';
import importDevice from './utils/importDevice';

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<{diaLog: boolean;}>(), {
    diaLog: false
});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const formRulesOperate = formRules();

const heightStatus = new HeightStatus(formRulesOperate);

const idcList = new IdcList(formRulesOperate, () => {
    // modelList.modelNameLoading.value = true;
    // modelList.getModelList();
});

// const modelList = new ModelList(idcList, formRulesOperate, heightStatus);

const upLoadOperate = new UpLoadOperate(formRulesOperate);

const downLoadOperate = downLoad();

const importDeviceOperate = importDevice(formRulesOperate, upLoadOperate, heightStatus, emitValue)

</script>
<style lang="scss">
@import './importDevice.scss';

</style>
