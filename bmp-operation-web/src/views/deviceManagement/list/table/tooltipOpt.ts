import {language} from 'utils/publicClass.ts';
import DeviceStaticData from 'staticData/device/index.ts';

class TooltipOpt {

    up = (item: {manageStatus: string}): boolean => {
        return ['in', 'putawayfail'].includes(item.manageStatus)
    };

    down = (item: {manageStatus: string}): boolean => {
        return ['putaway'].includes(item.manageStatus);
    };

    resetInstance = (instanceStatus: string): boolean => {
        return DeviceStaticData.resetInstanceStatus.includes(instanceStatus);
    };

    runningStatus = (instanceStatus: string): boolean => {
        return ['running'].includes(instanceStatus);
    };

    stoppedStatus = (instanceStatus: string): boolean => {
        return ['stopped'].includes(instanceStatus);
    };

    resetPasswordStatus = (instanceStatus: string): boolean => {
        return ['stopped', 'Resetpasswd failed'].includes(instanceStatus);
    };

    resetSystemStatus = (instanceStatus: string): boolean => {
        return ['stopped', 'Reinstall failed'].includes(instanceStatus);
    };

    deviceRemove = (instanceStatus: string): boolean => {
        return ['running', 'stopped', 'Resetpasswd failed', 'Reinstall failed'].includes(instanceStatus);
    };

    recoveryInstance = (item: {instanceStatus: string; locked: string;}): boolean => {
        return DeviceStaticData.recoveryInstance.includes(item.instanceStatus) && item.locked === 'unlocked';
    };

    recoveryInstanceContent = (locked: string) => {
        const text = new Map([
            ['locked', language.t('deviceList.tooltip.lock')]
        ]);
        return text.get(locked)?? language.t('deviceList.tooltip.recovery');
    };

    deviceDelete = (manageStatus: string): boolean => {
        return DeviceStaticData.deleteDevice.includes(manageStatus);
    };
};

export default TooltipOpt;
