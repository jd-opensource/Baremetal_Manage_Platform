/**
 * @file
 * @author
*/

import {AxiosError} from 'axios';
import {getDateMinutes, methodsTotal, msgTip} from 'utils/index.ts';
import {RouterOperate, CurrentInstance, language} from 'utils/publicClass.ts';
import MessageStaticData from 'staticData/message/index.ts';
import AllStaticData from 'staticData/staticData.ts';

class MessageDetailOpt {
    instanceProxy = new CurrentInstance().proxy;
    path: string = '';
    router;
    code: Ref<number> = ref<number>(0);
    monitorCode: Ref<number> = ref<number>(0);
    msgId: Ref<string> = ref<string>('');
    activeType: Ref<string> = ref<string>('');
    newEndTime: Ref<string> = ref<string>('');
    activeDate: Ref<string> = ref<string>('');
    name: Ref<string> = ref<string>('');
    userName: string = localStorage?.getItem('userName')?? '';
    nextId: Ref<string> = ref<string>('');
    isLoading: Ref<boolean> = ref<boolean>(true);
    prevId: Ref<string> = ref<string>('');
    reactiveArr = reactive({
        tableData: [
            {
                messageSubType: '',
                faultType: '',
                instanceName: '',
                sn: '',
                instanceId: '',
                idcName: '',
                deviceId: '',
                content: ''
            }
        ]
    });
    btnData: {id: string; text: string}[] = MessageStaticData.btnData;
    templateData: {prop: string; label: string; hasShow: string; hasTime: boolean; hasOther: boolean;}[] = MessageStaticData.templateData;
    instanceTemplateData: {prop: string; label: string;}[] = MessageStaticData.instanceTemplateData;

    constructor() {
        const obj = localStorage?.getItem('messageInfo')?? '';
        this.path = this.instanceProxy.$defInfo.routerPath('myMessage');
        this.router = new RouterOperate(this.path);
        if (!obj?.length) {
            this.jumpRouter(this.path);
            return;
        }
        const {id, type, newDate, endTime} = JSON.parse(decodeURIComponent(obj));
        this.activeType.value = type;
        this.msgId.value = id;
        this.newEndTime.value = endTime;
        this.activeDate.value = newDate;
        this.name.value = this.userName ? window.atob(this.userName) : '';
        this.init();
        this.readMessageClick();
        onUnmounted(() => localStorage.removeItem('messageInfo'));
    }

    readMessageClick = async () => {
        await this.instanceProxy.$messageApi.doReadAPI({messageIds: [this.msgId.value]});
    }

    setIdcInfo = () => {
        const {idcName} = this.#getTableData();
        return `${idcName}${language.t('messageDetail.optTips.instance')}`;
    }

    inMonitor = () => {
        const {instanceName, instanceId} = this.#getTableData();
        return `${instanceName}/${instanceId}`
    }

    monitorCount = () => {
        const {messageSubType, content} = this.#getTableData();
        return `${language.t('messageDetail.optTips.current')}${messageSubType}-${content}`
    }

    setTitle = () => {
        const {activeType, activeDate} = this;
        const {messageSubType, faultType} = this.#getTableData();
        const obj = {
            'optMessage': `${messageSubType?.length ? '【' + messageSubType + `】${language.t('messageDetail.failureNotification')}` : '--'}`,
            'unexpiredExpired': this.setDays(activeDate.value),
            'faultMessage': `${faultType?.length ? '【' + faultType + `】${language.t('messageDetail.remind')}` : '--'}`,
            'inbondAlert': `${messageSubType?.length ? '【' + messageSubType + `】${language.t('messageDetail.alarmTips')}` : '--'}`
        };
        return (obj[`${activeType.value}` as keyof typeof obj] as string)?? '';
    }

    setCount = () => {
        const {activeType, name} = this;
        const type = new Map([
            // 用户${name.value}您好！操作失败，请重新操作或联系管理员处理
            ['optMessage', `${language.t('messageDetail.tip', [name.value])}`],
            ['unexpiredExpired', this.setActiveDate()]
        ]);
        // `用户${name.value}您好！`
        return type.get(activeType.value)?? `${language.t('messageDetail.tip0', [name.value])}`;
    }

    jumpRouter = (path: string) => {
        this.router.router.push(path);
        localStorage.removeItem('messageInfo');
    }

    setActiveDate = () => {
        const {name, activeDate, newEndTime} = this;
        const type = new Map([
            ['0', `${language.t('messageDetail.tip0', [name.value])}${language.t('messageDetail.tip1')}<span>${language.t('messageDetail.tip11')}</span>${language.t('messageDetail.tip2')}`]
        ]);
        return type.get(activeDate.value)?? `
            ${language.t('messageDetail.tip0', [name.value])}
            ${language.t('messageDetail.tip3', [activeDate.value, getDateMinutes(newEndTime.value), getDateMinutes(newEndTime.value)])}
        `;
    }

