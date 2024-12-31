/**
 * @file
 * @author
*/

class ModelTemplateOpt {
    proxy;
    router;

    constructor(proxy: {$defInfo: {routerPath(arg0: string): void}}, router: {push(arg0: string): void}) {
        this.proxy = proxy;
        this.router = router;
    }

    addModelTemplateClick = <T extends {deviceType: string;}>(item: T) => {
        const path = this.proxy.$defInfo.routerPath('addModelTemplate')
        localStorage.setItem('model-item', JSON.stringify(item));
        this.router.push(`${path}?deviceType=${item.deviceType}`)
    }
}

export default ModelTemplateOpt;
