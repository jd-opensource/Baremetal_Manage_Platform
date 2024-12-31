import {CurrencyType} from '@utils/publicType';
import {language} from 'utils/publicClass.ts';

class SetTip {

    titleTip = (types: string) => {
        const titleTip: CurrencyType = {
            'open': language.t('batchOperate.openCloseRestart.open'),
            'close': language.t('batchOperate.openCloseRestart.close'),
            'restart': language.t('batchOperate.openCloseRestart.restart')
        };
        return titleTip[types];
    };

    /**
     * 开/关/重启操作返回对应的提示文案
     * @return {string} xxx 返回对应的提示信息
    */
    tooltipLeft = (types: string) => {
        const tooltipInfo: CurrencyType = {
            'open': language.t('batchOperate.openCloseRestart.open1'),
            'close': language.t('batchOperate.openCloseRestart.close1'),
            'restart': language.t('batchOperate.openCloseRestart.restart1')
        };
        return tooltipInfo[types];
    };
    
    tooltipRight = (types: string) => {
        const tooltipInfo: CurrencyType = {
            'open': language.t('batchOperate.openCloseRestart.open2'),
            'close': language.t('batchOperate.openCloseRestart.close2'),
            'restart': language.t('batchOperate.openCloseRestart.restart2')
        };
        return tooltipInfo[types];
    };
};

export default SetTip;