    setDesc = () => {
        const {activeType} = this;
        const type = new Map([
            ['optMessage', language.t('messageDetail.tip4')],
            ['unexpiredExpired', language.t('messageDetail.tip5')],
            ['inbondAlert', language.t('messageDetail.tip7')]
        ]);
        return type.get(activeType.value)?? language.t('messageDetail.tip6');
    }

    setName = () => {
        const {activeDate, activeType} = this;
        const {messageSubType, faultType, instanceName, sn, instanceId} = this.#getTableData();
        const obj = {
            'optMessage': `${messageSubType?.length ? '【' + messageSubType + `】${language.t('messageDetail.failureNotification')}${instanceName ? '（' + instanceName + '/' + sn + '）' : ''}${!instanceName ? '（' + sn + '）' : ''}` : '--'}`,
            'faultMessage': `${faultType?.length ? '【' + faultType + '】' + `${language.t('messageDetail.remind')}${sn ? '（' + sn + '）' : ''}` : '--'}`,
            'unexpiredExpired': this.setDays(activeDate.value),
            'inbondAlert': `${messageSubType?.length ? '【' + messageSubType + `】${language.t('messageDetail.alarmTips')}${instanceName ? '（' + instanceName + '/' + instanceId + '）' : ''}${!instanceName ? '（' + instanceId + '）' : ''}` : '--'}`,
        }
        return obj[`${activeType.value}` as keyof typeof obj];
    }

    jumpDeviceDetail = (path: string) => {
        const {sn, deviceId} = this.#getTableData();
        const query = `?sn=${sn}&type=performanceMonitoring&deviceId=${deviceId}`;
        this.router.router.push(`${path}${query}`);
    }

    #getTableData = () => {
        return this.reactiveArr.tableData[0];
    }

    setDays = (newDate: string) => {
        if (Object.is(newDate, null)) return '';
        const obj = new Map([
            ['90', language.t('messageDetail.day90')],
            ['30', language.t('messageDetail.day30')],
            ['15', language.t('messageDetail.day15')],
            ['0',  language.t('messageDetail.expired')]
        ]);
        return obj.get(newDate)?? '';
    }

    disabledSet = (id: string) => {
        const that = this;
        return !((this[`${id}` as keyof typeof that] as any).value !== '');
    }

    preNextClick = (id: string) => {
        const that = this;
        this.msgId.value = (this[`${id}` as keyof typeof that] as any).value;
        this.init();
        this.readMessageClick();
    }

    preClick = () => {
        this.msgId.value = this.prevId.value;
        this.init();
        this.readMessageClick();
    }

    nextClick = () => {
        this.msgId.value = this.nextId.value;
        this.init();
        this.readMessageClick();
    }

    setStatus = (time: number | undefined) => {
        if ([null, undefined].includes((time as undefined))) return '';
        const date: number = new Date().getTime() / 1000;
        if ((time as number) < date) {
            return language.t('messageDetail.status.expired');
        }
        return language.t('messageDetail.status.effective');
    }

    getQueryType = (messageType: string) => {
        const type = new Map([
            [language.t('messageDetail.faultMessages'), 'faultMessage'],
            [language.t('messageDetail.systemMessages'), 'unexpiredExpired'],
            [language.t('messageDetail.operationMessages'), 'optMessage'],
            [language.t('messageDetail.inbondAlert'), 'inbondAlert']
        ]);
        return type.get(messageType)
    }

    init = async () => {
        this.isLoading.value = true;
        try {
            const res = await this.instanceProxy.$messageApi.messageByIdAPI({messageId: this.msgId.value});
            if (res?.data?.result?.message) {
                const {message, nextMessageId, prevMessageId} = methodsTotal.lineConverting(res.data.result);
                this.nextId.value = nextMessageId;
                this.prevId.value = prevMessageId;
                this.activeType.value = (this.getQueryType(message.messageType) as string);
                this.activeDate.value = message.detail;
                this.newEndTime.value = message.licenseEndTime;
                this.reactiveArr.tableData = [message];
                const obj = {
                    type: (this.getQueryType(message.messageType) as string),
                    id: this.msgId.value,
                    newDate: message.detail
                }
                localStorage.setItem('messageInfo', encodeURIComponent(JSON.stringify(obj)));
                // if (Object.is(this.activeType.value, 'inbondAlert')) {
                //     this.getMonitor();
                // }
                return;
            }
        }
        catch (e) {
            const err = e as AxiosError;
            if (AllStaticData.noFoundData.includes(err?.message)) {
                this.jumpRouter(this.path);
                return;
            }
            err?.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
        }
    }
}

export default MessageDetailOpt;
