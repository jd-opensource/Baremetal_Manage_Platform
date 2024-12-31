/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const deviceNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.deviceManagement'),
        defaultImg: imgUrl('deviceDefault'),
        changeImg: imgUrl('device'),
        path: pathUrl('deviceList'),
        otherPath: [pathUrl('deviceList'), pathUrl('deviceDetail')]
    }
};

export default deviceNavData;
