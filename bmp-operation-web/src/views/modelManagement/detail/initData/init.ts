import store from 'store/index.ts';

const initData = (router: any) => {
    const {query} = router?.currentRoute?.value;
    const data = {
        routerType: query?.type,
        deviceTypeId: query?.deviceTypeId,
        ossStore: store.ossDataInfo()
    };

    return {
        query,
        ...data
    }
};

export default initData;
