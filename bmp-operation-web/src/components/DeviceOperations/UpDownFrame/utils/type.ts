interface TableDataType {
    sn: string;
    cabinet: string;
    uPosition: string;
    deviceTypeId: string;
    iloIp: string;
    privateIpv4: string;
};

type ParamsType = {
    deviceId: string;
    sns?: string;
};

interface PropsType {
    deviceId: string;
    title: string;
}

export {
    TableDataType,
    ParamsType,
    PropsType
};
