/**
 * @file
 * @author
*/

class AddModelOpt {
    proxy;
    router;

    constructor(proxy: {$defInfo: {routerPath(arg0: string): void}}, router: {push(arg0: string): void}) {
        this.proxy = proxy;
        this.router = router;
    }

    addModelClick = () => {
        localStorage.removeItem('model-item');
        const path = this.proxy.$defInfo.routerPath('addModel');
        this.router.push(`${path}`)
    }
}

export default AddModelOpt;
