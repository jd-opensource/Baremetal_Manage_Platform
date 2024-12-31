const ImageDetail = () => import(/* webpackChunkName: "ImageDetail" */ './imageDetail.vue');

export default {
    path: '/ImageManagement/imageDetail',
    name: 'ImageDetail',
    component: ImageDetail,
    meta: {
        type: 'imageDetail'
    }
};
