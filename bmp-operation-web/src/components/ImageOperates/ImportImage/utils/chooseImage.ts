import {msgTip} from 'utils/index.ts';
import {language, CurrentInstance} from 'utils/publicClass';
import {UpLoadsType, CurrencyType} from '@utils/publicType';

class ChooseImage {
    chooseImageFlag: Ref<boolean> = ref<boolean>(false);

    // 是否显示异常提示-默认false
    hasError: Ref<boolean> = ref<boolean>(false);
    // 上传
    uploads: Ref<UpLoadsType | undefined> = ref<UpLoadsType>();
    formRulesOperate: any;
    fn: any
    instanceProxy = new CurrentInstance().proxy;

    constructor (formRulesOperate: any, fn: any) {
        this.formRulesOperate = formRulesOperate;
        this.fn = fn;
    };
    
    /**
     * 上传的event事件
    */
    uploadEvent = (refs: {value: UpLoadsType}): void => {
        this.uploads.value = refs.value;
    };

    /**
     * 仅限上传一次，第二次上次替换第一次的值时-触发
     * @param {Object} files 上传的文件
    */
    handleExceed = (files: CurrencyType[]): void => {
        this.uploads.value!.clearFiles();
        const file: CurrencyType = files[0];
        this.uploads.value!.handleStart(file);
        this.hasError.value = false;
        this.chooseImageFlag.value = false;
    };

    handleChange = (_: unknown, data: Array<CurrencyType>) => {
        this.formRulesOperate.ruleForm.imageFile = data;
        this.uploads.value!.submit();
        this.chooseImageFlag.value = false;
    };

    /**
     * 点击删除上传信息时触发
    */
    removeChange = (): void => {
        // this.hasError.value = true;
        this.formRulesOperate.ruleForm.imageFile = [];
        this.chooseImageFlag.value = false;
    };

    /**
     * 上传成功的返回
    */
    uploadSuccess = (result: CurrencyType): void => {
        this.fn(result);
        // importImage.reactiveArr.fileParams = result;
        this.hasError.value = false;
        this.chooseImageFlag.value = false;
    };

    /**
     * 上传失败的返回
     * @param {Object} error 失败信息
     * @param {Object} uploadFile 上传的文件
     * @param {Object} uploadFiles 上传的文件信息
    */
    errorChange = (error: Error & string): void => {
        const status = JSON.parse(JSON.stringify(error)).status;
        if (status === 402) {
            const path: string = this.instanceProxy!.$defInfo.routerPath('login') as unknown as string;
            window.open(path, '_self');
            return;
        }
        const errorInfo: {error: {message: string;}} = JSON.parse(error.message);
        msgTip.error(errorInfo?.error?.message??language.t('importImage.errTip.error'));
        this.chooseImageFlag.value = false;
    };

    progressChange = (info: {percent: number | {toFixed(arg0: number): number;}}) => {
        this.hasError.value = false;
        this.chooseImageFlag.value = false;
        info.percent = Number(info.percent.toFixed(2));
    };
};

export default ChooseImage;
