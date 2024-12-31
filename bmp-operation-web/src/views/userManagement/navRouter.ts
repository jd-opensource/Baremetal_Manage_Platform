/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const userNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.userManagement'),
        defaultImg: imgUrl('userDefault'),
        changeImg: imgUrl('user'),
        path: pathUrl('user'),
        otherPath: [pathUrl('user')],
        hasAdmin: true,
        isF: true,
    }
};

export default userNavData;