import {describeAlertsExportAPI} from 'api/inBandMonitoring/historyAlarmInfo/request.ts';
import {ExportDataOperate} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class ExportFilterData {
    // 是否导出数据
    hasExportData: Ref<boolean> = ref<boolean>(true);
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    search: {
        ruleForm: {}
    };
    searchTip: {value: boolean};

    constructor(search: {ruleForm: {}}, searchTip: {value: boolean}) {
        this.search = search;
        this.searchTip = searchTip;
    };

    exportData = (that: never) => {
        new ExportDataOperate(describeAlertsExportAPI, {}, 'alert_history_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        const params = {};
        for (let index in this.search.ruleForm) {
            if (!index.includes('alarm')) {
                // @ts-ignore
                params[index] = this.search.ruleForm[index];
            }
        }
        const otherParams = this.searchTip.value ? params : {};
        this.filterEmptyInfo.deleteEmtpyData(otherParams);
        new ExportDataOperate(describeAlertsExportAPI, {
            ...otherParams,
        }, 'alert_history_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate,
    describeAlertsExportAPI
};
