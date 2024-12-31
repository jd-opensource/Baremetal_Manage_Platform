/**
 * @file
 * @author
*/

import {ExportDataOperate, CurrentInstance} from 'utils/publicClass.ts';

class ExportData {

    proxy = new CurrentInstance().proxy;

    exportData = (that: {isLoading: {value: boolean}}) => {
        new ExportDataOperate(this.proxy.$idcApi.idcListExportAPI).exportData(that);
    }
};

export {
    ExportData,
    ExportDataOperate
};
