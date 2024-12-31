import pinia from 'lib/pinia/pinia.ts'; // pina
import loginInfoStore from './modules/loginInfo';
import progressInfoStore from './modules/progress';
import customInfoStore from './modules/custom';
import sysPartitionStore from './modules/systemPartition';
import imagesStore from './modules/imagesData';
import filterEmptyStore from './modules/filterEmpty';
import navigationBarStore from './modules/navigationBar';
import ossDataStore from './modules/ossData';
import modeFormStore from './modules/modeForm';
import userPurviewStore from './modules/userPurview';
import idcStore from './modules/idc';
import faultLevelStore from './modules/faultLevel';
import faultSpecStore from './modules/faultSpec';
import deviceDetailErrorTipStore from './modules/deviceDetailErrorTip';

const store = {
    idcInfo: idcStore(pinia),
    loginInfo: loginInfoStore(pinia),
    progressInfo: progressInfoStore(),
    customInfo: customInfoStore(),
    sysPartitionInfo: sysPartitionStore(),
    imagesInfo: imagesStore(),
    filterEmpty: filterEmptyStore(),
    navigationBarStatus: navigationBarStore(),
    faultLevelInfo: faultLevelStore(),
    faultSpecInfo: faultSpecStore(),
    ossDataInfo: ossDataStore,
    modeFormInfo: modeFormStore,
    userPurviewInfo: userPurviewStore,
    deviceDetailErrorTipInfo: deviceDetailErrorTipStore()
};

export default store;
