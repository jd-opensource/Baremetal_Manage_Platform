import {describeRulesExportAPI} from 'api/inBandMonitoring/allAlarmRules/request.ts';
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
    filters;

    constructor(search: {ruleForm: {}}, searchTip: {value: boolean}, filters: any) {
        this.filters = filters;
        this.search = search;
        this.searchTip = searchTip;
    };

    exportData = (that: never) => {
        new ExportDataOperate(describeRulesExportAPI, {}, 'alert_rules_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        const otherParams = this.searchTip.value ? {...this.search.ruleForm,  ...this.filters.reactiveArr.filterParams} : {};
        this.filterEmptyInfo.deleteEmtpyData(otherParams);
        new ExportDataOperate(describeRulesExportAPI, {
            ...otherParams,
        }, 'alert_rules_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate,
    describeRulesExportAPI
};
