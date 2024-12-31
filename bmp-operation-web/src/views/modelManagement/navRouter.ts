/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const modelNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.modelManagement'),
        defaultImg: imgUrl('modelDefault'),
        changeImg: imgUrl('model'),
        path: pathUrl('modelList'),
        otherPath: [pathUrl('modelList'), pathUrl('modelDetail'), pathUrl('addModel'), pathUrl('editModel'), pathUrl('addModelTemplate')]
    }
};

export default modelNavData;
