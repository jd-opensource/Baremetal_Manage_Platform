const ImageList = () => import(/* webpackChunkName: "ImageList" */ './imageList.vue');

export default {
    path: '/ImageManagement/imageList',
    name: 'ImageList',
    component: ImageList,
    meta: {
        type: 'imageList'
    }
};
