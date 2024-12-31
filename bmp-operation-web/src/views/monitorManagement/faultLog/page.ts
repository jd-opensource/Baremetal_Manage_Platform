const FaultLog = () => import(/* webpackChunkName: "FaultLog" */ './faultLog.vue');

export default {
    path: '/MonitorManagement/faultLog',
    name: 'FaultLog',
    component: FaultLog,
    meta: {
        type: 'faultLog'
    }
};
