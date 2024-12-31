/**
 * @file
 * @author
*/

class EditModelOpt {
    proxy;
    router;

    constructor(proxy: {$defInfo: {routerPath(arg0: string): void}}, router: {push(arg0: string): void}) {
        this.proxy = proxy;
        this.router = router;
    }

    editModelClick = <T extends {deviceCount: number; deviceType: string;}>(item: T) => {
        if (item.deviceCount !== 0) return;
        const path = this.proxy.$defInfo.routerPath('editModel')
        localStorage.setItem('model-item', JSON.stringify(item));
        this.router.push(`${path}?deviceType=${item.deviceType}`)
    }
}

export default EditModelOpt;
