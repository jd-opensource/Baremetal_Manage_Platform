/**
 * @file
 * @author
*/

import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyType5, CurrencyStatusType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class IdcStaticData {
    static machineRoomGradeData: CurrencyType5[] = [
        {
            value: 1,
            label: 'T3'
        },
        {
            value: 2,
            label: 'T3+'
        },
        {
            value: 3,
            label: 'T4'
        },
        {
            value: 4,
            label: 'T4+'
        },
        {
            value: 5,
            label: global.t('editIdc.iptInfo.other')
        }
    ];

    static idcList = [
        {
            name: global.t('idcList.content.name'),
            selected: true,
            disabled: true,
            label: 'name',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.englishCode'),
            selected: true,
            disabled: false,
            label: 'nameEn',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.grade'),
            selected: true,
            disabled: false,
            label: 'level',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.address'),
            selected: true,
            disabled: false,
            label: 'address',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.createTime'),
            selected: true,
            disabled: false,
            label: 'createTime',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.createPeople'),
            selected: true,
            disabled: false,
            label: 'createdBy',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.updateTime'),
            selected: true,
            disabled: false,
            label: 'updateTime',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.updatePeople'),
            selected: true,
            disabled: false,
            label: 'updatedBy',
            showTooltip: false
        },
        {
            name: global.t('idcList.content.operate.name'),
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];

    static idcInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            address: status1, // 机房地址
            createTime: status1, // 创建时间
            createdBy: status1, // 创建人
            level: status1, // 机房等级
            name: status2, // 机房名称
            nameEn: status1, // 机房英文名称
            updateTime: status1, // 更新时间
            updatedBy: status1, // 更新人
            operate: status2 // 操作
        }
    };

    static levelData: string[] = ['T3', 'T3+', 'T4', 'T4+'];

    static otherData: string[] = ['其它', 'Other'];

    static idcTipData: string[] = [
        'nameTip',
        'nameEnTip',
        'levelTip',
        'addressTip',
        'createdByTip',
        'updatedByTip'
    ];

    static templateData = [
        {
            prop: 'name',
            align: 'center',
            minWidth: '110',
            ifSelect: 'name',
            fixed: true,
            label: global.t('idcList.content.name'),
            hasUserTemplate: true,
            hasClick: true,
            hasShowTooltip: true
        },
        {
            prop: 'nameEn',
            align: 'center',
            minWidth: '120',
            ifSelect: 'nameEn',
            fixed: false,
            label: global.t('idcList.content.englishCode'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: true
        },
        {
            prop: 'level',
            align: 'center',
            minWidth: '110',
            ifSelect: 'level',
            fixed: false,
            label: global.t('idcList.content.grade'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: true
        },
        {
            prop: 'address',
            align: 'center',
            minWidth: '140',
            ifSelect: 'address',
            fixed: false,
            label: global.t('idcList.content.address'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: true
        },
        {
            prop: 'createTime',
            align: 'center',
            minWidth: '150',
            ifSelect: 'createTime',
            fixed: false,
            label: global.t('idcList.content.createTime'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: false
        },
        {
            prop: 'createdBy',
            align: 'center',
            minWidth: '90',
            ifSelect: 'createdBy',
            fixed: false,
            label: global.t('idcList.content.createPeople'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: true
        },
        {
            prop: 'updateTime',
            align: 'center',
            minWidth: '150',
            ifSelect: 'updateTime',
            fixed: false,
            label: global.t('idcList.content.updateTime'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: false
        },
        {
            prop: 'updatedBy',
            align: 'center',
            minWidth: '90',
            ifSelect: 'updatedBy',
            fixed: false,
            label: global.t('idcList.content.updatePeople'),
            hasUserTemplate: true,
            hasClick: false,
            hasShowTooltip: true
        },
        {
            prop: 'operate',
            align: 'center',
            minWidth: '100',
            ifSelect: 'operate',
            fixed: 'right',
            label: global.t('idcList.content.operate.name'),
            hasUserTemplate: true,
            hasClick: true,
            hasShowTooltip: false
        }
    ];
}

export default IdcStaticData;
