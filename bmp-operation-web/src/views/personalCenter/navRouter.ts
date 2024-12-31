/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const personalCenterNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.userCenter'),
        defaultImg: imgUrl('userCenter'),
        changeImg: imgUrl('usercCenterActive'),
        path: [
            pathUrl('myProfile'),
            pathUrl('securitySettings'),
            pathUrl('apiKey'),
        ],
        otherPath: [''],
        firstIndex: 'PersonalCenter',
        childrenHasAdmin: true,
        isChildrenF: true,
        children: [
            {
                title: language.t('myProfile.header.myProfile'),
                defaultImg: imgUrl('userCenter'),
                changeImg: imgUrl('usercCenterActive'),
                path: pathUrl('myProfile'),
                otherPath: [pathUrl('myProfile')]
            },
            {
                title: language.t('securitySettings.header.securitySettings'),
                defaultImg: imgUrl('userCenter'),
                changeImg: imgUrl('usercCenterActive'),
                path: pathUrl('securitySettings'),
                otherPath: [pathUrl('securitySettings')]
            },
            {
                title: language.t('apiKey.header.apiKey'),
                defaultImg: imgUrl('userCenter'),
                changeImg: imgUrl('usercCenterActive'),
                path: pathUrl('apiKey'),
                otherPath: [pathUrl('apiKey')]
            }
        ]
    }
};

export default personalCenterNavData;
