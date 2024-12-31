const SecuritySettings = () => import(/* webpackChunkName: "SecuritySettings" */ './securitySettings.vue');

export default {
    path: '/PersonalCenter/securitySettings',
    name: 'SecuritySettings',
    component: SecuritySettings,
    meta: {
        type: 'securitySettings'
    }
};
