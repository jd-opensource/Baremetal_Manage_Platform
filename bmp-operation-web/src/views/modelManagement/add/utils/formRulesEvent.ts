/**
 * @file
 * @author
*/

import {RegExpCheck} from './regExpCheck';
import RegularContent from 'utils/regular.ts';
import {RuleFormRefType} from '@utils/publicType';
import {RulesType, RuleFormType} from '../typeManagement';

type FormVType = {
    name: string;ruleForm: RuleFormType;
    rules: RulesType;
    ruleFormRef: Ref<RuleFormRefType | undefined>
};

class FormRulesEvent {
    formRulesOperate: FormVType;
    isLongTip: Ref<boolean> = ref<boolean>(false);
    isShowLoading: Ref<boolean> = ref<boolean>(false);
    deviceTypeLoading: Ref<boolean> = ref<boolean>(false);
    dataVolumeSizeFlag: Ref<boolean> = ref<boolean>(false);
    systemVolumeSizeFlag: Ref<boolean> = ref<boolean>(false);

    constructor (formV: FormVType) {
        this.formRulesOperate = formV;
    };

    dataVolumeSizeBlur = () => {
        nextTick(async () => {
            await this.formRulesOperate.ruleFormRef.value!.validateField('dataVolumeSize', (valid: string) => !valid
                ? this.dataVolumeSizeFlag.value = false
                : this.dataVolumeSizeFlag.value = true
            );
        });
    };

    nameChange = () => {
        this.verifyRepeat(1);
        this.verifyNameDevice(this.isShowLoading, 'name');
    };

    nameBlur = () => {
        new RegExpCheck().nameFlag.value = true;
        const nameArr = [
            [
                (value: string) => !value,
                () => this.isLongTip.value = false
            ],
            [
                (value: string) => !(RegularContent.modelNameReg.test(value)),
                () => {
                    this.isLongTip.value = true;
                }
            ],
            [
                (value: string) => value,
                () => {
                    new RegExpCheck().nameFlag.value = false;
                    this.isLongTip.value = false;
                }
            ]
        ];
        for (const key of nameArr) {
            if (key[0](this.formRulesOperate.ruleForm.name)) {
                key[1](this.formRulesOperate.ruleForm.name);
                break;
            }
        }
    };

    deviceTypeChange = () => {
        this.verifyRepeat(0);
        this.verifyNameDevice(this.deviceTypeLoading, 'deviceType');
    };

    verifyRepeat = (type: number) => {
        const name = this.formRulesOperate.ruleForm.name;
        const deviceType = this.formRulesOperate.ruleForm.deviceType;
        if (name === deviceType && name !== '' && deviceType !== '') {
            sessionStorage.setItem('repeat', 'true');
            type && this.deviceTypeChange();
            return;
        }
        sessionStorage.removeItem('repeat');
    };

    verifyNameDevice = (loading: {value: boolean}, type: string) => {
        loading.value = true;
        this.formRulesOperate.ruleFormRef.value!.validateField(type, (_: string) => {
            loading.value = false;
        });
    }
};

export default FormRulesEvent;
