
import Custom from './custom/customList.vue';
import TableData from './table/tableData.vue';
import ExportData from './export/exportData.vue';
import Pagination from './pagination/pagination.vue';
import SystemReset from './resetSystem/systemReset.vue';
import LockOperate from './lockOperate/lockOperate.vue';
import UpDownDelete from './upDownDelete/upDownDelete.vue';
import DeviceImport from './deviceImport/deviceImport.vue';
import PasswordReset from './passwordReset/passwordReset.vue';
import InstanceResets from './instanceReset/instanceReset.vue';
import RecoveryRemove from './recoveryRemove/recoveryRemove.vue';
import ResetSystemConfirm from './resetSystem/resetSystemConfirm.vue';
import OpenCloseRestart from './openCloseRestart/openCloseRestart.vue';
import BatchPasswordResetOperate from './batchOperate/passwordReset/batchPasswordResetOperate.vue'
import BatchOpenCloseRestartOperate from './batchOperate/openCloseRestart/batchOpenCloseRestartOperate.vue';
import BatchRecoveryInstanceOperate from './batchOperate/recoveryInstance/batchRecoveryInstanceOperate.vue';
import BatchEditInstanceNameOperate from './batchOperate/editInstanceName/batchEditInstanceNameOperate.vue';

const pluginComp = {
    Custom,
    TableData,
    ExportData,
    Pagination,
    SystemReset,
    LockOperate,
    UpDownDelete,
    DeviceImport,
    PasswordReset,
    InstanceResets,
    RecoveryRemove,
    ResetSystemConfirm,
    OpenCloseRestart,
    BatchPasswordResetOperate,
    BatchOpenCloseRestartOperate,
    BatchRecoveryInstanceOperate,
    BatchEditInstanceNameOperate
};

export default pluginComp;
