import {language} from 'utils/publicClass.ts';
import {CurrencyType} from '@utils/publicType';

class SetInfo {
    title = (title: string) => {
        const titleData: CurrencyType = {
            'remove': language.t('removeRecovery.header.remove'),
            'recovery': language.t('removeRecovery.header.recovery')
        };
        return titleData[title];
    };

    /**
     * header头部标题
     * @return {string} xxx 对应标题的内容
    */
    headerTitle = (title: string) => {
        const headerTitle: CurrencyType = {
            'remove': language.t('removeRecovery.tooltip.remove'),
            'recovery': language.t('removeRecovery.tooltip.recovery')
        };
        return headerTitle[title];
    };

    tooltipCount = (title: string) => {
        const tip: CurrencyType = {
            'remove': language.t('removeRecovery.operateTip.remove'),
            'recovery': language.t('removeRecovery.operateTip.recovery')
        };
        return tip[title];
    };

    sureTip = (title: string) => {
        const tip: CurrencyType = {
            'remove': language.t('removeRecovery.checkbox.remove'),
            'recovery': language.t('removeRecovery.checkbox.recovery')
        };
        return tip[title];
    };
}

export default SetInfo;
