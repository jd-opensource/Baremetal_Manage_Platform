class PushEmailOpt {
    emailPushDiaLog: Ref<boolean> = ref<boolean>(false);
    operate: Ref<string> = ref<string>('');
    switchLoading: Ref<boolean> = ref<boolean>(false);
    emailPush: Ref<boolean> = ref<boolean>(false);
    fn;

    constructor(emailPush: boolean, fn: Function) {
        this.emailPush.value = emailPush;
        this.fn = fn;
    }

    switchBeforeChange = () => {
        this.switchLoading.value = true;
        return new Promise(() => {
            this.emailPushDiaLog.value = true;
        })
    }

    cancel = () => {
        this.switchLoading.value = false;
        this.emailPushDiaLog.value = false;
    }

    sure = () => {
        this.switchLoading.value = false;
        this.emailPushDiaLog.value = false;
        this.emailPush.value = !this.emailPush.value;
        this.fn();
    }

    errorClick = () => {
        this.switchLoading.value = false;
        this.emailPushDiaLog.value = false;
    }
}

export default PushEmailOpt;
