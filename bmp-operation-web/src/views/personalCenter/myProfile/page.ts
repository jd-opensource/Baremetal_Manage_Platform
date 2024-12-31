const MyProfile = () => import(/* webpackChunkName: "MyProfile" */ './myProfile.vue');

export default {
    path: '/PersonalCenter/myProfile',
    name: 'MyProfile',
    component: MyProfile,
    meta: {
        type: 'myProfile'
    }
};
