const MessageSettings = () => import(/* webpackChunkName: "MessageSettings" */ './messageSettings.vue');

export default {
    path: '/MessageNotification/messageSettings',
    name: 'MessageSettings',
    component: MessageSettings,
    meta: {
        type: 'messageSettings'
    }
};
