import mitt, {Emitter, EventType} from 'mitt'; // 事件总线-兄弟间传值

const mit: Emitter<Record<EventType, unknown>> = mitt();

/**
 * ComponentCustomProperties用于扩充组件实例类型以支持自定义全局属性，否则会报警告
*/
declare module 'vue' {
    export interface ComponentCustomProperties {
        $Bus: typeof mit;
    }
}

const mittCommunication: (app: any) => Emitter<Record<EventType, unknown>> = (app: any) => app.config.globalProperties.$Bus = mit;

export default mittCommunication;
