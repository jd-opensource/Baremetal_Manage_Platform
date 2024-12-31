<template>
    <div class="reset-system-information">
        <ui-form
            class="rule-form"
            label-position="left"
            label-width="105px"
            :model="formRulesOperate.ruleForm"
            :rules="formRulesOperate.rules"
            @rule-form-ref="formRulesOperate.getFormRef"
        >
            <!-- 镜像类型 -->
            <ui-form-item
                prop="imageType"
                class="custom-image"
                style="marginBottom: 25px"
                :label="$t('resetSystem.header.imageType')"
            >
                <ui-dropdown
                    popper-class="custom-card-select"
                    trigger="click"
                    v-for="(item, index) in formRulesOperate.reactiveArr.imageTypekeyData"
                    :key="index"
                    @visible-change="formRulesOperate.visibleChange"
                    
                >
                    <template #dropdownTitle>
                        <div
                            class="reset-system-information-count"
                            :style="formRulesOperate.currIndex.value === index ? {borderColor: '#108ee9'} : {borderColor: '#979797'}"
                            @click="formRulesOperate.visibleClick(item, index)"
                        >
                            <img
                                class="img"
                                :src="item.img"
                            />
                            <p class="info">
                                <span>{{item.name}}</span>
                                <span
                                    :style="[
                                        item.initName === $t('resetSystem.placeholder.select')
                                        ? {color: '#999'}
                                        : {}
                                    ]"
                                >
                                    {{item.initName}}
                                </span>
                            </p>
                            <ui-icon
                                :class="[
                                    'custom-icon',
                                    formRulesOperate.selectStatus.value && formRulesOperate.currIndex.value === index
                                    ? 'rotate180'
                                    : 'rotate0'
                                ]"
                                color="#333"
                                :size="14"
                            >
                                <arrow-down />
                            </ui-icon>
                            <div
                                class="triangle"
                                v-if="formRulesOperate.currIndex.value === index"
                            >
                            </div>
                            <ui-icon
                                color="#fff"
                                class="custom-yes-icon"
                                v-if="formRulesOperate.currIndex.value === index"
                                :size="9"
                            >
                                <Check />
                            </ui-icon>
                        </div>
                    </template>
                    <ui-dropdown-item
                        v-for="(ite, ind) in item.childrenData"
                        :key="ind"
                    >
                        <span
                            :class="[
                                'custom-drop-down-item',
                                formRulesOperate.ruleForm.imageType === ite.imageName
                                ? 'drop-down-item-active'
                                : ''
                            ]"
                            @click="formRulesOperate.dropdownItemChange(ite, index)"
                        >
                            {{ite.imageName}}
                        </span>
                    </ui-dropdown-item>
                </ui-dropdown>
                <div
                    :class="[
                        'active-bottom',
                        formRulesOperate.showRemaining.value ? 'icon-rotate180' : 'icon-rotate0'
                    ]"
                    v-if="formRulesOperate.reactiveArr.imageTypekeyData.length > 4"
                    @click="formRulesOperate.activeClick"
                >
                    <ui-icon><ArrowDownBold /></ui-icon>
                </div>
            </ui-form-item>
            <!-- <ui-form-item
                prop="imageType"
                v-else
                :label="$t('resetSystem.header.imageType')"
            >
                <span>{{$filter.emptyFilter('')}}</span>
            </ui-form-item> -->
            <!-- 系统卷 -->
            <ui-form-item
                prop="systemDisk"
                class="custom-system-disk"
                style="marginBottom: 20px"
                :label="$t('resetSystem.header.systemVolume')"
            >
            <!--  -->
                <ui-select
                    class="input-select"
                    v-model="formRulesOperate.ruleForm.systemVolumeRaidId"
                >
                    <ui-option
                        style="fontSize: 12px;"
                        v-for="(item, index) in formRulesOperate.reactiveArr.systemDiskData"

                        :key="index"
                        :label="item.volumeDetail"
                        :value="item.raidId"
                    />
                </ui-select>
                <!-- <div
                    style="display: flex; flex-wrap: wrap;"
                    v-if="formRulesOperate.reactiveArr.systemDiskData?.length"
                > -->
                    <!-- <div
                        class="system-disk"
                        v-for="(item, index) in formRulesOperate.reactiveArr.systemDiskData"
                        :key="index"
                        @click="formRulesOperate.systemDiskClick(item, index)"
                    >
                        <div
                            class="system-disk-header"
                            :style="formRulesOperate.systemCurrIndex.value === index ? {borderColor: '#108ee9', background: '#108ee9', color: '#fff'} : {borderColor: '#979797'}"
                        >
                            {{item.raidName}}
                        </div>
                        <div
                            class="system-disk-count"
                            :style="formRulesOperate.systemCurrIndex.value === index ? {borderColor: '#108ee9'} : {borderColor: '#979797'}"
                        >
                            <p
                                class="volume-detail"
                                :style="!item.diskType ? {marginTop: '8px'}: {}"
                            >
                                {{item.volumeDetail}}
                            </p>
                            <p v-if="item.diskType">{{item.diskType}}</p>
                        </div>
                        <div
                            class="triangle"
                            v-if="formRulesOperate.systemCurrIndex.value === index"
                        >
                        </div>
                        <ui-icon
                            color="#fff"
                            class="custom-yes-icon"
                            v-if="formRulesOperate.systemCurrIndex.value === index"
                            :size="9"
                        >
                            <Check />
                        </ui-icon>
                    </div> -->
                <!-- </div> -->
                <!-- <span v-else>{{$filter.emptyFilter('')}}</span> -->
                
            </ui-form-item>
            <!-- 系统卷分区 -->
            <ui-form-item
                prop="systemPartition"
                class="custom-label"
                style="marginBottom: 20px"
                :label="$t('resetSystem.header.systemVolumePartition')"
            >
                <div
                    class="system-partition"
                    v-for="(item, index) in formRulesOperate.reactiveArr.systemPartitionData"
                    :key="index"
                >
                    <div class="system-partition-header">
                        {{$filter.withClon($filter.emptyFilter(item.point))}}
                    </div>
                    <div class="system-partition-line"></div>
                    <div class="system-partition-num">
                        <span v-if="(typeof item.size === 'string')">
                            {{item.size}}
                        </span>
                        <span v-else>
                            {{item.size / 1024}}
                        </span>
                        <span v-if="(typeof item.size !== 'string')">
                            GiB
                        </span>
                    </div>
                    <div class="system-partition-unit">{{item.format}}</div>
                    <div class="system-partition-triangle"></div>
                    <div class="system-partition-triangle2"></div>
                </div>
            </ui-form-item>
            <!-- 引导模式 -->
            <ui-form-item
                prop="bootMode"
                class="boot-mode"
                :label="$t('resetSystem.header.bootMode')"
            >
                <div
                    v-for="(item, index) in formRulesOperate.reactiveArr.bootModeData"
                    :class="[
                        formRulesOperate.bootModeCurrIndex.value === index
                        ? 'custom-radio-btn-active'
                        : 'custom-radio-btn'
                    ]"
                    :key="index"
                    @click="formRulesOperate.radioChange(item, index)"
                >
                    {{item}}
                </div>
            </ui-form-item>
            <!-- HostName -->
            <ui-form-item
                prop="hostName"
                label="HostName"
                :class="['host-name', rulesCheck.hostNameFlag.value ? 'set-empty' : '']"
            >
                <ui-input
                    class="custom-ipt"
                    type="text"
                    v-model="formRulesOperate.ruleForm.hostName"
                    :placeholder="$t('resetSystem.placeholder.hostName')"
                />
            </ui-form-item>
            <!-- 用户名 -->
            <ui-form-item
                prop="userName"
                class="user-name"
                :label="$t('resetSystem.header.userName')"
            >
                <span>{{formRulesOperate.ruleForm.userName}}</span>
            </ui-form-item>
            <!-- 设置密码 -->
            <ui-form-item
                prop="userName"
                class="set-password radio-style1"
                :label="$t('resetSystem.header.setPassword')"
            >
                <ui-radio-group
                    v-model="formRulesOperate.ruleForm.setType"
                    @change="formRulesOperate.setPasswordChange"
                >
                    <ui-radio :label="'customPassword'">
                        {{$t('resetSystem.radio.customPassword')}}
                    </ui-radio>
                    <!-- <ui-radio :label="'sshKey'">
                        {{$t('resetSystem.radio.sshKeyLogin')}}
                    </ui-radio> -->
                </ui-radio-group>
            </ui-form-item>
            <!-- 密码 -->
            <ui-form-item
                prop="password"
                class="custom-password"
                v-if="!formRulesOperate.ruleForm.setType.localeCompare('customPassword')"
                :label="$t('resetSystem.header.password')"
                :show-message="false"
            >
                <ui-input
                    maxlength="30"
                    type="password"
                    class="custom-ipt"
                    placeholder=""
                    v-model.trim="formRulesOperate.ruleForm.password"
                />
                <p :class="[rulesCheck.colorStatus.value ? 'error-color-tip' : 'password-tip']">
                    {{rulesCheck.errorTip.value}}
                </p>
            </ui-form-item>
            <!-- 确认密码 -->
            <ui-form-item
                prop="confirmPassword"
                v-if="!formRulesOperate.ruleForm.setType.localeCompare('customPassword')"
                :label="$t('resetSystem.header.confirmPassword')"
                :class="[
                    setDiffClass('custom-label'),
                    rulesCheck.confirmPasswordFlag.value ? 'set-empty' : ''
                ]"
            >
                <ui-input
                    maxlength="30"
                    class="custom-ipt"
                    type="password"
                    v-model.trim="formRulesOperate.ruleForm.confirmPassword"
                    :placeholder="''"
                />
            </ui-form-item>
            <ui-form-item
                props="sshKeyId"
                v-else
            >
                <reset-system-key-table
                    :form-rules-operate="formRulesOperate"
                    :rules-check="rulesCheck"
                />
            </ui-form-item>
            <ui-form-item
                class="install-bmp-agent"
                prop="installBmpAgent"
                :label="$t('resetSystem.header.performanceMonitoring')"
            >
                <ui-checkbox
                    size="small"
                    v-model="formRulesOperate.ruleForm.installBmpAgent"
                >
                    {{$t('resetSystem.header.agent')}}
                </ui-checkbox>
                <ui-tooltip placement="bottom">
                    <template #content>
                        <div>{{$t('resetSystem.tip1')}}</div>
                    </template>
                    <img
                        alt=""
                        class="common-img"
                        :src="($defInfo.imagePath('toolTip') as unknown as string)"
                    />
                </ui-tooltip>
            </ui-form-item>
        </ui-form>
    </div>
</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts';
import {uiLocale} from 'utils/publicClass';

type PropsType = {
    formRulesOperate: any;
    rulesCheck: any;
};
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>