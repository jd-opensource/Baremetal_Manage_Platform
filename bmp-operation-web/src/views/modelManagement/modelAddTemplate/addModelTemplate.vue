<template>
    <div class="operate-management-detail model-operate">
        <!-- 头部信息 -->
        <detail-header
            :name="'no'"
            :title="'addModel'"
            :opt-flag="false"
            @header="router.push($defInfo.routerPath('modelList'))"
        />
        <main class="operate-management-detail-info custom-radio-button">
            <!-- 主体头部内容 -->
            <div class="info-header set-bottom">
                <span class="info-line"></span>
                <!-- 常规参数 -->
                <span class="info-title">
                    {{$t('modelForm.header.params.text1')}}
                </span>
            </div>
            <p class="template-tip">
                {{$t('modelForm.tooltip.tip', ['【' + $filter.emptyFilter(formRulesOperate.ruleForm.templateName) + '】'])}}：
            </p>
            <ui-form
                label-position="left"
                :class="['model-operate-rule-form', setDiffClass('set-right', 'default-right')]"
                :label-width="setDiffClass('140px', '110px')"
                :model="formRulesOperate.ruleForm"
                :rules="formRulesOperate.rules"
                @rule-form-ref="formRulesOperate.getFormRef"
            >
                <!-- 机房名称 -->
                <ui-form-item
                    prop="machineRoomName"
                    required
                    :label="$t('modelForm.label.computerRoomName')"
                    :class="[
                        dataOperate.reactiveData.computerRoomRadioBtn.length >= 4 ? 'set-label' : '',
                        setDiffClass('', 'model-operate-rule-form-computer-room')
                    ]"
                >
                    <ui-radio-group v-model="formRulesOperate.ruleForm.machineRoomName">
                        <!-- 机房名称 -->
                        <ui-tooltip
                            placement="bottom"
                            v-for="(item, index) in dataOperate.reactiveData.computerRoomRadioBtn"
                            v-model="item.showTooltip"
                            :disabled="!item.showTooltip"
                        >
                            <template #content>
                                {{item.title}}
                            </template>
                            <div
                                class="radio-button-text-ellipsis"
                                @mouseenter="hasShowTooltip($event, item, item.title, .8)"
                            >
                                <ui-radio-button
                                    :key="index"
                                    :label="item.title"
                                />
                            </div>
                        </ui-tooltip>
                    </ui-radio-group>
                </ui-form-item>
                <!-- 机型类型 -->
                <ui-form-item
                    prop="modelType"
                    required
                    :label="$t('modelForm.label.modelType')"
                >
                    <ui-radio-group v-model="formRulesOperate.ruleForm.modelType">
                        <!-- 机型类型 -->
                        <ui-radio-button
                            v-for="(item, index) in reactiveDataOperate.reactiveData.radioBtn"
                            :key="index"
                            :label="item"
                        />
                    </ui-radio-group>
                </ui-form-item>
                <!-- 机型名称 -->
                <ui-form-item
                    prop="name"
                    :class="[
                        'model-name',
                        locationItem.getLocationItem !== 'zh_CN' && formRulesEvent.isLongTip.value ? 'set-name-bottom' : '',
                        regExpCheck.nameFlag.value ? 'set-empty' : ''
                    ]"
                    :label="$t('modelForm.label.modelName')"
                >
                    <!-- 输入项 -->
                    <ui-input
                        style="width: 507px"
                        v-model="formRulesOperate.ruleForm.name"
                        :placeholder="$t('modelForm.placeholder.modelName')"
                        :disabled="regExpCheck.nameLoading.value"
                        :validate-event="false"
                        :has-suffix="true"
                        @change="formRulesEvent.nameChange"
                        @blur="formRulesEvent.nameBlur"
                    >
                        <template #suffix>
                            <ui-icon
                                class="is-loading"
                                v-if="regExpCheck.nameLoading.value"
                            >
                                <Loading />
                            </ui-icon> 
                        </template>
                    </ui-input>
                </ui-form-item>
                <!-- 体系架构 -->
                <ui-form-item
                    prop="architecture"
                    style="marginBottom: 17px"
                    :label="$t('modelForm.label.architecture')"
                >
                    <!-- 选择项 -->
                    <ui-select
                        style="width: 507px"
                        v-model="formRulesOperate.ruleForm.architecture"
                        :no-data-text="$t('modelForm.noData.tip')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="(item, index) in dataOperate.reactiveData.architecture"
                            :key="index"
                            :label="item"
                            :value="item"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 引导模式-仅添加机型支持 -->
                <ui-form-item
                    prop="bootMode"
                    required
                    class="boot-mode-count"
                    :label="$t('modelForm.label.bootMode')"
                    :show-message="true"
                >
                    <div class="checkbox-text-ellipsis flex-checkbox">
                        <div
                            style="marginLeft: 1px"
                            v-for="(item, index) in dataOperate.reactiveData.bootModeData"
                        >
                            <ui-checkbox
                                v-model="item.select"
                                :key="index"
                                :label="item.label"
                                :disabled="true"
                            >
                                {{item.label}}
                            </ui-checkbox>
                        </div>
                    </div>
                </ui-form-item>
                <!-- 机型规格 -->
                <div class="device-type-count">
                    <ui-form-item
                        prop="deviceType"
                        :class="[
                            'specifications',
                            regExpCheck.deviceTypeFlag.value ? 'set-empty1' : 'set-default'
                        ]"
                        :label="$t('modelForm.label.modelSpecifications')"
                    >
                        <!-- 输入项 -->
                        <ui-input
                            style="width: 507px"
                            v-model="formRulesOperate.ruleForm.deviceType"
                            :placeholder="$t('modelForm.placeholder.modelSpecifications')"
                            :disabled="regExpCheck.deviceTypeLoading.value"
                            :validate-event="false"
                            :has-suffix="true"
                            @change="formRulesEvent.deviceTypeChange"
                        >
                            <template #suffix>
                                <ui-icon
                                    class="is-loading"
                                    v-if="regExpCheck.deviceTypeLoading.value"
                                >
                                    <Loading />
                                </ui-icon> 
                                <el-tooltip
                                    placement="bottom"
                                    v-else
                                >
                                    <template #content>
                                        <div v-html="$t('modelForm.tooltip.title')"></div>
                                    </template>
                                    <img
                                        alt=""
                                        class="set-from-img"
                                        :src="($defInfo.imagePath('toolTip') as unknown as string)"
                                    />
                                </el-tooltip>
                            </template>
                        </ui-input>
                    </ui-form-item>
                </div>
                <!-- 描述 -->
                <ui-form-item
                    prop="description"
                    :class="['model-operate-rule-form-desc', setDiffClass('set-desc', '')]"
                    :label="$t('modelForm.label.desc')"
                >
                    <!-- 描述输入框 -->
                    <ui-input
                        type="textarea"
                        maxlength="256"
                        show-word-limit
                        v-model="formRulesOperate.ruleForm.description"
                        :placeholder="$t('modelForm.placeholder.desc')"
                    >
                    </ui-input>
                </ui-form-item>
                <div class="info-header set-bottom">
                    <span class="info-line"></span>
                    <!-- 配置参数 -->
                    <span class="info-title">
                        {{$t('modelForm.header.params.text2')}}
                    </span>
                </div>
                <!-- CPU -->
                <ui-form-item
                    prop="modelCPU"
                    required
                    class="model-operate-ruleForm-model-cpu"
                    :label="$t('modelForm.label.cpu')"
                >
                    <!-- CPU数据-预置规格-其它规格 -->
                    <ui-radio-group v-model="formRulesOperate.ruleForm.modelCPU">
                        <ui-radio-button
                            v-for="(item, index) in reactiveDataOperate.reactiveData.modelCPUBtn"
                            :key="index"
                            :label="item"
                        />
                    </ui-radio-group>
                </ui-form-item>
                <!-- CPU数据选项为预置规格时显示 -->
                <!-- 'hide-label', -->
                <ui-form-item
                    prop="cpuInfo"
                    id="cpuInfo"
                    v-if="ModelStaticData.modelSpecData.includes(formRulesOperate.ruleForm.modelCPU)"
                    :label="$t('modelForm.label.cpuParams')"
                    :class="[
                        formRulesOperate.ruleForm.cpuInfo ? '' : regExpCheck.cpuInfoFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <!-- 选择项 -->
                    <ui-select
                        v-model="formRulesOperate.ruleForm.cpuInfo"
                        style="width: 507px"
                        :no-data-text="$t('modelForm.noData.tip')"
                        :placeholder="$t('modelForm.placeholder.select')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in formRulesOperate.ruleForm.cpuData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.label"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 其它规格时显示 -->
                <div
                    style="marginBottom: 16px"
                    v-else
                >
                    <!-- 处理器厂商 -->
                    <ui-form-item
                        prop="cpuManufacturer"
                        id="cpuManufacturer"
                        :label="$t('modelForm.label.processorManufacturer')"
                        :class="[
                            'select-width',
                            formRulesOperate.ruleForm.cpuManufacturer ? '' : regExpCheck.cpuManufacturerFlag.value ? 'set-empty' : ''
                        ]"
                    >
                        <!-- 请选择 -->
                        <ui-select
                            v-model="formRulesOperate.ruleForm.cpuManufacturer"
                            :placeholder="$t('modelForm.placeholder.processorManufacturer')"
                        >
                            <ui-option
                                style="fontSize: 12px"
                                v-for="item in ModelStaticData.processorManufacturerData"
                                :key="item.value"
                                :label="item.label"
                                :value="item.label"
                            />
                        </ui-select>
                    </ui-form-item>
                    <!-- 型号 -->
                    <ui-form-item
                        prop="cpuModel"
                        id="cpuModel"
                        :class="['select-width', regExpCheck.cpuModelFlag.value ? 'set-empty' : '']"
                        :label="$t('modelForm.label.modelChoose')"
                    >
                        <!-- 型号输入项 -->
                        <ui-input
                            v-model="formRulesOperate.ruleForm.cpuModel"
                            :placeholder="$t('modelForm.placeholder.modelChoose')"
                        >
                        </ui-input>
                    </ui-form-item>
                    <!-- 物理核数 -->
                    <ui-form-item
                        prop="cpuCores"
                        id="cpuCores"
                        required
                        :class="[formRulesOperate.ruleForm.cpuCores === 1000 ? 'disabled-input-number' : 'change-input-number']"
                        :label="$t('modelForm.label.numberPhysicalCores')"
                    >
                        <!-- 物理核数输入项 -->
                        <ui-input-number
                            v-model="formRulesOperate.ruleForm.cpuCores"
                            :min="1"
                            :max="1000"
                            :precision="0"
                        />
                    </ui-form-item>
                    <!-- 主频GHz -->
                    <ui-form-item
                        prop="cpuFrequency"
                        id="cpuFrequency"
                        :class="['select-width', regExpCheck.cpuGHzFlag.value ? 'set-empty' : '']"
                        :label="$t('modelForm.label.dominantFrequency')"
                    >
                        <!-- 主频输入项 -->
                        <ui-input
                            v-model="formRulesOperate.ruleForm.cpuFrequency"
                            maxlength="64"
                            :placeholder="$t('modelForm.placeholder.dominantFrequency')"
                        >
                        </ui-input>
                    </ui-form-item>
                    <!-- cpu-数量 -->
                    <ui-form-item
                        prop="cpuAmount"
                        required
                        id="cpuAmount"
                        :class="[formRulesOperate.ruleForm.cpuAmount === 1000 ? 'disabled-input-number' : 'change-input-number']"
                        :label="$t('modelForm.label.numberOfRoutes')"
                    >
                        <!-- 数量输入项 -->
                        <ui-input-number
                            v-model="formRulesOperate.ruleForm.cpuAmount"
                            :min="1"
                            :max="1000"
                            :precision="0"
                        />
                    </ui-form-item>
                </div>
                <!-- 内存 -->
                <ui-form-item
                    prop="modelStorage"
                    required
                    class="model-operate-ruleForm-model-storage"
                    :label="$t('modelForm.label.storage')"
                >
                    <!-- 内存-预置规格-其它规格 -->
                    <ui-radio-group v-model="formRulesOperate.ruleForm.modelStorage">
                        <ui-radio-button
                            v-for="(item, index) in reactiveDataOperate.reactiveData.storageData"
                            :key="index"
                            :label="item"
                        />
                    </ui-radio-group>
                </ui-form-item>
                <!-- 'hide-label', -->
                <ui-form-item
                    prop="memInfo"
                    id="memInfo"
                    v-if="ModelStaticData.modelSpecData.includes(formRulesOperate.ruleForm.modelStorage)"
                    :label="$t('modelForm.label.storageParams')"
                    :class="[
                        formRulesOperate.ruleForm.memInfo ? '' : regExpCheck.memInfoFlag.value ? 'set-empty' : ''
                    ]"
                >
                    <!-- 选择项 -->
                    <ui-select
                        style="width: 507px"
                        v-model="formRulesOperate.ruleForm.memInfo"
                        :no-data-text="$t('modelForm.noData.tip')"
                        :placeholder="$t('modelForm.placeholder.select')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in formRulesOperate.ruleForm.memData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.label"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- 其它规格时显示 -->
                <div
                    style="marginBottom: 16px"
                    v-else
                >
                    <!-- 接口 -->
                    <ui-form-item
                        prop="memType"
                        id="memType"
                        :label="$t('modelForm.label.interface')"
                        :class="[
                            'select-width',
                            formRulesOperate.ruleForm.memType ? '' : regExpCheck.memTypeFlag.value ? 'set-empty' : ''
                        ]"
                    >
                        <!-- 选择 -->
                        <ui-select
                            v-model="formRulesOperate.ruleForm.memType"
                            :placeholder="$t('modelForm.placeholder.interface')"
                        >
                            <ui-option
                                style="fontSize: 12px"
                                v-for="item in ModelStaticData.interfaceData"
                                :key="item.value"
                                :label="item.label"
                                :value="item.label"
                            />
                        </ui-select>
                    </ui-form-item>
                    <!-- 主频(MHz) -->
                    <ui-form-item
                        prop="memFrequency"
                        id="memFrequency"
                        :class="[
                            'select-width',
                            formRulesOperate.ruleForm.memFrequency ? '' : regExpCheck.memFrequencyFlag.value ? 'set-empty' : ''
                        ]"
                        :label="$t('modelForm.label.dominantFrequencyMHz')"
                    >
                        <!-- 选择 -->
                        <ui-select
                            v-model="formRulesOperate.ruleForm.memFrequency"
                            :placeholder="$t('modelForm.placeholder.dominantFrequencyMHz')"
                        >
                            <ui-option
                                style="fontSize: 12px"
                                v-for="item in ModelStaticData.dominantFrequencyMHzData"
                                :key="item.value"
                                :label="item.label"
                                :value="item.label"
                            />
                        </ui-select>
                    </ui-form-item>
                    <!-- 容量(GB) -->
                    <ui-form-item
                        prop="memSize"
                        id="memSize"
                        :class="[
                            'select-width',
                            formRulesOperate.ruleForm.memSize ? '' : regExpCheck.memSizeFlag.value ? 'set-empty' : ''
                        ]"
                        :label="$t('modelForm.label.capacity')"
                    >
                        <!-- 选择 -->
                        <ui-select
                            v-model="formRulesOperate.ruleForm.memSize"
                            :placeholder="$t('modelForm.placeholder.capacity')"
                        >
                            <ui-option
                                style="fontSize: 12px"
                                v-for="item in ModelStaticData.capacityData"
                                :key="item.value"
                                :label="item.label"
                                :value="item.label"
                            />
                        </ui-select>
                    </ui-form-item>
                    <!-- 数量 -->
                    <ui-form-item
                        prop="memAmount"
                        id="memAmount"
                        required
                        :class="[formRulesOperate.ruleForm.memAmount === 1000 ? 'disabled-input-number' : 'change-input-number']"
                        :label="$t('modelForm.label.modelNumber')"
                    >
                        <!-- 数量输入项 -->
                        <ui-input-number
                            v-model="formRulesOperate.ruleForm.memAmount"
                            :min="1"
                            :max="1000"
                            :precision="0"
                        />
                    </ui-form-item>
                </div>
                <!-- 网卡速度(GE) -->
                <ui-form-item
                    required
                    prop="nicRate"
                    :class="[
                        'select-width',
                        formRulesOperate.ruleForm.nicRate ? '' : regExpCheck.nicRateFlag.value ? 'set-empty' : ''
                    ]"
                    :label="$t('modelForm.label.networkSpeed')"
                >
                    <!-- 选择 -->
                    <ui-select
                        v-model="formRulesOperate.ruleForm.nicRate"
                        :placeholder="$t('modelForm.placeholder.networkSpeed')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in ModelStaticData.networkSpeedData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.label"
                        />
                    </ui-select>
                </ui-form-item>
                <ui-form-item
                    required
                    prop="nicAmount"
                    :class="[formRulesOperate.ruleForm.memAmount === 2 ? 'disabled-input-number' : 'change-input-number']"
                    :label="$t('modelForm.label.networkNumber')"
                >
                    <!-- 网口数量 -->
                    <ui-input-number
                        v-model="formRulesOperate.ruleForm.nicAmount"
                        :min="1"
                        :max="2"
                        :precision="0"
                        @input="formRulesOperate.nicAmountIpt"
                    />
                </ui-form-item>
                <!-- 网络设置 -->
                <ui-form-item
                    prop="interfaceMode"
                    required
                    :class="[
                        'select-width',
                        formRulesOperate.ruleForm.interfaceMode ? '' : regExpCheck.interfaceModeFlag.value ? 'set-empty' : ''
                    ]"
                    :label="$t('modelForm.label.networkSettings')"
                >
                    <!-- 选择 -->
                    <ui-select
                        v-model="formRulesOperate.ruleForm.interfaceMode"
                        :placeholder="$t('modelForm.placeholder.networkSettings')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in formRulesOperate.ruleForm.nicAmount > 1 ? ModelStaticData.networkSetUpData1 : ModelStaticData.networkSetUpData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- GPU(品牌) -->
                <ui-form-item
                    prop="gpuManufacturer"
                    :class="['select-width', setDiffClass('set-english-width', 'set-width')]"
                    :label="$t('modelForm.label.gpuBrand')"
                >
                    <!-- 选择 -->
                    <ui-select
                        v-model="formRulesOperate.ruleForm.gpuManufacturer"
                        clearable
                        :placeholder="$t('modelForm.placeholder.gpuBrand')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in ModelStaticData.gpuBrandData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.label"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- GPU(型号) -->
                <ui-form-item
                    prop="gpuModel"
                    :class="['select-width', setDiffClass('set-english-width', 'set-width')]"
                    :label="$t('modelForm.label.gpuModel')"
                >
                    <!-- 选择 -->
                    <ui-select
                        v-model="formRulesOperate.ruleForm.gpuModel"
                        clearable
                        :placeholder="$t('modelForm.placeholder.gpuModel')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in ModelStaticData.gpuModelData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.label"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- GPU数量(个) -->
                <ui-form-item
                    prop="gpuAmount"
                    :class="[
                        formRulesOperate.ruleForm.gpuAmount === 1000 || !formRulesOperate.ruleForm.gpuModel || !formRulesOperate.ruleForm.gpuManufacturer
                            ? 'disabled-input-number'
                            : 'change-input-number',
                        setDiffClass('set-english-width', 'set-width')
                    ]"
                    :label="$t('modelForm.label.gpuNumber')"
                >
                    <ui-input-number
                        v-model="formRulesOperate.ruleForm.gpuAmount"
                        :min="formRulesOperate.amountMinValue()"
                        :max="formRulesOperate.ruleForm.gpuModel && formRulesOperate.ruleForm.gpuManufacturer ? 1000 : 0"
                        :precision="0"
                    />
                    <!-- GPU数量 -->
                </ui-form-item>
                <ui-form-item
                    prop="height"
                    :class="[
                        'select-width',
                        formRulesOperate.ruleForm.height ? '' : regExpCheck.heightFlag.value ? 'set-empty' : ''
                    ]"
                    :label="$t('modelForm.label.heightU')"
                >
                    <!-- 选择 -->
                    <ui-select
                        v-model="formRulesOperate.ruleForm.height"
                        :placeholder="$t('modelForm.placeholder.heightU')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in ModelStaticData.heightUData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.label"
                        />
                    </ui-select>
                </ui-form-item>
                <ui-form-item
                    prop="arrayCard"
                    class="radio-style1"
                    required
                    :label="$t('modelForm.label.card')"
                >
                    <ui-radio v-model="formRulesOperate.ruleForm.arrayCard" label="1">
                        {{$t('modelForm.config.configured')}}
                    </ui-radio>
                    <ui-radio v-model="formRulesOperate.ruleForm.arrayCard" label="2">
                        {{$t('modelForm.config.notConfigured')}}
                    </ui-radio>
                </ui-form-item>
                <ui-form-item
                    prop="volumeManagerTableData"
                    id="volumeManager"
                    v-if="formRulesOperate.ruleForm.arrayCard === '1'"
                    :label="$t('modelForm.label.volumeManage.title')"
                    :show-message="false"
                >
                    <ui-config-provider :locale="uiLocale.locale">
                        <ui-table
                            class="volume-manager-table"
                            style="width: 100%"
                            :data="formRulesOperate.ruleForm.volumeManagerTableData"
                            border
                            :class="tableClass(formRulesOperate.ruleForm.volumeManagerTableData, false)"
                            v-loading="false"
                        >
                            <ui-table-column
                                align="center"
                                fixed
                                :label="$t('modelForm.label.volumeManage.label.name')"
                                :min-width="setDiffClass('260', '170')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(scope.row.volumeName)}}
                                        </span>
                                        <ui-form v-else>
                                            <ui-form-item
                                                prop=""
                                                class="ipt-height"
                                                :label="''"
                                            >
                                            <ui-input
                                                placeholder="请输入"
                                                v-model="scope.row.volumeName"
                                                :class="[
                                                    ['success', void 0].includes(scope.row[`volumeNameFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                ]"
                                                @blur="formRulesOperate.volumeNameBlur(scope.row, scope.row.id)"
                                            />
                                            <p
                                                class="error-tip"
                                                v-if="['error', ''].includes(scope.row[`volumeNameFlag${scope.row.id}`])"
                                            >
                                                {{formRulesOperate.volumeNameTip.value}}
                                            </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                :min-width="setDiffClass('180', '140')"
                                :label="$t('modelForm.label.volumeManage.label.type')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(scope.row.volumeType)}}
                                        </span>
                                        <ui-form v-else>

                                            <ui-form-item
                                                prop=""
                                                class="ipt-height"
                                                :label="''"
                                            >
                                                <ui-select
                                                    v-model="scope.row.volumeType"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`volumeTypeFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.volumeTypeChange(scope.row, scope.row.id)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in dataOperate.reactiveData.volumeTypeData"
                                                        :key="index"
                                                        :label="item"
                                                        :value="item"
                                                        
                                                    />
                                                </ui-select>
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`volumeTypeFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.volumeTypeTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                :min-width="setDiffClass('170', '140')"
                                :label="$t('modelForm.label.volumeManage.label.raidCan')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(scope.row.raidCan)}}
                                        </span>
                                        <ui-form v-else>
                                            <ui-form-item
                                                prop=""
                                                :label="''"
                                                :class="[
                                                   'ipt-height'
                                                ]"
                                            >
                                                <ui-select
                                                    v-model="scope.row.raidCan"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`raidCanFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.raidCanChange(scope.row, scope.row.id)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in dataOperate.reactiveData.raidConfig"
                                                        :key="index"
                                                        :label="item"
                                                        :value="item"
                                                    />
                                                </ui-select>
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`raidCanFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.raidCanTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                min-width="200"
                                :label="$t('modelForm.label.volumeManage.label.raid')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <p v-if="scope.row.editFlag">
                                            <span v-if="Array.isArray(scope.row.sysRaid)">
                                                {{scope.row.sysRaid.join(',')}}
                                                <!-- {{formRulesOperate.setRaidInfo(scope.row.sysRaid, dataOperate.reactiveData.sysRaid)}} -->
                                            </span>
                                            <span v-else>
                                                {{$filter.emptyFilter(scope.row.sysRaid)}}
                                            </span>
                                        </p>
                                        <ui-form v-else>
                                            <ui-form-item
                                                prop=""
                                                class="multiple-choice ipt-height"
                                                :label="''"
                                            >
                                                <ui-select
                                                    v-if="scope.row.raidCan === 'NO RAID'"
                                                    v-model="scope.row.sysRaid"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`sysRaidFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.sysRaidChange(scope.row, scope.row.id, dataOperate.reactiveData.noSysRaid)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in dataOperate.reactiveData.noSysRaid"
                                                        :key="index"
                                                        :label="item.label"
                                                        :value="item.label"
                                                    />
                                                </ui-select>
                                                <ui-select
                                                    v-if="['', 'RAID'].includes(scope.row.raidCan)"
                                                    v-model="scope.row.sysRaid"
                                                    collapse-tags
                                                    multiple
                                                    popper-class="select-seleted-icon"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`sysRaidFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.sysRaidChange(scope.row, scope.row.id, dataOperate.reactiveData.sysRaid)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in scope.row.raidCan === 'RAID' ? dataOperate.reactiveData.sysRaid : []"
                                                        :key="index"
                                                        :label="item.label"
                                                        :value="item.label"
                                                    />
                                                </ui-select>
                                                <ui-select
                                                    v-if="scope.row.raidCan === $t('modelForm.raid')"
                                                    v-model="scope.row.sysRaid"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`sysRaidFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.sysRaidChange(scope.row, scope.row.id, dataOperate.reactiveData.sysDiskRaid0)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in dataOperate.reactiveData.sysDiskRaid0"
                                                        :key="index"
                                                        :label="item.label"
                                                        :value="item.label"
                                                    />
                                                </ui-select>
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`sysRaidFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.sysRaidTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                :min-width="setDiffClass('180', '140')"
                                :label="$t('modelForm.label.volumeManage.label.disk')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(formRulesOperate.setDiskInterfaceType(scope.row.diskType))}}
                                        </span>
                                        <ui-form v-else>
                                            <ui-form-item
                                                prop=""
                                                class="ipt-height"
                                                :label="''"
                                            >
                                                <ui-select
                                                    v-model="scope.row.diskType"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`diskTypeFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.diskTypeChange(scope.row, scope.row.id)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in dataOperate.reactiveData.diskTypeData"
                                                        :key="index"
                                                        :label="item.label"
                                                        :value="item.value"
                                                    />
                                                </ui-select>
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`diskTypeFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.diskTypeTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                :min-width="setDiffClass('180', '140')"
                                :label="$t('modelForm.label.volumeManage.label.interfaceType')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(formRulesOperate.setDiskInterfaceType(scope.row.interfaceType))}}
                                        </span>
                                        <ui-form v-else>
                                            <ui-form-item
                                                prop=""
                                                class="ipt-height"
                                                :label="''"
                                            >
                                                <ui-select
                                                    v-model="scope.row.interfaceType"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`interfaceTypeFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    @change="formRulesOperate.interfaceTypeChange(scope.row, scope.row.id)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in scope.row.raidCan === 'NO RAID' ? dataOperate.reactiveData.noRaidInterfaceData : dataOperate.reactiveData.interfaceData"
                                                        :key="index"
                                                        :label="item.label"
                                                        :value="item.value"
                                                    />
                                                </ui-select>
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`interfaceTypeFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.interfaceTypeTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                :min-width="setDiffClass('245', '230')"
                                :label="$t('modelForm.label.volumeManage.label.minNum')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(scope.row.minSizeNum)}}{{ scope.row.minSize }}
                                        </span>
                                        <ui-form
                                            class="min-size-count"
                                            v-else
                                        >
                                            <ui-form-item
                                                prop=""
                                                :label="''"
                                                :class="[
                                                    'ipt-height',
                                                    setDiffClass('set-data-sys-style', 'sys-single-capacity input-width')
                                                ]"
                                            >
                                                <ui-input
                                                    placeholder="请输入"
                                                    :class="[
                                                        ['success', void 0].includes(scope.row[`minSizeNumFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                                    ]"
                                                    v-model="scope.row.minSizeNum"
                                                    @blur="formRulesOperate.minSizeNumBlur(scope.row, scope.row.id)"
                                                >
                                                </ui-input>
                                                <ui-select
                                                    v-model="scope.row.minSize"
                                                    @change="formRulesOperate.minSizeChange(scope.row)"
                                                >
                                                    <ui-option
                                                        style="fontSize: 12px"
                                                        v-for="(item, index) in ['GB', 'TB']"
                                                        :key="index"
                                                        :label="item"
                                                        :value="item"
                                                    />
                                                </ui-select>
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`minSizeNumFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.minSizeNumTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                :min-width="setDiffClass('170', '150')"
                                :label="$t('modelForm.label.volumeManage.label.amount')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="scope.row.editFlag">
                                            {{$filter.emptyFilter(scope.row.amount)}}
                                        </span>
                                        <ui-form v-else>
                                            <!-- scope.row.amountDisabled ? 'disabled-input-number' : 'change-input-number', -->

                                            <ui-form-item
                                                prop=""
                                                :label="''"
                                                :class="[
                                                   'ipt-height',
                                                ]"
                                            >
                                                <ui-input-number
                                                    v-model="scope.row.amount"
                                                    :min="formRulesOperate.setAmountMinSize(scope.row, scope.row.id)"
                                                    :max="formRulesOperate.setAmountMaxSize(scope.row, scope.row.id)"
                                                    :precision="0"
                                                />
                                                <p
                                                    class="error-tip"
                                                    v-if="['error', ''].includes(scope.row[`amountFlag${scope.row.id}`])"
                                                >
                                                    {{formRulesOperate.amountTip.value}}
                                                </p>
                                            </ui-form-item>
                                        </ui-form>
                                    </div>
                                </template>
                            </ui-table-column>
                            <ui-table-column
                                align="center"
                                min-width="100"
                                fixed="right"
                                :label="$t('modelForm.label.volumeManage.label.operate')"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <p :class="['operate-btn', scope.row.editFlag ? 'edit-btn' : 'save-btn']">
                                        <span v-if="!scope.row.editFlag" @click="formRulesOperate.saveClick(scope.row, scope.row.id)">
                                            {{$t('modelForm.label.volumeManage.btn.save')}}
                                        </span>
                                        <span  v-else @click="formRulesOperate.editClick(scope.row)">
                                            {{$t('modelForm.label.volumeManage.btn.edit')}}
                                        </span>
                                        <span @click="formRulesOperate.deleteClick(scope.$index)">
                                            {{$t('modelForm.label.volumeManage.btn.delete')}}
                                        </span>
                                    </p>
                                </template>
                            </ui-table-column>
                        </ui-table>
                    </ui-config-provider>
                </ui-form-item>
                <NoConfiguration v-else :form-rules-operate="formRulesOperate" :data-operate="dataOperate"></NoConfiguration>

                <div v-if="formRulesOperate.ruleForm.arrayCard === '1'">
                    <p
                        v-if="formRulesOperate.tableHasError.value"
                        :class="setDiffClass('en-table-error-tip', 'table-error-tip')"
                    >
                        {{formRulesOperate.tableErrorTip.value}}
                    </p>
                    <p
                        :class="[
                            setDiffClass('add-volume-en', 'add-volume'),
                            formRulesOperate.setAddBtn(),
                            formRulesOperate.ruleForm.volumeManagerTableData?.length && formRulesOperate.hasAddFlag.value ? 'set-top' : ''
                        ]"
                        @click="formRulesOperate.addVolumeClick"
                    >
                        + {{$t('modelForm.label.volumeManage.btn.add')}}
                    </p>
                </div>
                <div v-else>
                    <p
                        v-if="formRulesOperate.tableHasError2.value"
                        :class="setDiffClass('en-table-error-tip', 'table-error-tip')"
                    >
                        {{formRulesOperate.tableErrorTip2.value}}
                    </p>
                    <p
                        :class="[
                            setDiffClass('add-volume-en', 'add-volume'),
                            formRulesOperate.setAddBtn2(),
                            formRulesOperate.ruleForm.noConfigurationData?.length && formRulesOperate.hasAddFlag2.value ? 'set-top' : ''
                        ]"
                        @click="formRulesOperate.addVolumeClick2"
                    >
                        + {{$t('modelForm.label.volumeManage.btn.add')}}
                    </p>
                </div>
            </ui-form>
        </main>
        <div class="footer-btn">
            <ui-button type="primary" :loading="formRulesOperate.isLoading.value" @click="formRulesOperate.sure">确定</ui-button>
            <ui-button class="custom-dialog-footer-cancel-btn" @click="formRulesOperate.cancel">取消</ui-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import {setDiffClass, hasShowTooltip, tableClass} from 'utils/index.ts';
// import {RuleFormRefType} from '@utils/publicType';
import {locationItem, uiLocale, RouterOperate} from 'utils/publicClass.ts';
// import {RuleFormType, RulesType} from './typeManagement'
import reactiveData from './utils/reactiveData';
import FormRulesEvent from './utils/formRulesEvent';
import formRules from './utils/newFormRules';
import {regExpCheck} from './utils/regExpCheck';
// import store from 'store/index.ts';
import DataOperate from './utils/dataOperate';
import ModelStaticData from 'staticData/model/index.ts';
import NoConfiguration from './noConfiguration.vue'

// 路由
const router = new RouterOperate().router;

const formRulesOperate = formRules();

const reactiveDataOperate = reactiveData();

const formRulesEvent = new FormRulesEvent(formRulesOperate as any);
const dataOperate = new DataOperate(formRulesOperate.ruleForm);
regExpCheck.nameFlag.value = false;
regExpCheck.deviceTypeFlag.value = false;

// const storageOperate = new StorageOperate(props);
// 

</script>
<style lang="scss">
@import 'assets/css/detail.scss';
@import './index';
</style>
