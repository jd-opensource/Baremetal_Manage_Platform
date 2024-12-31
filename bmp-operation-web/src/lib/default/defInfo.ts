/**
 * @file
 * @author
*/

import {AppType} from '../filters/filter';

const defInfo = (app: AppType) => app.config.globalProperties.$defInfo = {
    purview(value: string = '') {
        const purviewInfo: Map<string, string> = new Map([
            ['admin', 'role-admin-uuid-01'],
            ['operator', 'role-operator-uuid-01'],
            ['user', 'role-user-uuid-01']
        ]);
        return purviewInfo.get(value);
    },

    routerPath(value: string) {
        const url: Map<string, string> = new Map([
            ['login', '/Login/login'],
            ['idcList', '/IdcManagement/idcList'],
            ['idcDetail', '/IdcManagement/idcDetail'],
            ['modelList', '/ModelManagement/modelList'],
            ['modelDetail', '/ModelManagement/modelDetail'],
            ['addModel', '/ModelManagement/modelAdd'],
            ['editModel', '/ModelManagement/modelEdit'],
            ['addModelTemplate', '/ModelManagement/addModelTemplate'],
            ['imageList', '/ImageManagement/imageList'],
            ['imageDetail', '/ImageManagement/imageDetail'],
            ['deviceList', '/DeviceManagement/deviceList'],
            ['deviceDetail', '/DeviceManagement/deviceDetail'],
            ['hardwareStatus', '/MonitorManagement/hardwareStatus'],
            ['faultLog', '/MonitorManagement/faultLog'],
            ['faultRules', '/MonitorManagement/faultRules'],
            ['resourceMonitor', '/InBandMonitoringManagement/resourceMonitor'],
            ['historyAlarmInfo', '/InBandMonitoringManagement/historyAlarmInfo'],
            ['alarmRulesList', '/InBandMonitoringManagement/allAlarmRulesList'],
            ['alarmRulesDetail', '/InBandMonitoringManagement/allAlarmRulesDetail'],
            ['role', '/RoleManagement/roleList'],
            ['user', '/UserManagement/userList'],
            ['myMessage', '/MessageNotification/myMessage'],
            ['myMessageDetail', '/MessageNotification/messageDetail'],
            ['myMessageSettings', '/MessageNotification/messageSettings'],
            ['myProfile', '/PersonalCenter/myProfile'],
            ['securitySettings', '/PersonalCenter/securitySettings'],
            ['apiKey', '/PersonalCenter/apiKeyApi']
        ]);
        return url.get(value);
    },

    imagePath(value: string) {
        const imgUrl: Map<string, string> = new Map([
            ['loginOut', require('@/assets/img/loginImg/login-out.png')],
            ['imageManage', require('@/assets/img/navigationBarImg/image-manage.png')],
            ['descEdit', require('@/assets/img/uiImg/desc-edit.png')],
            ['warningTip', require('@/assets/img/diaLogImg/warning-tip.png')],
            ['device', require('@/assets/img/navigationBarImg/device.png')],
            ['lock', require('@/assets/img/uiImg/lock.png')],
            ['tip', require('@/assets/img/uiImg/tip.png')],
            ['centos', require('@/assets/img/uiImg/centos.png')],
            ['ubantu', require('@/assets/img/uiImg/ubantu.png')],
            ['debian', require('@/assets/img/uiImg/debian.png')],
            ['openEuler', require('@/assets/img/uiImg/openeuler.png')],
            ['windows', require('@/assets/img/uiImg/windows.png')],
            ['genericVersion', require('@/assets/img/uiImg/generic-version.png')],
            ['idcManage', require('@/assets/img/navigationBarImg/computer-room.png')],
            ['model', require('@/assets/img/navigationBarImg/model.png')],
            ['toolTip', require('@/assets/img/diaLogImg/tooltiip.png')],
            ['loginUser', require('@/assets/img/loginImg/user.png')],
            ['setUp', require('@/assets/img/listImg/set-up.png')],
            ['export', require('@/assets/img/listImg/export.png')],
            ['idcManageDef', require('@/assets/img/navigationBarImg/computer-room-x.png')],
            ['modelDefault', require('@/assets/img/navigationBarImg/default-model.png')],
            ['imageDefault', require('@/assets/img/navigationBarImg/default-image.png')],
            ['deviceDefault', require('@/assets/img/navigationBarImg/default-device.png')],
            ['surveillanceDef', require('@/assets/img/navigationBarImg/default-surveillance-manage-backups.png')],
            ['defaultInBandMonitoringDef', require('@/assets/img/navigationBarImg/default-in-band-monitoring.png')],
            ['defaultInBandMonitoringActive', require('@/assets/img/navigationBarImg/active-in-band-monitoring.png')],
            ['defaultCard', require('@/assets/img/navigationBarImg/default-card.png')],
            ['activeCard', require('@/assets/img/navigationBarImg/active-card.png')],
            ['roleDefault', require('@/assets/img/navigationBarImg/default-role.png')],
            ['userDefault', require('@/assets/img/navigationBarImg/default-user.png')],
            ['userCenter', require('@/assets/img/uiImg/user-center.png')],
            ['usercCenterActive', require('@/assets/img/uiImg/user-center-active.png')],
            ['password', require('@/assets/img/navigationBarImg/password.png')],
            ['surveillance', require('@/assets/img/navigationBarImg/surveillance-manage-backups.png')],
            ['searchIcon', require('@/assets/img/uiImg/search-icon.png')],
            ['clear', require('@/assets/img/uiImg/clear.png')],
            ['apiKey', require('@/assets/img/uiImg/api-key.png')],
            ['user', require('@/assets/img/navigationBarImg/user.png')],
            ['openEye', require('@/assets/img/loginImg/open-eye.png')],
            ['closeEye', require('@/assets/img/loginImg/close-eye.png')],
            ['loginClear', require('@/assets/img/loginImg/clear.png')],
            ['role', require('@/assets/img/navigationBarImg/role.png')],
            ['uiTooltip', require('@/assets/img/uiImg/tooltip.png')],
            ['arrowBottom', require('@/assets/img/listImg/arrow-bottom.png')],
            ['defArrowBottom', require('@/assets/img/uiImg/table-arrow-bottom.png')],
            ['userLogin', require('@/assets/img/loginImg/user-login.png')],
            ['shrinkExpand', require('@/assets/img/navigationBarImg/shrink-expand.png')],
            ['greenIcon', require('@/assets/img/uiImg/green.png')],
            ['redIcon', require('@/assets/img/uiImg/red.png')],
            ['overlayAactive', require('@/assets/img/uiImg/overlay-active.png')],
            ['overlay', require('@/assets/img/uiImg/overlay.png')],
            ['messageDef', require('@/assets/img/message/no-message.png')],
            ['messageTip', require('@/assets/img/message/message.png')],
            ['messageLight', require('@/assets/img/message/no-message-light.png')],
            ['collect', require('@/assets/img/uiImg/re-collect.png')],
            ['optSuccess', require('@/assets/img/uiImg/opt-success.png')],
            ['optFail', require('@/assets/img/uiImg/opt-fail.png')],
            ['optDoing', require('@/assets/img/uiImg/opt-doing.png')]
        ]);
        return imgUrl.get(value);
    },

    encryptDecrypt(num: number) {
        const info: Map<number, string[]> = new Map([
            [1, ['QAZPLMUNYRTVCBFG', '02835485DRFDHTYG']]
        ]);
        return info.get(num);
    }
};

export default defInfo;