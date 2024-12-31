<template>
    <!-- 上架下架删除操作 -->
<!-- 
    :set-up-title="upDownDel.setUpTitle.value"
        :set-title="upDownDel.setTitle.value"
        :next-loading="upDownDel.nextLoading" -->
        <!-- :has-set-up="title=== 'up'" -->

    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="[
            'currency up-down-frame',
            title === 'up' ? 'set-dia-log' : 'default-dia-log'
        ]"
        :is-loading="upDownDel.isLoading"
        :header-title="setInfo.titleComputed(title)"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="upDownDel.cancelClick"
        @determine-click="upDownDel.determineClick"
        @pre-click="upDownDel.preClick"
        @next-click="upDownDel.nextClick"
    >
        <div class="currency-count">
            <!-- <ui-steps
                align-center
                v-if="title === 'up'"
                :finish-status="upDownDel.active.value === 1 ? 'finish' : 'success'"
                :active="upDownDel.active.value"
            >
                <ui-step
                    :class="[
                        upDownDel.active.value > 1 ? 'first-step1' : 'first-step2'
                    ]"
                    :title="$t('upDownFrame.steps.title1')"
                />
                <ui-step
                    :class="[
                        'default-step',
                        upDownDel.active.value > 1 ? 'last-step1' : 'last-step2'
                    ]"
                    :title="$t('upDownFrame.steps.title2')"
                />
            </ui-steps> -->
            <p
                class="currency-count-title"
                v-if="title !== 'up'"
            >
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- countTitle设备提示 -->
                <span class="count-title-tip">
                    {{setInfo.countTitle(title)}}
                </span>
            </p>
            <!-- headerTitle提示 -->
            <!-- upDownDel.active.value > 1 ||  -->
            <p class="header-title">
                <span>{{setInfo.headerTitle(title)}}</span>
            </p>
            <!-- <ui-form>
            </ui-form> -->
            <div
                class="up-down-frame-count-table currency-count-table"
                v-if="!deleteStatus"
            >
                <!-- <ui-form
                    label-position="left"
                    label-width="100px"
                    :model="upFormOpt.ruleForm"
                    :rules="upFormOpt.rules"
                    v-if="title === 'up' && upDownDel.active.value === 1"
                    @submit.native.prevent
                    @rule-form-ref="upFormOpt.getFormRef"
                >
                    <ui-form-item
                        prop="modelName"
                        :class="[
                            upFormOpt.ruleForm.modelName ? '' : upDownDel.errorFlag.value ? 'set-empty' : ''
                        ]"
                        :label="$t('upDownFrame.select.model')"
                    >
                        <ui-select
                            style="width: 180px"
                            v-model="upFormOpt.ruleForm.modelName"
                            :no-data-text="$t('upDownFrame.noData.tip')"
                            :placeholder="$t('upDownFrame.select.placeholder')"
                            @change="upFormOpt.modelChange"
                        >
                            <ui-option
                                style="fontSize: 12px"
                                v-for="item in upFormOpt.reactiveArr.modelData"
                                :key="item.value"
                                :label="item.text"
                                :value="item.filterParams"
                            />
                            
                        </ui-select>
                        <p class="add-new-model">
                            <span @click="upFormOpt.jumpPage">
                                {{$t('upDownFrame.operate.addNewModel')}}
                            </span>
                        </p>
                    </ui-form-item>
                </ui-form> -->
                <!-- 表格数据 -->
                <!-- v-if="upDownDel.active.value > 1 || title !== 'up'"  -->
                <up-down-frame-table :table-data="upDownDel.upDownFrameTable"/>
                <!-- <ui-config-provider :locale="uiLocale.locale">
                    <ui-table
                        border
                        v-if="upDownDel.active.value === 1 && title === 'up'"
                        v-loading="upFormOpt.isTableLoading.value"
                        :data="upFormOpt.reactiveArr.tableData"
                        :class="tableClass(upFormOpt.reactiveArr.tableData, upFormOpt.isTableLoading.value)"
                    >
                        <ui-table-column
                            align="center"
                            prop="volumeName"
                            min-width="130"
                            fixed
                            :label="$t('upDownFrame.table.label.volumeName')"
                        >
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="volumeType"
                            min-width="100"
                            :label="$t('upDownFrame.table.label.volumeType')"
                        >
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            v-if="upFormOpt.arrayCard.value !== 'no_need_raid'"
                            :min-width="setDiffClass('160', '120')"
                            :label="$t('upDownFrame.table.label.raidCan')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span>
                                    {{$filter.emptyFilter(upFormOpt.setRaidCan(scope.row.raidCan))}}
                                </span>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            v-if="upFormOpt.arrayCard.value !== 'no_need_raid'"
                            :label="$t('upDownFrame.table.label.raid')"
                            :min-width="upFormOpt.reactiveArr.tableData.length ? '180' : '110'"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <p>
                                    {{$filter.emptyFilter(upFormOpt.setRaid(scope.row.raids))}}
                                </p>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="diskType"
                            min-width="100"
                            :label="$t('upDownFrame.table.label.diskType')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <p>
                                    {{$filter.emptyFilter(upFormOpt.setDiskInterfaceType(scope.row.diskType))}}
                                </p>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="interfaceType"
                            min-width="100"
                            :label="$t('upDownFrame.table.label.interfaceType')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <p>
                                    {{$filter.emptyFilter(upFormOpt.setDiskInterfaceType(scope.row.interfaceType))}}
                                </p>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            :min-width="setDiffClass('210', '120')"
                            :label="$t('upDownFrame.table.label.minNum')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span>{{$filter.emptyFilter(scope.row.volumeSize)}}{{scope.row.volumeUnit}}</span>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="volumeAmount"
                            :min-width="setDiffClass('180', '120')"
                            :label="$t('upDownFrame.table.label.amount')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span>{{$filter.emptyFilter(scope.row.volumeAmount)}}</span>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            fixed="right"
                            :min-width="upFormOpt.reactiveArr.tableData.length ? '380' : '240'"
                            :label="$t('upDownFrame.table.label.associatedDisk')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <div>
                                    <ui-form v-if="['RAID0-stripping', 'NO RAID'].indexOf(scope.row.raidCan) == -1">
                                        <ui-form-item :class="['multiple-choice', scope.row.disksDataFlag ? 'error-tip' : 'def-tip']">
                                            <ui-select
                                                v-model="scope.row.disks"
                                                collapse-tags
                                                multiple
                                                :placeholder="scope.row.disksDataFlag ? $t('upDownFrame.placeholder.tip2') : $t('associatedModel.content.table.diskPlaceholder')"
                                                @change="upFormOpt.disksDataChange(scope.row, scope.$index)"
                                            >
                                                <ui-option
                                                    style="fontSize: 12px"
                                                    v-for="(item, index) in upFormOpt.reactiveArr.disksData[scope.$index]"
                                                    :key="index"
                                                    :label="upFormOpt.setDiskLabel(item)"
                                                    :value="item.diskId"
                                                    :disabled="item.select"
                                                />
                                            </ui-select>
                                        </ui-form-item>
                                    </ui-form>
                                    <ui-form v-else>
                                        <ui-form-item :class="['def-select', scope.row.disksDataFlag ? 'error-tip' : 'def-tip']">
                                            <ui-select
                                                v-model="scope.row.disks"
                                                :placeholder="scope.row.disksDataFlag ? $t('upDownFrame.placeholder.tip2') : $t('associatedModel.content.table.diskPlaceholder')"
                                                @change="upFormOpt.disksDataChange(scope.row, scope.$index)"
                                            >
                                                <ui-option
                                                    style="fontSize: 12px"
                                                    v-for="(item, index) in upFormOpt.reactiveArr.disksData[scope.$index]"
                                                    :key="index"
                                                    :label="upFormOpt.setDiskLabel(item)"
                                                    :value="item.diskId"
                                                    :disabled="item.select"
                                                />
                                            </ui-select>
                                        </ui-form-item>
                                    </ui-form>
                                </div>
                            </template>
                        </ui-table-column>
                    </ui-table>
                </ui-config-provider> -->
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import { tableClass, methodsTotal, setDiffClass} from 'utils/index.ts';
import SetInfo from './utils/setInfo';
import UpDownDel from './utils/upDownDeleteOpt';
import {uiLocale} from 'utils/publicClass.ts';
import UpFormOpt from './utils/upForm';
import store from 'store/index.ts';
// import {RuleFormRefType} from '@utils/publicType';
// import {ElForm} from 'element-plus';
// props 类
interface PropsType {
    deviceId: string;
    diaLog: boolean;
    title: string;
    deleteStatus?: boolean;
};

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

const setInfo = new SetInfo();

const upFormOpt = new UpFormOpt(props);

const upDownDel = new UpDownDel(props, emitValue, upFormOpt);


const useListenMsg = methodsTotal.listenMsg((info: {type: string, content: string}) => {
    const {type, content} = info;
    if (['model-add-success', ' model-edit-success', 'model-delete-success', 'model-add-template-success'].indexOf(type) > -1) {
        if (content === 'success') {
            store.ossDataInfo().getModelListRes().then((data: {text: string; filterParams: string; value: number}[]) => {
                upFormOpt.reactiveArr.modelData = data;
            })
        }
    }
})


onUnmounted(useListenMsg);

</script>
<style lang="scss">
@import './upDownFrame.scss';

</style>
