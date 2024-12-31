import {auditLogsExportAPI} from 'api/device/request.ts';
import {ExportDataOperate} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class ExportFilterData {
    // 是否导出数据
    // hasExportData: Ref<boolean> = ref<boolean>(true);
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    // list: any;
    ruleFormOperate: {
        ruleForm: {}
    };
    filter: {
        filterParams: {}
    };
    sn;

    constructor(sn: string, ruleFormOperate: {
        ruleForm: {}
    }, filter: {
        filterParams: {}
    }) {
        this.sn = sn;
        this.ruleFormOperate = ruleFormOperate;
        this.filter = filter;
    };

    exportData = (that: never) => {
        new ExportDataOperate(auditLogsExportAPI, {sn: this.sn}, 'audit_logs_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        this.filterEmptyInfo.deleteEmtpyData({...this.filter.filterParams, ...this.ruleFormOperate.ruleForm});
        new ExportDataOperate(auditLogsExportAPI, {
            sn: this.sn,
            ...this.ruleFormOperate.ruleForm,
            ...this.filter.filterParams
        }, 'audit_logs_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate,
    auditLogsExportAPI
};
