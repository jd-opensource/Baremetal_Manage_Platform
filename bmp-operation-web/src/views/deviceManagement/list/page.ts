const DeviceList = () => import(/* webpackChunkName: "DeviceList" */ './deviceList.vue');

export default {
    path: '/DeviceManagement/deviceList',
    name: 'DeviceList',
    component: DeviceList,
    meta: {
        type: 'deviceList'
    }
};
