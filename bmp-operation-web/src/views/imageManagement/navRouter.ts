/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const imageNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.imageManagement'),
        defaultImg: imgUrl('imageDefault'),
        changeImg: imgUrl('imageManage'),
        path: pathUrl('imageList'),
        otherPath: [pathUrl('imageList'), pathUrl('imageDetail')]
    }
};

export default imageNavData;
