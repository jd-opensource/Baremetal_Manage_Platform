import {faultRulesExportAPI} from 'api/surveillance/faultRule/request.ts';
import {ExportDataOperate} from 'utils/publicClass.ts';
import {methodsTotal} from 'utils/index.ts';
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
    searchTip: {value: boolean}

    constructor(search: {ruleForm: {}}, searchTip: {value: boolean}) {
        this.search = search;
        this.searchTip = searchTip;
    };

    exportData = (that: never) => {
        new ExportDataOperate(faultRulesExportAPI, {}, 'alert_rule_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        const otherParams = this.searchTip.value ? methodsTotal.toLine(this.search.ruleForm) : {};
        this.filterEmptyInfo.deleteEmtpyData(otherParams);
        new ExportDataOperate(faultRulesExportAPI, {
            ...otherParams,
        }, 'alert_rule_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate,
    faultRulesExportAPI
};
