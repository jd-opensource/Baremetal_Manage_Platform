import custom from './custom/custom';
import resetOperate from './table/reset';
import SearchOperate from './search/search';
import DeviceList from './table/deviceList';
import refreshOperate from './table/refresh';
import FilterOperate from './table/tableFilter';
import tableScrollOperate from './table/tableScroll';
import upDownDelete from './upDownDelete/upDownDelete';
import deviceImport from './deviceImport/deviceImport';
import lockDeblocking from './lockOperate/lockOperate';
import resetInstance from './instanceReset/instanceReset';
import removeRecovery from './recoveryRemove/recoveryRemove';
import passwordReset from './passwordReset/passwordReset';
import openCloseRestart from './openCloseRestart/openCloseRestart';
import batchPasswordReset from './batchOperate/passwordReset/operate';
import batchOpenCloseRestart from './batchOperate/openCloseRestart/operate';
import batchRecoveryInstance from './batchOperate/recoveryInstance/operate';
import batchEditInstanceName from './batchOperate/editInstanceName/operate';
import resetSystemConfirmOperate from './resetSystem/resetSystemConfirm';
import resetSystemOperate from './resetSystem/systemReset';
import {paginationOperate} from 'utils/publicClass.ts';
// import {paginationOperate, customPagination} from './pagination/pagination';
import {ExportFilterData, ExportDataOperate, devicesListExportAPI} from './export/export';

export {
    custom,
    DeviceList,
    FilterOperate,
    SearchOperate,
    ExportFilterData,
    paginationOperate,
    ExportDataOperate,
    devicesListExportAPI,
    resetOperate,
    upDownDelete,
    deviceImport,
    passwordReset,
    resetInstance,
    refreshOperate,
    removeRecovery,
    lockDeblocking,
    openCloseRestart,
    // customPagination,
    tableScrollOperate,
    resetSystemOperate,
    batchPasswordReset,
    batchOpenCloseRestart,
    batchRecoveryInstance,
    batchEditInstanceName,
    resetSystemConfirmOperate
};
