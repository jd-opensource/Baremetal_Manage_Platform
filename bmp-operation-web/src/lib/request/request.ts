/**
 * @file
 * @author
*/

import idcApiCount from 'api/idc/request.ts';
import modelApiCount from 'api/model/request.ts';
import roleApiCount from 'api/role/request.ts';
import userApiCount from 'api/user/request.ts';
import loginUserApiCount from 'api/login/request.ts';
import faultLogApiCount from 'api/surveillance/faultLog/request.ts';
import hardwareStatusApiCount from 'api/surveillance/hardwareStatus/request.ts';
import licenseApiCount from 'api/userCenter/license/request.ts';
import messageApiCount from 'api/message/request.ts';

interface AppType {
    config: {
        globalProperties: {
            $idcApi: Function;
            $modelApi: Parameters<typeof modelApiCount>
            $roleApi: Function;
            $userApi: Function;
            $faultLogApi: Function;
            $loginUserApi: Function;
            $hardwareStatusApi: Function;
            $licenseApi: Parameters< typeof licenseApiCount>;
            $messageApi: Function;
        }
    }
};

export default (app: AppType) => {
    app.config.globalProperties.$idcApi = idcApiCount;
    app.config.globalProperties.$modelApi = modelApiCount;
    app.config.globalProperties.$roleApi = roleApiCount;
    app.config.globalProperties.$userApi = userApiCount;
    app.config.globalProperties.$faultLogApi = faultLogApiCount;
    app.config.globalProperties.$loginUserApi = loginUserApiCount;
    app.config.globalProperties.$hardwareStatusApi = hardwareStatusApiCount;
    app.config.globalProperties.$licenseApi = licenseApiCount;
    app.config.globalProperties.$messageApi = messageApiCount;
};
