import {language} from 'utils/publicClass.ts';
import RegularContent from 'utils/regular.ts';

class RulesCheck {
    iptTipVal: Ref<string> = ref<string>(language.t('batchOperate.instanceName.iptTip.text'));
    errorFlag: Ref<boolean> = ref<boolean>(false);
    rules: any = reactive<any>({
        instanceName: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ]
    });

    constructor() {
        this.rules.instanceName[0].validator = this.instanceNameCheck;
    }

    instanceNameCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
    
        const nameCheck = [
            [
                (value: string) => !value?.length,
                () => {
                    this.errorFlag.value = true;
                    this.iptTipVal.value = language.t('batchOperate.instanceName.iptTip.empty')
                    callback(new Error());
                }
            ],
            [
                (value: string) => (!RegularContent.instanceNameReg.test(value)),
                ()=> {
                    this.errorFlag.value = true;
                    this.iptTipVal.value = language.t('batchOperate.instanceName.iptTip.text');
                    callback(new Error());
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.iptTipVal.value = language.t('batchOperate.instanceName.iptTip.text');
                    this.errorFlag.value = false;
                    callback();
                }
            ]
        ]
    
        for (const key of nameCheck) {
            if (key[0](value)) {
                key[1](value)
                break;
            }
        }
    };
};

export default RulesCheck;
