const MessageDetail = () => import(/* webpackChunkName: "MessageDetail" */ './messageDetail.vue');

export default {
    path: '/MessageNotification/messageDetail',
    name: 'MessageDetail',
    component: MessageDetail,
    meta: {
        type: 'messageDetail'
    }
};
