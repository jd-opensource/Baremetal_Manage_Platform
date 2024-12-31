/**
 * @file
 * @author zhangjingyuan02
*/

import store from 'store/index.ts';
import {ExportDataOperate, CurrentInstance} from 'utils/publicClass.ts';

class ExportFilterData {
    // 是否导出数据
    hasExportData: Ref<boolean> = ref<boolean>(true);
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    // list: any;
    search: {
        searchParams: {},
        radioGroup: {}
    };
    filter: {
        filterParams: {}
    };
    proxy = new CurrentInstance().proxy;

    constructor(search: {
        searchParams: {},
        radioGroup: {}
    }, filter: {
        filterParams: {}
    }) {
        // this.list = list;
        this.search = search;
        this.filter = filter;
    };

    exportData = (that: never) => {
        new ExportDataOperate(this.proxy.$messageApi.messageExportAPI, {}, 'message_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        this.filterEmptyInfo.deleteEmtpyData({...this.filter.filterParams, ...this.search.searchParams, ...this.search.radioGroup});
        new ExportDataOperate(this.proxy.$messageApi.messageExportAPI, {
            ...this.search.searchParams,
            ...this.filter.filterParams,
            ...this.search.radioGroup
        }, 'message_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate
};
