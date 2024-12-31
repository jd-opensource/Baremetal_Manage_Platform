<template>
    <!-- 导入镜像操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'import-image'"
        :is-loading="importImage.isLoading"
        :header-title="$t('importImage.header.title')"
        :src="($defInfo.imagePath('imageManage') as unknown as string)"
        @dia-log-close="importImage.cancelClick"
        @determine-click="importImage.determineClick"
    >
        <!-- 主体内容-表单操作 -->
        <ui-form
            class="import-image-ruleForm"
            label-position="left"
            label-width="110px"
            :model="formRulesOperate.ruleForm"
            :rules="rulesVerification.rules"
            @rule-form-ref="formRulesOperate.getFormRef"
        >
            <!-- 镜像名称 -->
            <ui-form-item
                prop="imageName"
                :class="[
                    'import-image-name', rulesVerification.imageNameFlag.value ? 'set-empty' : ''
                ]"
                :label="$t('importImage.content.label.imageName')"
            >
                <ui-input
                    type="text"
                    v-model.trim="formRulesOperate.ruleForm.imageName"
                    :placeholder="$t('importImage.content.placeholder.imageName')"
                    :disabled="rulesVerification.isShowLoading.value"
                />
                <ui-icon
                    style="right: 28px"
                    class="is-loading"
                    v-if="rulesVerification.isShowLoading.value"
                >
                    <Loading />
                </ui-icon>
            </ui-form-item>
            <!-- 体系架构 -->
            <ui-form-item
                prop="architecture"
                :label="$t('importImage.content.label.architecture')"
            >
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.architecture"
                    :no-data-text="$t('importImage.emptyTip.noData')"
                    :placeholder="$t('importImage.content.placeholder.architecture')"
                    @change="importImage.architectureChange"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in importImage.reactiveArr.architecture"
                        :key="item.value"
                        :label="item.label"
                        :value="item.label"
                    />
                </ui-select>
            </ui-form-item>
            <!-- 操作系统平台 -->
            <ui-form-item
                prop="osType"
                :class="[formRulesOperate.ruleForm.osType ? '' : importImage.hasClick.value ? 'set-empty' : '']"
                :label="$t('importImage.content.label.operateSysPlatform')"
            >
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.osType"
                    :no-data-text="$t('importImage.emptyTip.noData')"
                    :placeholder="$t('importImage.content.placeholder.operateSysPlatform')"
                    @change="importImage.osTypeChange"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="(item, index) in importImage.reactiveArr.osType"
                        :key="index"
                        :label="item.label"
                        :value="item.label"
                    />
                </ui-select>
            </ui-form-item>
            <!-- 自定义操作系统平台 -->
            <ui-form-item
                prop="customOperatePlatform"
                v-if="ImageStaticData.otherData.includes(formRulesOperate.ruleForm.osType)"
                :class="[rulesVerification.customOSType.value ? 'set-empty' : '']"
            >
                <ui-input
                    v-model="formRulesOperate.ruleForm.customOperatePlatform"
                    maxlength="128"
                    :placeholder="$t('importImage.content.placeholder.customOperatePlatform')"
                />
            </ui-form-item>
            <!-- 自定义版本 -->
            <ui-form-item
                prop="customVersion"
                v-if="ImageStaticData.otherData.includes(formRulesOperate.ruleForm.osType)"
                :class="[rulesVerification.customOSVersion.value ? 'set-empty' : '']"
            >
                <ui-input
                    v-model="formRulesOperate.ruleForm.customVersion"
                    maxlength="128"
                    :placeholder="$t('importImage.content.placeholder.customVersion')"
                />
            </ui-form-item>
            <!-- 操作系统版本 -->
            <ui-form-item
                prop="version"
                v-else
                :class="[
                    formRulesOperate.ruleForm.version ? 'set-default' : importImage.hasClick.value ? 'set-empty' : 'set-default'
                ]"
                :label="$t('importImage.content.label.operateSysVersion')"
            >
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.version"
                    :no-data-text="$t('importImage.emptyTip.noData')"
                    :placeholder="$t('importImage.content.placeholder.operateSysVersion')"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in importImage.reactiveArr.osVersion"
                        :key="item.value"
                        :label="item.label"
                        :value="item.label"
                    />
                </ui-select>
            </ui-form-item>
            <!-- 镜像格式 -->
            <ui-form-item
                prop="imageFormat"
                :class="[
                    formRulesOperate.ruleForm.imageFormat ? 'set-default' : importImage.imageFormatClick.value ? 'set-empty' : 'set-default'
                ]"
                :label="$t('importImage.content.label.imageFormat')"
            >
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.imageFormat"
                    :no-data-text="$t('importImage.emptyTip.noData')"
                    :placeholder="$t('importImage.content.placeholder.imageFormat')"
                    @change="importImage.imageFormatChange"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in importImage.reactiveArr.imageFormatData"
                        :key="item"
                        :label="item"
                        :value="item"
                    />
                </ui-select>
            </ui-form-item>
            <!-- 引导模式 -->
            <ui-form-item
                prop="bootMode"
                :class="[
                    'multiple-choice',
                    importImage.hasBootModeFlag.value ? 'set-empty' : 'set-default'
                ]"
                v-if="formRulesOperate.ruleForm.imageFormat === 'tar'"
                :label="$t('importImage.content.label.bootMode')"
            >
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.bootMode"
                    :no-data-text="$t('importImage.emptyTip.noData')"
                    :multiple="true"
                    :placeholder="$t('importImage.content.placeholder.bootMode')"
                    @change="importImage.bootModeChange"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in importImage.reactiveArr.bootModeData"
                        :key="item"
                        :label="item"
                        :value="item"
                    />
                </ui-select>
            </ui-form-item>
            <ui-form-item
                prop="bootMode"
                :class="[
                    'multiple-choice',
                    importImage.hasBootModeFlag.value ? 'set-empty' : 'set-default'
                ]"
                v-else
                :label="$t('importImage.content.label.bootMode')"
            >
            <!-- formRulesOperate.ruleForm.imageFormat === 'tar' -->
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.bootMode"
                    :no-data-text="$t('importImage.emptyTip.noData')"
                    :placeholder="$t('importImage.content.placeholder.bootMode')"
                    @change="importImage.bootModeChange"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in importImage.reactiveArr.bootModeData"
                        :key="item"
                        :label="item"
                        :value="item"
                    />
                </ui-select>
            </ui-form-item>
            <!-- 系统盘分区模板 -->
            <ui-form-item
                prop="systemPartition"
                style="marginTop: -5px; marginBottom: 10px"
                :class="setDiffClass('sys-partition', '')"
                :label="$t('importImage.content.label.systemPartitionTemplate')"
            >
                <span v-html="store.sysPartitionInfo.systemPartitionHtml(formRulesOperate.ruleForm.systemPartition)"></span>
                <!-- <span v-if="ImageStaticData.otherSystemData.includes(formRulesOperate.ruleForm.osType)">
                    --
                </span>
                <span v-else v-html="store.sysPartitionInfo.systemPartitionHtml(formRulesOperate.ruleForm.systemPartition)"></span> -->
            </ui-form-item>
            <!-- 镜像描述 -->
            <ui-form-item
                prop="description"
                class="import-image-ruleForm-desc"
                :label="$t('importImage.content.label.imageDesc')"
            >
                <ui-input
                    type="textarea"
                    maxlength="256"
                    v-model.trim="formRulesOperate.ruleForm.description"
                    show-word-limit
                    :placeholder="$t('importImage.content.placeholder.imageDesc')"
                >
                </ui-input>
            </ui-form-item>
            <!-- 选择镜像 -->
            <ui-form-item
                prop="imageFile"
                :class="[
                    'import-image-ruleForm-choose-image', chooseImageOperate.chooseImageFlag.value ? 'set-empty-chooseimage' : ''
                ]"
                :show-message="false"
                :label="$t('importImage.content.label.chooseImage')"
            >
                <ui-upload
                    class="upload-demo"
                    action="/operation-web/images/upload"
                    name="imageFile"
                    :limit="1"
                    :auto-upload="false"
                    :on-exceed="chooseImageOperate.handleExceed"
                    :on-change="chooseImageOperate.handleChange"
                    :file-list="formRulesOperate.ruleForm.imageFile"
                    :on-error="chooseImageOperate.errorChange"
                    :on-progress="chooseImageOperate.progressChange"
                    :on-remove="chooseImageOperate.removeChange"
                    :on-success="chooseImageOperate.uploadSuccess"
                    @upload-event="chooseImageOperate.uploadEvent"
                >
                    <template #trigger>
                        <!-- 上传 -->
                        <ui-button
                            type="primary"
                            class="import-image-footer-confirm-btn"
                        >
                            {{$t('importImage.btn.upLoad')}}
                        </ui-button>
                    </template>
                    <!-- 异常提示 -->
                    <template
                        #tip
                        v-if="chooseImageOperate.hasError.value"
                    >
                        <div
                            name="tip"
                            class="error-tip"
                        >
                            {{$t('importImage.errTip.numberLimit')}}
                        </div>
                    </template>
                </ui-upload>
            </ui-form-item>
        </ui-form>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts'; // 设置不同class
import store from 'store/index.ts';
import ImageStaticData from 'staticData/image/index.ts';
import RulesVerification from './utils/rulesVerification';
import formRules from './utils/formRules';
import ChooseImage from './utils/chooseImage';
import ImportImageOperate from './utils/importImage';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
};
/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const rulesVerification = new RulesVerification((status: boolean) => {
    importImage.hasBootModeFlag.value = status;
});

const formRulesOperate = formRules();
const chooseImageOperate = new ChooseImage(formRulesOperate, (res: any) => {
    importImage.reactiveArr.fileParams = res;
});

const importImage = new ImportImageOperate(formRulesOperate, rulesVerification, chooseImageOperate, emitValue);

</script>
<style lang="scss">
@import './importImage.scss';
</style>
