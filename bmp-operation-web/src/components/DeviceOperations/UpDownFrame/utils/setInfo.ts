import {language} from 'utils/publicClass.ts';

class SetInfo {

    titleComputed = (title: string) => {
        const titleData: {[x: string]: string} = {
            'up': language.t('upDownFrame.titleTip.up'),
            'down': language.t('upDownFrame.titleTip.down'),
            'delete': language.t('upDownFrame.titleTip.delete')
        };
        return titleData[title];
    };

    /**
     * 主体内容
     * @return {string} xxx 对应的主体内容
    */
    countTitle = (title: string) => {
        const countTitle: {[x: string]: string | null} = {
            'up': null,
            'down': language.t('upDownFrame.countTitleTip.down'),
            'delete': language.t('upDownFrame.countTitleTip.delete')
        };
        return countTitle[title];
    };

    /**
     * header头部标题
     * @return {string} xxx 对应标题的内容
    */
    headerTitle = (title: string) => {
        const headerTitle: {[x: string]: string} = {
            'up': language.t('upDownFrame.headerTitleTip.up'),
            'down': language.t('upDownFrame.headerTitleTip.down'),
            'delete': language.t('upDownFrame.headerTitleTip.delete')
        };
        return headerTitle[title];
    };
};

export default SetInfo;
