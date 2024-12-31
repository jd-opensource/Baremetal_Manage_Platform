import DeviceStaticData from 'staticData/device/index.ts';

class DropDownOpt {

    up = (manageStatus: string): boolean => {
        return ['in', 'putawayfail'].includes(manageStatus)
    };

    down = (manageStatus: string): boolean => {
        return ['putaway'].includes(manageStatus);
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

    running = (instanceStatus: string): boolean => {
        return ['running'].includes(instanceStatus);
    };

    resetInstance = (instanceStatus: string) => {
        return DeviceStaticData.resetInstance.includes(instanceStatus);
    };

    deviceRemove = (instanceStatus: string): boolean => {
        return ['running', 'stopped', 'Resetpasswd failed', 'Reinstall failed'].includes(instanceStatus)
    };

    recoveryInstance = (instanceStatus: string, locked: string): boolean => {
        return DeviceStaticData.recoveryInstance.includes(instanceStatus) && locked === 'unlocked';
    };

    deviceDelete = (manageStatus: string): boolean => {
        return ['in', 'putawayfail', 'removed'].includes(manageStatus);
    };
};

export default DropDownOpt;
