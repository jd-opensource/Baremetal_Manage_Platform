/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const idcNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.computerRoomManagement'),
        defaultImg: imgUrl('idcManageDef'),
        changeImg: imgUrl('idcManage'),
        path: pathUrl('idcList'),
        otherPath: [
            pathUrl('idcList'),
            pathUrl('idcDetail')
        ]
    }
};

export default idcNavData;
