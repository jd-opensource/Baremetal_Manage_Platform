import i18n from 'lib/i18n/index.ts'; // 国际化
import {UserType, RoleNameDataType} from './type';
import {CurrencyStatusType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class UserStaticData {
    // 手机号
    static cellPhoneType: UserType[] = [
        {
            value: 1,
            text: '086',
            label: global.t('addUser.phoneData.China') // 中国大陆
        },
        {
            value: 2,
            text: '852',
            label: global.t('addUser.phoneData.HongKong') // 中国香港
        },
        {
            value: 3,
            text: '853',
            label: global.t('addUser.phoneData.MacaMacao') // 中国澳门
        },
        {
            value: 4,
            text: '886',
            label: global.t('addUser.phoneData.Taiwan') // 中国台湾
        },
        {
            value: 5,
            text: '54',
            label: global.t('addUser.phoneData.Argentina') // 阿根廷
        },
        {
            value: 6,
            text: '61',
            label: global.t('addUser.phoneData.Australia') // 澳大利亚
        },
        {
            value: 7,
            text: '55',
            label: global.t('addUser.phoneData.Brazil') // 巴西
        },
        {
            value: 8,
            text: '49',
            label: global.t('addUser.phoneData.Germany') // 德国
        },
        {
            value: 9,
            text: '33',
            label: global.t('addUser.phoneData.France') // 法国
        },
        {
            value: 10,
            text: '82',
            label: global.t('addUser.phoneData.SouthKorea') // 韩国
        },
        {
            value: 11,
            text: '01',
            label: global.t('addUser.phoneData.Canada') // 加拿大
        },
        {
            value: 12,
            text: '1',
            label: global.t('addUser.phoneData.USA') // 美国
        },
        {
            value: 13,
            text: '52',
            label: global.t('addUser.phoneData.Mexico') // 墨西哥
        },
        {
            value: 14,
            text: '27',
            label: global.t('addUser.phoneData.SouthAfrica') // 南非
        },
        {
            value: 15,
            text: '81',
            label: global.t('addUser.phoneData.Japan') // 日本
        },
        {
            value: 16,
            text: '966',
            label: global.t('addUser.phoneData.SaudiArabia') // 沙特阿拉伯
        },
        {
            value: 17,
            text: '90',
            label: global.t('addUser.phoneData.Turkey') // 土耳其
        },
        {
            value: 18,
            text: '39',
            label: global.t('addUser.phoneData.Italy') // 意大利
        },
        {
            value: 19,
            text: '91',
            label: global.t('addUser.phoneData.India') // 印度
        },
        {
            value: 20,
            text: '62',
            label: global.t('addUser.phoneData.Indonesia') // 印度尼西亚
        },
        {
            value: 21,
            text: '44',
            label: global.t('addUser.phoneData.UnitedKiongdom') // 英国
        }
    ];

    // 用户筛选数据
    static userFilterData: RoleNameDataType[] = [
        {
            filterParams: 'role-admin-uuid-01',
            showInfo: 'admin',
            text: global.t('addUser.role.owner'), // 平台拥有者
            value: 1
        },
        {
            filterParams: 'role-user-uuid-01',
            showInfo: 'user',
            text: global.t('addUser.role.user'), // 普通用户
            value: 2
        },
        {
            filterParams: 'role-operator-uuid-01',
            showInfo: 'operator',
            text: global.t('addUser.role.operator'), // 运营人员
            value: 3
        }
    ];

    static phoneType: {type: string; name: string;}[] = [
        {
            type: '086',
            name: global.t('editUser.phoneData.China'),
        },
        {
            type: '852',
            name: global.t('editUser.phoneData.HongKong')
        },
    
        {
            type: '853',
            name: global.t('editUser.phoneData.MacaMacao'),
        },
        {
            type: '886',
            name: global.t('editUser.phoneData.Taiwan')
        },
        {
            type: '54',
            name: global.t('editUser.phoneData.Argentina'),
        },
        {
            type: '61',
            name: global.t('editUser.phoneData.Australia')
        },
        {
            type: '55',
            name: global.t('editUser.phoneData.Brazil'),
        },
        {
            type: '49',
            name: global.t('editUser.phoneData.Germany')
        },
    
        {
            type: '33',
            name: global.t('editUser.phoneData.France'),
        },
        {
            type: '82',
            name: global.t('editUser.phoneData.SouthKorea')
        },
        {
            type: '01',
            name: global.t('editUser.phoneData.Canada'),
        },
        {
            type: '1',
            name: global.t('editUser.phoneData.USA')
        },
        {
            type: '52',
            name: global.t('editUser.phoneData.Mexico'),
        },
        {
            type: '27',
            name: global.t('editUser.phoneData.SouthAfrica')
        },
    
        {
            type: '81',
            name: global.t('editUser.phoneData.Japan'),
        },
        {
            type: '966',
            name: global.t('editUser.phoneData.SaudiArabia')
        },
        {
            type: '90',
            name: global.t('editUser.phoneData.Turkey'),
        },
        {
            type: '39',
            name: global.t('editUser.phoneData.Italy')
        },
        {
            type: '91',
            name: global.t('editUser.phoneData.India'),
        },
        {
            type: '62',
            name: global.t('editUser.phoneData.Indonesia')
        },
    
        {
            type: '44',
            name: global.t('editUser.phoneData.UnitedKiongdom'),
        }
    ];

    static userList = [
        {
            name: global.t('userList.label.userName'),
            selected: true,
            disabled: true,
            label: 'userName',
            showTooltip: false
        },
        {
            name: global.t('userList.label.userId'),
            selected: true,
            disabled: false,
            label: 'userId',
            showTooltip: false
        },
        {
            name: global.t('userList.label.role'),
            selected: true,
            disabled: false,
            label: 'roleName',
            showTooltip: false
        },
        {
            name: global.t('userList.label.itemNum'),
            selected: true,
            disabled: false,
            label: 'projectCount',
            showTooltip: false
        },
        {
            name: global.t('userList.label.cellPhone'),
            selected: true,
            disabled: false,
            label: 'phone',
            showTooltip: false
        },
        {
            name: global.t('userList.label.email'),
            selected: true,
            disabled: false,
            label: 'email',
            showTooltip: false
        },
        {
            name: global.t('userList.label.desc'),
            selected: true,
            disabled: false,
            label: 'remark',
            showTooltip: false
        },
        {
            name: global.t('userList.label.createTime'),
            selected: true,
            disabled: false,
            label: 'createTime',
            showTooltip: false
        },
        {
            name: global.t('userList.label.operate.name'),
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];

    static userInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            createTime: status1,
            email: status1,
            operate: status2,
            phone: status1,
            projectCount: status1,
            remark: status1,
            userId: status1,
            roleName: status1,
            userName: status2
        }
    };

    static userTipData: string[] = [
        'userNameTip',
        'userIdTip',
        'phoneNumberTip',
        'emailTip',
        'descriptionTip'
    ];
};
export default UserStaticData;
