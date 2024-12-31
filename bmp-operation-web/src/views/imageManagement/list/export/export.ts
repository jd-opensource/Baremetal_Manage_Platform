import {imagesExportAPI} from 'api/image/request.ts';
import {ExportDataOperate} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class ExportFilterData {
    // 是否导出数据
    hasExportData: Ref<boolean> = ref<boolean>(true);
    filterEmptyInfo: {
        deleteEmtpyData(arg0?: {}): Function;
    }  = store.filterEmpty;
    // list: any;
    search: {
        searchParams: {}
    };
    filter: {
        filterParams: {}
    };

    constructor(search: {
        searchParams: {}
    }, filter: {
        filterParams: {}
    }) {
        // this.list = list;
        this.search = search;
        this.filter = filter;
    };

    exportData = (that: never) => {
        new ExportDataOperate(imagesExportAPI, {}, 'images_list').exportData(that);
    };

    /**
     * 参数处理
    */
    paramsProcessing = (that: never) => {
        this.filterEmptyInfo.deleteEmtpyData({...this.filter.filterParams, ...this.search.searchParams});
        new ExportDataOperate(imagesExportAPI, {
            ...this.search.searchParams,
            ...this.filter.filterParams
        }, 'images_list').exportData(that);
    }
};

export {
    ExportFilterData,
    ExportDataOperate,
    imagesExportAPI
};
