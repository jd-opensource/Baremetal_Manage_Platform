<template>
    <div>
        <p class="batch-tip">
            {{ $t('list.content.batchContent') }}
            <span>{{props.multipleSelection?.length}}</span>
            {{ batchOperationTip[props.batchTip] }}
        </p>
        <el-table 
            border
            ref="tableRef"
            :data="props.multipleSelection" 
            style="width: 100%"
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
            <!-- 带外IP -->
            <el-table-column 
                prop="iloIp" 
                :label="$t('instance.list.outBandIp')" 
            >
                <template v-slot="scope">
                    {{scope.row.iloIp || $filter.emptyFilter()}}
                </template>
            </el-table-column>
            <!-- 内网IPv4 -->
            <el-table-column 
                prop="privateIpv4" 
                :label="$t('instance.list.intranetIpv4')" 
            >
                <template v-slot="scope">
                    {{scope.row.privateIpv4 || $filter.emptyFilter()}}
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup lang="ts">
import {
    hasShowTooltip, // 是否显示提示
} from 'utils/index.ts';
import {
    batchOperationTip, // 批量操作提示
} from 'utils/constants.ts'; 

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    batchTip: string;
    multipleSelection: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    batchTip: '',
    multipleSelection: {}
});
</script>

<style lang="scss" scoped>
.batch-tip {
    font-size: 12px;
    margin-bottom: 20px;
    span {
        color: #108ef9;
    }
}
</style>