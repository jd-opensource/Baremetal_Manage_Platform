import {consolelogsExportAPI} from 'api/device/request.ts';
import {ExportDataOperate, RouterOperate} from 'utils/publicClass.ts';

class ExportData {
    query = new RouterOperate().router.currentRoute.value.query

    exportData = (that: never) => {
        new ExportDataOperate(consolelogsExportAPI, {sn: this.query.sn}, 'alarm_log_list').exportData(that);
    }
};

// 实例化-导出数据操作
const exportDataOperate = new ExportDataOperate(consolelogsExportAPI);

export {
    ExportData,
    exportDataOperate
};

