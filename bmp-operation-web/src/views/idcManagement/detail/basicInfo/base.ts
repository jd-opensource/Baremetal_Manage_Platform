/**
 * @file
 * @author
*/
import {language} from 'utils/publicClass.ts';

class BaseInfoOpt {
    static baseInfo = [
        {
            label: language.t('idcDetail.basicInfo.label.name'),
            info: 'name',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.englishCode'),
            info: 'nameEn',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.grade'),
            info: 'level',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.address'),
            info: 'address',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.createTime'),
            info: 'createTime',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.createPeople'),
            info: 'createdBy',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.updateTime'),
            info: 'updateTime',
            class: 'set-wrap info-name'
        },
        {
            label: language.t('idcDetail.basicInfo.label.updatePeople'),
            info: 'updatedBy',
            class: 'set-wrap info-name'
        }
    ]
}

export default BaseInfoOpt;
