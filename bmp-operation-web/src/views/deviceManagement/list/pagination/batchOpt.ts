import {language} from 'utils/publicClass.ts';

class BatchOperate {
    selectStatus: Ref<boolean> = ref<boolean>(false);
    emitValue: any;

    constructor (emitValue: any) {
        this.emitValue = emitValue;
    };

    selectHover = (val: boolean): boolean => {
       return this.selectStatus.value = val;
    };

    openTip = (status: number) => {
        return this.#resultTip(status, 'turnOn');
    };

    closeTip = (status: number) => {
        return this.#resultTip(status, 'turnOff');
    };

    restartTip = (status: number) => {
        return this.#resultTip(status, 'restart');
    };

    resetPasswordTip = (deviceList: any) => {
        const status: number = deviceList.resetPasswordBntDisabled.value && deviceList.reactiveArr.selectArr?.length;
        return this.#resultTip(status, 'resetPassword');
    };

    editInstanceNameTip = (deviceList: any) => {
        const status: number = deviceList.instanceNameBntDisabled.value && deviceList.reactiveArr.selectArr?.length;
        return this.#resultTip(status, 'instanceName');
    };

    #resultTip = (status: number, type: string) => {
        const text = new Map([
            [0, language.t('deviceList.batchOperate.tip.default')]
        ]);
        return text.get(status)?? language.t(`deviceList.batchOperate.tip.${type}`);
    };

    recoveryInstanceTip = (selectArr: any[]) => {
        const lockStatus: boolean = selectArr.every((item: {locked: string}) => !item.locked.localeCompare('locked'));
        const content = [
            [
                () => !selectArr.length,
                () => language.t('deviceList.batchOperate.tip.default')
            ],
            [
                () => lockStatus,
                () => language.t('deviceList.batchOperate.tip.lock')
            ],
            [
                () => !lockStatus,
                () => language.t('deviceList.batchOperate.tip.recovery')
            ]
        ];

        for (const key of content) {
            if (key[0]()) {
                return key[1]();
            }
        }
    };

    setClass = (status: boolean) => {
        const className = new Map([
            [true, 'drop-down-disabled']
        ]);
        return className.get(status)?? 'drop-down-operate-content';
    };

    openCloseRestartClick = (status: boolean, type: string) => {
        if (!status) return;
        this.emitValue('open-close-restart', type)
    };

    recoveryInstanceClick = (status: boolean) => {
        if (!status) return;
        this.emitValue('recovery-instance')
    };
};

export default BatchOperate;
