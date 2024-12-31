/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const userNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.messageCenter'),
        defaultImg: imgUrl('messageDef'),
        changeImg: imgUrl('messageLight'),
        path: pathUrl('myMessage'),
        otherPath: [pathUrl('myMessage'), pathUrl('myMessageDetail')],
        hasAdmin: false,
        isF: true,
    }
};

const adminUserNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.messageCenter'),
        defaultImg: imgUrl('messageDef'),
        changeImg: imgUrl('messageLight'),
        path: [
            pathUrl('myMessage'),
            pathUrl('myMessageSettings'),
            pathUrl('myMessageDetail')
        ],
        otherPath: [''],
        firstIndex: 'MessageNotification',
        childrenHasAdmin: false,
        isChildrenF: true,
        children: [
            {
                title: language.t('navigationBar.children.myMessages'),
                defaultImg: imgUrl('messageDef'),
                changeImg: imgUrl('messageLight'),
                path: pathUrl('myMessage'),
                otherPath: [pathUrl('myMessage'), pathUrl('myMessageDetail')],
            },
            {
                title: language.t('navigationBar.children.messageSet'),
                defaultImg: imgUrl('messageDef'),
                changeImg: imgUrl('messageLight'),
                path: pathUrl('myMessageSettings'),
                otherPath: [pathUrl('myMessageSettings')]
            }
        ]
    }
};

export {
    userNavData,
    adminUserNavData
}