import store from 'store/index.ts';
import {methodsTotal} from 'utils/index.ts';
import {ExportDataOperate, CurrentInstance} from 'utils/publicClass.ts';

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
    proxy = new CurrentInstance().proxy;

    constructor(search: {ruleForm: {}}, searchTip: {value: boolean}) {
        this.search = search;
        this.searchTip = searchTip;
    };

    exportData = (that: never) => {
        const api = this.proxy.$hardwareStatusApi.hardwareStatusExportAPI;
        new ExportDataOperate(api, {}, 'device_status_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        const api = this.proxy.$hardwareStatusApi.hardwareStatusExportAPI;
        const otherParams = this.searchTip.value ? methodsTotal.toLine(this.search.ruleForm) : {};
        this.filterEmptyInfo.deleteEmtpyData(otherParams);
        new ExportDataOperate(api, {
            ...otherParams,
        }, 'device_status_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate
};
