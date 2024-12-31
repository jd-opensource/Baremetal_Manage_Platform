class StorageOperate {

    constructor (props: {name: string; deviceType: string;}) {
        sessionStorage.setItem('propsName', props.name);
        sessionStorage.setItem('propsDeviceType', props.deviceType);
    };

    remove = () => {
        sessionStorage?.removeItem('repeat')
        sessionStorage?.removeItem('propsName');
        sessionStorage?.removeItem('propsDeviceType');
    };
};

export default StorageOperate;
