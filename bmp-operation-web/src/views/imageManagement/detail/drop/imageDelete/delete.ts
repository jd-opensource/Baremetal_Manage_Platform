import {CurrentInstance} from 'utils/publicClass.ts';

const imageDelete = (router: any, baseInfo: any) => {
    const deleteImageDiaLog: Ref<boolean> = ref<boolean>(false);
    const instanceProxy = new CurrentInstance().proxy;

    const deleteImageClick = () => {
        deleteImageDiaLog.value = !deleteImageDiaLog.value;
    };

    const deleteImageCancel = (type: boolean) => {
        return deleteImageDiaLog.value = type;
    };

    const deleteImageSure = () => {
        const path: string = instanceProxy.$defInfo.routerPath('imageList');
        router.push(path);
    };

    const imageDeleteError = () => {
        baseInfo.getImageDetail();
    };

    return {
        deleteImageDiaLog,
        deleteImageClick,
        deleteImageCancel,
        deleteImageSure,
        imageDeleteError
    };
};

export default imageDelete;
