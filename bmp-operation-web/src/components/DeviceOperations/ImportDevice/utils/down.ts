const downLoad = () => {

    const templateDownLoad = () => {
        const agreement: string = window.location.protocol;
        const url: string = `${agreement}//${window.location.host}/data/device_import.xlsx`;
        window.open(url, '_self');
    };

    return {
        templateDownLoad
    };
};

export default downLoad;
