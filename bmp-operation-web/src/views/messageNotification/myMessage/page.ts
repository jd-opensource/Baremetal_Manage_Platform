const MyMessage = () => import(/* webpackChunkName: "MyMessage" */ './myMessage.vue');

export default {
    path: '/MessageNotification/myMessage',
    name: 'MyMessage',
    component: MyMessage,
    meta: {
        type: 'myMessage'
    }
};
