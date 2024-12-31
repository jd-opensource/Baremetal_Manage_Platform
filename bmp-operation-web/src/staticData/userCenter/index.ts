import {language} from 'utils/publicClass.ts';
import {TabsDataType} from '@utils/publicType';

class UserCenterDataStatic {
    static tabsStaticData: TabsDataType['tabsPangeData'] = [
        {
            label: language.t('userCenter.tabs.myProfile'),
            name: 'myProfile'
        },
        {
            label: language.t('userCenter.tabs.securitySettings'),
            name: 'securitySettings'
        },
        {
            label: language.t('userCenter.tabs.apiKey'),
            name: 'apiKey'
        }
    ];

    static apiKeyTipData: string[] = ['nameTip'];
};

export default UserCenterDataStatic;
