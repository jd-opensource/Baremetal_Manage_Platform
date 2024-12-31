/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const roleNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.roleManagement'),
        defaultImg: imgUrl('roleDefault'),
        changeImg: imgUrl('role'),
        path: pathUrl('role'),
        otherPath: [pathUrl('role')],
        hasAdmin: true,
        isF: true,
    }
};

export default roleNavData;