const DeviceDetail = () => import(/* webpackChunkName: "DeviceDetail" */ './deviceDetail.vue');

export default {
    path: '/DeviceManagement/deviceDetail',
    name: 'DeviceDetail',
    component: DeviceDetail,
    meta: {
        type: 'deviceDetail'
    }
};
