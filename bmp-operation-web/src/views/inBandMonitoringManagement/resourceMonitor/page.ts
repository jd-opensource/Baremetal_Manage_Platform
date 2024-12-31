const ResourceMonitor = () => import(/* webpackChunkName: "ResourceMonitor" */ './resourceMonitor.vue');

export default {
    path: '/InBandMonitoringManagement/resourceMonitor',
    name: 'ResourceMonitor',
    component: ResourceMonitor,
    meta: {
        type: 'resourceMonitor'
    }
};
