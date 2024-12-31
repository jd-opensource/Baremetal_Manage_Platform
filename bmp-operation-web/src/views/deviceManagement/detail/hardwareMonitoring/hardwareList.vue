<template>
    <div>
        <!-- header信息 设备信息-->
        <title-info :title="$t('deviceDetail.hardwareMonitoring.hardwareDeviceStatus')"/>
        <!-- 硬件设备状态 -->
        <device-status
            :surveillance-opt="surveillanceOpt"
            @search-ref="tableS.getSearchRef"
        />
        <title-info :title="$t('deviceDetail.hardwareMonitoring.alarmLogList')"/>
        <div class="search-tip">
            <!-- 操作-刷新&设置&导出 -->
            <refresh-custom-export
                :default-class="'operate-management-content-operate'"
                :has-btn-operate="false"
                @refresh="faultLogTable.refresh"
                @custom="customOperate.customClickOperate"
                @export="exportOpt.exportDataOperate.diaLogClick"
                @operate-ref="tableS.getOperateRef"
            />
        </div>
        <table-list
            :fault-log-table="faultLogTable"
            :table-scroll-operate="tableS"
            :custom-operate="customOperate"
        />
        <pagination
            :fault-log-table="faultLogTable"
            :active-operate="activeOperate"
            :table-detail="surveillanceOpt"
        />
        <export-data/>
        <list-custom :custom-operate="customOperate"/>
    </div>
</template>
<script lang="ts" setup>
import methods from './methods';
import pluginComp from './module';

const [customOpt, exportOpt, PaginationOperate, SurveillanceOpt, FaultLogTable] = methods;

interface PropsType {
    query: any;
    activeOperate: any;
    tableS: any;
    deviceDetail: any;
}

// defineEmits API 用来代替emit
const emitValue = defineEmits(['surveillance-opt']);

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

const [ListCustom, DeviceStatus, ExportData, Pagination, TableList, titleInfo] = pluginComp;

const {CustomOperate, name, deviceFaultLogInfo, deviceFaultLogList} = customOpt.default;

const faultLogTable = new FaultLogTable.default(props.query);
// 实例化-自定义操作
const customOperate = new CustomOperate(name, deviceFaultLogInfo, deviceFaultLogList);

const surveillanceOpt = new SurveillanceOpt.default(props, faultLogTable);

new PaginationOperate.default(faultLogTable);

</script>
