<template>
    <div class="disks-detail">
        <!-- 物理盘视图 -->
        <div class="no-table-img-url">
            <!-- header信息 物理盘视图-->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 物理盘视图 -->
                <span class="info-title">
                    {{$t('deviceDetail.diskDetail.physicalDiskView')}}
                </span>
            </div>
            <ui-config-provider :locale="locale">
                <ui-table
                    border
                    style="width: 100%"
                    v-loading="diskTableOpt.isLoading.value"
                    :max-height="500"
                    :class="[tableClass(diskTableOpt.reactiveArr.disksData, diskTableOpt.isLoading.value)]"
                    :data="diskTableOpt.reactiveArr.disksData"
                >
                    <!-- 磁盘 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        fixed
                        :label="$t('deviceDetail.diskDetail.table.label.disk')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.name)}}
                        </template>
                    </ui-table-column>
                    <!-- 单盘容量 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.table.label.singleDiskCapacity')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.size)}}{{scope.row.sizeUnit}}
                        </template>
                    </ui-table-column>
                    <!-- 磁盘类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.table.label.diskType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <!-- {{scope.row.mediaType === 'notLimited' ? $t('modelForm.unrestricted') : scope.row.mediaType}} -->
                            {{$filter.emptyFilter(diskTableOpt.setDiskInterfaceType(scope.row.mediaType))}}
                        </template>
                    </ui-table-column>
                    <!-- 接口类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.table.label.interfaceType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(diskTableOpt.setDiskInterfaceType(scope.row.pdType))}}
                        </template>
                    </ui-table-column>
                    <!-- 背板号(enclosure) -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.table.label.enclosure')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.enclosure)}}
                        </template>
                    </ui-table-column>
                    <!-- 槽位(slot) -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        fixed="right"
                        :label="$t('deviceDetail.diskDetail.table.label.slot')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <span>
                                <!-- 后端为0不返回这个字段，前端处理 -->
                                {{[void 0, ''].includes(scope.row.slot) ? 0 : scope.row.slot}}
                            </span>
                            <!-- {{$filter.emptyFilter(scope.row.slot)}} -->
                        </template>
                    </ui-table-column>
                </ui-table>
            </ui-config-provider>
        </div>
        <!-- 逻辑盘视图 -->
        <div class="no-table-img-url">
            <!-- header信息 逻辑盘视图-->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 逻辑盘视图 -->
                <span class="info-title">
                    {{$t('deviceDetail.diskDetail.logicalDriveView')}}
                </span>
            </div>
            <ui-config-provider :locale="locale">
                <ui-table
                    border
                    style="width: 100%"
                    v-loading="diskTableOpt.isLoading.value"
                    :max-height="500"
                    :class="[tableClass(diskTableOpt.reactiveArr.tableData, diskTableOpt.isLoading.value)]"
                    :data="diskTableOpt.reactiveArr.tableData"
                >
                    <!-- 盘符名称 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        fixed
                        :label="$t('deviceDetail.diskDetail.table2.label.driveLetterName')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.name)}}
                        </template>
                    </ui-table-column>
                    <!-- 单盘容量 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.table2.label.singleDiskCapacity')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.size)}}{{scope.row.sizeUnit}}
                        </template>
                    </ui-table-column>
                    <!-- 磁盘类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.table2.label.diskType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(diskTableOpt.setDiskInterfaceType(scope.row.mediaType))}}
                        </template>
                    </ui-table-column>
                    <!-- 接口类型 -->
                    <!-- <ui-table-column
                        align="center"
                        min-width="100"
                        fixed="right"
                        :label="$t('deviceDetail.diskDetail.table2.label.interfaceType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(diskTableOpt.setDiskInterfaceType(scope.row.pdType))}}
                        </template>
                    </ui-table-column> -->
                </ui-table>
            </ui-config-provider>
        </div>
        <!-- 机型信息 -->
        <div class="no-table-img-url">
            <!-- header信息 机型信息-->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 机型信息 -->
                <span class="info-title">
                    {{$t('deviceDetail.diskDetail.modelRolInfo')}}
                </span>
                <span v-if="detail?.deviceTypeName">（{{$t('deviceDetail.diskDetail.modelName')}}：{{$filter.emptyFilter(detail?.deviceTypeName)}}）</span>
            </div>
            <ui-tooltip
                placement="bottom"
                :disabled="['in'].includes(detail.manageStatus)"
                :content="$t('deviceDetail.tooltip.associatedModels')"
            >
                <ui-button
                    type="primary"
                    class="btn"
                    :disabled="!['in'].includes(detail.manageStatus)"
                    @click="emitValue('add-models')"
                >
                    {{$t('deviceDetail.diskDetail.btn')}}
                </ui-button>
            </ui-tooltip>
            <ui-config-provider :locale="locale">
                <ui-table
                    border
                    style="width: 100%"
                    v-loading="diskTableOpt.isLoading.value"
                    :max-height="500"
                    :class="[tableClass(diskTableOpt.reactiveArr.volumesData, diskTableOpt.isLoading.value)]"
                    :data="diskTableOpt.reactiveArr.volumesData"
                >
                    <!-- 卷名称 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        fixed
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.volumeName')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeName)}}
                        </template>
                    </ui-table-column>
                    <!-- 卷类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.volumeType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeType)}}
                        </template>
                    </ui-table-column>
                    <!-- RAID模式 -->
                    <ui-table-column
                        align="center"
                        v-if="detail.isNeedRaid !== 'no_need_raid'"
                        :min-width="setDiffClass('180', '100')"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.raid')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <div>
                                <p v-if="Array.isArray(scope.row.raids) && scope.row.raids.length > 0">
                                    <!-- {{scope.row.raids.map(item =).join(',')}} -->
                                    {{$filter.emptyFilter(diskTableOpt.setRaids(scope.row.raids))}}
                                </p>
                                <p v-else>
                                    {{$filter.emptyFilter('')}}
                                </p>
                            </div>
                        </template>
                    </ui-table-column>
                    <!-- 磁盘类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.diskType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(diskTableOpt.setDiskInterfaceType(scope.row.diskType))}}
                        </template>
                    </ui-table-column>
                    <!-- 接口类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.interfaceType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(diskTableOpt.setDiskInterfaceType(scope.row.interfaceType))}}
                        </template>
                    </ui-table-column>
                    <!-- 单盘最小容量 -->
                    <ui-table-column
                        align="center"
                        :min-width="setDiffClass('200', '100')"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.minNum')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeSize)}}{{scope.row.volumeUnit}}
                        </template>
                    </ui-table-column>
                    <!-- 最低数量（块） -->
                    <ui-table-column
                        align="center"
                        :min-width="setDiffClass('180', '100')"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.minimumQuantity')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeAmount)}}
                        </template>
                    </ui-table-column>
                    <!-- 关联磁盘 -->
                    <ui-table-column
                        align="center"
                        min-width="150"
                        fixed="right"
                        :label="$t('deviceDetail.diskDetail.volumeManager.table.label.associatedDisk')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <div>
                                <div v-if="Array.isArray(scope.row.disks) && scope.row.disks.length > 0">
                                    <ui-tooltip
                                        placement="left"
                                        v-model="scope.row[`newDisksValTip${scope.$index}`].showTooltip"
                                        :disabled="!scope.row[`newDisksValTip${scope.$index}`].showTooltip"
                                    >
                                        <template #content>
                                            <div class="set-tooltip-width">
                                                <span>
                                                    {{scope.row.newDisksVal}}
                                                </span>
                                            </div>
                                        </template>
                                        <div
                                            class="more-text-ellipsis"
                                            @mouseenter="hasShowTooltip($event, scope.row[`newDisksValTip${scope.$index}`], scope.row.newDisksVal)"
                                        >
                                            <span
                                                style="marginRight: 22px"
                                                :class="setTextClass(false)"
                                            >
                                                {{scope.row.newDisksVal}}
                                            </span>
                                        </div>
                                    </ui-tooltip>
                                </div>
                                <p v-else>
                                    {{$filter.emptyFilter('')}}
                                </p>
                            </div>
                        </template>
                    </ui-table-column>
                </ui-table>
            </ui-config-provider>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {tableClass, languageSwitch, hasShowTooltip, setTextClass, setDiffClass} from 'utils/index.ts';

interface PropsType {
    diskTableOpt: any;
    detail: {deviceTypeName: string; manageStatus: string; isNeedRaid: string};
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['add-models']);
const locale = languageSwitch();
</script>
<style lang="scss" scoped>
.disks-detail {
    padding: 0 30px;
    margin-top: 5px;
}
</style>